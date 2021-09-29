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

import (
	"reflect"
)

type CArray struct {
	// count is the number of real data in ring buffer
	count int
	// cap is the cap of buf, and must be the power of two.
	cap int
	// buf is a ring-buffer structure with a fixed size.
	buf []interface{}
	// _type is the type of the buf's element which is determined when CArray is initialized.
	_type reflect.Type
	// front is the position of header element in buf.
	front int
	// tail is the position of tail element in buf.
	tail int
	// remove is the function whether to remove the element with data, return true if need removed.
	remove func (elem *(interface{}), data interface{}, arr *CArray) bool
}

func MakeCArray(cap int, remove func (elem *(interface{}), data interface{}, arr *CArray) bool) *CArray {
	if cap <= 0 {
		panic(plainError("MakeCArray: size must be a positive number"))
	}
	cap = LowPowerOfTwo(cap)
	return &CArray{
		count: 0,
		cap: cap,
		buf: make([]interface{}, cap, cap),
		remove: remove,
	}
}

func (arr *CArray) Empty() bool {
	return arr.count == 0
}

func (arr *CArray) Full() bool {
	return arr.count == arr.cap
}

func (arr *CArray) Front() *CArray {
	if arr.Empty() { return nil }
	return arr.buf[arr.front]
}

func (arr *CArray) Tail()) *CArray {
	if arr.Empty() { return nil }
	return arr.buf[arr.tail]
}

func (arr *CArray) PushHeader(elem interface{}) bool {
	if !arr.checkElem(elem) {
		panic(plainError("CArray.PushHeader: type of elem is different from the first"))
	}
	if arr.Full() {
		return false
	}
	if !arr.Empty() {
		arr.front = arr.realIndex(arr.front - 1)
	}
	arr.buf[arr.front] = elem
	arr.count++
	return true
}

func (arr *CArray) PushTail(elem interface{}) bool {
	if !arr.checkElem(elem) {
		panic(plainError("CArray.PushHeader: type of elem is different from the first"))
	}
	if arr.Full() {
		return false
	}
	if !arr.Empty() {
		arr.tail = arr.realIndex(arr.tail + 1)
	}
	arr.buf[arr.tail] = elem
	arr.count++
	return true
}

func (arr *CArray) PopFront() interface{} {
	if arr.Empty() {
		return nil
	}
	defer func () {
		if !(arr.count == 1) {
			arr.front = arr.realIndex(arr.front + 1)
		}
		arr.count--
	}()
	return arr.buf[arr.front]
}

func (arr *CArray) PopTail() interface{} {
	if arr.Empty() {
		return nil
	}
	defer func () {
		if !(arr.count == 1) {
			arr.tail = arr.realIndex(arr.tail - 1)
		}
		arr.count--
	}()
	return arr.buf[arr.tail]
}

func (arr *CArray) Remove(data interface{}, front bool) (interface{}, bool) {
	if arr.Empty() {
		return nil, false
	}
	elem := arr.buf[arr.front]
	if !front {
		elem = arr.buf[arr.tail]
	}
	if arr.remove(&elem, data, arr) {
		if front {
			arr.PopFront()
		} else {
			arr.PopTail()
		}
		return elem, true
	}
	return nil, false
}

func (arr *CArray) realIndex(index int) int {
	return index & (arr.cap - 1)
}

func (arr *CArray) checkElem(elem interface{}) bool {
	if arr._type != nil && arr._type.Kind() != reflect.TypeOf(elem).Kind() {
		return false
	}
	if arr._type == nil {
		arr._type = reflect.TypeOf(elem)
	}
	return true
}
