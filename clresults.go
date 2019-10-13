/*
 *          Copyright 2019, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

package main

import (
	"github.com/vbsw/osargs"
)

type clResults struct {
	help      []osargs.Param
	version   []osargs.Param
	copyright []osargs.Param
}

func (results *clResults) infoAvailable() bool {
	return len(results.help) > 0 || len(results.version) > 0 || len(results.copyright) > 0
}

func (results *clResults) oneParamHasMultipleResults() bool {
	for _, result := range results.toArray() {
		if len(result) > 1 {
			return true
		}
	}
	return false
}

func (results *clResults) toArray() [][]osargs.Param {
	resultsList := make([][]osargs.Param, 3)
	resultsList[0] = results.help
	resultsList[1] = results.version
	resultsList[2] = results.copyright
	return resultsList
}
