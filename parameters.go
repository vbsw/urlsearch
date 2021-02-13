/*
 *        Copyright 2020, 2021 Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package main

import (
	"errors"
	"github.com/vbsw/cmdl"
)

// parameters holds parsed arguments.
type parameters struct {
	help       *cmdl.Parameter
	version    *cmdl.Parameter
	copyright  *cmdl.Parameter
	port       *cmdl.Parameter
	title      *cmdl.Parameter
	workingDir *cmdl.Parameter
	background *cmdl.Parameter
}

func parametersFromArgs(osArgs []string) (*parameters, error) {
	var err error
	cl := cmdl.NewFrom(osArgs)
	asgOp := cmdl.NewAsgOp(true, true, "=")
	args := new(parameters)

	args.help = cl.NewParam().Parse("-h", "--help", "-help", "help")
	args.version = cl.NewParam().Parse("-v", "--version", "-version", "version")
	args.copyright = cl.NewParam().Parse("--copyright", "-copyright", "copyright")
	args.port = cl.NewParam().ParsePairs(asgOp, "-p", "--port", "-port", "port")
	args.title = cl.NewParam().ParsePairs(asgOp, "-t", "--title", "-title", "title")
	args.workingDir = cl.NewParam().ParsePairs(asgOp, "-d", "--dir", "-dir", "dir")
	args.background = cl.NewParam().ParsePairs(asgOp, "-b", "--background", "-background", "background")

	unparsedArgs := cl.UnparsedArgs()

	if len(unparsedArgs) > 0 {
		err = errors.New("unknown argument \"" + unparsedArgs[0] + "\"")
	}
	return args, err
}

func (args *parameters) incompatibleParameters() bool {
	other := args.port.Available() || args.title.Available() || args.workingDir.Available() || args.background.Available()

	if args.help.Available() && (args.version.Available() || args.copyright.Available() || other) {
		return true

	} else if args.version.Available() && (args.help.Available() || args.copyright.Available() || other) {
		return true

	} else if args.copyright.Available() && (args.help.Available() || args.version.Available() || other) {
		return true
	}
	return false
}

func (args *parameters) oneParamHasMultipleResults() bool {
	return args.help.Count() > 1 || args.version.Count() > 1 || args.copyright.Count() > 1 || args.port.Count() > 1 || args.title.Count() > 1 || args.workingDir.Count() > 1 || args.background.Count() > 1
}

func argsToStringArray(param *cmdl.Parameter) []string {
	strings := make([]string, 0, param.Count())
	for _, value := range param.Values() {
		if len(value) > 0 {
			strings = append(strings, value)
		}
	}
	return strings
}
