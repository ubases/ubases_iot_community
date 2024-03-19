// Code generated by protoc,2022-04-27 17:47:11. DO NOT EDIT.
// versions: v1.0.0
// 如果确实需要修改,请修改后改文件名,可以移除文件名中的.gen,以防下次生成覆盖

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: cms_document_images_model.gen.proto

package protosService

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

//来源于数据表
type CmsDocumentImages struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id" gorm:"default:文档id;comment:文档id;size:19;"
	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: json:"documentId" gorm:"default:文档类型id;comment:文档类型id;size:19;"
	DocumentId int64 `protobuf:"varint,101,opt,name=documentId,proto3" json:"documentId,omitempty"`
	// @inject_tag: json:"imgTitle" gorm:"default:内容描述;comment:内容描述;size:200;"
	ImgTitle string `protobuf:"bytes,102,opt,name=imgTitle,proto3" json:"imgTitle,omitempty"`
	// @inject_tag: json:"imgUrl" gorm:"default:图片地址;comment:图片地址;size:255;"
	ImgUrl string `protobuf:"bytes,103,opt,name=imgUrl,proto3" json:"imgUrl,omitempty"`
}

func (x *CmsDocumentImages) Reset() {
	*x = CmsDocumentImages{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cms_document_images_model_gen_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmsDocumentImages) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmsDocumentImages) ProtoMessage() {}

func (x *CmsDocumentImages) ProtoReflect() protoreflect.Message {
	mi := &file_cms_document_images_model_gen_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmsDocumentImages.ProtoReflect.Descriptor instead.
func (*CmsDocumentImages) Descriptor() ([]byte, []int) {
	return file_cms_document_images_model_gen_proto_rawDescGZIP(), []int{0}
}

func (x *CmsDocumentImages) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CmsDocumentImages) GetDocumentId() int64 {
	if x != nil {
		return x.DocumentId
	}
	return 0
}

func (x *CmsDocumentImages) GetImgTitle() string {
	if x != nil {
		return x.ImgTitle
	}
	return ""
}

func (x *CmsDocumentImages) GetImgUrl() string {
	if x != nil {
		return x.ImgUrl
	}
	return ""
}

type CmsDocumentImagesFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id" gorm:"default:文档id;comment:文档id;size:19;"
	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: json:"documentId" gorm:"default:文档类型id;comment:文档类型id;size:19;"
	DocumentId int64 `protobuf:"varint,101,opt,name=documentId,proto3" json:"documentId,omitempty"`
	// @inject_tag: json:"imgTitle" gorm:"default:内容描述;comment:内容描述;size:200;"
	ImgTitle string `protobuf:"bytes,102,opt,name=imgTitle,proto3" json:"imgTitle,omitempty"`
	// @inject_tag: json:"imgUrl" gorm:"default:图片地址;comment:图片地址;size:255;"
	ImgUrl string `protobuf:"bytes,103,opt,name=imgUrl,proto3" json:"imgUrl,omitempty"`
}

func (x *CmsDocumentImagesFilter) Reset() {
	*x = CmsDocumentImagesFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cms_document_images_model_gen_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmsDocumentImagesFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmsDocumentImagesFilter) ProtoMessage() {}

func (x *CmsDocumentImagesFilter) ProtoReflect() protoreflect.Message {
	mi := &file_cms_document_images_model_gen_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmsDocumentImagesFilter.ProtoReflect.Descriptor instead.
func (*CmsDocumentImagesFilter) Descriptor() ([]byte, []int) {
	return file_cms_document_images_model_gen_proto_rawDescGZIP(), []int{1}
}

func (x *CmsDocumentImagesFilter) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CmsDocumentImagesFilter) GetDocumentId() int64 {
	if x != nil {
		return x.DocumentId
	}
	return 0
}

func (x *CmsDocumentImagesFilter) GetImgTitle() string {
	if x != nil {
		return x.ImgTitle
	}
	return ""
}

func (x *CmsDocumentImagesFilter) GetImgUrl() string {
	if x != nil {
		return x.ImgUrl
	}
	return ""
}

type CmsDocumentImagesListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query     *CmsDocumentImages `protobuf:"bytes,101,opt,name=query,proto3" json:"query,omitempty"`
	Page      int64              `protobuf:"varint,102,opt,name=page,proto3" json:"page,omitempty"`
	PageSize  int64              `protobuf:"varint,103,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	OrderKey  string             `protobuf:"bytes,104,opt,name=orderKey,proto3" json:"orderKey,omitempty"`
	OrderDesc string             `protobuf:"bytes,105,opt,name=orderDesc,proto3" json:"orderDesc,omitempty"`
	SearchKey string             `protobuf:"bytes,106,opt,name=searchKey,proto3" json:"searchKey,omitempty"`
}

func (x *CmsDocumentImagesListRequest) Reset() {
	*x = CmsDocumentImagesListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cms_document_images_model_gen_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmsDocumentImagesListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmsDocumentImagesListRequest) ProtoMessage() {}

func (x *CmsDocumentImagesListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cms_document_images_model_gen_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmsDocumentImagesListRequest.ProtoReflect.Descriptor instead.
func (*CmsDocumentImagesListRequest) Descriptor() ([]byte, []int) {
	return file_cms_document_images_model_gen_proto_rawDescGZIP(), []int{2}
}

func (x *CmsDocumentImagesListRequest) GetQuery() *CmsDocumentImages {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *CmsDocumentImagesListRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *CmsDocumentImagesListRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *CmsDocumentImagesListRequest) GetOrderKey() string {
	if x != nil {
		return x.OrderKey
	}
	return ""
}

func (x *CmsDocumentImagesListRequest) GetOrderDesc() string {
	if x != nil {
		return x.OrderDesc
	}
	return ""
}

func (x *CmsDocumentImagesListRequest) GetSearchKey() string {
	if x != nil {
		return x.SearchKey
	}
	return ""
}

type CmsDocumentImagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32                `protobuf:"varint,101,opt,name=code,proto3" json:"code,omitempty"`
	Message string               `protobuf:"bytes,102,opt,name=message,proto3" json:"message,omitempty"`
	Total   int64                `protobuf:"varint,103,opt,name=total,proto3" json:"total,omitempty"`
	Data    []*CmsDocumentImages `protobuf:"bytes,104,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *CmsDocumentImagesResponse) Reset() {
	*x = CmsDocumentImagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cms_document_images_model_gen_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmsDocumentImagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmsDocumentImagesResponse) ProtoMessage() {}

func (x *CmsDocumentImagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cms_document_images_model_gen_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmsDocumentImagesResponse.ProtoReflect.Descriptor instead.
func (*CmsDocumentImagesResponse) Descriptor() ([]byte, []int) {
	return file_cms_document_images_model_gen_proto_rawDescGZIP(), []int{3}
}

func (x *CmsDocumentImagesResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CmsDocumentImagesResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CmsDocumentImagesResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *CmsDocumentImagesResponse) GetData() []*CmsDocumentImages {
	if x != nil {
		return x.Data
	}
	return nil
}

type CmsDocumentImagesUpdateFieldsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields []string           `protobuf:"bytes,101,rep,name=fields,proto3" json:"fields,omitempty"`
	Data   *CmsDocumentImages `protobuf:"bytes,102,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CmsDocumentImagesUpdateFieldsRequest) Reset() {
	*x = CmsDocumentImagesUpdateFieldsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cms_document_images_model_gen_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmsDocumentImagesUpdateFieldsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmsDocumentImagesUpdateFieldsRequest) ProtoMessage() {}

func (x *CmsDocumentImagesUpdateFieldsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cms_document_images_model_gen_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmsDocumentImagesUpdateFieldsRequest.ProtoReflect.Descriptor instead.
func (*CmsDocumentImagesUpdateFieldsRequest) Descriptor() ([]byte, []int) {
	return file_cms_document_images_model_gen_proto_rawDescGZIP(), []int{4}
}

func (x *CmsDocumentImagesUpdateFieldsRequest) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *CmsDocumentImagesUpdateFieldsRequest) GetData() *CmsDocumentImages {
	if x != nil {
		return x.Data
	}
	return nil
}

// 表模型主键，只能是整数或字符串，若生成错误，请检查数据库表主键是否正确。支持组合主键
type CmsDocumentImagesPrimarykey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,100,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CmsDocumentImagesPrimarykey) Reset() {
	*x = CmsDocumentImagesPrimarykey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cms_document_images_model_gen_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmsDocumentImagesPrimarykey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmsDocumentImagesPrimarykey) ProtoMessage() {}

func (x *CmsDocumentImagesPrimarykey) ProtoReflect() protoreflect.Message {
	mi := &file_cms_document_images_model_gen_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmsDocumentImagesPrimarykey.ProtoReflect.Descriptor instead.
func (*CmsDocumentImagesPrimarykey) Descriptor() ([]byte, []int) {
	return file_cms_document_images_model_gen_proto_rawDescGZIP(), []int{5}
}

func (x *CmsDocumentImagesPrimarykey) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CmsDocumentImagesBatchDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys []*CmsDocumentImagesPrimarykey `protobuf:"bytes,101,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *CmsDocumentImagesBatchDeleteRequest) Reset() {
	*x = CmsDocumentImagesBatchDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cms_document_images_model_gen_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmsDocumentImagesBatchDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmsDocumentImagesBatchDeleteRequest) ProtoMessage() {}

func (x *CmsDocumentImagesBatchDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cms_document_images_model_gen_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmsDocumentImagesBatchDeleteRequest.ProtoReflect.Descriptor instead.
func (*CmsDocumentImagesBatchDeleteRequest) Descriptor() ([]byte, []int) {
	return file_cms_document_images_model_gen_proto_rawDescGZIP(), []int{6}
}

func (x *CmsDocumentImagesBatchDeleteRequest) GetKeys() []*CmsDocumentImagesPrimarykey {
	if x != nil {
		return x.Keys
	}
	return nil
}

var File_cms_document_images_model_gen_proto protoreflect.FileDescriptor

var file_cms_document_images_model_gen_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x6d, 0x73, 0x5f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x67, 0x65, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x77,
	0x0a, 0x11, 0x43, 0x6d, 0x73, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x18, 0x65, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6d, 0x67, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x67, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x69, 0x6d, 0x67, 0x55, 0x72, 0x6c, 0x18, 0x67, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x69, 0x6d, 0x67, 0x55, 0x72, 0x6c, 0x22, 0x7d, 0x0a, 0x17, 0x43, 0x6d, 0x73, 0x44, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x18, 0x65, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6d, 0x67, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x66,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x67, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x69, 0x6d, 0x67, 0x55, 0x72, 0x6c, 0x18, 0x67, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x69, 0x6d, 0x67, 0x55, 0x72, 0x6c, 0x22, 0xd8, 0x01, 0x0a, 0x1c, 0x43, 0x6d, 0x73, 0x44, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x43, 0x6d, 0x73, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x73, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x67, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x4b, 0x65, 0x79, 0x18, 0x68, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65,
	0x73, 0x63, 0x18, 0x69, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44,
	0x65, 0x73, 0x63, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4b, 0x65, 0x79,
	0x18, 0x6a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4b, 0x65,
	0x79, 0x22, 0x8f, 0x01, 0x0a, 0x19, 0x43, 0x6d, 0x73, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x66,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x67, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x68, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6d, 0x73, 0x44,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x6e, 0x0a, 0x24, 0x43, 0x6d, 0x73, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x65, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x66, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6d, 0x73, 0x44,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x2d, 0x0a, 0x1b, 0x43, 0x6d, 0x73, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x6b,
	0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x5f, 0x0a, 0x23, 0x43, 0x6d, 0x73, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x04, 0x6b, 0x65, 0x79,
	0x73, 0x18, 0x65, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x43, 0x6d, 0x73, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x73, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x6b, 0x65, 0x79, 0x52, 0x04, 0x6b,
	0x65, 0x79, 0x73, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cms_document_images_model_gen_proto_rawDescOnce sync.Once
	file_cms_document_images_model_gen_proto_rawDescData = file_cms_document_images_model_gen_proto_rawDesc
)

func file_cms_document_images_model_gen_proto_rawDescGZIP() []byte {
	file_cms_document_images_model_gen_proto_rawDescOnce.Do(func() {
		file_cms_document_images_model_gen_proto_rawDescData = protoimpl.X.CompressGZIP(file_cms_document_images_model_gen_proto_rawDescData)
	})
	return file_cms_document_images_model_gen_proto_rawDescData
}

var file_cms_document_images_model_gen_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_cms_document_images_model_gen_proto_goTypes = []interface{}{
	(*CmsDocumentImages)(nil),                    // 0: service.CmsDocumentImages
	(*CmsDocumentImagesFilter)(nil),              // 1: service.CmsDocumentImagesFilter
	(*CmsDocumentImagesListRequest)(nil),         // 2: service.CmsDocumentImagesListRequest
	(*CmsDocumentImagesResponse)(nil),            // 3: service.CmsDocumentImagesResponse
	(*CmsDocumentImagesUpdateFieldsRequest)(nil), // 4: service.CmsDocumentImagesUpdateFieldsRequest
	(*CmsDocumentImagesPrimarykey)(nil),          // 5: service.CmsDocumentImagesPrimarykey
	(*CmsDocumentImagesBatchDeleteRequest)(nil),  // 6: service.CmsDocumentImagesBatchDeleteRequest
}
var file_cms_document_images_model_gen_proto_depIdxs = []int32{
	0, // 0: service.CmsDocumentImagesListRequest.query:type_name -> service.CmsDocumentImages
	0, // 1: service.CmsDocumentImagesResponse.data:type_name -> service.CmsDocumentImages
	0, // 2: service.CmsDocumentImagesUpdateFieldsRequest.data:type_name -> service.CmsDocumentImages
	5, // 3: service.CmsDocumentImagesBatchDeleteRequest.keys:type_name -> service.CmsDocumentImagesPrimarykey
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_cms_document_images_model_gen_proto_init() }
func file_cms_document_images_model_gen_proto_init() {
	if File_cms_document_images_model_gen_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cms_document_images_model_gen_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmsDocumentImages); i {
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
		file_cms_document_images_model_gen_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmsDocumentImagesFilter); i {
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
		file_cms_document_images_model_gen_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmsDocumentImagesListRequest); i {
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
		file_cms_document_images_model_gen_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmsDocumentImagesResponse); i {
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
		file_cms_document_images_model_gen_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmsDocumentImagesUpdateFieldsRequest); i {
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
		file_cms_document_images_model_gen_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmsDocumentImagesPrimarykey); i {
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
		file_cms_document_images_model_gen_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmsDocumentImagesBatchDeleteRequest); i {
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
			RawDescriptor: file_cms_document_images_model_gen_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cms_document_images_model_gen_proto_goTypes,
		DependencyIndexes: file_cms_document_images_model_gen_proto_depIdxs,
		MessageInfos:      file_cms_document_images_model_gen_proto_msgTypes,
	}.Build()
	File_cms_document_images_model_gen_proto = out.File
	file_cms_document_images_model_gen_proto_rawDesc = nil
	file_cms_document_images_model_gen_proto_goTypes = nil
	file_cms_document_images_model_gen_proto_depIdxs = nil
}
