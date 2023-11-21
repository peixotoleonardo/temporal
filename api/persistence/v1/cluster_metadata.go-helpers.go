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

// Marshal an object of type ClusterMetadata to the protobuf v3 wire format
func (val *ClusterMetadata) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ClusterMetadata from the protobuf v3 wire format
func (val *ClusterMetadata) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ClusterMetadata) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ClusterMetadata values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ClusterMetadata) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ClusterMetadata
	switch t := that.(type) {
	case *ClusterMetadata:
		that1 = t
	case ClusterMetadata:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type IndexSearchAttributes to the protobuf v3 wire format
func (val *IndexSearchAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type IndexSearchAttributes from the protobuf v3 wire format
func (val *IndexSearchAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *IndexSearchAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two IndexSearchAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *IndexSearchAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *IndexSearchAttributes
	switch t := that.(type) {
	case *IndexSearchAttributes:
		that1 = t
	case IndexSearchAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
