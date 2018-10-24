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

var (
	BoolFunc       func() bool  = boolGen
	IntFunc        func() int   = intGen
	Int8Func       func() int8  = int8Gen
	Int16Func      func() int16 = int16Gen
	Int32Func      func() int32 = int32Gen
	Int64Func      func() int64 = int64Gen
	UintFunc       func() uint
	Uint8Func      func() uint8
	Uint16Func     func() uint16
	Uint32Func     func() uint32
	Uint64Func     func() uint64
	UintptrFunc    func() uintptr
	Float32Func    func() float32
	Float64Func    func() float64
	Complex64Func  func() complex64
	Complex128Func func() complex128
	StrFunc        func() string = strGen
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func boolGen() bool {
	return rand.Int31() > (math.MaxInt32 >> 1)
}

// intGen() return a random value of int
func intGen() int {
	return rand.Int()
}

func int8Gen() int8 {
	return int8(intGen())
}

func int16Gen() int16 {
	return int16(intGen())
}

// int32Gen() return a random value of int32
func int32Gen() int32 {
	return rand.Int31()
}

func int64Gen() int64 {
	return rand.Int63()
}

//func uintGen() uint {
//	return rand.UInt()
//}
//
//func uint8Gen() uint8 {
//	return int8(intGen())
//}
//
//func uint16Gen() uint16 {
//	return int16(intGen())
//}
//
//func uint32Gen() uint32 {
//	return rand.Int31()
//}
//
//func uint64Gen() uint64 {
//	return rand.Int63()
//}

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

// Supplier is a function that returns a values of certain type
type Supplier func() interface{}

// wrap a function f as supplier
func funcWrap(f interface{}) Supplier {
	rv := reflect.ValueOf(f)
	if rv.Kind() != reflect.Func {
		panic("cannot wrap f. f is not a function")
	}

	return func() interface{} {
		results := rv.Call([]reflect.Value{})
		return results[0].Interface()
	}
}
