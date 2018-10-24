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
	It("should generate random bool", func() {
		var value bool

		err := Create(&value)
		Expect(err).To(BeNil())

		testType(&value)
		fmt.Println(": ", value)
	})
	It("should generate random int", func() {
		var value int

		err := Create(&value)
		Expect(err).To(BeNil())

		testType(&value)
		fmt.Println(": ", value)
	})

	It("should generate random int32", func() {
		var value int32

		err := Create(&value)
		Expect(err).To(BeNil())

		testType(&value)
		fmt.Println(": ", value)
	})

	It("should generate random string", func() {
		var value string

		err := Create(&value)
		Expect(err).To(BeNil())

		testType(&value)
		fmt.Println(": ", value)
	})

	It("should generate random array", func() {
		var value [8]int8

		err := Create(&value)
		Expect(err).To(BeNil())

		testType(&value)
		fmt.Println(": ", value)
	})

	It("should generate random slice", func() {
		var value []int16

		err := Create(&value)
		Expect(err).To(BeNil())

		testType(&value)
		fmt.Println(": ", value)
	})

	It("should generate random map", func() {
		var value map[int8]string

		err := Create(&value)
		Expect(err).To(BeNil())

		testType(&value)
		fmt.Println(": ", value)
	})
})

func testType(x interface{}) {
	v := reflect.ValueOf(x)
	switch v.Kind() {
	case reflect.Bool:
		fmt.Printf("bool: %v\t", v.Bool())
	case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64:
		fmt.Printf("int: %v\t", v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint32, reflect.Uint64:
		fmt.Printf("uint: %v\t", v.Uint())
	case reflect.Float32, reflect.Float64:
		fmt.Printf("float: %v\t", v.Float())
	case reflect.String:
		fmt.Printf("string: %v\t", v.String())
	case reflect.Slice:
		fmt.Printf("slice: len=%d, %v\t", v.Len(), v.Interface())
	case reflect.Map:
		fmt.Printf("map: %v\t", v.Interface())
	case reflect.Chan:
		fmt.Printf("chan %v\t", v.Interface())
	case reflect.Ptr:
		fmt.Printf("ptr %v\t", v.Type())
	default:
		fmt.Println(x)
	}
}
