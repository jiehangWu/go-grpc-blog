// Code generated by protoc-gen-go. DO NOT EDIT.
// source: blog.proto

package genproto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Article struct {
	ArticleId            int64                `protobuf:"varint,1,opt,name=article_id,json=articleId,proto3" json:"article_id,omitempty"`
	ArticleName          string               `protobuf:"bytes,2,opt,name=article_name,json=articleName,proto3" json:"article_name,omitempty"`
	ArticleContent       string               `protobuf:"bytes,3,opt,name=article_content,json=articleContent,proto3" json:"article_content,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Article) Reset()         { *m = Article{} }
func (m *Article) String() string { return proto.CompactTextString(m) }
func (*Article) ProtoMessage()    {}
func (*Article) Descriptor() ([]byte, []int) {
	return fileDescriptor_6745b25902462fb1, []int{0}
}

func (m *Article) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Article.Unmarshal(m, b)
}
func (m *Article) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Article.Marshal(b, m, deterministic)
}
func (m *Article) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Article.Merge(m, src)
}
func (m *Article) XXX_Size() int {
	return xxx_messageInfo_Article.Size(m)
}
func (m *Article) XXX_DiscardUnknown() {
	xxx_messageInfo_Article.DiscardUnknown(m)
}

var xxx_messageInfo_Article proto.InternalMessageInfo

func (m *Article) GetArticleId() int64 {
	if m != nil {
		return m.ArticleId
	}
	return 0
}

func (m *Article) GetArticleName() string {
	if m != nil {
		return m.ArticleName
	}
	return ""
}

func (m *Article) GetArticleContent() string {
	if m != nil {
		return m.ArticleContent
	}
	return ""
}

func (m *Article) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func init() {
	proto.RegisterType((*Article)(nil), "blog.Article")
}

func init() {
	proto.RegisterFile("blog.proto", fileDescriptor_6745b25902462fb1)
}

var fileDescriptor_6745b25902462fb1 = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xca, 0xc9, 0x4f,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0xa5, 0xe4, 0xd3, 0xf3, 0xf3, 0xd3,
	0x73, 0x52, 0xf5, 0xc1, 0x62, 0x49, 0xa5, 0x69, 0xfa, 0x25, 0x99, 0xb9, 0xa9, 0xc5, 0x25, 0x89,
	0xb9, 0x05, 0x10, 0x65, 0x4a, 0xeb, 0x18, 0xb9, 0xd8, 0x1d, 0x8b, 0x4a, 0x32, 0x93, 0x73, 0x52,
	0x85, 0x64, 0xb9, 0xb8, 0x12, 0x21, 0xcc, 0xf8, 0xcc, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xe6,
	0x20, 0x4e, 0xa8, 0x88, 0x67, 0x8a, 0x90, 0x22, 0x17, 0x0f, 0x4c, 0x3a, 0x2f, 0x31, 0x37, 0x55,
	0x82, 0x49, 0x81, 0x51, 0x83, 0x33, 0x88, 0x1b, 0x2a, 0xe6, 0x97, 0x98, 0x9b, 0x2a, 0xa4, 0xce,
	0xc5, 0x0f, 0x53, 0x92, 0x9c, 0x9f, 0x57, 0x92, 0x9a, 0x57, 0x22, 0xc1, 0x0c, 0x56, 0xc5, 0x07,
	0x15, 0x76, 0x86, 0x88, 0x0a, 0x59, 0x70, 0x71, 0xc2, 0x5d, 0x22, 0xc1, 0xa2, 0xc0, 0xa8, 0xc1,
	0x6d, 0x24, 0xa5, 0x07, 0x71, 0xab, 0x1e, 0xcc, 0xad, 0x7a, 0x21, 0x30, 0x15, 0x41, 0x08, 0xc5,
	0x4e, 0xdc, 0x51, 0x9c, 0xfa, 0xe9, 0xa9, 0x79, 0x10, 0x45, 0x6c, 0x60, 0xca, 0x18, 0x10, 0x00,
	0x00, 0xff, 0xff, 0xed, 0x8f, 0x08, 0x0f, 0xf9, 0x00, 0x00, 0x00,
}
