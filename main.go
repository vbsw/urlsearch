/*
 *          Copyright 2020, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

// Package urlsearch is compiled to an executable. It is a server to save and search URLs by keywords.
package main

import (
	"errors"
	"fmt"
	"github.com/vbsw/bgproc"
	"os"
)

type messageGenerator struct {
}

func main() {
	msgGen := new(messageGenerator)
	cmd, err := commandFromOSArgs(msgGen)

	if err == nil {
		if cmd.Info {
			fmt.Println(cmd.InfoMessage)

		} else if cmd.background {
			err = startHTTPServerInBackground(cmd)

		} else {
			err = initPreferences(cmd)

			if err == nil {
				initLogger()
				initURLs()
				startHTTPServer()
			}
		}
	}
	if err != nil {
		fmt.Println("error:", err.Error())
	}
}

func startHTTPServerInBackground(cmd *command) error {
	proc := bgproc.New(cmd.programCall)
	proc.AddArg("--port=" + cmd.port)
	proc.AddArg("--title=" + cmd.title)
	proc.AddArg("--dir=" + cmd.workingDir)
	err := proc.Start()

	if err != nil {
		err = errors.New("start detached process - " + err.Error())
	}
	return err
}

func fileExists(path string) bool {
	fileInfo, err := os.Stat(path)
	return (err == nil || !os.IsNotExist(err)) && fileInfo != nil && !fileInfo.IsDir()
}

func directoryExists(path string) bool {
	fileInfo, err := os.Stat(path)
	return (err == nil || !os.IsNotExist(err)) && fileInfo != nil && fileInfo.IsDir()
}

func (msg *messageGenerator) ShortInfo() string {
	return "Run \"urlsearch --help\" for usage."
}

func (msg *messageGenerator) Help() string {
	message := "URL Search is a server to save and search URLs by keywords.\n\n"
	message += "USAGE\n"
	message += "  urlsearch (INFO | {SERVER-PARAM})\n\n"
	message += "INFO\n"
	message += "  -h, --help         print this help\n"
	message += "  -v, --version      print version\n"
	message += "  --copyright        print copyright\n\n"
	message += "SERVER-PARAM\n"
	message += "  -p=N, --port=N     port number (N: integer)\n"
	message += "  -t=S, --title=S    page title (S: string)\n"
	message += "  -d=S, --dir=S      working directory (S: string)\n"
	message += "  -b, --background   run in background"
	return message
}

func (msg *messageGenerator) Version() string {
	return "0.1.0"
}

func (msg *messageGenerator) Copyright() string {
	message := "Copyright 2020, Vitali Baumtrok (vbsw@mailbox.org).\n"
	message += "Distributed under the Boost Software License, Version 1.0."
	return message
}
