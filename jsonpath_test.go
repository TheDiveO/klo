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

package klo

import (
	"bytes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("JSONPath printer", func() {

	It("prints JSON using JSONPath", func() {
		_, err := NewJSONPathPrinter("{.Foo")
		Expect(err).Should(HaveOccurred())

		type foo struct {
			Foo string
		}
		f := foo{Foo: "bar"}
		var out bytes.Buffer
		p, err := NewJSONPathPrinter("{.Foo}")
		Expect(err).ShouldNot(HaveOccurred())
		Expect(p.Fprint(&out, f)).Should(Succeed())
		Expect(out.String()).Should(Equal(`bar`))

		out.Reset()
		p, err = NewJSONPathPrinter("{.Nothing}")
		Expect(err).ShouldNot(HaveOccurred())
		Expect(p.Fprint(&out, f)).ShouldNot(Succeed())
	})

})