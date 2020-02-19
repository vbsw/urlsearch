/*
 *          Copyright 2020, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package main

import (
	"os"
	"path/filepath"
)

type preferences struct {
	logPath    string
	prefPath   string
	urlsPath   string
	port       string
	title      string
	workingDir string
}

var pref preferences

func initPreferences(cmd *command) error {
	err := initLogPath(cmd.workingDir, logFileName)
	err = initPrefPath(cmd.workingDir, prefFileName, err)
	err = initURLsPath(cmd.workingDir, urlsDirName, err)

	if err == nil {
		pref.port = cmd.port
		pref.title = cmd.title
		pref.workingDir = cmd.workingDir
	}
	return err
}

func initLogPath(workingDir, fileName string) error {
	var err error
	pref.logPath = filepath.Join(workingDir, fileName)

	if !fileExists(pref.logPath) {
		var file *os.File
		file, err = os.Create(pref.logPath)

		if err == nil {
			file.Close()
		}
	}
	return err
}

func initPrefPath(workingDir, fileName string, err error) error {
	if err == nil {
		pref.prefPath = filepath.Join(workingDir, fileName)

		if !fileExists(pref.prefPath) {
			var file *os.File
			file, err = os.Create(pref.prefPath)

			if err == nil {
				file.Close()
			}
		}
	}
	return err
}

func initURLsPath(workingDir, dirName string, err error) error {
	if err == nil {
		pref.urlsPath = filepath.Join(workingDir, dirName)

		if !directoryExists(pref.urlsPath) {
			err = os.MkdirAll(pref.urlsPath, os.ModePerm)
		}
	}
	return err
}
