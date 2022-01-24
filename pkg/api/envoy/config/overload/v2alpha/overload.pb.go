// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: envoy/config/overload/v2alpha/overload.proto

package v2alpha

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	any "github.com/golang/protobuf/ptypes/any"
	duration "github.com/golang/protobuf/ptypes/duration"
	_struct "github.com/golang/protobuf/ptypes/struct"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ResourceMonitor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the resource monitor to instantiate. Must match a registered
	// resource monitor type. The built-in resource monitors are:
	//
	// * :ref:`envoy.resource_monitors.fixed_heap
	//   <envoy_api_msg_config.resource_monitor.fixed_heap.v2alpha.FixedHeapConfig>`
	// * :ref:`envoy.resource_monitors.injected_resource
	//   <envoy_api_msg_config.resource_monitor.injected_resource.v2alpha.InjectedResourceConfig>`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Configuration for the resource monitor being instantiated.
	//
	// Types that are assignable to ConfigType:
	//	*ResourceMonitor_Config
	//	*ResourceMonitor_TypedConfig
	ConfigType isResourceMonitor_ConfigType `protobuf_oneof:"config_type"`
}

func (x *ResourceMonitor) Reset() {
	*x = ResourceMonitor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_overload_v2alpha_overload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceMonitor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceMonitor) ProtoMessage() {}

func (x *ResourceMonitor) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_overload_v2alpha_overload_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceMonitor.ProtoReflect.Descriptor instead.
func (*ResourceMonitor) Descriptor() ([]byte, []int) {
	return file_envoy_config_overload_v2alpha_overload_proto_rawDescGZIP(), []int{0}
}

func (x *ResourceMonitor) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (m *ResourceMonitor) GetConfigType() isResourceMonitor_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

// Deprecated: Do not use.
func (x *ResourceMonitor) GetConfig() *_struct.Struct {
	if x, ok := x.GetConfigType().(*ResourceMonitor_Config); ok {
		return x.Config
	}
	return nil
}

func (x *ResourceMonitor) GetTypedConfig() *any.Any {
	if x, ok := x.GetConfigType().(*ResourceMonitor_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

type isResourceMonitor_ConfigType interface {
	isResourceMonitor_ConfigType()
}

type ResourceMonitor_Config struct {
	// Deprecated: Do not use.
	Config *_struct.Struct `protobuf:"bytes,2,opt,name=config,proto3,oneof"`
}

type ResourceMonitor_TypedConfig struct {
	TypedConfig *any.Any `protobuf:"bytes,3,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

func (*ResourceMonitor_Config) isResourceMonitor_ConfigType() {}

func (*ResourceMonitor_TypedConfig) isResourceMonitor_ConfigType() {}

type ThresholdTrigger struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If the resource pressure is greater than or equal to this value, the trigger
	// will fire.
	Value float64 `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ThresholdTrigger) Reset() {
	*x = ThresholdTrigger{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_overload_v2alpha_overload_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThresholdTrigger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThresholdTrigger) ProtoMessage() {}

func (x *ThresholdTrigger) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_overload_v2alpha_overload_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThresholdTrigger.ProtoReflect.Descriptor instead.
func (*ThresholdTrigger) Descriptor() ([]byte, []int) {
	return file_envoy_config_overload_v2alpha_overload_proto_rawDescGZIP(), []int{1}
}

func (x *ThresholdTrigger) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type Trigger struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the resource this is a trigger for.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are assignable to TriggerOneof:
	//	*Trigger_Threshold
	TriggerOneof isTrigger_TriggerOneof `protobuf_oneof:"trigger_oneof"`
}

func (x *Trigger) Reset() {
	*x = Trigger{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_overload_v2alpha_overload_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trigger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trigger) ProtoMessage() {}

func (x *Trigger) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_overload_v2alpha_overload_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trigger.ProtoReflect.Descriptor instead.
func (*Trigger) Descriptor() ([]byte, []int) {
	return file_envoy_config_overload_v2alpha_overload_proto_rawDescGZIP(), []int{2}
}

func (x *Trigger) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (m *Trigger) GetTriggerOneof() isTrigger_TriggerOneof {
	if m != nil {
		return m.TriggerOneof
	}
	return nil
}

func (x *Trigger) GetThreshold() *ThresholdTrigger {
	if x, ok := x.GetTriggerOneof().(*Trigger_Threshold); ok {
		return x.Threshold
	}
	return nil
}

type isTrigger_TriggerOneof interface {
	isTrigger_TriggerOneof()
}

type Trigger_Threshold struct {
	Threshold *ThresholdTrigger `protobuf:"bytes,2,opt,name=threshold,proto3,oneof"`
}

func (*Trigger_Threshold) isTrigger_TriggerOneof() {}

type OverloadAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the overload action. This is just a well-known string that listeners can
	// use for registering callbacks. Custom overload actions should be named using reverse
	// DNS to ensure uniqueness.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A set of triggers for this action. If any of these triggers fire the overload action
	// is activated. Listeners are notified when the overload action transitions from
	// inactivated to activated, or vice versa.
	Triggers []*Trigger `protobuf:"bytes,2,rep,name=triggers,proto3" json:"triggers,omitempty"`
}

func (x *OverloadAction) Reset() {
	*x = OverloadAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_overload_v2alpha_overload_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OverloadAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OverloadAction) ProtoMessage() {}

func (x *OverloadAction) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_overload_v2alpha_overload_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OverloadAction.ProtoReflect.Descriptor instead.
func (*OverloadAction) Descriptor() ([]byte, []int) {
	return file_envoy_config_overload_v2alpha_overload_proto_rawDescGZIP(), []int{3}
}

func (x *OverloadAction) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OverloadAction) GetTriggers() []*Trigger {
	if x != nil {
		return x.Triggers
	}
	return nil
}

type OverloadManager struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The interval for refreshing resource usage.
	RefreshInterval *duration.Duration `protobuf:"bytes,1,opt,name=refresh_interval,json=refreshInterval,proto3" json:"refresh_interval,omitempty"`
	// The set of resources to monitor.
	ResourceMonitors []*ResourceMonitor `protobuf:"bytes,2,rep,name=resource_monitors,json=resourceMonitors,proto3" json:"resource_monitors,omitempty"`
	// The set of overload actions.
	Actions []*OverloadAction `protobuf:"bytes,3,rep,name=actions,proto3" json:"actions,omitempty"`
}

func (x *OverloadManager) Reset() {
	*x = OverloadManager{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_overload_v2alpha_overload_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OverloadManager) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OverloadManager) ProtoMessage() {}

func (x *OverloadManager) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_overload_v2alpha_overload_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OverloadManager.ProtoReflect.Descriptor instead.
func (*OverloadManager) Descriptor() ([]byte, []int) {
	return file_envoy_config_overload_v2alpha_overload_proto_rawDescGZIP(), []int{4}
}

func (x *OverloadManager) GetRefreshInterval() *duration.Duration {
	if x != nil {
		return x.RefreshInterval
	}
	return nil
}

func (x *OverloadManager) GetResourceMonitors() []*ResourceMonitor {
	if x != nil {
		return x.ResourceMonitors
	}
	return nil
}

func (x *OverloadManager) GetActions() []*OverloadAction {
	if x != nil {
		return x.Actions
	}
	return nil
}

var File_envoy_config_overload_v2alpha_overload_proto protoreflect.FileDescriptor

var file_envoy_config_overload_v2alpha_overload_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6f,
	0x76, 0x65, 0x72, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f,
	0x6f, 0x76, 0x65, 0x72, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6f, 0x76, 0x65,
	0x72, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61,
	0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf,
	0x01, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x35, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x42, 0x02, 0x18, 0x01, 0x48, 0x00, 0x52, 0x06,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x39, 0x0a, 0x0c, 0x74, 0x79, 0x70, 0x65, 0x64, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41,
	0x6e, 0x79, 0x48, 0x00, 0x52, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x42, 0x0d, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x22, 0x41, 0x0a, 0x10, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x54, 0x72, 0x69,
	0x67, 0x67, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x01, 0x42, 0x17, 0xfa, 0x42, 0x14, 0x12, 0x12, 0x19, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0xf0, 0x3f, 0x29, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x8d, 0x01, 0x0a, 0x07, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x12,
	0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x4f, 0x0a, 0x09,
	0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6f,
	0x76, 0x65, 0x72, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72,
	0x48, 0x00, 0x52, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x42, 0x14, 0x0a,
	0x0d, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x12, 0x03,
	0xf8, 0x42, 0x01, 0x22, 0x7b, 0x0a, 0x0e, 0x4f, 0x76, 0x65, 0x72, 0x6c, 0x6f, 0x61, 0x64, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x4c, 0x0a, 0x08, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x32, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2e, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x42, 0x08, 0xfa, 0x42,
	0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x08, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73,
	0x22, 0x87, 0x02, 0x0a, 0x0f, 0x4f, 0x76, 0x65, 0x72, 0x6c, 0x6f, 0x61, 0x64, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x12, 0x44, 0x0a, 0x10, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x72, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x65, 0x0a, 0x11, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x32,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52,
	0x10, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x73, 0x12, 0x47, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2e, 0x4f, 0x76, 0x65, 0x72, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x8c, 0x01, 0x0a, 0x2b, 0x69,
	0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x6f,
	0x61, 0x64, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x0d, 0x4f, 0x76, 0x65, 0x72,
	0x6c, 0x6f, 0x61, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c,
	0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2f, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_envoy_config_overload_v2alpha_overload_proto_rawDescOnce sync.Once
	file_envoy_config_overload_v2alpha_overload_proto_rawDescData = file_envoy_config_overload_v2alpha_overload_proto_rawDesc
)

func file_envoy_config_overload_v2alpha_overload_proto_rawDescGZIP() []byte {
	file_envoy_config_overload_v2alpha_overload_proto_rawDescOnce.Do(func() {
		file_envoy_config_overload_v2alpha_overload_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_overload_v2alpha_overload_proto_rawDescData)
	})
	return file_envoy_config_overload_v2alpha_overload_proto_rawDescData
}

var file_envoy_config_overload_v2alpha_overload_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_envoy_config_overload_v2alpha_overload_proto_goTypes = []interface{}{
	(*ResourceMonitor)(nil),   // 0: envoy.config.overload.v2alpha.ResourceMonitor
	(*ThresholdTrigger)(nil),  // 1: envoy.config.overload.v2alpha.ThresholdTrigger
	(*Trigger)(nil),           // 2: envoy.config.overload.v2alpha.Trigger
	(*OverloadAction)(nil),    // 3: envoy.config.overload.v2alpha.OverloadAction
	(*OverloadManager)(nil),   // 4: envoy.config.overload.v2alpha.OverloadManager
	(*_struct.Struct)(nil),    // 5: google.protobuf.Struct
	(*any.Any)(nil),           // 6: google.protobuf.Any
	(*duration.Duration)(nil), // 7: google.protobuf.Duration
}
var file_envoy_config_overload_v2alpha_overload_proto_depIdxs = []int32{
	5, // 0: envoy.config.overload.v2alpha.ResourceMonitor.config:type_name -> google.protobuf.Struct
	6, // 1: envoy.config.overload.v2alpha.ResourceMonitor.typed_config:type_name -> google.protobuf.Any
	1, // 2: envoy.config.overload.v2alpha.Trigger.threshold:type_name -> envoy.config.overload.v2alpha.ThresholdTrigger
	2, // 3: envoy.config.overload.v2alpha.OverloadAction.triggers:type_name -> envoy.config.overload.v2alpha.Trigger
	7, // 4: envoy.config.overload.v2alpha.OverloadManager.refresh_interval:type_name -> google.protobuf.Duration
	0, // 5: envoy.config.overload.v2alpha.OverloadManager.resource_monitors:type_name -> envoy.config.overload.v2alpha.ResourceMonitor
	3, // 6: envoy.config.overload.v2alpha.OverloadManager.actions:type_name -> envoy.config.overload.v2alpha.OverloadAction
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_envoy_config_overload_v2alpha_overload_proto_init() }
func file_envoy_config_overload_v2alpha_overload_proto_init() {
	if File_envoy_config_overload_v2alpha_overload_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_overload_v2alpha_overload_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceMonitor); i {
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
		file_envoy_config_overload_v2alpha_overload_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThresholdTrigger); i {
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
		file_envoy_config_overload_v2alpha_overload_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Trigger); i {
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
		file_envoy_config_overload_v2alpha_overload_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OverloadAction); i {
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
		file_envoy_config_overload_v2alpha_overload_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OverloadManager); i {
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
	file_envoy_config_overload_v2alpha_overload_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ResourceMonitor_Config)(nil),
		(*ResourceMonitor_TypedConfig)(nil),
	}
	file_envoy_config_overload_v2alpha_overload_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Trigger_Threshold)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_config_overload_v2alpha_overload_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_overload_v2alpha_overload_proto_goTypes,
		DependencyIndexes: file_envoy_config_overload_v2alpha_overload_proto_depIdxs,
		MessageInfos:      file_envoy_config_overload_v2alpha_overload_proto_msgTypes,
	}.Build()
	File_envoy_config_overload_v2alpha_overload_proto = out.File
	file_envoy_config_overload_v2alpha_overload_proto_rawDesc = nil
	file_envoy_config_overload_v2alpha_overload_proto_goTypes = nil
	file_envoy_config_overload_v2alpha_overload_proto_depIdxs = nil
}
