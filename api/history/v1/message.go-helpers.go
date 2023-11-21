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

package history

import (
	"google.golang.org/protobuf/proto"
)

// Marshal an object of type TransientWorkflowTaskInfo to the protobuf v3 wire format
func (val *TransientWorkflowTaskInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type TransientWorkflowTaskInfo from the protobuf v3 wire format
func (val *TransientWorkflowTaskInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *TransientWorkflowTaskInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two TransientWorkflowTaskInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *TransientWorkflowTaskInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *TransientWorkflowTaskInfo
	switch t := that.(type) {
	case *TransientWorkflowTaskInfo:
		that1 = t
	case TransientWorkflowTaskInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type VersionHistoryItem to the protobuf v3 wire format
func (val *VersionHistoryItem) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type VersionHistoryItem from the protobuf v3 wire format
func (val *VersionHistoryItem) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *VersionHistoryItem) Size() int {
	return proto.Size(val)
}

// Equal returns whether two VersionHistoryItem values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *VersionHistoryItem) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *VersionHistoryItem
	switch t := that.(type) {
	case *VersionHistoryItem:
		that1 = t
	case VersionHistoryItem:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type VersionHistory to the protobuf v3 wire format
func (val *VersionHistory) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type VersionHistory from the protobuf v3 wire format
func (val *VersionHistory) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *VersionHistory) Size() int {
	return proto.Size(val)
}

// Equal returns whether two VersionHistory values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *VersionHistory) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *VersionHistory
	switch t := that.(type) {
	case *VersionHistory:
		that1 = t
	case VersionHistory:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type VersionHistories to the protobuf v3 wire format
func (val *VersionHistories) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type VersionHistories from the protobuf v3 wire format
func (val *VersionHistories) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *VersionHistories) Size() int {
	return proto.Size(val)
}

// Equal returns whether two VersionHistories values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *VersionHistories) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *VersionHistories
	switch t := that.(type) {
	case *VersionHistories:
		that1 = t
	case VersionHistories:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type TaskKey to the protobuf v3 wire format
func (val *TaskKey) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type TaskKey from the protobuf v3 wire format
func (val *TaskKey) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *TaskKey) Size() int {
	return proto.Size(val)
}

// Equal returns whether two TaskKey values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *TaskKey) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *TaskKey
	switch t := that.(type) {
	case *TaskKey:
		that1 = t
	case TaskKey:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type TaskRange to the protobuf v3 wire format
func (val *TaskRange) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type TaskRange from the protobuf v3 wire format
func (val *TaskRange) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *TaskRange) Size() int {
	return proto.Size(val)
}

// Equal returns whether two TaskRange values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *TaskRange) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *TaskRange
	switch t := that.(type) {
	case *TaskRange:
		that1 = t
	case TaskRange:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
