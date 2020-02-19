/*
 *          Copyright 2020, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func startHTTPServerInBackground(cmd *command) error {
	var err error

	if runtime.GOOS == "windows" {
		err = startInBackgroundWindows(cmd)

	} else if runtime.GOOS == "linux" {
		err = startInBackgroundLinux(cmd)

	} else {
		err = startInBackgroundOther(cmd)
	}
	return err
}

func startInBackgroundWindows(cmd *command) error {
	scriptPath, err := createVBSkript(cmd)

	if err == nil {
		osCmd := exec.Command("wscript", scriptPath)
		err = osCmd.Start()
	}
	return err
}

func startInBackgroundLinux(cmd *command) error {
	args := argsWOBackgroundFlag(cmd)
	prog := args[0]
	params := make([]string, 0, 6+len(args)-1)
	params = append(params, "--start")
	params = append(params, "--background")
	params = append(params, "--exec")
	params = append(params, prog)
	params = append(params, args[1:]...)

	// I couldn't get syscall.ForkExec() to work :(
	osCmd := exec.Command("start-stop-daemon", params...)
	err := osCmd.Start()

	return err
}

func startInBackgroundOther(cmd *command) error {
	args := argsWOBackgroundFlag(cmd)
	osCmd := exec.Command(args[0], args[1:]...)
	err := osCmd.Start()
	return err
}

func argsWOBackgroundFlag(cmd *command) []string {
	args := make([]string, 4)
	args[0] = cmd.programCall
	args[1] = "--port=" + cmd.port
	args[2] = "--title=" + cmd.title
	args[3] = "--dir=" + cmd.workingDir
	return args
}

func createVBSkript(cmd *command) (string, error) {
	script := "Set WshShell = CreateObject(\"WScript.Shell\")\r\n"
	script += "WshShell.Run \"\"\"" + cmd.programCall + "\"\" \"\"--port=" + cmd.port + "\"\" \"\"--title=" + cmd.title + "\"\" \"\"--dir=" + cmd.workingDir + "\"\"\", 0\r\n"
	script += "Set WshShell = Nothing\r\n"

	skriptPath := filepath.Join(cmd.workingDir, vbScriptName)
	skriptFile, err := os.OpenFile(skriptPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)

	if err == nil {
		defer skriptFile.Close()
		_, err = skriptFile.Write([]byte(script))
	}

	return skriptPath, err
}
