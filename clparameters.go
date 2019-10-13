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

type clParameters struct {
	help      []osargs.Param
	version   []osargs.Param
	copyright []osargs.Param
}

func (parameters *clParameters) infoAvailable() bool {
	return len(parameters.help) > 0 || len(parameters.version) > 0 || len(parameters.copyright) > 0
}

func (parameters *clParameters) anyParameterMultiple() bool {
	for _, result := range parameters.toArray() {
		if len(result) > 1 {
			return true
		}
	}
	return false
}

func (parameters *clParameters) toArray() [][]osargs.Param {
	parametersArray := make([][]osargs.Param, 3)
	parametersArray[0] = parameters.help
	parametersArray[1] = parameters.version
	parametersArray[2] = parameters.copyright
	return parametersArray
}
