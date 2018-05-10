/*
 * Radon
 *
 * Copyright 2018 The Radon Authors.
 * Code is licensed under the GPLv3.
 *
 */

package main

import (
	"github.com/thinkdb/radon/src/build"
	"github.com/thinkdb/radon/src/config"
	"github.com/thinkdb/radon/src/ctl"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"github.com/thinkdb/radon/src/proxy"
	"runtime"
	"syscall"

	"github.com/xelabs/go-mysqlstack/xlog"
)

var (
	flag_conf string
)

func init() {
	flag.StringVar(&flag_conf, "c", "", "radon config file")
	flag.StringVar(&flag_conf, "config", "", "radon config file")
}

func usage() {
	fmt.Println("Usage: " + os.Args[0] + " [-c|--config] <radon-config-file>")
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log := xlog.NewStdLog(xlog.Level(xlog.DEBUG))

	buildInfo := build.GetInfo()
	fmt.Printf("radon:[%+v]\n", buildInfo)

	// config
	flag.Usage = func() { usage() }
	flag.Parse()
	if flag_conf == "" {
		usage()
		os.Exit(0)
	}

	conf, err := config.LoadConfig(flag_conf)
	if err != nil {
		log.Panic("radon.load.config.error[%v]", err)
	}
	log.SetLevel(conf.Log.Level)

	// Proxy.
	proxyProxy := proxy.NewProxy(log, flag_conf, conf)
	proxyProxy.Start()

	// Admin portal.
	admin := ctl.NewAdmin(log, proxyProxy)
	admin.Start()

	// Handle SIGINT and SIGTERM.
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Info("radon.signal:%+v", <-ch)

	// Stop the proxy and httpserver.
	proxyProxy.Stop()
	admin.Stop()
}
