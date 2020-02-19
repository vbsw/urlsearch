/*
 *          Copyright 2019, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package main

import (
	"fmt"
	"net/http"
	"time"
)

var (
	serverPort   string
	websiteTitle string
	workingDir   string
)

func startHTTPServer(port, title, wrkDir string) {
	serverPort = port
	websiteTitle = title
	workingDir = wrkDir

	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+serverPort, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	dateStr := dateString()
	fmt.Fprintln(w, "<!DOCTYPE HTML>")
	fmt.Fprintln(w, "<html>")
	fmt.Fprintln(w, "<head>")
	fmt.Fprintln(w, "\t<meta http-equiv=\"content-type\" content=\"text/html;charset=UTF-8\">")
	fmt.Fprintln(w, "\t<meta http-equiv=\"content-style-type\" content=\"text/css\">")
	fmt.Fprintln(w, "\t<meta http-equiv=\"content-script-type\" content=\"text/javascript\">")
	fmt.Fprintln(w, "\t<meta name=\"author\" content=\"Vitali Baumtrok\">")
	fmt.Fprintln(w, "\t<meta name=\"date\" content=\""+dateStr+"\">")
	fmt.Fprintln(w, "\t<meta name=\"viewport\" content=\"width=device-width\">")
	fmt.Fprintln(w, "\t<title>"+websiteTitle+"</title>")
	fmt.Fprintln(w, "</head>")
	fmt.Fprintln(w, "<body>")
	fmt.Fprintln(w, "\t<h1>"+websiteTitle+"</h1>")
	fmt.Fprintln(w, "\t<p>working directory:<br>"+workingDir+"</p>")
	fmt.Fprintln(w, "</body>")
	fmt.Fprintln(w, "</html>")
}

func dateString() string {
	t := time.Now()
	/* https://golang.org/src/time/format.go */
	dateStr := t.Format("2006-01-02")
	return dateStr
}
