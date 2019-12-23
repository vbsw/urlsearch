# URL Search

[![GoDoc](https://godoc.org/github.com/vbsw/urlsearch?status.svg)](https://godoc.org/github.com/vbsw/urlsearch) [![Go Report Card](https://goreportcard.com/badge/github.com/vbsw/urlsearch)](https://goreportcard.com/report/github.com/vbsw/urlsearch) [![Stability: Experimental](https://masterminds.github.io/stability/experimental.svg)](https://masterminds.github.io/stability/experimental.html)

## About
URL Search is a server to save and search URLs by keywords. URL Search is published on <https://github.com/vbsw/urlsearch>.

## Copyright
Copyright 2019, Vitali Baumtrok (vbsw@mailbox.org).

URL Search is distributed under the Boost Software License, version 1.0. (See accompanying file LICENSE or copy at http://www.boost.org/LICENSE_1_0.txt)

URL Search is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the Boost Software License for more details.

## Usage

	urlsearch (INFO | {SERVER-PARAM})

	INFO
		-h           print this help
		-v           print version
		--copyright  print copyright

	SERVER-PARAM
		-p=N         port number (N is an integer)
		-t=S         page title (S is a string)
		-d=S         working directory (S is a string)

## Run on Windows
To run program without Cmd create file "start.vbs" with following lines:

	Set WshShell = CreateObject("WScript.Shell") 
	WshShell.Run """<path>\urlsearch.exe"" ""-tMy Title"" ""-p8080""", 0
	Set WshShell = Nothing

## References

- https://golang.org/doc/install
- https://git-scm.com/book/en/v2/Getting-Started-Installing-Git
