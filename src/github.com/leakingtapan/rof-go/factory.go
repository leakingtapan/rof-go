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
	"reflect"
)

// Top level interface
// Examples:
type ObjectFactory interface {
	Create(v interface{}) error
}

var defaultFactory ObjectFactory = &primitiveFactory{
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

func Create(v interface{}) error {
	return defaultFactory.Create(v)
}
