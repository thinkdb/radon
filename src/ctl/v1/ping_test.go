/*
 * Radon
 *
 * Copyright 2018 The Radon Authors.
 * Code is licensed under the GPLv3.
 *
 */

package v1

import (
	"errors"
	"github.com/thinkdb/radon/src/proxy"
	"testing"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/ant0ine/go-json-rest/rest/test"
	"github.com/xelabs/go-mysqlstack/sqlparser/depends/sqltypes"
	"github.com/xelabs/go-mysqlstack/xlog"
)

func TestCtlV1Ping(t *testing.T) {
	log := xlog.NewStdLog(xlog.Level(xlog.PANIC))
	fakeDbs, proxyNew, cleanup := proxy.MockProxy(log)
	defer cleanup()

	// fakeDbs.
	{
		fakeDbs.AddQueryPattern("select .*", &sqltypes.Result{})
	}

	{
		// server
		api := rest.NewApi()
		router, _ := rest.MakeRouter(
			rest.Get("/v1/radon/ping", PingHandler(log, proxyNew)),
		)
		api.SetApp(router)
		handler := api.MakeHandler()

		// client
		recorded := test.RunRequest(t, handler, test.MakeSimpleRequest("GET", "http://localhost/v1/radon/ping", nil))
		recorded.CodeIs(200)
	}
}

func TestCtlV1PingError(t *testing.T) {
	log := xlog.NewStdLog(xlog.Level(xlog.PANIC))
	fakeDbs, proxyNew, cleanup := proxy.MockProxy(log)
	defer cleanup()

	// fakeDbs.
	{
		fakeDbs.AddQueryError("select 1", errors.New("mock.ping.error"))
	}

	// server
	api := rest.NewApi()
	router, _ := rest.MakeRouter(
		rest.Get("/v1/radon/ping", PingHandler(log, proxyNew)),
	)
	api.SetApp(router)
	handler := api.MakeHandler()

	// 405.
	{
		recorded := test.RunRequest(t, handler, test.MakeSimpleRequest("POST", "http://localhost/v1/radon/ping", nil))
		recorded.CodeIs(405)
	}

	// 503.
	{
		recorded := test.RunRequest(t, handler, test.MakeSimpleRequest("GET", "http://localhost/v1/radon/ping", nil))
		recorded.CodeIs(503)
	}
}
