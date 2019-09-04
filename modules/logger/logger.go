// Copyright 2018 cg33.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package logger

import (
	"github.com/chenhg5/go-admin/context"
	"github.com/mgutz/ansi"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strconv"
)

var (
	manager = map[string]*logrus.Logger{
		"info":   logrus.New(),
		"error":  logrus.New(),
		"access": logrus.New(),
	}
	sqlLogOpen = false
)

func init() {
	for _, l := range manager {
		l.Out = os.Stdout
	}
}

func SetInfoLogger(path string, debug bool) {
	SetLogger("info", path, debug)
}

func SetErrorLogger(path string, debug bool) {
	SetLogger("error", path, debug)
}

func SetAccessLogger(path string, debug bool) {
	SetLogger("access", path, debug)
}

func SetLogger(kind, path string, debug bool) {
	if debug {
		manager[kind].Out = io.MultiWriter(openFile(path), os.Stdout)
	} else {
		manager[kind].Out = openFile(path)
	}
}

func openFile(path string) *os.File {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return file
}

func OpenSqlLog() {
	sqlLogOpen = true
}

func Error(err ...interface{}) {
	manager["error"].Errorln(err...)
}

func Info(info ...interface{}) {
	manager["info"].Infoln(info...)
}

func Warn(info ...interface{}) {
	manager["info"].Warnln(info...)
}

func Access(ctx *context.Context) {
	manager["access"].Println("[GoAdmin]",
		ansi.Color(" "+strconv.Itoa(ctx.Response.StatusCode)+" ", "white:blue"),
		ansi.Color(" "+string(ctx.Method()[:])+"   ", "white:blue+h"),
		ctx.Path())
}

func LogSql(info ...interface{}) {
	if sqlLogOpen {
		manager["info"].Infoln(info...)
	}
}
