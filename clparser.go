/*
 *          Copyright 2019, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package main

import (
	"github.com/vbsw/osargs"
)

type parseResult struct {
	command executionCommand
	message string
}

type executionCommand int

const (
	none  executionCommand = 0
	info  executionCommand = 1
	wrong executionCommand = 2
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
		parameters := parseFlags(osArgs)
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
	result.command = info
	result.message = "Run 'urlsearch --help' for usage."
}

func interpretOneParameter(result *parseResult, parameters *clParameters) {
	if len(parameters.help) > 0 {
		result.command = info
		result.message = "fsplit splits files into many, or combines them back to one\n\n"
		result.message = result.message + "USAGE\n"
		result.message = result.message + "  ursearch INFO\n\n"
		result.message = result.message + "INFO\n"
		result.message = result.message + "  -h           print this help\n"
		result.message = result.message + "  -v           print version\n"
		result.message = result.message + "  --copyright  print copyright\n\n"

	} else if len(parameters.version) > 0 {
		result.command = info
		result.message = version.String()

	} else if len(parameters.copyright) > 0 {
		result.command = info
		result.message = "Copyright 2019, Vitali Baumtrok (vbsw@mailbox.org).\n"
		result.message = result.message + "Distributed under the Boost Software License, version 1.0."

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
		result.command = wrong
		result.message = "unknown state (2)"
	}
}

func setWrongArgumentUsage(result *parseResult) {
	result.command = wrong
	result.message = "wrong argument usage"
}

func parseFlags(osArgs *osargs.OSArgs) *clParameters {
	parameters := new(clParameters)

	parameters.help = osArgs.Parse("-h", "--help", "-help", "help")
	parameters.version = osArgs.Parse("-v", "--version", "-version", "version")
	parameters.copyright = osArgs.Parse("--copyright", "-copyright", "copyright")

	return parameters
}
