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

// Copyright (c) 2020 Temporal Technologies, Inc.
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// source: temporal/server/api/persistence/v1/cluster_metadata.proto

package persistence

import (
	reflect "reflect"
	sync "sync"

	v11 "go.temporal.io/api/enums/v1"
	v1 "go.temporal.io/api/version/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// data column
type ClusterMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterName              string                            `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	HistoryShardCount        int32                             `protobuf:"varint,2,opt,name=history_shard_count,json=historyShardCount,proto3" json:"history_shard_count,omitempty"`
	ClusterId                string                            `protobuf:"bytes,3,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	VersionInfo              *v1.VersionInfo                   `protobuf:"bytes,4,opt,name=version_info,json=versionInfo,proto3" json:"version_info,omitempty"`
	IndexSearchAttributes    map[string]*IndexSearchAttributes `protobuf:"bytes,5,rep,name=index_search_attributes,json=indexSearchAttributes,proto3" json:"index_search_attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ClusterAddress           string                            `protobuf:"bytes,6,opt,name=cluster_address,json=clusterAddress,proto3" json:"cluster_address,omitempty"`
	FailoverVersionIncrement int64                             `protobuf:"varint,7,opt,name=failover_version_increment,json=failoverVersionIncrement,proto3" json:"failover_version_increment,omitempty"`
	InitialFailoverVersion   int64                             `protobuf:"varint,8,opt,name=initial_failover_version,json=initialFailoverVersion,proto3" json:"initial_failover_version,omitempty"`
	IsGlobalNamespaceEnabled bool                              `protobuf:"varint,9,opt,name=is_global_namespace_enabled,json=isGlobalNamespaceEnabled,proto3" json:"is_global_namespace_enabled,omitempty"`
	IsConnectionEnabled      bool                              `protobuf:"varint,10,opt,name=is_connection_enabled,json=isConnectionEnabled,proto3" json:"is_connection_enabled,omitempty"`
	UseClusterIdMembership   bool                              `protobuf:"varint,11,opt,name=use_cluster_id_membership,json=useClusterIdMembership,proto3" json:"use_cluster_id_membership,omitempty"`
	Tags                     map[string]string                 `protobuf:"bytes,12,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ClusterMetadata) Reset() {
	*x = ClusterMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_server_api_persistence_v1_cluster_metadata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterMetadata) ProtoMessage() {}

func (x *ClusterMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_server_api_persistence_v1_cluster_metadata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterMetadata.ProtoReflect.Descriptor instead.
func (*ClusterMetadata) Descriptor() ([]byte, []int) {
	return file_temporal_server_api_persistence_v1_cluster_metadata_proto_rawDescGZIP(), []int{0}
}

func (x *ClusterMetadata) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *ClusterMetadata) GetHistoryShardCount() int32 {
	if x != nil {
		return x.HistoryShardCount
	}
	return 0
}

func (x *ClusterMetadata) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *ClusterMetadata) GetVersionInfo() *v1.VersionInfo {
	if x != nil {
		return x.VersionInfo
	}
	return nil
}

func (x *ClusterMetadata) GetIndexSearchAttributes() map[string]*IndexSearchAttributes {
	if x != nil {
		return x.IndexSearchAttributes
	}
	return nil
}

func (x *ClusterMetadata) GetClusterAddress() string {
	if x != nil {
		return x.ClusterAddress
	}
	return ""
}

func (x *ClusterMetadata) GetFailoverVersionIncrement() int64 {
	if x != nil {
		return x.FailoverVersionIncrement
	}
	return 0
}

func (x *ClusterMetadata) GetInitialFailoverVersion() int64 {
	if x != nil {
		return x.InitialFailoverVersion
	}
	return 0
}

func (x *ClusterMetadata) GetIsGlobalNamespaceEnabled() bool {
	if x != nil {
		return x.IsGlobalNamespaceEnabled
	}
	return false
}

func (x *ClusterMetadata) GetIsConnectionEnabled() bool {
	if x != nil {
		return x.IsConnectionEnabled
	}
	return false
}

func (x *ClusterMetadata) GetUseClusterIdMembership() bool {
	if x != nil {
		return x.UseClusterIdMembership
	}
	return false
}

func (x *ClusterMetadata) GetTags() map[string]string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type IndexSearchAttributes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomSearchAttributes map[string]v11.IndexedValueType `protobuf:"bytes,1,rep,name=custom_search_attributes,json=customSearchAttributes,proto3" json:"custom_search_attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=temporal.api.enums.v1.IndexedValueType"`
}

func (x *IndexSearchAttributes) Reset() {
	*x = IndexSearchAttributes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_server_api_persistence_v1_cluster_metadata_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexSearchAttributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexSearchAttributes) ProtoMessage() {}

func (x *IndexSearchAttributes) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_server_api_persistence_v1_cluster_metadata_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexSearchAttributes.ProtoReflect.Descriptor instead.
func (*IndexSearchAttributes) Descriptor() ([]byte, []int) {
	return file_temporal_server_api_persistence_v1_cluster_metadata_proto_rawDescGZIP(), []int{1}
}

func (x *IndexSearchAttributes) GetCustomSearchAttributes() map[string]v11.IndexedValueType {
	if x != nil {
		return x.CustomSearchAttributes
	}
	return nil
}

var File_temporal_server_api_persistence_v1_cluster_metadata_proto protoreflect.FileDescriptor

var file_temporal_server_api_persistence_v1_cluster_metadata_proto_rawDesc = []byte{
	0x0a, 0x39, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x74, 0x65, 0x6d,
	0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a,
	0x22, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x07, 0x0a, 0x0f, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x21,
	0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x2e, 0x0a, 0x13, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x68, 0x61,
	0x72, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11,
	0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x68, 0x61, 0x72, 0x64, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x47, 0x0a, 0x0c, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61,
	0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x86, 0x01, 0x0a, 0x17, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4e, 0x2e, 0x74, 0x65,
	0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x15, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x3c, 0x0a, 0x1a, 0x66,
	0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x18, 0x66, 0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x18, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x61, 0x6c, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x16, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x61, 0x6c, 0x46, 0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x3d, 0x0a, 0x1b, 0x69, 0x73, 0x5f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x18, 0x69, 0x73, 0x47, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x45, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x12, 0x32, 0x0a, 0x15, 0x69, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x13, 0x69, 0x73, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x39, 0x0a, 0x19, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73,
	0x68, 0x69, 0x70, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x75, 0x73, 0x65, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69,
	0x70, 0x12, 0x51, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x3d, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x1a, 0x83, 0x01, 0x0a, 0x1a, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x4f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x69,
	0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x37, 0x0a, 0x09, 0x54, 0x61,
	0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x9d, 0x02, 0x0a, 0x15, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x8f, 0x01,
	0x0a, 0x18, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f,
	0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x55, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x16, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x1a,
	0x72, 0x0a, 0x1b, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x3d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x27, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x64, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72,
	0x61, 0x6c, 0x2e, 0x69, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b,
	0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_temporal_server_api_persistence_v1_cluster_metadata_proto_rawDescOnce sync.Once
	file_temporal_server_api_persistence_v1_cluster_metadata_proto_rawDescData = file_temporal_server_api_persistence_v1_cluster_metadata_proto_rawDesc
)

func file_temporal_server_api_persistence_v1_cluster_metadata_proto_rawDescGZIP() []byte {
	file_temporal_server_api_persistence_v1_cluster_metadata_proto_rawDescOnce.Do(func() {
		file_temporal_server_api_persistence_v1_cluster_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(file_temporal_server_api_persistence_v1_cluster_metadata_proto_rawDescData)
	})
	return file_temporal_server_api_persistence_v1_cluster_metadata_proto_rawDescData
}

var file_temporal_server_api_persistence_v1_cluster_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_temporal_server_api_persistence_v1_cluster_metadata_proto_goTypes = []interface{}{
	(*ClusterMetadata)(nil),       // 0: temporal.server.api.persistence.v1.ClusterMetadata
	(*IndexSearchAttributes)(nil), // 1: temporal.server.api.persistence.v1.IndexSearchAttributes
	nil,                           // 2: temporal.server.api.persistence.v1.ClusterMetadata.IndexSearchAttributesEntry
	nil,                           // 3: temporal.server.api.persistence.v1.ClusterMetadata.TagsEntry
	nil,                           // 4: temporal.server.api.persistence.v1.IndexSearchAttributes.CustomSearchAttributesEntry
	(*v1.VersionInfo)(nil),        // 5: temporal.api.version.v1.VersionInfo
	(v11.IndexedValueType)(0),     // 6: temporal.api.enums.v1.IndexedValueType
}
var file_temporal_server_api_persistence_v1_cluster_metadata_proto_depIdxs = []int32{
	5, // 0: temporal.server.api.persistence.v1.ClusterMetadata.version_info:type_name -> temporal.api.version.v1.VersionInfo
	2, // 1: temporal.server.api.persistence.v1.ClusterMetadata.index_search_attributes:type_name -> temporal.server.api.persistence.v1.ClusterMetadata.IndexSearchAttributesEntry
	3, // 2: temporal.server.api.persistence.v1.ClusterMetadata.tags:type_name -> temporal.server.api.persistence.v1.ClusterMetadata.TagsEntry
	4, // 3: temporal.server.api.persistence.v1.IndexSearchAttributes.custom_search_attributes:type_name -> temporal.server.api.persistence.v1.IndexSearchAttributes.CustomSearchAttributesEntry
	1, // 4: temporal.server.api.persistence.v1.ClusterMetadata.IndexSearchAttributesEntry.value:type_name -> temporal.server.api.persistence.v1.IndexSearchAttributes
	6, // 5: temporal.server.api.persistence.v1.IndexSearchAttributes.CustomSearchAttributesEntry.value:type_name -> temporal.api.enums.v1.IndexedValueType
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_temporal_server_api_persistence_v1_cluster_metadata_proto_init() }
func file_temporal_server_api_persistence_v1_cluster_metadata_proto_init() {
	if File_temporal_server_api_persistence_v1_cluster_metadata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_temporal_server_api_persistence_v1_cluster_metadata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterMetadata); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_temporal_server_api_persistence_v1_cluster_metadata_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexSearchAttributes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_temporal_server_api_persistence_v1_cluster_metadata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_server_api_persistence_v1_cluster_metadata_proto_goTypes,
		DependencyIndexes: file_temporal_server_api_persistence_v1_cluster_metadata_proto_depIdxs,
		MessageInfos:      file_temporal_server_api_persistence_v1_cluster_metadata_proto_msgTypes,
	}.Build()
	File_temporal_server_api_persistence_v1_cluster_metadata_proto = out.File
	file_temporal_server_api_persistence_v1_cluster_metadata_proto_rawDesc = nil
	file_temporal_server_api_persistence_v1_cluster_metadata_proto_goTypes = nil
	file_temporal_server_api_persistence_v1_cluster_metadata_proto_depIdxs = nil
}
