/*
 * Radon
 *
 * Copyright 2018 The Radon Authors.
 * Code is licensed under the GPLv3.
 *
 */

package cmd

import (
	"github.com/thinkdb/radon/src/ctl"
	"github.com/thinkdb/radon/src/proxy"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCmdTwopc(t *testing.T) {
	_, proxyProxy, cleanup := proxy.MockProxy(log)
	defer cleanup()

	admin := ctl.NewAdmin(log, proxyProxy)
	admin.Start()
	defer admin.Stop()
	time.Sleep(100)

	// enable.
	{
		cmd := NewTwopcCommand()
		_, err := executeCommand(cmd, "enable")
		assert.Nil(t, err)
	}
	// disable.
	{
		cmd := NewTwopcCommand()
		_, err := executeCommand(cmd, "disable")
		assert.Nil(t, err)
	}
}
