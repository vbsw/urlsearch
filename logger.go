/*
 *          Copyright 2020, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	loggerPath string
)

func initLogger(workingDir string) error {
	var err error
	loggerPath = filepath.Join(workingDir, loggerFileName)

	if !fileExists(loggerPath) {
		var file *os.File
		file, err = os.Create(loggerPath)

		if err == nil {
			file.Close()
		}
	}
	return err
}

func logError(msg string) {
	logMessage(dateTimeString() + " error: " + msg + "\n")
}

func logWarning(msg string) {
	logMessage(dateTimeString() + " warning: " + msg + "\n")
}

func logInfo(msg string) {
	logMessage(dateTimeString() + " info: " + msg + "\n")
}

func logMessage(msg string) {
	file, err := os.OpenFile(loggerPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err == nil {
		defer file.Close()
		_, err = file.Write([]byte(msg))

	} else {
		fmt.Println("error:", err.Error())
	}
}
