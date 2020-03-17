// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/campaign_label.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Represents a relationship between a campaign and a label.
type CampaignLabel struct {
	// Name of the resource.
	// Campaign label resource names have the form:
	// `customers/{customer_id}/campaignLabels/{campaign_id}~{label_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The campaign to which the label is attached.
	Campaign *wrappers.StringValue `protobuf:"bytes,2,opt,name=campaign,proto3" json:"campaign,omitempty"`
	// The label assigned to the campaign.
	Label                *wrappers.StringValue `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CampaignLabel) Reset()         { *m = CampaignLabel{} }
func (m *CampaignLabel) String() string { return proto.CompactTextString(m) }
func (*CampaignLabel) ProtoMessage()    {}
func (*CampaignLabel) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1b047812890bf04, []int{0}
}

func (m *CampaignLabel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignLabel.Unmarshal(m, b)
}
func (m *CampaignLabel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignLabel.Marshal(b, m, deterministic)
}
func (m *CampaignLabel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignLabel.Merge(m, src)
}
func (m *CampaignLabel) XXX_Size() int {
	return xxx_messageInfo_CampaignLabel.Size(m)
}
func (m *CampaignLabel) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignLabel.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignLabel proto.InternalMessageInfo

func (m *CampaignLabel) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CampaignLabel) GetCampaign() *wrappers.StringValue {
	if m != nil {
		return m.Campaign
	}
	return nil
}

func (m *CampaignLabel) GetLabel() *wrappers.StringValue {
	if m != nil {
		return m.Label
	}
	return nil
}

func init() {
	proto.RegisterType((*CampaignLabel)(nil), "google.ads.googleads.v1.resources.CampaignLabel")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/campaign_label.proto", fileDescriptor_a1b047812890bf04)
}

var fileDescriptor_a1b047812890bf04 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4a, 0xc3, 0x30,
	0x1c, 0xc6, 0x69, 0x87, 0xa2, 0xd5, 0x5d, 0x7a, 0x1a, 0x63, 0xc8, 0xa6, 0x0c, 0x76, 0x4a, 0xe8,
	0x04, 0x91, 0x78, 0xea, 0x3c, 0x0c, 0x44, 0x64, 0x4c, 0xe8, 0x41, 0x0a, 0xe3, 0xbf, 0x35, 0x86,
	0x42, 0x9b, 0x84, 0xa4, 0x9d, 0xef, 0xb3, 0xa3, 0x8f, 0xe2, 0xa3, 0xf8, 0x12, 0x4a, 0x9b, 0x26,
	0xe0, 0x45, 0xbd, 0x7d, 0xed, 0xff, 0xf7, 0x7d, 0xdf, 0x3f, 0x49, 0x70, 0xc3, 0x84, 0x60, 0x05,
	0xc5, 0x90, 0x69, 0x6c, 0x64, 0xa3, 0xf6, 0x11, 0x56, 0x54, 0x8b, 0x5a, 0xed, 0xa8, 0xc6, 0x3b,
	0x28, 0x25, 0xe4, 0x8c, 0x6f, 0x0a, 0xd8, 0xd2, 0x02, 0x49, 0x25, 0x2a, 0x11, 0x4e, 0x0c, 0x8c,
	0x20, 0xd3, 0xc8, 0xf9, 0xd0, 0x3e, 0x42, 0xce, 0x37, 0x1c, 0xd9, 0x68, 0x99, 0x63, 0xe0, 0x5c,
	0x54, 0x50, 0xe5, 0x82, 0x6b, 0x13, 0x30, 0xbc, 0xe8, 0xa6, 0xed, 0xd7, 0xb6, 0x7e, 0xc5, 0x6f,
	0x0a, 0xa4, 0xa4, 0xaa, 0x9b, 0x5f, 0x1e, 0xbc, 0xa0, 0x7f, 0xdf, 0x35, 0x3f, 0x36, 0xc5, 0xe1,
	0x55, 0xd0, 0xb7, 0xe1, 0x1b, 0x0e, 0x25, 0x1d, 0x78, 0x63, 0x6f, 0x76, 0xba, 0x3e, 0xb7, 0x3f,
	0x9f, 0xa0, 0xa4, 0xe1, 0x6d, 0x70, 0x62, 0xf7, 0x1d, 0xf8, 0x63, 0x6f, 0x76, 0x36, 0x1f, 0x75,
	0xfb, 0x21, 0xdb, 0x84, 0x9e, 0x2b, 0x95, 0x73, 0x96, 0x40, 0x51, 0xd3, 0xb5, 0xa3, 0xc3, 0x79,
	0x70, 0xd4, 0x1e, 0x70, 0xd0, 0xfb, 0x87, 0xcd, 0xa0, 0x8b, 0x2f, 0x2f, 0x98, 0xee, 0x44, 0x89,
	0xfe, 0xbc, 0x8c, 0x45, 0xf8, 0xe3, 0x2c, 0xab, 0x26, 0x73, 0xe5, 0xbd, 0x3c, 0x74, 0x46, 0x26,
	0x0a, 0xe0, 0x0c, 0x09, 0xc5, 0x30, 0xa3, 0xbc, 0x6d, 0xb4, 0xaf, 0x21, 0x73, 0xfd, 0xcb, 0xe3,
	0xdc, 0x39, 0x75, 0xf0, 0x7b, 0xcb, 0x38, 0x7e, 0xf7, 0x27, 0x4b, 0x13, 0x19, 0x67, 0x1a, 0x19,
	0xd9, 0xa8, 0x24, 0x42, 0x6b, 0x4b, 0x7e, 0x58, 0x26, 0x8d, 0x33, 0x9d, 0x3a, 0x26, 0x4d, 0xa2,
	0xd4, 0x31, 0x9f, 0xfe, 0xd4, 0x0c, 0x08, 0x89, 0x33, 0x4d, 0x88, 0xa3, 0x08, 0x49, 0x22, 0x42,
	0x1c, 0xb7, 0x3d, 0x6e, 0x97, 0xbd, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x53, 0xed, 0x07, 0xb0,
	0x48, 0x02, 0x00, 0x00,
}