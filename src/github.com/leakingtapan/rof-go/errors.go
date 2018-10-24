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

import "fmt"

type InvalidInputError struct {
	msg string
}

func (e *InvalidInputError) Error() string {
	return fmt.Sprintf("invalid input: %s", e.msg)
}

type UnknownTypeError struct {
	v interface{}
}

func (e *UnknownTypeError) Error() string {
	return fmt.Sprintf("type is unknown: [%v]", e.v)
}
