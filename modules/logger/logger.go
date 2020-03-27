// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package logger

import (
	"github.com/GoAdminGroup/go-admin/context"
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

// SetInfoLogger set the info logger.
func SetInfoLogger(path string, debug, isInfoLogOn bool) {
	if path != "" {
		SetLogger("info", path, debug)
	}
	infoLogOff = isInfoLogOn
}

// SetErrorLogger set the error logger.
func SetErrorLogger(path string, debug, isErrorLogOn bool) {
	if path != "" {
		SetLogger("error", path, debug)
	}
	errorLogOff = isErrorLogOn
}

// SetAccessLogger set the access logger.
func SetAccessLogger(path string, debug, isAccessLogOn bool) {
	if path != "" {
		SetLogger("access", path, debug)
	}
	accessLogOff = isAccessLogOn
}

// SetLogger set the logger.
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

// OpenSQLLog set the sqlLogOpen true.
func OpenSQLLog() {
	sqlLogOpen = true
}

// Error print the error message.
func Error(err ...interface{}) {
	if !errorLogOff {
		manager["error"].Errorln(err...)
	}
}

// Info print the info message.
func Info(info ...interface{}) {
	if !infoLogOff {
		manager["info"].Infoln(info...)
	}
}

// Warn print the warning message.
func Warn(info ...interface{}) {
	manager["info"].Warnln(info...)
}

// Access print the access message.
func Access(ctx *context.Context) {
	if !accessLogOff {
		manager["access"].Println("[GoAdmin]",
			ansi.Color(" "+strconv.Itoa(ctx.Response.StatusCode)+" ", "white:blue"),
			ansi.Color(" "+string(ctx.Method()[:])+"   ", "white:blue+h"),
			ctx.Path())
	}
}

// LogSQL print the sql info message.
func LogSQL(statement string, args []interface{}) {
	if sqlLogOpen && statement != "" {
		manager["info"].Infoln("[GoAdmin]", "statement", statement, "args", args)
	}
}
