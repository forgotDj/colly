// Copyright 2018 Adam Tauber
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package colly

import (
	"strconv"
	"testing"
)

func TestContextIteration(t *testing.T) {
	ctx := NewContext()
	for i := 0; i < 10; i++ {
		ctx.Put(strconv.Itoa(i), i)
	}
	values := ctx.ForEach(func(k string, v interface{}) interface{} {
		return v.(int)
	})
	if len(values) != 10 {
		t.Fatal("fail to iterate context")
	}
	for _, i := range values {
		v := i.(int)
		if v != ctx.GetAny(strconv.Itoa(v)).(int) {
			t.Fatal("value not equal")
		}
	}
}

func TestContextClone(t *testing.T) {
	ctxOrg := NewContext()
	for i := 0; i < 10; i++ {
		ctxOrg.Put(strconv.Itoa(i), i)
	}

	ctx := ctxOrg.Clone()
	values := ctx.ForEach(func(k string, v interface{}) interface{} {
		return v.(int)
	})
	if len(values) != 10 {
		t.Fatal("fail to iterate context")
	}
	for _, i := range values {
		v := i.(int)
		if v != ctx.GetAny(strconv.Itoa(v)).(int) {
			t.Fatal("value not equal")
		}
	}
}
