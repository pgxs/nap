// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a Apache-style
// license that can be found in the LICENSE file.

// copy from https://github.com/chai2010/logger

package libconfd

import (
	"sync/atomic"

	log2 "pgxs.io/chassis/log"
)

var pkgLogger atomic.Value

//todo set on logger
func GetLogger() *log2.Entry {
	logger := log2.New()
	return logger.Category("lib").Component("confd")
}
