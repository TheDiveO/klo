// Copyright 2019 Harald Albrecht.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package testutil

import (
	. "github.com/onsi/gomega"
)

type PASSFAILS []PASSFAIL
type PASSFAIL interface {
	Description() string
	Actual() interface{}
}

// FAIL this testcase.
type FAIL struct {
	D string // test case description
	A interface{}
}

func (f FAIL) Description() string { return f.D }
func (f FAIL) Actual() interface{} { return f.A }

// PASS is just another type of FAIL, after all. (No, this hasn't been ripped
// off straight from one of Nico Semsrott's satirical revues.)
type PASS FAIL

func (p PASS) Description() string { return p.D }
func (p PASS) Actual() interface{} { return p.A }

// PassFail runs a series of PASS and/or FAIL tests, checking for either
// success or the absense of success (failure might be to harsh a word).
func PassFail(tests PASSFAILS) {
	for _, t := range tests {
		if _, ok := t.(PASS); ok {
			ExpectWithOffset(1, t.Actual()).
				Should(Succeed(), t.Description())
		} else {
			ExpectWithOffset(1, t.Actual()).
				ShouldNot(Succeed(), t.Description())
		}
	}
}

// Err returns only the last actual result of a multi-result function call,
// which typically is an error result. Err is most useful in writing simple
// PASS/FAIL test cases for multi-result function calls.
func Err(actual interface{}, extras ...interface{}) interface{} {
	if len(extras) > 0 {
		return extras[len(extras)-1]
	}
	return actual
}
