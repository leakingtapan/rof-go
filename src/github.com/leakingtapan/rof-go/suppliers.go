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
	"math"
	"math/rand"
	"reflect"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())

	defaultSuppliers = map[reflect.Type]supplier{}
	for _, f := range funcs {
		defaultSuppliers.setFunc(f)
	}
}

// SetFunc sets the default function if it already exists
// Or a new one will be added
func SetFunc(f interface{}) {
	defaultSuppliers.setFunc(f)
}

// Default functions used to generate random values
var (
	funcs = []interface{}{
		boolGen,
		intGen,
		int8Gen,
		int16Gen,
		int32Gen,
		int64Gen,
		uintGen,
		uint8Gen,
		uint16Gen,
		uint32Gen,
		uint64Gen,
		float32Gen,
		float64Gen,
		complex64Gen,
		complex128Gen,
		strGen,
	}
)

//////////////////////////////
// Default built-in functions
//////////////////////////////

func boolGen() bool {
	return rand.Int31() > (math.MaxInt32 >> 1)
}

func intGen() int {
	return rand.Int()
}

func int8Gen() int8 {
	return int8(intGen())
}

func int16Gen() int16 {
	return int16(intGen())
}

func int32Gen() int32 {
	return rand.Int31()
}

func int64Gen() int64 {
	return rand.Int63()
}

func uintGen() uint {
	return uint(intGen())
}

func uint8Gen() uint8 {
	return uint8(int8Gen())
}

func uint16Gen() uint16 {
	return uint16(int16Gen())
}

func uint32Gen() uint32 {
	return uint32(int32Gen())
}

func uint64Gen() uint64 {
	return uint64(int64Gen())
}

func float32Gen() float32 {
	return rand.Float32()
}

func float64Gen() float64 {
	return rand.Float64()
}

func complex64Gen() complex64 {
	return complex(float32Gen(), float32Gen())
}

func complex128Gen() complex128 {
	return complex(float64Gen(), float64Gen())
}

const (
	alphanumerics = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func strGen() string {
	size := 8
	res := make([]byte, size)
	for i := 0; i < size; i++ {
		id := rand.Intn(len(alphanumerics))
		res[i] = alphanumerics[id]
	}
	return string(res)
}

// supplier is a function that returns a values of certain type
type supplier func() interface{}

var defaultSuppliers suppliers

type suppliers map[reflect.Type]supplier

func (s suppliers) setFunc(f interface{}) {
	rf := reflect.ValueOf(f)
	if rf.IsNil() || rf.Kind() != reflect.Func {
		panic("f is nil or is not function")
	}

	fTyp := rf.Type()
	if fTyp.NumIn() != 0 || fTyp.NumOut() != 1 {
		panic("function f is not of type func f() interface{}")
	}

	outType := fTyp.Out(0)
	s[outType] = funcWrap(f)
}

// wrap a function f as supplier
func funcWrap(f interface{}) supplier {
	rv := reflect.ValueOf(f)
	if rv.Kind() != reflect.Func {
		panic("cannot wrap f. f is not a function")
	}

	return func() interface{} {
		results := rv.Call([]reflect.Value{})
		return results[0].Interface()
	}
}
