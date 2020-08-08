/*
 *          Copyright 2020, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package main

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"
)

type command struct {
	Info        bool
	InfoMessage string
	programCall string
	port        string
	title       string
	workingDir  string
	background  bool
}

func commandFromOSArgs(msgGen *messageGenerator) (*command, error) {
	cmd, err := commandFromArgs(os.Args[1:], msgGen)

	if err == nil {
		cmd.programCall, err = filepath.Abs(os.Args[0])
	}
	return cmd, err
}

func commandFromArgs(osArgs []string, msgGen *messageGenerator) (*command, error) {
	var cmd *command
	args, err := argumentsFromArgs(osArgs)

	if err == nil {
		if args == nil {
			cmd = new(command)
			cmd.Info = true
			cmd.InfoMessage = msgGen.ShortInfo()

		} else if args.incompatibleArguments() {
			err = errors.New("wrong argument usage")

		} else if args.oneParamHasMultipleResults() {
			err = errors.New("wrong argument usage")

		} else {
			cmd = new(command)
			err = cmd.setValidcommand(args, msgGen)
		}
	}
	return cmd, err
}

func (cmd *command) setValidcommand(args *arguments, msgGen *messageGenerator) error {
	var err error

	if len(args.help) > 0 {
		cmd.Info = true
		cmd.InfoMessage = msgGen.Help()

	} else if len(args.version) > 0 {
		cmd.Info = true
		cmd.InfoMessage = msgGen.Version()

	} else if len(args.copyright) > 0 {
		cmd.Info = true
		cmd.InfoMessage = msgGen.Copyright()

	} else {
		cmd.background = len(args.background) > 0
		err = cmd.interpretTitle(args, err)
		err = cmd.interpretPort(args, err)
		err = cmd.interpretWorkingDir(args, err)
	}
	return err
}

func (cmd *command) interpretTitle(args *arguments, err error) error {
	if err == nil {
		if len(args.title) > 0 {
			cmd.title = args.title[0].Value
		} else {
			cmd.title = "URL Search"
		}
	}
	return err
}

func (cmd *command) interpretPort(args *arguments, err error) error {
	if err == nil {
		if len(args.port) > 0 {
			var port int
			port, err = strconv.Atoi(args.port[0].Value)

			// TODO: port checks
			if err == nil && port > 0 {
				cmd.port = strconv.Itoa(port)
			} else {
				err = errors.New("bad port number \"" + args.port[0].Value + "\"")
			}
		} else {
			cmd.port = "8080"
		}
	}
	return err
}

func (cmd *command) interpretWorkingDir(args *arguments, err error) error {
	if err == nil {
		cmd.workingDir, err = workingDirectory(args)

		if err == nil {
			fileInfo, fileErr := os.Stat(cmd.workingDir)

			if fileErr == nil || !os.IsNotExist(fileErr) {
				if fileInfo != nil {
					if fileInfo.IsDir() {
						cmd.workingDir, err = filepath.Abs(cmd.workingDir)

					} else {
						err = errors.New("working directory is a file, but must be a directory")
					}
				} else {
					err = errors.New("working directory wrong syntax")
				}
			} else if os.IsNotExist(fileErr) {
				err = os.MkdirAll(cmd.workingDir, os.ModePerm)

			} else {
				err = errors.New("working directory does not exist")
			}
		}
	}
	return err
}

func workingDirectory(args *arguments) (string, error) {
	var path string
	var err error

	if len(args.workingDir) > 0 {
		path = args.workingDir[0].Value

	} else {
		path, err = os.UserHomeDir()

		if err == nil {
			path = filepath.Join(path, "urlsearch")
		} else {
			err = errors.New("default working directory - " + err.Error())
		}
	}
	return path, err
}
