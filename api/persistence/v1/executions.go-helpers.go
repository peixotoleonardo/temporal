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

// Marshal an object of type ShardInfo to the protobuf v3 wire format
func (val *ShardInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ShardInfo from the protobuf v3 wire format
func (val *ShardInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ShardInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ShardInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ShardInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ShardInfo
	switch t := that.(type) {
	case *ShardInfo:
		that1 = t
	case ShardInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type WorkflowExecutionInfo to the protobuf v3 wire format
func (val *WorkflowExecutionInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type WorkflowExecutionInfo from the protobuf v3 wire format
func (val *WorkflowExecutionInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *WorkflowExecutionInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two WorkflowExecutionInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *WorkflowExecutionInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *WorkflowExecutionInfo
	switch t := that.(type) {
	case *WorkflowExecutionInfo:
		that1 = t
	case WorkflowExecutionInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ExecutionStats to the protobuf v3 wire format
func (val *ExecutionStats) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ExecutionStats from the protobuf v3 wire format
func (val *ExecutionStats) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ExecutionStats) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ExecutionStats values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ExecutionStats) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ExecutionStats
	switch t := that.(type) {
	case *ExecutionStats:
		that1 = t
	case ExecutionStats:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type WorkflowExecutionState to the protobuf v3 wire format
func (val *WorkflowExecutionState) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type WorkflowExecutionState from the protobuf v3 wire format
func (val *WorkflowExecutionState) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *WorkflowExecutionState) Size() int {
	return proto.Size(val)
}

// Equal returns whether two WorkflowExecutionState values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *WorkflowExecutionState) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *WorkflowExecutionState
	switch t := that.(type) {
	case *WorkflowExecutionState:
		that1 = t
	case WorkflowExecutionState:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type TransferTaskInfo to the protobuf v3 wire format
func (val *TransferTaskInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type TransferTaskInfo from the protobuf v3 wire format
func (val *TransferTaskInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *TransferTaskInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two TransferTaskInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *TransferTaskInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *TransferTaskInfo
	switch t := that.(type) {
	case *TransferTaskInfo:
		that1 = t
	case TransferTaskInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ReplicationTaskInfo to the protobuf v3 wire format
func (val *ReplicationTaskInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ReplicationTaskInfo from the protobuf v3 wire format
func (val *ReplicationTaskInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ReplicationTaskInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ReplicationTaskInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ReplicationTaskInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ReplicationTaskInfo
	switch t := that.(type) {
	case *ReplicationTaskInfo:
		that1 = t
	case ReplicationTaskInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type VisibilityTaskInfo to the protobuf v3 wire format
func (val *VisibilityTaskInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type VisibilityTaskInfo from the protobuf v3 wire format
func (val *VisibilityTaskInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *VisibilityTaskInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two VisibilityTaskInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *VisibilityTaskInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *VisibilityTaskInfo
	switch t := that.(type) {
	case *VisibilityTaskInfo:
		that1 = t
	case VisibilityTaskInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type TimerTaskInfo to the protobuf v3 wire format
func (val *TimerTaskInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type TimerTaskInfo from the protobuf v3 wire format
func (val *TimerTaskInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *TimerTaskInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two TimerTaskInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *TimerTaskInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *TimerTaskInfo
	switch t := that.(type) {
	case *TimerTaskInfo:
		that1 = t
	case TimerTaskInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ArchivalTaskInfo to the protobuf v3 wire format
func (val *ArchivalTaskInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ArchivalTaskInfo from the protobuf v3 wire format
func (val *ArchivalTaskInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ArchivalTaskInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ArchivalTaskInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ArchivalTaskInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ArchivalTaskInfo
	switch t := that.(type) {
	case *ArchivalTaskInfo:
		that1 = t
	case ArchivalTaskInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ActivityInfo to the protobuf v3 wire format
func (val *ActivityInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ActivityInfo from the protobuf v3 wire format
func (val *ActivityInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ActivityInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ActivityInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ActivityInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ActivityInfo
	switch t := that.(type) {
	case *ActivityInfo:
		that1 = t
	case ActivityInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type TimerInfo to the protobuf v3 wire format
func (val *TimerInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type TimerInfo from the protobuf v3 wire format
func (val *TimerInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *TimerInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two TimerInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *TimerInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *TimerInfo
	switch t := that.(type) {
	case *TimerInfo:
		that1 = t
	case TimerInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ChildExecutionInfo to the protobuf v3 wire format
func (val *ChildExecutionInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ChildExecutionInfo from the protobuf v3 wire format
func (val *ChildExecutionInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ChildExecutionInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ChildExecutionInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ChildExecutionInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ChildExecutionInfo
	switch t := that.(type) {
	case *ChildExecutionInfo:
		that1 = t
	case ChildExecutionInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RequestCancelInfo to the protobuf v3 wire format
func (val *RequestCancelInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RequestCancelInfo from the protobuf v3 wire format
func (val *RequestCancelInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RequestCancelInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RequestCancelInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RequestCancelInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RequestCancelInfo
	switch t := that.(type) {
	case *RequestCancelInfo:
		that1 = t
	case RequestCancelInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type SignalInfo to the protobuf v3 wire format
func (val *SignalInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type SignalInfo from the protobuf v3 wire format
func (val *SignalInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *SignalInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two SignalInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *SignalInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *SignalInfo
	switch t := that.(type) {
	case *SignalInfo:
		that1 = t
	case SignalInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type Checksum to the protobuf v3 wire format
func (val *Checksum) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type Checksum from the protobuf v3 wire format
func (val *Checksum) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *Checksum) Size() int {
	return proto.Size(val)
}

// Equal returns whether two Checksum values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *Checksum) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *Checksum
	switch t := that.(type) {
	case *Checksum:
		that1 = t
	case Checksum:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
