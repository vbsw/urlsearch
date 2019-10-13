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

type cmdParser struct {
	cmdType command
	message string
}

type command int

const (
	none  command = 0
	info  command = 1
	wrong command = 2
)

func newCmdParser() *cmdParser {
	cmd := new(cmdParser)
	cmd.cmdType = none
	return cmd
}

func (cmd *cmdParser) parseOSArgs() {
	osArgs := osargs.New()

	if len(osArgs.Args) == 0 {
		cmd.interpretZeroArguments()

	} else {
		results := parseFlaggedParameters(osArgs)
		restParams := osArgs.Rest(results.toArray())

		if len(restParams) == 0 {
			if len(osArgs.Args) == 1 {
				cmd.interpretOneArgument(results)

			} else {
				cmd.interpretManyArguments(results)
			}

		} else {
			cmd.cmdType = wrong
			cmd.message = "unknown argument \"" + restParams[0].Value + "\""
		}
	}
}

func (cmd *cmdParser) interpretZeroArguments() {
	cmd.cmdType = info
	cmd.message = "Run 'urlsearch --help' for usage."
}

func (cmd *cmdParser) interpretOneArgument(results *clResults) {
	if len(results.help) > 0 {
		cmd.cmdType = info
		cmd.message = "fsplit splits files into many, or combines them back to one\n\n"
		cmd.message = cmd.message + "USAGE\n"
		cmd.message = cmd.message + "  ursearch INFO\n\n"
		cmd.message = cmd.message + "INFO\n"
		cmd.message = cmd.message + "  -h           print this help\n"
		cmd.message = cmd.message + "  -v           print version\n"
		cmd.message = cmd.message + "  --copyright  print copyright\n\n"

	} else if len(results.version) > 0 {
		cmd.cmdType = info
		cmd.message = version.String()

	} else if len(results.copyright) > 0 {
		cmd.cmdType = info
		cmd.message = "Copyright 2019, Vitali Baumtrok (vbsw@mailbox.org).\n"
		cmd.message = cmd.message + "Distributed under the Boost Software License, version 1.0."

	} else {
		cmd.cmdType = wrong
		cmd.message = "unknown state (1)"
	}
}

func (cmd *cmdParser) interpretManyArguments(results *clResults) {
	if results.infoAvailable() {
		cmd.setWrongArgumentUsage()

	} else if results.oneParamHasMultipleResults() {
		cmd.setWrongArgumentUsage()

	} else {
		cmd.cmdType = wrong
		cmd.message = "unknown state (2)"
	}
}

func (cmd *cmdParser) setWrongArgumentUsage() {
	cmd.cmdType = wrong
	cmd.message = "wrong argument usage"
}

func parseFlaggedParameters(osArgs *osargs.OSArgs) *clResults {
	results := new(clResults)

	results.help = osArgs.Parse("-h", "--help", "-help", "help")
	results.version = osArgs.Parse("-v", "--version", "-version", "version")
	results.copyright = osArgs.Parse("--copyright", "-copyright", "copyright")

	return results
}
