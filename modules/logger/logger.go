// Copyright 2019 GoAdmin Core Team.  All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package logger

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
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
	sqlLogOpen   = false
	accessLogOff = false
	infoLogOff   = false
	errorLogOff  = false
)

func init() {
	for _, l := range manager {
		l.Out = os.Stdout
	}
}

func SetInfoLogger(path string, debug, isInfoLogOn bool) {
	if path != "" {
		SetLogger("info", path, debug)
	}
	infoLogOff = isInfoLogOn
}

func SetErrorLogger(path string, debug, isErrorLogOn bool) {
	if path != "" {
		SetLogger("error", path, debug)
	}
	errorLogOff = isErrorLogOn
}

func SetAccessLogger(path string, debug, isAccessLogOn bool) {
	if path != "" {
		SetLogger("access", path, debug)
	}
	accessLogOff = isAccessLogOn
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
	if !errorLogOff {
		manager["error"].Errorln(err...)
	}
}

func Info(info ...interface{}) {
	if !infoLogOff {
		manager["info"].Infoln(info...)
	}
}

func Warn(info ...interface{}) {
	manager["info"].Warnln(info...)
}

func Access(ctx *context.Context) {
	if !accessLogOff {
		manager["access"].Println("["+constant.Title+"]",
			ansi.Color(" "+strconv.Itoa(ctx.Response.StatusCode)+" ", "white:blue"),
			ansi.Color(" "+string(ctx.Method()[:])+"   ", "white:blue+h"),
			ctx.Path())
	}
}

func LogSql(statement string, args []interface{}) {
	if sqlLogOpen && statement != "" {
		manager["info"].Infoln("["+constant.Title+"]", "statement", statement, "args", args)
	}
}
