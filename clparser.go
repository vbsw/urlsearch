/*
 *          Copyright 2019, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package main

import (
	"github.com/vbsw/osargs"
	"os"
	"runtime"
	"strconv"
)

type parseResult struct {
	command executionCommand
	message string
	port    int
	title   string
	dir     string
}

type executionCommand int

const (
	none  executionCommand = 0
	info  executionCommand = 1
	start executionCommand = 2
	wrong executionCommand = 3
)

func newParseResult() *parseResult {
	result := new(parseResult)
	result.command = none
	return result
}

func parseOSArgs() *parseResult {
	osArgs := osargs.New()
	result := newParseResult()

	if len(osArgs.Args) == 0 {
		interpretZeroArguments(result)

	} else {
		parameters := parseParameters(osArgs)
		restParams := osArgs.Rest(parameters.toArray())

		if len(restParams) == 0 {
			if len(osArgs.Args) == 1 {
				interpretOneParameter(result, parameters)

			} else {
				interpretManyParameters(result, parameters)
			}

		} else {
			result.command = wrong
			result.message = "unknown argument \"" + restParams[0].Value + "\""
		}
	}
	return result
}

func interpretZeroArguments(result *parseResult) {
	result.command = start
	result.port = defaultPort
	result.title = defaultTitle
}

func interpretOneParameter(result *parseResult, parameters *clParameters) {
	if len(parameters.help) > 0 {
		result.command = info
		result.message = "fsplit splits files into many, or combines them back to one\n\n"
		result.message = result.message + "USAGE\n"
		result.message = result.message + "  ursearch (INFO | {SERVER-PARAM}\n\n"
		result.message = result.message + "INFO\n"
		result.message = result.message + "  -h           print this help\n"
		result.message = result.message + "  -v           print version\n"
		result.message = result.message + "  --copyright  print copyright\n\n"
		result.message = result.message + "SERVER-PARAM\n"
		result.message = result.message + "  -p=N         port number (N is an integer)\n"
		result.message = result.message + "  -t=S         page title (S is a string)\n\n"

	} else if len(parameters.version) > 0 {
		result.command = info
		result.message = version.String()

	} else if len(parameters.copyright) > 0 {
		result.command = info
		result.message = "Copyright 2019, Vitali Baumtrok (vbsw@mailbox.org).\n"
		result.message = result.message + "Distributed under the Boost Software License, version 1.0."

	} else if len(parameters.port) > 0 {
		interpretPath(result, parameters)
		interpretPort(result, parameters)
		interpretTitle(result, parameters)

		if result.command == none {
			result.command = start
		}

	} else {
		result.command = wrong
		result.message = "unknown state (1)"
	}
}

func interpretManyParameters(result *parseResult, parameters *clParameters) {
	if parameters.infoAvailable() {
		setWrongArgumentUsage(result)

	} else if parameters.anyParameterMultiple() {
		setWrongArgumentUsage(result)

	} else {
		interpretPath(result, parameters)
		interpretPort(result, parameters)
		interpretTitle(result, parameters)

		if result.command == none {
			result.command = start
		}
	}
}

func setWrongArgumentUsage(result *parseResult) {
	result.command = wrong
	result.message = "wrong argument usage"
}

func parseParameters(osArgs *osargs.OSArgs) *clParameters {
	parameters := new(clParameters)
	operator := osargs.NewAsgOp(" ", "", "=")

	parameters.help = osArgs.Parse("-h", "--help", "-help", "help")
	parameters.version = osArgs.Parse("-v", "--version", "-version", "version")
	parameters.copyright = osArgs.Parse("--copyright", "-copyright", "copyright")
	parameters.port = osArgs.ParsePairs(operator, "-p", "--port", "-port", "port")
	parameters.title = osArgs.ParsePairs(operator, "-t", "--title", "-title", "title")
	parameters.dir = osArgs.ParsePairs(operator, "-d", "--dir", "-dir", "dir")

	return parameters
}

func interpretPath(result *parseResult, parameters *clParameters) {
	if result.command == none {
		if len(parameters.dir) > 0 {
			dir := parameters.dir[0].Value

			if directoryExists(dir) {
				result.dir = dir

			} else {
				interpretPathError(result, dir)
			}

		} else {
			createDefaultWorkingDirectory(result)
		}
	}
}

func createDefaultWorkingDirectory(result *parseResult) {
	dir, err := os.UserHomeDir()

	if err == nil {
		if directoryExists(dir) {
			if runtime.GOOS == "windows" {
				winDir := dir + "/Documents"

				if directoryExists(winDir) {
					dir = winDir
				}
			}
			dir = dir + "/urls"

			if directoryExists(dir) {
				result.dir = unifySlashes(dir)

			} else {
				err = os.Mkdir(dir, os.ModeDir)

				if err == nil {
					result.dir = unifySlashes(dir)

				} else {
					result.command = wrong
					result.message = err.Error()
				}
			}

		} else {
			interpretPathError(result, dir)
		}

	} else {
		result.command = wrong
		result.message = err.Error()
	}
}

func unifySlashes(path string) string {
	pathBytes := []byte(path)
	var oldB, newB byte

	for _, b := range pathBytes {
		if b == '\\' {
			oldB = '/'
			newB = '\\'
			break

		} else if b == '/' {
			oldB = '\\'
			newB = '/'
			break
		}
	}
	for i, b := range pathBytes {
		if b == oldB {
			pathBytes[i] = newB
		}
	}
	return string(pathBytes)
}

func interpretPathError(result *parseResult, dir string) {
	if fileExists(dir) {
		result.command = wrong
		result.message = "working directory is not a directory \"" + dir + "\""

	} else {
		result.command = wrong
		result.message = "working directory does not exist \"" + dir + "\""
	}
}

func interpretPort(result *parseResult, parameters *clParameters) {
	if result.command == none {
		if len(parameters.port) > 0 {
			port, err := strconv.Atoi(parameters.port[0].Value)
			if err == nil {
				result.port = abs(port)
			} else {
				result.command = wrong
				result.message = "can't parse port number"
			}
		} else {
			result.port = defaultPort
		}
	}
}

func interpretTitle(result *parseResult, parameters *clParameters) {
	if result.command == none {
		if len(parameters.title) > 0 {
			result.title = parameters.title[0].Value
		} else {
			result.title = defaultTitle
		}
	}
}

func abs(value int) int {
	if value > 0 {
		return value
	}
	return -value
}

func directoryExists(path string) bool {
	fileInfo, err := os.Stat(path)
	return (err == nil || !os.IsNotExist(err)) && fileInfo.IsDir()
}

func fileExists(path string) bool {
	fileInfo, err := os.Stat(path)
	return (err == nil || !os.IsNotExist(err)) && !fileInfo.IsDir()
}
