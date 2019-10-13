/*
 *          Copyright 2019, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

// Package fsplit is compiled to an executable. It splits one file into many or concatenates them back to one.
package main

import (
	"fmt"
	"github.com/vbsw/semver"
)

var version semver.Version

func main() {
	version = semver.New(0, 1, 0)
	cmd := newCmdParser()

	cmd.parseOSArgs()

	switch cmd.cmdType {
	case none:
		cmd.message = "unknown state"
		printError(cmd)
	case info:
		printInfo(cmd)
	default:
		printError(cmd)
	}
}

func printInfo(cmd *cmdParser) {
	fmt.Println(cmd.message)
}

func printError(cmd *cmdParser) {
	fmt.Println("error:", cmd.message)
}
