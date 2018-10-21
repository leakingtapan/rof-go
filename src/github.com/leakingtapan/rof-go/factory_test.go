package rof

import (
	"fmt"
	"reflect"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("object factory", func() {
	It("should generate random int", func() {
		var a int

		err := Create(&a)
		Expect(err).To(BeNil())

		test(&a)
		fmt.Println("int: ", a)
	})

	It("should generate random int", func() {
		var a int32

		err := Create(&a)
		Expect(err).To(BeNil())

		test(&a)
		fmt.Println("int32: ", a)
	})
})

func test(x interface{}) {
	v := reflect.ValueOf(x)
	switch v.Kind() {
	case reflect.Bool:
		fmt.Printf("bool: %v\n", v.Bool())
	case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64:
		fmt.Printf("int: %v\n", v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint32, reflect.Uint64:
		fmt.Printf("uint: %v\n", v.Uint())
	case reflect.Float32, reflect.Float64:
		fmt.Printf("float: %v\n", v.Float())
	case reflect.String:
		fmt.Printf("string: %v\n", v.String())
	case reflect.Slice:
		fmt.Printf("slice: len=%d, %v\n", v.Len(), v.Interface())
	case reflect.Map:
		fmt.Printf("map: %v\n", v.Interface())
	case reflect.Chan:
		fmt.Printf("chan %v\n", v.Interface())
	case reflect.Ptr:
		fmt.Printf("ptr %v\n", v.Type())
	default:
		fmt.Println(x)
	}
}
