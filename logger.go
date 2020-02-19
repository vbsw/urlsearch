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
)

func initLogger() {
}

func logError(msg string) {
	logMessage(dateTimeString() + " ERROR " + msg + "\n")
}

func logWarning(msg string) {
	logMessage(dateTimeString() + " WARNING " + msg + "\n")
}

func logInfo(msg string) {
	logMessage(dateTimeString() + " INFO " + msg + "\n")
}

func logMessage(msg string) {
	file, err := os.OpenFile(pref.logPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err == nil {
		defer file.Close()
		_, err = file.Write([]byte(msg))

	} else {
		fmt.Println("error:", err.Error())
	}
}
