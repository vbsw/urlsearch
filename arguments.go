/*
 *          Copyright 2020, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package main

import (
	"errors"
	"github.com/vbsw/cl"
)

// arguments holds parsed arguments.
type arguments struct {
	help       []cl.Argument
	version    []cl.Argument
	copyright  []cl.Argument
	port       []cl.Argument
	title      []cl.Argument
	workingDir []cl.Argument
	background []cl.Argument
}

func argumentsFromArgs(osArgs []string) (*arguments, error) {
	var err error
	ops := []string{" ", "=", ""}
	cmdLine := cl.New(osArgs)
	args := new(arguments)

	args.help = cmdLine.Parse("-h", "--help", "-help", "help")
	args.version = cmdLine.Parse("-v", "--version", "-version", "version")
	args.copyright = cmdLine.Parse("--copyright", "-copyright", "copyright")
	args.port = cmdLine.ParsePairs(ops, "-p", "--port", "-port", "port")
	args.title = cmdLine.ParsePairs(ops, "-t", "--title", "-title", "title")
	args.workingDir = cmdLine.ParsePairs(ops, "-d", "--dir", "-dir", "dir")
	args.background = cmdLine.ParsePairs(ops, "-b", "--background", "-background", "background")

	unparsedArgs := cmdLine.UnparsedArgsIndices()

	if len(unparsedArgs) > 0 {
		err = errors.New("unknown argument \"" + cmdLine.Args[unparsedArgs[0]] + "\"")
	}
	return args, err
}

func (args *arguments) incompatibleArguments() bool {
	other := len(args.port) > 0 || len(args.title) > 0 || len(args.workingDir) > 0 || len(args.background) > 0

	if len(args.help) > 0 && (len(args.version) > 0 || len(args.copyright) > 0 || other) {
		return true

	} else if len(args.version) > 0 && (len(args.help) > 0 || len(args.copyright) > 0 || other) {
		return true

	} else if len(args.copyright) > 0 && (len(args.help) > 0 || len(args.version) > 0 || other) {
		return true
	}
	return false
}

func (args *arguments) oneParamHasMultipleResults() bool {
	return len(args.help) > 1 || len(args.version) > 1 || len(args.copyright) > 1 || len(args.port) > 1 || len(args.title) > 1 || len(args.workingDir) > 1 || len(args.background) > 1
}

func argsToStringArray(clArgs []cl.Argument) []string {
	strings := make([]string, 0, len(clArgs))
	for _, param := range clArgs {
		if len(param.Value) > 0 {
			strings = append(strings, param.Value)
		}
	}
	return strings
}
