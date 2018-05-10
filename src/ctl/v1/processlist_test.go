/*
 * Radon
 *
 * Copyright 2018 The Radon Authors.
 * Code is licensed under the GPLv3.
 *
 */

package v1

import (
	"github.com/thinkdb/radon/src/proxy"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/ant0ine/go-json-rest/rest/test"
	"github.com/stretchr/testify/assert"
	"github.com/xelabs/go-mysqlstack/driver"
	"github.com/xelabs/go-mysqlstack/sqlparser/depends/sqltypes"
	"github.com/xelabs/go-mysqlstack/xlog"
)

func TestCtlV1Processlist(t *testing.T) {
	log := xlog.NewStdLog(xlog.Level(xlog.PANIC))
	fakeDbs, proxyNew, cleanup := proxy.MockProxy(log)
	defer cleanup()
	address := proxyNew.Address()

	// fakeDbs.
	{
		fakeDbs.AddQueryPattern("create table .*", &sqltypes.Result{})
		fakeDbs.AddQueryPattern("select .*", &sqltypes.Result{})
		fakeDbs.AddQueryDelay("select * from test.t1_0000 as t1", &sqltypes.Result{}, 1000)
	}

	// create test table.
	{
		client, err := driver.NewConn("mock", "mock", address, "", "utf8")
		assert.Nil(t, err)
		query := "create table test.t1(id int, b int) partition by hash(id)"
		_, err = client.FetchAll(query, -1)
		assert.Nil(t, err)
	}

	var wg sync.WaitGroup
	{
		wg.Add(2)
		go func() {
			defer wg.Done()
			client, err := driver.NewConn("mock", "mock", address, "", "utf8")
			assert.Nil(t, err)
			query := "select * from test.t1"
			_, err = client.FetchAll(query, -1)
			assert.Nil(t, err)
		}()
		go func() {
			defer wg.Done()
			client, err := driver.NewConn("mock", "mock", address, "", "utf8")
			assert.Nil(t, err)
			query := "select * from test.t1"
			_, err = client.FetchAll(query, -1)
			assert.Nil(t, err)
		}()
	}
	time.Sleep(time.Millisecond * 100)

	{
		api := rest.NewApi()
		router, _ := rest.MakeRouter(
			rest.Get("/v1/debug/processlist", ProcesslistHandler(log, proxyNew)),
		)
		api.SetApp(router)
		handler := api.MakeHandler()

		recorded := test.RunRequest(t, handler, test.MakeSimpleRequest("GET", "http://localhost/v1/debug/processlist", nil))
		recorded.CodeIs(200)

		got := recorded.Recorder.Body.String()
		log.Debug(got)
		assert.True(t, strings.Contains(got, "select * from test.t1"))
	}
	wg.Wait()
}
