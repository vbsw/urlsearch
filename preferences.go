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
	prefPath string
)

func initPreferences(workingDir string) {
	var err error
	prefPath = filepath.Join(workingDir, prefFileName)

	if !fileExists(prefPath) {
		var file *os.File
		file, err = os.Create(prefPath)

		if err == nil {
			file.Close()

		} else {
			logError(err.Error())
		}
	}
}
