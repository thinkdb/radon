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
	"testing"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/ant0ine/go-json-rest/rest/test"
	"github.com/stretchr/testify/assert"
	"github.com/xelabs/go-mysqlstack/driver"
	"github.com/xelabs/go-mysqlstack/sqlparser/depends/sqltypes"
	"github.com/xelabs/go-mysqlstack/xlog"
)

func TestCtlV1Explain(t *testing.T) {
	log := xlog.NewStdLog(xlog.Level(xlog.PANIC))
	fakeDbs, proxyNew, cleanup := proxy.MockProxy(log)
	defer cleanup()
	address := proxyNew.Address()

	// fakedbs.
	{
		fakeDbs.AddQueryPattern("create table .*", &sqltypes.Result{})
	}

	// create test table.
	{
		client, err := driver.NewConn("mock", "mock", address, "", "utf8")
		assert.Nil(t, err)
		query := "create table test.t1(id int, b int) partition by hash(id)"
		_, err = client.FetchAll(query, -1)
		assert.Nil(t, err)
	}

	{
		// server
		api := rest.NewApi()
		router, _ := rest.MakeRouter(
			rest.Post("/v1/radon/explain", ExplainHandler(log, proxyNew)),
		)
		api.SetApp(router)
		handler := api.MakeHandler()

		p := &explainParams{
			Query: "select id, k, avg, c, count from test.t1 group by id order by c limit 1",
		}
		recorded := test.RunRequest(t, handler, test.MakeSimpleRequest("POST", "http://localhost/v1/radon/explain", p))
		recorded.CodeIs(200)
	}
}
