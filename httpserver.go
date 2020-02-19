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

func startHTTPServer() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/archives", handleArchives)
	http.HandleFunc("/preferences", handlePreferences)
	http.ListenAndServe(":"+pref.port, nil)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	writeHead(w, r, pref.title)
	fmt.Fprintln(w, "\t\t\t<li id=\"selected\"><a href=\"/\">home</a></li>")
	fmt.Fprintln(w, "\t\t\t<li><a href=\"archives\">archives</a></li>")
	fmt.Fprintln(w, "\t\t\t<li><a href=\"preferences\">preferences</a></li>")
	fmt.Fprintln(w, "\t\t</ul>")
	fmt.Fprintln(w, "\t\t<form style=\"margin:1em 0 0 0\" action=\"/\">")
	fmt.Fprintln(w, "\t\t\t<input style=\"padding:0 0.2em\" type=\"text\" name=\"search\" size=\"50\" placeholder=\"search terms\" autofocus>")
	fmt.Fprintln(w, "\t\t\t<input style=\"padding:0 1.5em\" type=\"submit\" value=\"search\">")
	fmt.Fprintln(w, "\t\t</form>")
	writeFoot(w, r)
}

func handleArchives(w http.ResponseWriter, r *http.Request) {
	writeHead(w, r, pref.title+" - archives")
	fmt.Fprintln(w, "\t\t\t<li><a href=\"/\">home</a></li>")
	fmt.Fprintln(w, "\t\t\t<li id=\"selected\"><a href=\"archives\">archives</a></li>")
	fmt.Fprintln(w, "\t\t\t<li><a href=\"preferences\">preferences</a></li>")
	fmt.Fprintln(w, "\t\t</ul>")
	writeFoot(w, r)
}

func handlePreferences(w http.ResponseWriter, r *http.Request) {
	writeHead(w, r, pref.title+" - preferences")
	fmt.Fprintln(w, "\t\t\t<li><a href=\"/\">home</a></li>")
	fmt.Fprintln(w, "\t\t\t<li><a href=\"archives\">archives</a></li>")
	fmt.Fprintln(w, "\t\t\t<li id=\"selected\"><a href=\"preferences\">preferences</a></li>")
	fmt.Fprintln(w, "\t\t</ul>")
	fmt.Fprintln(w, "\t\t<p style=\"margin:1em 0 0 0\"><b>working directory</b><br>"+pref.workingDir+"</p>")
	writeFoot(w, r)
}

func writeHead(w http.ResponseWriter, r *http.Request, pageTitle string) {
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
	fmt.Fprintln(w, "\t<title>"+pageTitle+"</title>")
	fmt.Fprintln(w, "\t<style>")
	fmt.Fprintln(w, "\t\t* {margin:0;padding:0;text-decoration:none;border-collapse: collapse;list-style-type:none;}")
	fmt.Fprintln(w, "\t\ta, a:link, a:visited {color:black;}")
	fmt.Fprintln(w, "\t\tul#nav {height:1em;margin:0.5em 0 0 0;}")
	fmt.Fprintln(w, "\t\tul#nav li {float:left;}")
	fmt.Fprintln(w, "\t\tul#nav li a {padding:0 1em;}")
	fmt.Fprintln(w, "\t\tul#nav li a:hover, ul#nav li a:active {color:#ddd;background-color:#444;}")
	fmt.Fprintln(w, "\t\tul#nav li#selected {background-color:#ddd;}")
	fmt.Fprintln(w, "\t</style>")
	fmt.Fprintln(w, "</head>")
	fmt.Fprintln(w, "<body>")
	fmt.Fprintln(w, "\t<div style=\"padding:2em 0 0 5em\">")
	fmt.Fprintln(w, "\t\t<h1>"+pref.title+"</h1>")
	fmt.Fprintln(w, "\t\t<ul id=\"nav\">")
}
func writeFoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "\t</div>")
	fmt.Fprintln(w, "</body>")
	fmt.Fprintln(w, "</html>")
}

func dateString() string {
	t := time.Now()
	/* https://golang.org/src/time/format.go */
	dateStr := t.Format("2006-01-02")
	return dateStr
}

func dateTimeString() string {
	t := time.Now()
	/* https://golang.org/src/time/format.go */
	dateStr := t.Format("2006-01-02 03:04:05")
	return dateStr
}
