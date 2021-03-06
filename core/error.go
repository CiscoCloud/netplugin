/***
Copyright 2014 Cisco Systems Inc. All rights reserved.

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

package core

import "strings"

type Error struct {
	Desc string
}

func (e *Error) Error() string {
	return e.Desc
}

func ErrIfKeyExists(err error) error {
	if err == nil || strings.Contains(err.Error(), "Key not found") {
		return nil
	} else {
		return err
	}
}
