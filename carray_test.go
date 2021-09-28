// MIT License
//
// Copyright (c) 2021 The gnet Authors
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package carray

import "testing"

func TestCapOfCArray(t *testing.T) {
	var capList = []int {1, 2, 3, 4, 5, 6, 7, 8, 9}
	var realCapList = []int {1, 2, 4, 4, 8, 8, 8, 8, 16}
	for index, cap := range capList {
		if realCap := MakeCArray(cap, nil).cap; realCap != realCapList[index] {
			t.Errorf("Wrong cap of CArray: %d, %d", cap, realCap)
		}
	}
}

func TestPushInt(t *testing.T) {
	arr := MakeCArray(4, nil)
	arr.PushHeader(4)
}

func TestPushWrongType(t *testing.T) {
	defer func () {
		r := recover()
		if _, ok := r.(plainError); !ok {
			t.Errorf("Push Wrong Type but NO ERROR")
		}
	}()
	arr := MakeCArray(4, nil)
	arr.PushHeader(4)
	arr.PushHeader("123")
}

func TestPushHeaderOverFlow(t *testing.T) {
	arr := MakeCArray(4, nil)
	if res := arr.PushHeader(1); !res {
		t.Errorf("Cap-4 CArray Push Header Time-1 Error")
		return
	}
	if arr.front != 0 {
		t.Errorf("Cap-4 CArray Push Header Time-1 Front Index: %d Error", arr.front)
		return
	}
	if arr.count != 1 {
		t.Errorf("Cap-4 CArray Push Header Time-1 Count: %d Error", arr.count)
		return
	}
	if res := arr.PushHeader(2); !res {
		t.Errorf("Cap-4 CArray Push Header Time-2 Error")
		return
	}
	if arr.front != 3 {
		t.Errorf("Cap-4 CArray Push Header Time-2 Front Index: %d Error", arr.front)
		return
	}
	if arr.count != 2 {
		t.Errorf("Cap-4 CArray Push Header Time-2 Count: %d Error", arr.count)
		return
	}
	if res := arr.PushHeader(3); !res {
		t.Errorf("Cap-4 CArray Push Header Time-3 Error")
		return
	}
	if arr.front != 2 {
		t.Errorf("Cap-4 CArray Push Header Time-3 Front Index: %d Error", arr.front)
		return
	}
	if arr.count != 3 {
		t.Errorf("Cap-4 CArray Push Header Time-3 Count: %d Error", arr.count)
		return
	}
	if res := arr.PushHeader(4); !res {
		t.Errorf("Cap-4 CArray Push Header Time-4 Error")
		return
	}
	if arr.front != 1 {
		t.Errorf("Cap-4 CArray Push Header Time-4 Front Index: %d Error", arr.front)
		return
	}
	if arr.count != 4 {
		t.Errorf("Cap-4 CArray Push Header Time-4 Count: %d Error", arr.count)
		return
	}
	if res := arr.PushHeader(5); res {
		t.Errorf("Cap-4 CArray Push Header Time-5 Success")
		return
	}
	if arr.front != 1 {
		t.Errorf("Cap-4 CArray Push Header Time-5 Front Index: %d Error", arr.front)
		return
	}
	if arr.count != 4 {
		t.Errorf("Cap-4 CArray Push Header Time-5 Count: %d Error", arr.count)
		return
	}
}

func TestPushTailOverFlow(t *testing.T) {
	arr := MakeCArray(4, nil)
	if res := arr.PushTail(1); !res {
		t.Errorf("Cap-4 CArray Push Tail Time-1 Error")
		return
	}
	if arr.tail != 0 {
		t.Errorf("Cap-4 CArray Push Tail Time-1 Tail Index: %d Error", arr.front)
		return
	}
	if arr.count != 1 {
		t.Errorf("Cap-4 CArray Push Tail Time-1 Count: %d Error", arr.count)
		return
	}
	if res := arr.PushTail(2); !res {
		t.Errorf("Cap-4 CArray Push Tail Time-2 Error")
		return
	}
	if arr.tail != 1 {
		t.Errorf("Cap-4 CArray Push Tail Time-2 Tail Index: %d Error", arr.front)
		return
	}
	if arr.count != 2 {
		t.Errorf("Cap-4 CArray Push Tail Time-2 Count: %d Error", arr.count)
		return
	}
	if res := arr.PushTail(3); !res {
		t.Errorf("Cap-4 CArray Push Tail Time-3 Error")
		return
	}
	if arr.tail != 2 {
		t.Errorf("Cap-4 CArray Push Tail Time-3 Tail Index: %d Error", arr.front)
		return
	}
	if arr.count != 3 {
		t.Errorf("Cap-4 CArray Push Tail Time-3 Count: %d Error", arr.count)
		return
	}
	if res := arr.PushTail(4); !res {
		t.Errorf("Cap-4 CArray Push Tail Time-4 Error")
		return
	}
	if arr.tail != 3 {
		t.Errorf("Cap-4 CArray Push Tail Time-4 Tail Index: %d Error", arr.front)
		return
	}
	if arr.count != 4 {
		t.Errorf("Cap-4 CArray Push Tail Time-4 Count: %d Error", arr.count)
		return
	}
	if res := arr.PushTail(5); res {
		t.Errorf("Cap-4 CArray Push Tail Time-5 Success")
		return
	}
	if arr.tail != 3 {
		t.Errorf("Cap-4 CArray Push Header Time-5 Tail Index: %d Error", arr.front)
		return
	}
	if arr.count != 4 {
		t.Errorf("Cap-4 CArray Push Header Time-5 Count: %d Error", arr.count)
		return
	}
}

func TestRemove(t *testing.T) {
	arr := MakeCArray(4, func(elem *(interface{}), data interface{}) bool {
		realElem, ok := (*elem).(int)
		if !ok {
			t.Errorf("Wrong CArray Element Type")
			return false
		}
		realElem += data.(int)
		*elem = realElem
		if realElem < 10 {
			return false
		}
		return true
	})
	arr.PushHeader(9)
	elem, ok := arr.Remove(1, true) 
	if !ok{
		t.Errorf("not remove data")
	}
	if elem.(int) != 10 {
		t.Errorf("the removed data is wrong: %d", elem)
	}
}
