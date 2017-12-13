// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/m3db/m3db/digest (interfaces: ReaderWithDigest)

package digest

import (
	gomock "github.com/golang/mock/gomock"
	hash "hash"
	io "io"
)

// Mock of ReaderWithDigest interface
type MockReaderWithDigest struct {
	ctrl     *gomock.Controller
	recorder *_MockReaderWithDigestRecorder
}

// Recorder for MockReaderWithDigest (not exported)
type _MockReaderWithDigestRecorder struct {
	mock *MockReaderWithDigest
}

func NewMockReaderWithDigest(ctrl *gomock.Controller) *MockReaderWithDigest {
	mock := &MockReaderWithDigest{ctrl: ctrl}
	mock.recorder = &_MockReaderWithDigestRecorder{mock}
	return mock
}

func (_m *MockReaderWithDigest) EXPECT() *_MockReaderWithDigestRecorder {
	return _m.recorder
}

func (_m *MockReaderWithDigest) Digest() hash.Hash32 {
	ret := _m.ctrl.Call(_m, "Digest")
	ret0, _ := ret[0].(hash.Hash32)
	return ret0
}

func (_mr *_MockReaderWithDigestRecorder) Digest() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Digest")
}

func (_m *MockReaderWithDigest) Read(_param0 []byte) (int, error) {
	ret := _m.ctrl.Call(_m, "Read", _param0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockReaderWithDigestRecorder) Read(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Read", arg0)
}

func (_m *MockReaderWithDigest) Reset(_param0 io.Reader) {
	_m.ctrl.Call(_m, "Reset", _param0)
}

func (_mr *_MockReaderWithDigestRecorder) Reset(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Reset", arg0)
}

func (_m *MockReaderWithDigest) Validate(_param0 uint32) error {
	ret := _m.ctrl.Call(_m, "Validate", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockReaderWithDigestRecorder) Validate(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Validate", arg0)
}
