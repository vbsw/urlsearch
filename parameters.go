/*
 *          Copyright 2020, Vitali Baumtrok.
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
	help       []osargs.Parameter
	version    []osargs.Parameter
	copyright  []osargs.Parameter
	port       []osargs.Parameter
	title      []osargs.Parameter
	workingDir []osargs.Parameter
	background []osargs.Parameter
}

func parametersFromArgs(args []string) (*parameters, error) {
	var err error
	osArgs := osargs.NewFromArgs(args, " ", "=", "")
	params := new(parameters)

	params.help = osArgs.Parse("-h", "--help", "-help", "help")
	params.version = osArgs.Parse("-v", "--version", "-version", "version")
	params.copyright = osArgs.Parse("--copyright", "-copyright", "copyright")
	params.port = osArgs.ParsePairs("-p", "--port", "-port", "port")
	params.title = osArgs.ParsePairs("-t", "--title", "-title", "title")
	params.workingDir = osArgs.ParsePairs("-d", "--dir", "-dir", "dir")
	params.background = osArgs.ParsePairs("-b", "--background", "-background", "background")

	unparsedArgs := osArgs.Rest(params.help, params.version, params.copyright, params.port, params.title, params.workingDir, params.background)

	if len(unparsedArgs) > 0 {
		err = errors.New("unknown argument \"" + osArgs.Str[unparsedArgs[0]] + "\"")
	}
	return params, err
}

func (params *parameters) incompatibleArguments() bool {
	other := len(params.port) > 0 || len(params.title) > 0 || len(params.workingDir) > 0 || len(params.background) > 0

	if len(params.help) > 0 && (len(params.version) > 0 || len(params.copyright) > 0 || other) {
		return true

	} else if len(params.version) > 0 && (len(params.help) > 0 || len(params.copyright) > 0 || other) {
		return true

	} else if len(params.copyright) > 0 && (len(params.help) > 0 || len(params.version) > 0 || other) {
		return true
	}
	return false
}

func (params *parameters) oneParamHasMultipleResults() bool {
	return len(params.help) > 1 || len(params.version) > 1 || len(params.copyright) > 1 || len(params.port) > 1 || len(params.title) > 1 || len(params.workingDir) > 1 || len(params.background) > 1
}

func paramsToStringArray(params []osargs.Parameter) []string {
	strings := make([]string, 0, len(params))
	for _, param := range params {
		if len(param.Value) > 0 {
			strings = append(strings, param.Value)
		}
	}
	return strings
}
