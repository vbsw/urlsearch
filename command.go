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

func commandFromArgs(args []string, msgGen *messageGenerator) (*command, error) {
	var cmd *command
	params, err := parametersFromArgs(args)

	if err == nil {
		if params == nil {
			cmd = new(command)
			cmd.Info = true
			cmd.InfoMessage = msgGen.ShortInfo()

		} else if params.incompatibleArguments() {
			err = errors.New("wrong argument usage")

		} else if params.oneParamHasMultipleResults() {
			err = errors.New("wrong argument usage")

		} else {
			cmd = new(command)
			err = cmd.setValidcommand(params, msgGen)
		}
	}
	return cmd, err
}

func (cmd *command) setValidcommand(params *parameters, msgGen *messageGenerator) error {
	var err error

	if len(params.help) > 0 {
		cmd.Info = true
		cmd.InfoMessage = msgGen.Help()

	} else if len(params.version) > 0 {
		cmd.Info = true
		cmd.InfoMessage = msgGen.Version()

	} else if len(params.copyright) > 0 {
		cmd.Info = true
		cmd.InfoMessage = msgGen.Copyright()

	} else {
		cmd.background = len(params.background) > 0
		err = cmd.interpretTitle(params, err)
		err = cmd.interpretPort(params, err)
		err = cmd.interpretWorkingDir(params, err)
	}
	return err
}

func (cmd *command) interpretTitle(params *parameters, err error) error {
	if err == nil {
		if len(params.title) > 0 {
			cmd.title = params.title[0].Value
		} else {
			cmd.title = "URL Search"
		}
	}
	return err
}

func (cmd *command) interpretPort(params *parameters, err error) error {
	if err == nil {
		if len(params.port) > 0 {
			var port int
			port, err = strconv.Atoi(params.port[0].Value)

			// TODO: port checks
			if err == nil && port > 0 {
				cmd.port = strconv.Itoa(port)
			} else {
				err = errors.New("bad port number \"" + params.port[0].Value + "\"")
			}
		} else {
			cmd.port = "8080"
		}
	}
	return err
}

func (cmd *command) interpretWorkingDir(params *parameters, err error) error {
	if err == nil {
		cmd.workingDir, err = workingDirectory(params)

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

func workingDirectory(params *parameters) (string, error) {
	var path string
	var err error

	if len(params.workingDir) > 0 {
		path = params.workingDir[0].Value

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
