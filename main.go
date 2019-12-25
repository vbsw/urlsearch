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
	"github.com/vbsw/contains"
	"github.com/vbsw/remove"
	"github.com/vbsw/semver"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
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
	case background:
		startInBackground()
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

func startInBackground() {
	if runtime.GOOS == "windows" {
		startInBackgroundWindows()
	} else {
		startInBackgroundDefault()
	}
}

func removeElements(list []string, elements ...string) []string {
	listWOElement := list
	for i, arg := range list {
		if contains.String(elements, arg) {
			listWOElement = remove.String(list, i)
			break
		}
	}
	return listWOElement
}

func startInBackgroundWindows() {
}

func startInBackgroundDefault() {
	progPathAbs, errPathAbs := filepath.Abs(os.Args[0])

	if errPathAbs == nil {
		args := removeElements(os.Args[1:], "-b", "--background", "-background", "background")
		params := make([]string, 0, 6+len(args))
		dir := filepath.Dir(progPathAbs)
		prog := filepath.Base(progPathAbs)
		params = append(params, "--start")
		params = append(params, "--background")
		params = append(params, "--chdir")
		params = append(params, dir)
		params = append(params, "--exec")
		params = append(params, prog)
		params = append(params, args...)
		/* I couldn't get syscall.ForkExec() to work :( */
		osCmd := exec.Command("start-stop-daemon", params...)
		errCmd := osCmd.Start()

		if errCmd != nil {
			println(errCmd.Error())
		}
	} else {
		fmt.Println(errPathAbs.Error())
	}
}
