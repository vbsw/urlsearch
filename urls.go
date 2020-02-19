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

var (
	urlsPath string
)

func initURLs(workingDir string) {
	urlsPath = filepath.Join(workingDir, urlsDirName)

	if !directoryExists(urlsPath) {
		err := os.MkdirAll(urlsPath, os.ModePerm)

		if err != nil {
			logError(err.Error())
		}
	}
}
