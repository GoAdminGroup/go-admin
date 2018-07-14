package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/reuseport"
	"goAdmin/config"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/signal"
	"path/filepath"
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
	return string(rs[1: length-0])
}

func fsHandlerPortable(ctx *fasthttp.RequestCtx) {
	path := string(ctx.Path())
	data, err := Asset("resources" + path)
	if err != nil {
		fmt.Println(err)
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		ctx.SetContentType("application/json")
		ctx.WriteString(`{"code":404, "msg":"route not found"}`)
	} else {
		ctx.Response.Header.Set("Content-Type", "text/"+GetFileSuffix(path)+"; charset=utf-8")
		ctx.Response.SetStatusCode(200)
		ctx.Write(data)
	}
}

var fsHandler fasthttp.RequestHandler

func NotFoundHandler(ctx *fasthttp.RequestCtx) {

	defer controller.GlobalDeferHandler(ctx)

	if !config.EnvConfig["PORTABLE"].(bool) {
		if !PathExist("./resources" + string(ctx.Path())) {
			ctx.SetStatusCode(fasthttp.StatusNotFound)
			ctx.SetContentType("application/json")
			ctx.WriteString(`{"code":404, "msg":"route not found"}`)
		} else {
			fsHandler(ctx)
		}
	} else {
		fsHandlerPortable(ctx)
	}
}

func PathExist(_path string) bool {
	_, err := os.Stat(_path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
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

	router := InitRouter()
	router.NotFound = NotFoundHandler

	go func() {
		fasthttp.Serve(graceful, router.Handler)
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
