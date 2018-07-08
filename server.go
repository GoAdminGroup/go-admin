package main

import (
	"fmt"
	"github.com/mgutz/ansi"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/reuseport"
	"goAdmin/auth"
	"goAdmin/config"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"sync/atomic"
	"time"
	"goAdmin/controllers"
)

type GracefulListener struct {
	// inner listener
	ln net.Listener

	// maximum wait time for graceful shutdown
	maxWaitTime time.Duration

	// this channel is closed during graceful shutdown on zero open connections.
	done chan struct{}

	// the number of open connections
	connsCount uint64

	// becomes non-zero when graceful shutdown starts
	shutdown uint64
}

// NewGracefulListener wraps the given listener into 'graceful shutdown' listener.
func newGracefulListener(ln net.Listener, maxWaitTime time.Duration) net.Listener {
	return &GracefulListener{
		ln:          ln,
		maxWaitTime: maxWaitTime,
		done:        make(chan struct{}),
	}
}

func (ln *GracefulListener) Accept() (net.Conn, error) {
	c, err := ln.ln.Accept()

	if err != nil {
		return nil, err
	}

	atomic.AddUint64(&ln.connsCount, 1)

	return &gracefulConn{
		Conn: c,
		ln:   ln,
	}, nil
}

func (ln *GracefulListener) Addr() net.Addr {
	return ln.ln.Addr()
}

// Close closes the inner listener and waits until all the pending open connections
// are closed before returning.
func (ln *GracefulListener) Close() error {
	err := ln.ln.Close()

	if err != nil {
		return nil
	}

	return ln.waitForZeroConns()
}

func (ln *GracefulListener) waitForZeroConns() error {
	atomic.AddUint64(&ln.shutdown, 1)

	if atomic.LoadUint64(&ln.connsCount) == 0 {
		close(ln.done)
		return nil
	}

	select {
	case <-ln.done:
		return nil
	case <-time.After(ln.maxWaitTime):
		return fmt.Errorf("cannot complete graceful shutdown in %s", ln.maxWaitTime)
	}

	return nil
}

func (ln *GracefulListener) closeConn() {
	connsCount := atomic.AddUint64(&ln.connsCount, ^uint64(0))

	if atomic.LoadUint64(&ln.shutdown) != 0 && connsCount == 0 {
		close(ln.done)
	}
}

type gracefulConn struct {
	net.Conn
	ln *GracefulListener
}

func (c *gracefulConn) Close() error {
	err := c.Conn.Close()

	if err != nil {
		return err
	}

	c.ln.closeConn()

	return nil
}

func GetFileSuffix(path string) string {
	suffix := filepath.Ext(path)
	rs := []rune(suffix)
	length := len(rs)
	return string(rs[1 : length-0])
}

func InitServer(addr string) {
	// create a fast listener ;)
	ln, err := reuseport.Listen("tcp4", addr)
	if err != nil {
		log.Fatalf("error in reuseport listener: %s", err)
	}

	// create a graceful shutdown listener
	duration := 30 * time.Second
	graceful := newGracefulListener(ln, duration)

	var fsHandler fasthttp.RequestHandler
	if !config.EnvConfig["PORTABLE"].(bool) {
		fs := &fasthttp.FS{
			Root:               "./resources",
			IndexNames:         []string{"index.html"},
			GenerateIndexPages: true,
			Compress:           false,
			AcceptByteRange:    false,
		}
		fsHandler = fs.NewRequestHandler()
	}

	go func() {
		fasthttp.Serve(graceful, func(ctx *fasthttp.RequestCtx) {

			// TODO: 区分静态，动态

			path := string(ctx.Path())

			if _, ok := GlobalRouter[path]; ok {
				if user, ok := auth.Filter(ctx); ok {
					GlobalRouter[path].Handler(ctx, path, GlobalRouter[path].Prefix, user)
				} else {
					ctx.Response.Header.Add("Location", "/login")
					ctx.Response.SetStatusCode(302)
				}
			} else if path == "/login" {
				controller.ShowLogin(ctx)
			} else if path == "/signup" {
				controller.Auth(ctx)
			} else if path == "/logout" {
				controller.Logout(ctx)
			} else if path == "/install" {
				controller.ShowInstall(ctx)
			} else if path == "/menu" {
				if user, ok := auth.Filter(ctx); ok {
					controller.ShowMenu(ctx, path, user)
				} else {
					ctx.Response.Header.Add("Location", "/login")
					ctx.Response.SetStatusCode(302)
				}
			} else if path == "/menu/new" {
				controller.NewMenu(ctx)
			} else if path == "/menu/delete" {
				controller.DeleteMenu(ctx)
			} else if path == "/menu/edit" {
				controller.EditMenu(ctx)
			} else if path == "/menu/edit/show" {
				controller.ShowEditMenu(ctx)
			} else if path == "/menu/order" {

			} else {
				if !config.EnvConfig["PORTABLE"].(bool) {
					fsHandler(ctx)
				} else {
					data, err := Asset("resources" + path)
					if err != nil {
						fmt.Println(err)
						ctx.Response.SetStatusCode(500)
						ctx.WriteString("error")
					} else {
						ctx.Response.Header.Set("Content-Type", "text/"+GetFileSuffix(path)+"; charset=utf-8")
						ctx.Response.SetStatusCode(200)
						ctx.Write(data)
					}
				}
			}
			log.Println("[GoAdmin]",
				ansi.Color(" "+strconv.Itoa(ctx.Response.StatusCode())+" ", "white:blue"),
				ansi.Color(" "+string(ctx.Method()[:])+"   ", "white:blue+h"),
				path)
		})
	}()

	pid := fmt.Sprintf("%d", os.Getpid())
	_, openErr := os.OpenFile("pid", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if openErr == nil {
		ioutil.WriteFile("pid", []byte(pid), 0)
	}

	osSignals := make(chan os.Signal)
	signal.Notify(osSignals, os.Interrupt)

	<-osSignals

	log.Printf("graceful shutdown signal received.\n")

	if err := graceful.Close(); err != nil {
		log.Fatalf("error with graceful close: %s", err)
	}
}
