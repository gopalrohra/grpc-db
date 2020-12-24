// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.1
// source: grpc_database.proto

package grpc_db

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Database struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dbname string `protobuf:"bytes,1,opt,name=dbname,proto3" json:"dbname,omitempty"`
	Dbtype string `protobuf:"bytes,2,opt,name=dbtype,proto3" json:"dbtype,omitempty"`
}

func (x *Database) Reset() {
	*x = Database{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_database_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Database) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Database) ProtoMessage() {}

func (x *Database) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_database_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Database.ProtoReflect.Descriptor instead.
func (*Database) Descriptor() ([]byte, []int) {
	return file_grpc_database_proto_rawDescGZIP(), []int{0}
}

func (x *Database) GetDbname() string {
	if x != nil {
		return x.Dbname
	}
	return ""
}

func (x *Database) GetDbtype() string {
	if x != nil {
		return x.Dbtype
	}
	return ""
}

type DatabaseInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User     string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Dbname   string `protobuf:"bytes,3,opt,name=dbname,proto3" json:"dbname,omitempty"`
}

func (x *DatabaseInfo) Reset() {
	*x = DatabaseInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_database_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatabaseInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatabaseInfo) ProtoMessage() {}

func (x *DatabaseInfo) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_database_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatabaseInfo.ProtoReflect.Descriptor instead.
func (*DatabaseInfo) Descriptor() ([]byte, []int) {
	return file_grpc_database_proto_rawDescGZIP(), []int{1}
}

func (x *DatabaseInfo) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *DatabaseInfo) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *DatabaseInfo) GetDbname() string {
	if x != nil {
		return x.Dbname
	}
	return ""
}

type DatabaseResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status      string        `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Description string        `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Info        *DatabaseInfo `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *DatabaseResult) Reset() {
	*x = DatabaseResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_database_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatabaseResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatabaseResult) ProtoMessage() {}

func (x *DatabaseResult) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_database_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatabaseResult.ProtoReflect.Descriptor instead.
func (*DatabaseResult) Descriptor() ([]byte, []int) {
	return file_grpc_database_proto_rawDescGZIP(), []int{2}
}

func (x *DatabaseResult) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *DatabaseResult) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *DatabaseResult) GetInfo() *DatabaseInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type SelectQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info      *DatabaseInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	TableName string        `protobuf:"bytes,2,opt,name=tableName,proto3" json:"tableName,omitempty"`
	Fields    []string      `protobuf:"bytes,3,rep,name=fields,proto3" json:"fields,omitempty"`
	Clauses   []string      `protobuf:"bytes,4,rep,name=clauses,proto3" json:"clauses,omitempty"`
}

func (x *SelectQuery) Reset() {
	*x = SelectQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_database_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectQuery) ProtoMessage() {}

func (x *SelectQuery) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_database_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectQuery.ProtoReflect.Descriptor instead.
func (*SelectQuery) Descriptor() ([]byte, []int) {
	return file_grpc_database_proto_rawDescGZIP(), []int{3}
}

func (x *SelectQuery) GetInfo() *DatabaseInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *SelectQuery) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

func (x *SelectQuery) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *SelectQuery) GetClauses() []string {
	if x != nil {
		return x.Clauses
	}
	return nil
}

type SelectQueryResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status      string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Description string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Records     *Records `protobuf:"bytes,3,opt,name=records,proto3" json:"records,omitempty"`
}

func (x *SelectQueryResult) Reset() {
	*x = SelectQueryResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_database_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectQueryResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectQueryResult) ProtoMessage() {}

func (x *SelectQueryResult) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_database_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectQueryResult.ProtoReflect.Descriptor instead.
func (*SelectQueryResult) Descriptor() ([]byte, []int) {
	return file_grpc_database_proto_rawDescGZIP(), []int{4}
}

func (x *SelectQueryResult) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *SelectQueryResult) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SelectQueryResult) GetRecords() *Records {
	if x != nil {
		return x.Records
	}
	return nil
}

type Records struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rows []*Row `protobuf:"bytes,1,rep,name=rows,proto3" json:"rows,omitempty"`
}

func (x *Records) Reset() {
	*x = Records{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_database_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Records) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Records) ProtoMessage() {}

func (x *Records) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_database_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Records.ProtoReflect.Descriptor instead.
func (*Records) Descriptor() ([]byte, []int) {
	return file_grpc_database_proto_rawDescGZIP(), []int{5}
}

func (x *Records) GetRows() []*Row {
	if x != nil {
		return x.Rows
	}
	return nil
}

type Row struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Row) Reset() {
	*x = Row{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_database_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Row) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Row) ProtoMessage() {}

func (x *Row) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_database_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Row.ProtoReflect.Descriptor instead.
func (*Row) Descriptor() ([]byte, []int) {
	return file_grpc_database_proto_rawDescGZIP(), []int{6}
}

type TableRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info      *DatabaseInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Name      string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ColumnDef []string      `protobuf:"bytes,3,rep,name=columnDef,proto3" json:"columnDef,omitempty"`
}

func (x *TableRequest) Reset() {
	*x = TableRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_database_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TableRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableRequest) ProtoMessage() {}

func (x *TableRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_database_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableRequest.ProtoReflect.Descriptor instead.
func (*TableRequest) Descriptor() ([]byte, []int) {
	return file_grpc_database_proto_rawDescGZIP(), []int{7}
}

func (x *TableRequest) GetInfo() *DatabaseInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *TableRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TableRequest) GetColumnDef() []string {
	if x != nil {
		return x.ColumnDef
	}
	return nil
}

type TableResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status      string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *TableResponse) Reset() {
	*x = TableResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_database_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TableResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableResponse) ProtoMessage() {}

func (x *TableResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_database_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableResponse.ProtoReflect.Descriptor instead.
func (*TableResponse) Descriptor() ([]byte, []int) {
	return file_grpc_database_proto_rawDescGZIP(), []int{8}
}

func (x *TableResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *TableResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_grpc_database_proto protoreflect.FileDescriptor

var file_grpc_database_proto_rawDesc = []byte{
	0x0a, 0x13, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x22, 0x3a, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x64, 0x62, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x64, 0x62, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x62, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x62, 0x74, 0x79, 0x70, 0x65,
	0x22, 0x56, 0x0a, 0x0c, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x64, 0x62, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x64, 0x62, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x7b, 0x0a, 0x0e, 0x44, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x8e, 0x01, 0x0a, 0x0b, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x2f, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6c, 0x61, 0x75, 0x73, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6c, 0x61, 0x75, 0x73, 0x65, 0x73, 0x22, 0x7f, 0x0a, 0x11, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x07,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x22, 0x31, 0x0a, 0x07, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x73, 0x12, 0x26, 0x0a, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x52, 0x6f, 0x77, 0x52, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x22, 0x05, 0x0a, 0x03, 0x52, 0x6f,
	0x77, 0x22, 0x71, 0x0a, 0x0c, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2f, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e,
	0x44, 0x65, 0x66, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6c, 0x75, 0x6d,
	0x6e, 0x44, 0x65, 0x66, 0x22, 0x49, 0x0a, 0x0d, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x32,
	0xf7, 0x01, 0x0a, 0x0c, 0x47, 0x52, 0x50, 0x43, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x12, 0x4a, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x12, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x1a, 0x1d, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0b,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1b, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0d, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x12, 0x1a, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x20, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42, 0x1f, 0x5a, 0x1d, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x70, 0x61, 0x6c, 0x72, 0x6f, 0x68,
	0x72, 0x61, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x64, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_grpc_database_proto_rawDescOnce sync.Once
	file_grpc_database_proto_rawDescData = file_grpc_database_proto_rawDesc
)

func file_grpc_database_proto_rawDescGZIP() []byte {
	file_grpc_database_proto_rawDescOnce.Do(func() {
		file_grpc_database_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_database_proto_rawDescData)
	})
	return file_grpc_database_proto_rawDescData
}

var file_grpc_database_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_grpc_database_proto_goTypes = []interface{}{
	(*Database)(nil),          // 0: grpc_database.Database
	(*DatabaseInfo)(nil),      // 1: grpc_database.DatabaseInfo
	(*DatabaseResult)(nil),    // 2: grpc_database.DatabaseResult
	(*SelectQuery)(nil),       // 3: grpc_database.SelectQuery
	(*SelectQueryResult)(nil), // 4: grpc_database.SelectQueryResult
	(*Records)(nil),           // 5: grpc_database.Records
	(*Row)(nil),               // 6: grpc_database.Row
	(*TableRequest)(nil),      // 7: grpc_database.TableRequest
	(*TableResponse)(nil),     // 8: grpc_database.TableResponse
}
var file_grpc_database_proto_depIdxs = []int32{
	1, // 0: grpc_database.DatabaseResult.info:type_name -> grpc_database.DatabaseInfo
	1, // 1: grpc_database.SelectQuery.info:type_name -> grpc_database.DatabaseInfo
	5, // 2: grpc_database.SelectQueryResult.records:type_name -> grpc_database.Records
	6, // 3: grpc_database.Records.rows:type_name -> grpc_database.Row
	1, // 4: grpc_database.TableRequest.info:type_name -> grpc_database.DatabaseInfo
	0, // 5: grpc_database.GRPCDatabase.CreateDatabase:input_type -> grpc_database.Database
	7, // 6: grpc_database.GRPCDatabase.CreateTable:input_type -> grpc_database.TableRequest
	3, // 7: grpc_database.GRPCDatabase.ExecuteSelect:input_type -> grpc_database.SelectQuery
	2, // 8: grpc_database.GRPCDatabase.CreateDatabase:output_type -> grpc_database.DatabaseResult
	8, // 9: grpc_database.GRPCDatabase.CreateTable:output_type -> grpc_database.TableResponse
	4, // 10: grpc_database.GRPCDatabase.ExecuteSelect:output_type -> grpc_database.SelectQueryResult
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_grpc_database_proto_init() }
func file_grpc_database_proto_init() {
	if File_grpc_database_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_database_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Database); i {
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
		file_grpc_database_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatabaseInfo); i {
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
		file_grpc_database_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatabaseResult); i {
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
		file_grpc_database_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelectQuery); i {
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
		file_grpc_database_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SelectQueryResult); i {
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
		file_grpc_database_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Records); i {
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
		file_grpc_database_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Row); i {
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
		file_grpc_database_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TableRequest); i {
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
		file_grpc_database_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TableResponse); i {
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
			RawDescriptor: file_grpc_database_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_database_proto_goTypes,
		DependencyIndexes: file_grpc_database_proto_depIdxs,
		MessageInfos:      file_grpc_database_proto_msgTypes,
	}.Build()
	File_grpc_database_proto = out.File
	file_grpc_database_proto_rawDesc = nil
	file_grpc_database_proto_goTypes = nil
	file_grpc_database_proto_depIdxs = nil
}
