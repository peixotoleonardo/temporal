// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
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

package persistence

import (
	"google.golang.org/protobuf/proto"
)

// Marshal an object of type HistoryTreeInfo to the protobuf v3 wire format
func (val *HistoryTreeInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type HistoryTreeInfo from the protobuf v3 wire format
func (val *HistoryTreeInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *HistoryTreeInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two HistoryTreeInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *HistoryTreeInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *HistoryTreeInfo
	switch t := that.(type) {
	case *HistoryTreeInfo:
		that1 = t
	case HistoryTreeInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type HistoryBranch to the protobuf v3 wire format
func (val *HistoryBranch) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type HistoryBranch from the protobuf v3 wire format
func (val *HistoryBranch) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *HistoryBranch) Size() int {
	return proto.Size(val)
}

// Equal returns whether two HistoryBranch values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *HistoryBranch) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *HistoryBranch
	switch t := that.(type) {
	case *HistoryBranch:
		that1 = t
	case HistoryBranch:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type HistoryBranchRange to the protobuf v3 wire format
func (val *HistoryBranchRange) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type HistoryBranchRange from the protobuf v3 wire format
func (val *HistoryBranchRange) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *HistoryBranchRange) Size() int {
	return proto.Size(val)
}

// Equal returns whether two HistoryBranchRange values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *HistoryBranchRange) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *HistoryBranchRange
	switch t := that.(type) {
	case *HistoryBranchRange:
		that1 = t
	case HistoryBranchRange:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
