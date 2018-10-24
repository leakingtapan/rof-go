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

type primitiveFactory struct {
	suppliers map[reflect.Kind]Supplier
}

// v is a pointer to the value
func (f *primitiveFactory) Create(v interface{}) error {
	ptrValue := reflect.ValueOf(v)
	if ptrValue.Kind() != reflect.Ptr || ptrValue.IsNil() {
		return &InvalidInputError{v}
	}

	value := reflect.Indirect(ptrValue)
	supplier, exist := f.suppliers[value.Kind()]
	if exist {
		value.Set(reflect.ValueOf(supplier()))
		return nil
	}

	switch value.Kind() {
	case reflect.Array:
		size := value.Len()
		typ := value.Type()
		for i := 0; i < size; i++ {
			elemValue := f.createFrom(typ.Elem())
			value.Index(i).Set(elemValue)
		}
		return nil
	case reflect.Slice:
		typ := value.Type()
		//TODO: configurable size
		size := 10
		s := reflect.MakeSlice(typ, 0, size)
		for i := 0; i < size; i++ {
			elemValue := f.createFrom(typ.Elem())
			s = reflect.Append(s, elemValue)
		}
		value.Set(s)
		return nil
	case reflect.Map:
		typ := value.Type()
		keyTyp := typ.Key()
		valueTyp := typ.Elem()
		size := 10
		m := reflect.MakeMap(typ)
		for i := 0; i < size; i++ {
			keyValue := f.createFrom(keyTyp)
			elemValue := f.createFrom(valueTyp)
			fmt.Println(keyValue)
			m.MapIndex(keyValue).Set(elemValue)
		}
		value.Set(m)
		return nil
	case reflect.Struct:
		return nil
	}

	// handle Array, Slice, Map and Struct
	return &UnknownTypeError{value}
}

// creates a reflect.Value for given reflect.Type
func (f *primitiveFactory) createFrom(t reflect.Type) reflect.Value {
	rv := reflect.New(t)
	err := f.Create(rv.Interface())
	if err != nil {
		// happens when type is not supported
		// eg. Pointer type *int
		panic(err)
	}

	return reflect.Indirect(rv)
}
