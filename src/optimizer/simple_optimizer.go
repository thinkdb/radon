/*
 * Radon
 *
 * Copyright 2018 The Radon Authors.
 * Code is licensed under the GPLv3.
 *
 */

package optimizer

import (
	"github.com/thinkdb/radon/src/planner"
	"github.com/thinkdb/radon/src/router"

	"github.com/pkg/errors"

	"github.com/xelabs/go-mysqlstack/sqlparser"
	"github.com/xelabs/go-mysqlstack/xlog"
)

var (
	_ Optimizer = &SimpleOptimizer{}
)

// SimpleOptimizer is a simple optimizer who dispatches the plans
type SimpleOptimizer struct {
	log      *xlog.Log
	database string
	query    string
	node     sqlparser.Statement
	router   *router.Router
}

// NewSimpleOptimizer creates the new simple optimizer.
func NewSimpleOptimizer(log *xlog.Log, database string, query string, node sqlparser.Statement, router *router.Router) *SimpleOptimizer {
	return &SimpleOptimizer{
		log:      log,
		database: database,
		query:    query,
		node:     node,
		router:   router,
	}
}

// BuildPlanTree used to build plan trees for the query.
func (so *SimpleOptimizer) BuildPlanTree() (*planner.PlanTree, error) {
	log := so.log
	database := so.database
	query := so.query
	node := so.node
	routerNew := so.router

	plans := planner.NewPlanTree()
	switch node.(type) {
	case *sqlparser.DDL:
		node := planner.NewDDLPlan(log, database, query, node.(*sqlparser.DDL), routerNew)
		plans.Add(node)
	case *sqlparser.Insert:
		node := planner.NewInsertPlan(log, database, query, node.(*sqlparser.Insert), routerNew)
		plans.Add(node)
	case *sqlparser.Delete:
		node := planner.NewDeletePlan(log, database, query, node.(*sqlparser.Delete), routerNew)
		plans.Add(node)
	case *sqlparser.Update:
		node := planner.NewUpdatePlan(log, database, query, node.(*sqlparser.Update), routerNew)
		plans.Add(node)
	case *sqlparser.Select:
		nod := node.(*sqlparser.Select)
		selectNode := planner.NewSelectPlan(log, database, query, nod, routerNew)
		plans.Add(selectNode)
	default:
		return nil, errors.Errorf("optimizer.unsupported.query.type[%+v]", node)
	}

	// Build plantree.
	if err := plans.Build(); err != nil {
		return nil, err
	}
	return plans, nil
}
