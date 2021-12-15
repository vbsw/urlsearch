/*
 *        Copyright 2020, 2021 Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package main

import (
	"errors"
	"github.com/vbsw/osargs"
)

// parameters holds parsed arguments.
type parameters struct {
	help       *osargs.Result
	version    *osargs.Result
	copyright  *osargs.Result
	port       *osargs.Result
	title      *osargs.Result
	workingDir *osargs.Result
	background *osargs.Result
}

func parametersFromArgs(osArgs []string) (*parameters, error) {
	var err error
	cl := new(osargs.Arguments)
	cl.Values = osArgs
	cl.Parsed = make([]bool, len(cl.Values))
	delimiter := osargs.NewDelimiter(true, true, "=")
	args := new(parameters)

	args.help = cl.Parse("-h", "--help", "-help", "help")
	args.version = cl.Parse("-v", "--version", "-version", "version")
	args.copyright = cl.Parse("--copyright", "-copyright", "copyright")
	args.port = cl.ParsePairs(delimiter, "-p", "--port", "-port", "port")
	args.title = cl.ParsePairs(delimiter, "-t", "--title", "-title", "title")
	args.workingDir = cl.ParsePairs(delimiter, "-d", "--dir", "-dir", "dir")
	args.background = cl.ParsePairs(delimiter, "-b", "--background", "-background", "background")

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

func argsToStringArray(param *osargs.Result) []string {
	strings := make([]string, 0, param.Count())
	for _, value := range param.Values {
		if len(value) > 0 {
			strings = append(strings, value)
		}
	}
	return strings
}
