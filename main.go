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
	"github.com/vbsw/remove"
	"github.com/vbsw/semver"
	"os"
	"os/exec"
)

var version semver.Version

func main() {
	version = semver.New(0, 1, 0)
	result := parseOSArgs()

	switch result.command {
	case none:
		result.message = "unknown state"
		printError(result)
	case info:
		printInfo(result)
	case daemon:
		startDeamon()
	case start:
		configHTTPServer(result)
		startHTTPServer(result)
	default:
		printError(result)
	}
}

func printInfo(result *parseResult) {
	fmt.Println(result.message)
}

func printError(result *parseResult) {
	fmt.Println("error:", result.message)
}

func startDeamon() {
	args := removeElement(os.Args[1:], "--daemon")
	prog := os.Args[0]
	for _, arg := range os.Args {
		println(arg)
	}
	println("-----")
	println(prog)
	for _, arg := range args {
		println(arg)
	}
	osCmd := exec.Command(prog, args...)
	err := osCmd.Start()
	println(err.Error())
	os.Exit(0)
}

func removeElement(list []string, element string) []string {
	listWOElement := list
	for i, arg := range list {
		if arg == element {
			listWOElement = remove.String(list, i)
			break
		}
	}
	return listWOElement
}
