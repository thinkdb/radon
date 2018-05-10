/*
 * Radon
 *
 * Copyright 2018 The Radon Authors.
 * Code is licensed under the GPLv3.
 *
 */

package cmd

import (
	"github.com/thinkdb/radon/src/build"
	"fmt"

	"github.com/spf13/cobra"
)

func NewVersionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of radon client",
		Run:   versionCommandFn,
	}

	return cmd
}

func versionCommandFn(cmd *cobra.Command, args []string) {
	buildInfo := build.GetInfo()
	fmt.Printf("radoncli:[%+v]\n", buildInfo)
}
