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
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// intGen() return a random value of int
func intGen() int {
	res := rand.Int()
	return res
}

// int32Gen() return a random value of int32
func int32Gen() int32 {
	res := rand.Int31()
	return res
}

// AUTO gen? or use reflection?
func intItf(f func() int) func() interface{} {
	return func() interface{} {
		return f()
	}
}

func int32Itf(f func() int32) func() interface{} {
	return func() interface{} {
		return f()
	}
}
