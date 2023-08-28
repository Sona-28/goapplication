// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: customer/customer.proto

package customer

import (
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

type CustomerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId   int32  `protobuf:"varint,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	CustomerName string `protobuf:"bytes,2,opt,name=customer_name,json=customerName,proto3" json:"customer_name,omitempty"`
	AccountId    int32  `protobuf:"varint,3,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	AccountType  string `protobuf:"bytes,4,opt,name=account_type,json=accountType,proto3" json:"account_type,omitempty"`
}

func (x *CustomerRequest) Reset() {
	*x = CustomerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_customer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerRequest) ProtoMessage() {}

func (x *CustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customer_customer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerRequest.ProtoReflect.Descriptor instead.
func (*CustomerRequest) Descriptor() ([]byte, []int) {
	return file_customer_customer_proto_rawDescGZIP(), []int{0}
}

func (x *CustomerRequest) GetCustomerId() int32 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *CustomerRequest) GetCustomerName() string {
	if x != nil {
		return x.CustomerName
	}
	return ""
}

func (x *CustomerRequest) GetAccountId() int32 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *CustomerRequest) GetAccountType() string {
	if x != nil {
		return x.AccountType
	}
	return ""
}

type CustomerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CustomerResponse) Reset() {
	*x = CustomerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_customer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerResponse) ProtoMessage() {}

func (x *CustomerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customer_customer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerResponse.ProtoReflect.Descriptor instead.
func (*CustomerResponse) Descriptor() ([]byte, []int) {
	return file_customer_customer_proto_rawDescGZIP(), []int{1}
}

func (x *CustomerResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type CustList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cust []*CustomerRequest `protobuf:"bytes,1,rep,name=cust,proto3" json:"cust,omitempty"`
}

func (x *CustList) Reset() {
	*x = CustList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_customer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustList) ProtoMessage() {}

func (x *CustList) ProtoReflect() protoreflect.Message {
	mi := &file_customer_customer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustList.ProtoReflect.Descriptor instead.
func (*CustList) Descriptor() ([]byte, []int) {
	return file_customer_customer_proto_rawDescGZIP(), []int{2}
}

func (x *CustList) GetCust() []*CustomerRequest {
	if x != nil {
		return x.Cust
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_customer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_customer_customer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_customer_customer_proto_rawDescGZIP(), []int{3}
}

var File_customer_customer_proto protoreflect.FileDescriptor

var file_customer_customer_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x22, 0x99, 0x01, 0x0a, 0x0f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22,
	0x2a, 0x0a, 0x10, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x39, 0x0a, 0x08, 0x43,
	0x75, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x04, 0x63, 0x75, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x04, 0x63, 0x75, 0x73, 0x74, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32,
	0x8b, 0x01, 0x0a, 0x0f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x12, 0x19, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x0f, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x31, 0x5a,
	0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x6f, 0x6e, 0x61,
	0x2d, 0x32, 0x38, 0x2f, 0x67, 0x6f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x33, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_customer_customer_proto_rawDescOnce sync.Once
	file_customer_customer_proto_rawDescData = file_customer_customer_proto_rawDesc
)

func file_customer_customer_proto_rawDescGZIP() []byte {
	file_customer_customer_proto_rawDescOnce.Do(func() {
		file_customer_customer_proto_rawDescData = protoimpl.X.CompressGZIP(file_customer_customer_proto_rawDescData)
	})
	return file_customer_customer_proto_rawDescData
}

var file_customer_customer_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_customer_customer_proto_goTypes = []interface{}{
	(*CustomerRequest)(nil),  // 0: customer.CustomerRequest
	(*CustomerResponse)(nil), // 1: customer.CustomerResponse
	(*CustList)(nil),         // 2: customer.CustList
	(*Empty)(nil),            // 3: customer.Empty
}
var file_customer_customer_proto_depIdxs = []int32{
	0, // 0: customer.CustList.cust:type_name -> customer.CustomerRequest
	0, // 1: customer.CustomerService.AddCustomer:input_type -> customer.CustomerRequest
	3, // 2: customer.CustomerService.GetCustomer:input_type -> customer.Empty
	1, // 3: customer.CustomerService.AddCustomer:output_type -> customer.CustomerResponse
	2, // 4: customer.CustomerService.GetCustomer:output_type -> customer.CustList
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_customer_customer_proto_init() }
func file_customer_customer_proto_init() {
	if File_customer_customer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_customer_customer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerRequest); i {
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
		file_customer_customer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerResponse); i {
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
		file_customer_customer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustList); i {
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
		file_customer_customer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_customer_customer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_customer_customer_proto_goTypes,
		DependencyIndexes: file_customer_customer_proto_depIdxs,
		MessageInfos:      file_customer_customer_proto_msgTypes,
	}.Build()
	File_customer_customer_proto = out.File
	file_customer_customer_proto_rawDesc = nil
	file_customer_customer_proto_goTypes = nil
	file_customer_customer_proto_depIdxs = nil
}