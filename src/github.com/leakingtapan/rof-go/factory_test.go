/*
Copyright 2018.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package rof

import (
	"fmt"
	"reflect"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("object factory", func() {
	It("should return error for nil", func() {
		err := Create(nil)
		Expect(err).ToNot(BeNil())
	})

	It("should generate random bool", func() {
		var value bool

		err := Create(&value)
		Expect(err).To(BeNil())

		testType(value)
	})

	It("should generate random int", func() {
		var value int

		err := Create(&value)
		Expect(err).To(BeNil())

		testType(value)
	})

	It("should generate random int8", func() {
		var value int8

		err := Create(&value)
		Expect(err).To(BeNil())

		testType(value)
	})

	It("should generate random int16", func() {
		var value int16

		err := Create(&value)
		Expect(err).To(BeNil())

		testType(value)
	})

	It("should generate random int32", func() {
		var value int32

		err := Create(&value)
		Expect(err).To(BeNil())

		testType(value)
	})

	It("should generate random int64", func() {
		var value int64

		err := Create(&value)
		Expect(err).To(BeNil())

		testType(value)
	})

	It("should generate random string", func() {
		var value string

		err := Create(&value)
		Expect(err).To(BeNil())

		testType(value)
	})

	It("should generate random array", func() {
		var value [8]int8

		err := Create(&value)
		Expect(err).To(BeNil())

		testType(value)
	})

	It("should generate random slice", func() {
		var value []int16

		err := Create(&value)
		Expect(err).To(BeNil())

		testType(value)
	})

	It("should generate random map", func() {
		var value map[int]string

		err := Create(&value)
		Expect(err).To(BeNil())

		testType(value)
	})
})

func testType(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Printf("%s: %v\n", v.Kind(), v.Interface())
}
