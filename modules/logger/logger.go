// Copyright 2018 cg33.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package logger

import (
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
)

var (
	InfoLogger   = logrus.New()
	ErrorLogger  = logrus.New()
	AccessLogger = logrus.New()
)

func init() {
	InfoLogger.Out = os.Stdout
	ErrorLogger.Out = os.Stdout
	AccessLogger.Out = os.Stdout
}

func SetInfoLogger(path string, debug bool) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	if debug {
		InfoLogger.Out = io.MultiWriter(file, os.Stdout)
	} else {
		InfoLogger.Out = file
	}
}

func SetErrorLogger(path string, debug bool) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	if debug {
		ErrorLogger.Out = io.MultiWriter(file, os.Stdout)
	} else {
		ErrorLogger.Out = file
	}
}

func SetAccessLogger(path string, debug bool) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	if debug {
		AccessLogger.Out = io.MultiWriter(file, os.Stdout)
	} else {
		AccessLogger.Out = file
	}
}

func Error(err ...interface{}) {
	ErrorLogger.Errorln(err...)
}

func Info(info ...interface{}) {
	InfoLogger.Infoln(info...)
}

func Warn(info ...interface{}) {
	InfoLogger.Warnln(info...)
}
