package rof

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

var r *rand.Rand

func init() {
	source := rand.NewSource(time.Now().UnixNano())
	r = rand.New(source)
}

type InvalidInputError struct {
	v interface{}
}

func (e *InvalidInputError) Error() string {
	return fmt.Sprintf("Invalid input: [%v]", e.v)
}

type ObjectFactory interface {
	Create(v interface{}) error
}

type primitiveObjectFactory struct {
	suppliers map[reflect.Kind]func() interface{}
}

func (f *primitiveObjectFactory) Create(v interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return &InvalidInputError{v}
	}

	switch t := v.(type) {
	case *int:
		*t = *intGen().(*int)
	case *int32:
		*t = r.Int31()
		//case reflect.Slice:
		//  t = make([]int, 10)
	}
	return nil
}

//type Supplier interface {
//	Get() interface{}
//}

func intGen() interface{} {
	res := r.Int()
	return &res
}

var rof ObjectFactory = &primitiveObjectFactory{
	suppliers: map[reflect.Kind]func() interface{}{
		reflect.Int:   intGen,
		reflect.Int32: intGen,
	},
}

func Create(v interface{}) error {
	return rof.Create(v)
}

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
