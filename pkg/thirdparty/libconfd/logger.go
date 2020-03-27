// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache-style
// license that can be found in the LICENSE file.

// copy from https://github.com/chai2010/logger

package libconfd

import (
	"strings"
	"sync/atomic"

	log2 "pgxs.io/chassis/log"
)

var pkgLogger atomic.Value

//todo set on logger
func GetLogger() *log2.Entry {
	logger := log2.New()
	logger.SetLevel(5)
	return logger.Category("xx")
	//return pkgLogger.Load().(Logger)
}

type logLevelType uint32

const (
	logPanicLevel logLevelType = iota // invalid
	logFatalLevel
	logErrorLevel
	logWarnLevel
	logInfoLevel
	logDebugLevel
)

func (level logLevelType) Valid() bool {
	return level >= logPanicLevel && level <= logDebugLevel
}

func newLogLevel(name string) logLevelType {
	switch strings.ToUpper(name) {
	case "DEBUG", "":
		return logDebugLevel
	case "INFO":
		return logInfoLevel
	case "WARN":
		return logWarnLevel
	case "ERROR":
		return logErrorLevel
	case "PANIC":
		return logPanicLevel
	case "FATAL":
		return logFatalLevel
	}
	return logPanicLevel
}
