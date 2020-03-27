// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package libconfd

import (
	//"reflect"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestNewDefaultConfig(t *testing.T) {
	_ = newDefaultConfig()
}

func TestConfig(t *testing.T) {
	tConfig := newDefaultConfig()

	p, err := LoadConfig("conf.test.yml")
	assert.NoError(t, err)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, tConfig.ConfDir, p.ConfDir)
	if !reflect.DeepEqual(p, tConfig) {
		t.Fatalf("expect = %#v, got = %#v", tConfig, p)
	}
}
