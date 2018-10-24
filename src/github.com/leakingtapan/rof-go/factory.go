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
)

func Create(v interface{}) error {
	return defaultFactory.Create(v)
}

// Top level interface
// Examples:
type ObjectFactory interface {
	Create(v interface{}) error
}

var defaultFactory ObjectFactory = &defaultObjectFactory{
	suppliers: map[reflect.Kind]Supplier{
		reflect.Bool:  funcWrap(BoolFunc),
		reflect.Int:   funcWrap(IntFunc),
		reflect.Int8:  funcWrap(Int8Func),
		reflect.Int16: funcWrap(Int16Func),
		reflect.Int32: funcWrap(Int32Func),
		reflect.Int64: funcWrap(Int64Func),
		//reflect.Uint:   funcWrap(intFunc),
		//reflect.Uint8:  funcWrap(int8Func),
		//reflect.Uint32: funcWrap(int32Func),
		//reflect.Uint64: funcWrap(int64Func),
		reflect.String: funcWrap(strGen),
	},
}

type defaultObjectFactory struct {
	suppliers map[reflect.Kind]Supplier
}

// ptr is a pointer to a value that is to be initialized by the factory
func (f *defaultObjectFactory) Create(ptr interface{}) error {
	ptrValue := reflect.ValueOf(ptr)
	if ptrValue.Kind() != reflect.Ptr || ptrValue.IsNil() {
		return &InvalidInputError{fmt.Sprintf("%v is not a pointer or is nil", ptr)}
	}

	value := reflect.Indirect(ptrValue)

	// create for primitive type
	// TODO: match by type so that custom type is configurable
	supplier, exist := f.suppliers[value.Kind()]
	if exist {
		value.Set(reflect.ValueOf(supplier()))
		return nil
	}

	// create for composite type
	switch value.Kind() {
	case reflect.Array:
		f.createArray(value)
	case reflect.Slice:
		f.createSlice(value)
	case reflect.Map:
		f.createMap(value)
	case reflect.Struct:
	default:
		return &UnknownTypeError{value}
	}

	return nil
}

// creates a reflect.Value for given reflect.Type
func (f *defaultObjectFactory) createFrom(t reflect.Type) reflect.Value {
	rv := reflect.New(t)
	err := f.Create(rv.Interface())
	if err != nil {
		// happens when type is not supported
		// eg. Pointer type *int
		panic(err)
	}

	return reflect.Indirect(rv)
}

func (f *defaultObjectFactory) createArray(value reflect.Value) {
	size := value.Len()
	typ := value.Type()
	for i := 0; i < size; i++ {
		elemValue := f.createFrom(typ.Elem())
		value.Index(i).Set(elemValue)
	}
}

func (f *defaultObjectFactory) createSlice(value reflect.Value) {
	typ := value.Type()
	//TODO: configurable size
	size := 10
	s := reflect.MakeSlice(typ, 0, size)
	for i := 0; i < size; i++ {
		elemValue := f.createFrom(typ.Elem())
		s = reflect.Append(s, elemValue)
	}
	value.Set(s)

}

func (f *defaultObjectFactory) createMap(value reflect.Value) {
	typ := value.Type()
	size := 10
	m := reflect.MakeMap(typ)
	value.Set(m)
	for i := 0; i < size; i++ {
		keyValue := f.createFrom(typ.Key())
		elemValue := f.createFrom(typ.Elem())
		value.SetMapIndex(keyValue, elemValue)
	}
}
