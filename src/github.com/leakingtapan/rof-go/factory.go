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

// Supplier is a function that returns a values of certain type
type Supplier func() interface{}

var defaultFactory ObjectFactory = &primitiveFactory{
	suppliers: map[reflect.Kind]Supplier{
		reflect.Int:   intItf(intGen),
		reflect.Int32: int32Itf(int32Gen),
	},
}

func Create(v interface{}) error {
	return defaultFactory.Create(v)
}
