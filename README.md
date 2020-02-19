# URL Search

[![GoDoc](https://godoc.org/github.com/vbsw/urlsearch?status.svg)](https://godoc.org/github.com/vbsw/urlsearch) [![Go Report Card](https://goreportcard.com/badge/github.com/vbsw/urlsearch)](https://goreportcard.com/report/github.com/vbsw/urlsearch) [![Stability: Experimental](https://masterminds.github.io/stability/experimental.svg)](https://masterminds.github.io/stability/experimental.html)

## About
URL Search is a server to save and search URLs by keywords. URL Search is published on <https://github.com/vbsw/urlsearch>.

## Copyright
Copyright 2020, Vitali Baumtrok (vbsw@mailbox.org).

URL Search is distributed under the Boost Software License, version 1.0. (See accompanying file LICENSE or copy at http://www.boost.org/LICENSE_1_0.txt)

URL Search is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the Boost Software License for more details.

## Usage

	urlsearch (INFO | {SERVER-PARAM})

	INFO
		-h, --help         print this help
		-v, --version      print version
		--copyright        print copyright

	SERVER-PARAM
		-p=N, --port=N     port number (N: integer)
		-t=S, --title=S    page title (S: string)
		-d=S, --dir=S      working directory (S: string)
		-b, --background   run in background

## References

- https://golang.org/doc/install
- https://git-scm.com/book/en/v2/Getting-Started-Installing-Git
