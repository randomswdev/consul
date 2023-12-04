// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: pbmesh/v2beta1/gatewayclassconfig.proto

package meshv2beta1

import (
	_ "github.com/hashicorp/consul/proto-public/pbresource"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	v1 "k8s.io/api/core/v1"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a Resource type.
type GatewayClassConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// +kubebuilder:validation:Enum=ClusterIP;NodePort;LoadBalancer
	ServiceType *string `protobuf:"bytes,1,opt,name=service_type,json=serviceType,proto3,oneof" json:"service_type,omitempty"`
	// NodeSelector is a selector which must be true for the pod to fit on a node.
	// Selector which must match a node's labels for the pod to be scheduled on that node.
	// More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
	NodeSelector map[string]string `protobuf:"bytes,2,rep,name=node_selector,json=nodeSelector,proto3" json:"node_selector,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Tolerations allow the scheduler to schedule nodes with matching taints.
	// More Info: https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/
	Tolerations       []*v1.Toleration `protobuf:"bytes,3,rep,name=tolerations,proto3" json:"tolerations,omitempty"`
	Deployment        *Deployment      `protobuf:"bytes,4,opt,name=deployment,proto3,oneof" json:"deployment,omitempty"`
	CopyAnnotations   *CopyAnnotations `protobuf:"bytes,5,opt,name=copy_annotations,json=copyAnnotations,proto3,oneof" json:"copy_annotations,omitempty"`
	PodSecurityPolicy *string          `protobuf:"bytes,6,opt,name=pod_security_policy,json=podSecurityPolicy,proto3,oneof" json:"pod_security_policy,omitempty"`
	OpenshiftSccName  *string          `protobuf:"bytes,7,opt,name=openshift_scc_name,json=openshiftSccName,proto3,oneof" json:"openshift_scc_name,omitempty"`
	UseHostPorts      *bool            `protobuf:"varint,8,opt,name=use_host_ports,json=useHostPorts,proto3,oneof" json:"use_host_ports,omitempty"`
}

func (x *GatewayClassConfig) Reset() {
	*x = GatewayClassConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbmesh_v2beta1_gatewayclassconfig_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatewayClassConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayClassConfig) ProtoMessage() {}

func (x *GatewayClassConfig) ProtoReflect() protoreflect.Message {
	mi := &file_pbmesh_v2beta1_gatewayclassconfig_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayClassConfig.ProtoReflect.Descriptor instead.
func (*GatewayClassConfig) Descriptor() ([]byte, []int) {
	return file_pbmesh_v2beta1_gatewayclassconfig_proto_rawDescGZIP(), []int{0}
}

func (x *GatewayClassConfig) GetServiceType() string {
	if x != nil && x.ServiceType != nil {
		return *x.ServiceType
	}
	return ""
}

func (x *GatewayClassConfig) GetNodeSelector() map[string]string {
	if x != nil {
		return x.NodeSelector
	}
	return nil
}

func (x *GatewayClassConfig) GetTolerations() []*v1.Toleration {
	if x != nil {
		return x.Tolerations
	}
	return nil
}

func (x *GatewayClassConfig) GetDeployment() *Deployment {
	if x != nil {
		return x.Deployment
	}
	return nil
}

func (x *GatewayClassConfig) GetCopyAnnotations() *CopyAnnotations {
	if x != nil {
		return x.CopyAnnotations
	}
	return nil
}

func (x *GatewayClassConfig) GetPodSecurityPolicy() string {
	if x != nil && x.PodSecurityPolicy != nil {
		return *x.PodSecurityPolicy
	}
	return ""
}

func (x *GatewayClassConfig) GetOpenshiftSccName() string {
	if x != nil && x.OpenshiftSccName != nil {
		return *x.OpenshiftSccName
	}
	return ""
}

func (x *GatewayClassConfig) GetUseHostPorts() bool {
	if x != nil && x.UseHostPorts != nil {
		return *x.UseHostPorts
	}
	return false
}

type CopyAnnotations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service []string `protobuf:"bytes,1,rep,name=service,proto3" json:"service,omitempty"`
}

func (x *CopyAnnotations) Reset() {
	*x = CopyAnnotations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbmesh_v2beta1_gatewayclassconfig_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CopyAnnotations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CopyAnnotations) ProtoMessage() {}

func (x *CopyAnnotations) ProtoReflect() protoreflect.Message {
	mi := &file_pbmesh_v2beta1_gatewayclassconfig_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CopyAnnotations.ProtoReflect.Descriptor instead.
func (*CopyAnnotations) Descriptor() ([]byte, []int) {
	return file_pbmesh_v2beta1_gatewayclassconfig_proto_rawDescGZIP(), []int{1}
}

func (x *CopyAnnotations) GetService() []string {
	if x != nil {
		return x.Service
	}
	return nil
}

type Deployment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// +kubebuilder:default:=1
	// +kubebuilder:validation:Maximum=8
	// +kubebuilder:validation:Minimum=1
	// Number of gateway instances that should be deployed by default
	DefaultInstances *uint32 `protobuf:"varint,1,opt,name=default_instances,json=defaultInstances,proto3,oneof" json:"default_instances,omitempty"`
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Maximum=8
	// +kubebuilder:validation:Minimum=1
	// Minimum allowed number of gateway instances
	MinInstances *uint32 `protobuf:"varint,2,opt,name=min_instances,json=minInstances,proto3,oneof" json:"min_instances,omitempty"`
	// +kubebuilder:default:=8
	// +kubebuilder:validation:Maximum=8
	// +kubebuilder:validation:Minimum=1
	// Max allowed number of gateway instances
	MaxInstances *uint32 `protobuf:"varint,3,opt,name=max_instances,json=maxInstances,proto3,oneof" json:"max_instances,omitempty"`
	// Resources defines the resource requirements for the gateway.
	Resources *v1.ResourceRequirements `protobuf:"bytes,4,opt,name=resources,proto3" json:"resources,omitempty"`
}

func (x *Deployment) Reset() {
	*x = Deployment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbmesh_v2beta1_gatewayclassconfig_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Deployment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deployment) ProtoMessage() {}

func (x *Deployment) ProtoReflect() protoreflect.Message {
	mi := &file_pbmesh_v2beta1_gatewayclassconfig_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Deployment.ProtoReflect.Descriptor instead.
func (*Deployment) Descriptor() ([]byte, []int) {
	return file_pbmesh_v2beta1_gatewayclassconfig_proto_rawDescGZIP(), []int{2}
}

func (x *Deployment) GetDefaultInstances() uint32 {
	if x != nil && x.DefaultInstances != nil {
		return *x.DefaultInstances
	}
	return 0
}

func (x *Deployment) GetMinInstances() uint32 {
	if x != nil && x.MinInstances != nil {
		return *x.MinInstances
	}
	return 0
}

func (x *Deployment) GetMaxInstances() uint32 {
	if x != nil && x.MaxInstances != nil {
		return *x.MaxInstances
	}
	return 0
}

func (x *Deployment) GetResources() *v1.ResourceRequirements {
	if x != nil {
		return x.Resources
	}
	return nil
}

var File_pbmesh_v2beta1_gatewayclassconfig_proto protoreflect.FileDescriptor

var file_pbmesh_v2beta1_gatewayclassconfig_proto_rawDesc = []byte{
	0x0a, 0x27, 0x70, 0x62, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x68, 0x61, 0x73, 0x68, 0x69,
	0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x68,
	0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x22, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x70, 0x62,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xeb, 0x05, 0x0a, 0x12, 0x47,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x26, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x68, 0x0a, 0x0d, 0x6e, 0x6f, 0x64,
	0x65, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x43, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e,
	0x73, 0x75, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x6e, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x12, 0x40, 0x0a, 0x0b, 0x74, 0x6f, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69,
	0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f,
	0x6c, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x6f, 0x6c, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x4e, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x68, 0x61, 0x73, 0x68,
	0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6d, 0x65, 0x73,
	0x68, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x48, 0x01, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x5e, 0x0a, 0x10, 0x63, 0x6f, 0x70, 0x79, 0x5f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2e, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73,
	0x75, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x43, 0x6f, 0x70, 0x79, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x48,
	0x02, 0x52, 0x0f, 0x63, 0x6f, 0x70, 0x79, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x88, 0x01, 0x01, 0x12, 0x33, 0x0a, 0x13, 0x70, 0x6f, 0x64, 0x5f, 0x73, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x03, 0x52, 0x11, 0x70, 0x6f, 0x64, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x88, 0x01, 0x01, 0x12, 0x31, 0x0a, 0x12, 0x6f, 0x70,
	0x65, 0x6e, 0x73, 0x68, 0x69, 0x66, 0x74, 0x5f, 0x73, 0x63, 0x63, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x10, 0x6f, 0x70, 0x65, 0x6e, 0x73, 0x68,
	0x69, 0x66, 0x74, 0x53, 0x63, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a,
	0x0e, 0x75, 0x73, 0x65, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x08, 0x48, 0x05, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x48, 0x6f, 0x73, 0x74,
	0x50, 0x6f, 0x72, 0x74, 0x73, 0x88, 0x01, 0x01, 0x1a, 0x3f, 0x0a, 0x11, 0x4e, 0x6f, 0x64, 0x65,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0x06, 0xa2, 0x93, 0x04, 0x02, 0x08,
	0x01, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x63, 0x6f, 0x70, 0x79, 0x5f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x16, 0x0a, 0x14, 0x5f, 0x70, 0x6f, 0x64, 0x5f, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0x15,
	0x0a, 0x13, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x73, 0x68, 0x69, 0x66, 0x74, 0x5f, 0x73, 0x63, 0x63,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x75, 0x73, 0x65, 0x5f, 0x68, 0x6f,
	0x73, 0x74, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x22, 0x2b, 0x0a, 0x0f, 0x43, 0x6f, 0x70, 0x79,
	0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x94, 0x02, 0x0a, 0x0a, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x11, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x48,
	0x00, 0x52, 0x10, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x73, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x0d, 0x6d, 0x69, 0x6e, 0x5f, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x01, 0x52,
	0x0c, 0x6d, 0x69, 0x6e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x88, 0x01, 0x01,
	0x12, 0x28, 0x0a, 0x0d, 0x6d, 0x61, 0x78, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x02, 0x52, 0x0c, 0x6d, 0x61, 0x78, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x88, 0x01, 0x01, 0x12, 0x46, 0x0a, 0x09, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e,
	0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x6d, 0x69, 0x6e,
	0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x6d,
	0x61, 0x78, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x42, 0x98, 0x02, 0x0a,
	0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63,
	0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x42, 0x17, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63,
	0x6f, 0x72, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x70, 0x62, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76,
	0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x6d, 0x65, 0x73, 0x68, 0x76, 0x32, 0x62, 0x65, 0x74,
	0x61, 0x31, 0xa2, 0x02, 0x03, 0x48, 0x43, 0x4d, 0xaa, 0x02, 0x1d, 0x48, 0x61, 0x73, 0x68, 0x69,
	0x63, 0x6f, 0x72, 0x70, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x4d, 0x65, 0x73, 0x68,
	0x2e, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x1d, 0x48, 0x61, 0x73, 0x68, 0x69,
	0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x5c, 0x4d, 0x65, 0x73, 0x68,
	0x5c, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x29, 0x48, 0x61, 0x73, 0x68, 0x69,
	0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x5c, 0x4d, 0x65, 0x73, 0x68,
	0x5c, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x20, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70,
	0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x3a, 0x3a, 0x4d, 0x65, 0x73, 0x68, 0x3a, 0x3a,
	0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pbmesh_v2beta1_gatewayclassconfig_proto_rawDescOnce sync.Once
	file_pbmesh_v2beta1_gatewayclassconfig_proto_rawDescData = file_pbmesh_v2beta1_gatewayclassconfig_proto_rawDesc
)

func file_pbmesh_v2beta1_gatewayclassconfig_proto_rawDescGZIP() []byte {
	file_pbmesh_v2beta1_gatewayclassconfig_proto_rawDescOnce.Do(func() {
		file_pbmesh_v2beta1_gatewayclassconfig_proto_rawDescData = protoimpl.X.CompressGZIP(file_pbmesh_v2beta1_gatewayclassconfig_proto_rawDescData)
	})
	return file_pbmesh_v2beta1_gatewayclassconfig_proto_rawDescData
}

var file_pbmesh_v2beta1_gatewayclassconfig_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pbmesh_v2beta1_gatewayclassconfig_proto_goTypes = []interface{}{
	(*GatewayClassConfig)(nil),      // 0: hashicorp.consul.mesh.v2beta1.GatewayClassConfig
	(*CopyAnnotations)(nil),         // 1: hashicorp.consul.mesh.v2beta1.CopyAnnotations
	(*Deployment)(nil),              // 2: hashicorp.consul.mesh.v2beta1.Deployment
	nil,                             // 3: hashicorp.consul.mesh.v2beta1.GatewayClassConfig.NodeSelectorEntry
	(*v1.Toleration)(nil),           // 4: k8s.io.api.core.v1.Toleration
	(*v1.ResourceRequirements)(nil), // 5: k8s.io.api.core.v1.ResourceRequirements
}
var file_pbmesh_v2beta1_gatewayclassconfig_proto_depIdxs = []int32{
	3, // 0: hashicorp.consul.mesh.v2beta1.GatewayClassConfig.node_selector:type_name -> hashicorp.consul.mesh.v2beta1.GatewayClassConfig.NodeSelectorEntry
	4, // 1: hashicorp.consul.mesh.v2beta1.GatewayClassConfig.tolerations:type_name -> k8s.io.api.core.v1.Toleration
	2, // 2: hashicorp.consul.mesh.v2beta1.GatewayClassConfig.deployment:type_name -> hashicorp.consul.mesh.v2beta1.Deployment
	1, // 3: hashicorp.consul.mesh.v2beta1.GatewayClassConfig.copy_annotations:type_name -> hashicorp.consul.mesh.v2beta1.CopyAnnotations
	5, // 4: hashicorp.consul.mesh.v2beta1.Deployment.resources:type_name -> k8s.io.api.core.v1.ResourceRequirements
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_pbmesh_v2beta1_gatewayclassconfig_proto_init() }
func file_pbmesh_v2beta1_gatewayclassconfig_proto_init() {
	if File_pbmesh_v2beta1_gatewayclassconfig_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pbmesh_v2beta1_gatewayclassconfig_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatewayClassConfig); i {
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
		file_pbmesh_v2beta1_gatewayclassconfig_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CopyAnnotations); i {
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
		file_pbmesh_v2beta1_gatewayclassconfig_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Deployment); i {
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
	file_pbmesh_v2beta1_gatewayclassconfig_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_pbmesh_v2beta1_gatewayclassconfig_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pbmesh_v2beta1_gatewayclassconfig_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pbmesh_v2beta1_gatewayclassconfig_proto_goTypes,
		DependencyIndexes: file_pbmesh_v2beta1_gatewayclassconfig_proto_depIdxs,
		MessageInfos:      file_pbmesh_v2beta1_gatewayclassconfig_proto_msgTypes,
	}.Build()
	File_pbmesh_v2beta1_gatewayclassconfig_proto = out.File
	file_pbmesh_v2beta1_gatewayclassconfig_proto_rawDesc = nil
	file_pbmesh_v2beta1_gatewayclassconfig_proto_goTypes = nil
	file_pbmesh_v2beta1_gatewayclassconfig_proto_depIdxs = nil
}
