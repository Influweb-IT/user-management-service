// Code generated by protoc-gen-go. DO NOT EDIT.
// source: study-service-api.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type NewStudyRequest struct {
	Token                *TokenInfos `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Study                *Study      `protobuf:"bytes,2,opt,name=study,proto3" json:"study,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *NewStudyRequest) Reset()         { *m = NewStudyRequest{} }
func (m *NewStudyRequest) String() string { return proto.CompactTextString(m) }
func (*NewStudyRequest) ProtoMessage()    {}
func (*NewStudyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_81f0d8f98f9be15c, []int{0}
}

func (m *NewStudyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewStudyRequest.Unmarshal(m, b)
}
func (m *NewStudyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewStudyRequest.Marshal(b, m, deterministic)
}
func (m *NewStudyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewStudyRequest.Merge(m, src)
}
func (m *NewStudyRequest) XXX_Size() int {
	return xxx_messageInfo_NewStudyRequest.Size(m)
}
func (m *NewStudyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NewStudyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NewStudyRequest proto.InternalMessageInfo

func (m *NewStudyRequest) GetToken() *TokenInfos {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *NewStudyRequest) GetStudy() *Study {
	if m != nil {
		return m.Study
	}
	return nil
}

type SurveyAndContext struct {
	Survey               *Survey         `protobuf:"bytes,1,opt,name=survey,proto3" json:"survey,omitempty"`
	Context              *SurveyContext  `protobuf:"bytes,2,opt,name=context,proto3" json:"context,omitempty"`
	Prefill              *SurveyResponse `protobuf:"bytes,3,opt,name=prefill,proto3" json:"prefill,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *SurveyAndContext) Reset()         { *m = SurveyAndContext{} }
func (m *SurveyAndContext) String() string { return proto.CompactTextString(m) }
func (*SurveyAndContext) ProtoMessage()    {}
func (*SurveyAndContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_81f0d8f98f9be15c, []int{1}
}

func (m *SurveyAndContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SurveyAndContext.Unmarshal(m, b)
}
func (m *SurveyAndContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SurveyAndContext.Marshal(b, m, deterministic)
}
func (m *SurveyAndContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SurveyAndContext.Merge(m, src)
}
func (m *SurveyAndContext) XXX_Size() int {
	return xxx_messageInfo_SurveyAndContext.Size(m)
}
func (m *SurveyAndContext) XXX_DiscardUnknown() {
	xxx_messageInfo_SurveyAndContext.DiscardUnknown(m)
}

var xxx_messageInfo_SurveyAndContext proto.InternalMessageInfo

func (m *SurveyAndContext) GetSurvey() *Survey {
	if m != nil {
		return m.Survey
	}
	return nil
}

func (m *SurveyAndContext) GetContext() *SurveyContext {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *SurveyAndContext) GetPrefill() *SurveyResponse {
	if m != nil {
		return m.Prefill
	}
	return nil
}

type StudyReferenceReq struct {
	Token                *TokenInfos `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	StudyKey             string      `protobuf:"bytes,2,opt,name=study_key,json=studyKey,proto3" json:"study_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *StudyReferenceReq) Reset()         { *m = StudyReferenceReq{} }
func (m *StudyReferenceReq) String() string { return proto.CompactTextString(m) }
func (*StudyReferenceReq) ProtoMessage()    {}
func (*StudyReferenceReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_81f0d8f98f9be15c, []int{2}
}

func (m *StudyReferenceReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StudyReferenceReq.Unmarshal(m, b)
}
func (m *StudyReferenceReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StudyReferenceReq.Marshal(b, m, deterministic)
}
func (m *StudyReferenceReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StudyReferenceReq.Merge(m, src)
}
func (m *StudyReferenceReq) XXX_Size() int {
	return xxx_messageInfo_StudyReferenceReq.Size(m)
}
func (m *StudyReferenceReq) XXX_DiscardUnknown() {
	xxx_messageInfo_StudyReferenceReq.DiscardUnknown(m)
}

var xxx_messageInfo_StudyReferenceReq proto.InternalMessageInfo

func (m *StudyReferenceReq) GetToken() *TokenInfos {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *StudyReferenceReq) GetStudyKey() string {
	if m != nil {
		return m.StudyKey
	}
	return ""
}

type SurveyInfoResp struct {
	Infos                []*SurveyInfoResp_SurveyInfo `protobuf:"bytes,1,rep,name=infos,proto3" json:"infos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *SurveyInfoResp) Reset()         { *m = SurveyInfoResp{} }
func (m *SurveyInfoResp) String() string { return proto.CompactTextString(m) }
func (*SurveyInfoResp) ProtoMessage()    {}
func (*SurveyInfoResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_81f0d8f98f9be15c, []int{3}
}

func (m *SurveyInfoResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SurveyInfoResp.Unmarshal(m, b)
}
func (m *SurveyInfoResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SurveyInfoResp.Marshal(b, m, deterministic)
}
func (m *SurveyInfoResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SurveyInfoResp.Merge(m, src)
}
func (m *SurveyInfoResp) XXX_Size() int {
	return xxx_messageInfo_SurveyInfoResp.Size(m)
}
func (m *SurveyInfoResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SurveyInfoResp.DiscardUnknown(m)
}

var xxx_messageInfo_SurveyInfoResp proto.InternalMessageInfo

func (m *SurveyInfoResp) GetInfos() []*SurveyInfoResp_SurveyInfo {
	if m != nil {
		return m.Infos
	}
	return nil
}

type SurveyInfoResp_SurveyInfo struct {
	Key                  string             `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Name                 []*LocalisedObject `protobuf:"bytes,2,rep,name=name,proto3" json:"name,omitempty"`
	Description          []*LocalisedObject `protobuf:"bytes,3,rep,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *SurveyInfoResp_SurveyInfo) Reset()         { *m = SurveyInfoResp_SurveyInfo{} }
func (m *SurveyInfoResp_SurveyInfo) String() string { return proto.CompactTextString(m) }
func (*SurveyInfoResp_SurveyInfo) ProtoMessage()    {}
func (*SurveyInfoResp_SurveyInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_81f0d8f98f9be15c, []int{3, 0}
}

func (m *SurveyInfoResp_SurveyInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SurveyInfoResp_SurveyInfo.Unmarshal(m, b)
}
func (m *SurveyInfoResp_SurveyInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SurveyInfoResp_SurveyInfo.Marshal(b, m, deterministic)
}
func (m *SurveyInfoResp_SurveyInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SurveyInfoResp_SurveyInfo.Merge(m, src)
}
func (m *SurveyInfoResp_SurveyInfo) XXX_Size() int {
	return xxx_messageInfo_SurveyInfoResp_SurveyInfo.Size(m)
}
func (m *SurveyInfoResp_SurveyInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SurveyInfoResp_SurveyInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SurveyInfoResp_SurveyInfo proto.InternalMessageInfo

func (m *SurveyInfoResp_SurveyInfo) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SurveyInfoResp_SurveyInfo) GetName() []*LocalisedObject {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *SurveyInfoResp_SurveyInfo) GetDescription() []*LocalisedObject {
	if m != nil {
		return m.Description
	}
	return nil
}

type AddSurveyReq struct {
	Token                *TokenInfos `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	StudyKey             string      `protobuf:"bytes,2,opt,name=study_key,json=studyKey,proto3" json:"study_key,omitempty"`
	Survey               *Survey     `protobuf:"bytes,3,opt,name=survey,proto3" json:"survey,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AddSurveyReq) Reset()         { *m = AddSurveyReq{} }
func (m *AddSurveyReq) String() string { return proto.CompactTextString(m) }
func (*AddSurveyReq) ProtoMessage()    {}
func (*AddSurveyReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_81f0d8f98f9be15c, []int{4}
}

func (m *AddSurveyReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddSurveyReq.Unmarshal(m, b)
}
func (m *AddSurveyReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddSurveyReq.Marshal(b, m, deterministic)
}
func (m *AddSurveyReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddSurveyReq.Merge(m, src)
}
func (m *AddSurveyReq) XXX_Size() int {
	return xxx_messageInfo_AddSurveyReq.Size(m)
}
func (m *AddSurveyReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AddSurveyReq.DiscardUnknown(m)
}

var xxx_messageInfo_AddSurveyReq proto.InternalMessageInfo

func (m *AddSurveyReq) GetToken() *TokenInfos {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *AddSurveyReq) GetStudyKey() string {
	if m != nil {
		return m.StudyKey
	}
	return ""
}

func (m *AddSurveyReq) GetSurvey() *Survey {
	if m != nil {
		return m.Survey
	}
	return nil
}

type SubmitResponseReq struct {
	Token                *TokenInfos     `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	StudyKey             string          `protobuf:"bytes,2,opt,name=study_key,json=studyKey,proto3" json:"study_key,omitempty"`
	Response             *SurveyResponse `protobuf:"bytes,3,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *SubmitResponseReq) Reset()         { *m = SubmitResponseReq{} }
func (m *SubmitResponseReq) String() string { return proto.CompactTextString(m) }
func (*SubmitResponseReq) ProtoMessage()    {}
func (*SubmitResponseReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_81f0d8f98f9be15c, []int{5}
}

func (m *SubmitResponseReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubmitResponseReq.Unmarshal(m, b)
}
func (m *SubmitResponseReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubmitResponseReq.Marshal(b, m, deterministic)
}
func (m *SubmitResponseReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmitResponseReq.Merge(m, src)
}
func (m *SubmitResponseReq) XXX_Size() int {
	return xxx_messageInfo_SubmitResponseReq.Size(m)
}
func (m *SubmitResponseReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmitResponseReq.DiscardUnknown(m)
}

var xxx_messageInfo_SubmitResponseReq proto.InternalMessageInfo

func (m *SubmitResponseReq) GetToken() *TokenInfos {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *SubmitResponseReq) GetStudyKey() string {
	if m != nil {
		return m.StudyKey
	}
	return ""
}

func (m *SubmitResponseReq) GetResponse() *SurveyResponse {
	if m != nil {
		return m.Response
	}
	return nil
}

type EnterStudyRequest struct {
	Token                *TokenInfos `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	StudyKey             string      `protobuf:"bytes,2,opt,name=study_key,json=studyKey,proto3" json:"study_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *EnterStudyRequest) Reset()         { *m = EnterStudyRequest{} }
func (m *EnterStudyRequest) String() string { return proto.CompactTextString(m) }
func (*EnterStudyRequest) ProtoMessage()    {}
func (*EnterStudyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_81f0d8f98f9be15c, []int{6}
}

func (m *EnterStudyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnterStudyRequest.Unmarshal(m, b)
}
func (m *EnterStudyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnterStudyRequest.Marshal(b, m, deterministic)
}
func (m *EnterStudyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnterStudyRequest.Merge(m, src)
}
func (m *EnterStudyRequest) XXX_Size() int {
	return xxx_messageInfo_EnterStudyRequest.Size(m)
}
func (m *EnterStudyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EnterStudyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EnterStudyRequest proto.InternalMessageInfo

func (m *EnterStudyRequest) GetToken() *TokenInfos {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *EnterStudyRequest) GetStudyKey() string {
	if m != nil {
		return m.StudyKey
	}
	return ""
}

type SurveyReferenceRequest struct {
	Token                *TokenInfos `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	StudyKey             string      `protobuf:"bytes,2,opt,name=study_key,json=studyKey,proto3" json:"study_key,omitempty"`
	SurveyKey            string      `protobuf:"bytes,3,opt,name=survey_key,json=surveyKey,proto3" json:"survey_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SurveyReferenceRequest) Reset()         { *m = SurveyReferenceRequest{} }
func (m *SurveyReferenceRequest) String() string { return proto.CompactTextString(m) }
func (*SurveyReferenceRequest) ProtoMessage()    {}
func (*SurveyReferenceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_81f0d8f98f9be15c, []int{7}
}

func (m *SurveyReferenceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SurveyReferenceRequest.Unmarshal(m, b)
}
func (m *SurveyReferenceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SurveyReferenceRequest.Marshal(b, m, deterministic)
}
func (m *SurveyReferenceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SurveyReferenceRequest.Merge(m, src)
}
func (m *SurveyReferenceRequest) XXX_Size() int {
	return xxx_messageInfo_SurveyReferenceRequest.Size(m)
}
func (m *SurveyReferenceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SurveyReferenceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SurveyReferenceRequest proto.InternalMessageInfo

func (m *SurveyReferenceRequest) GetToken() *TokenInfos {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *SurveyReferenceRequest) GetStudyKey() string {
	if m != nil {
		return m.StudyKey
	}
	return ""
}

func (m *SurveyReferenceRequest) GetSurveyKey() string {
	if m != nil {
		return m.SurveyKey
	}
	return ""
}

type StatusReportRequest struct {
	Token                *TokenInfos     `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	StatusSurvey         *SurveyResponse `protobuf:"bytes,2,opt,name=status_survey,json=statusSurvey,proto3" json:"status_survey,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *StatusReportRequest) Reset()         { *m = StatusReportRequest{} }
func (m *StatusReportRequest) String() string { return proto.CompactTextString(m) }
func (*StatusReportRequest) ProtoMessage()    {}
func (*StatusReportRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_81f0d8f98f9be15c, []int{8}
}

func (m *StatusReportRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusReportRequest.Unmarshal(m, b)
}
func (m *StatusReportRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusReportRequest.Marshal(b, m, deterministic)
}
func (m *StatusReportRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusReportRequest.Merge(m, src)
}
func (m *StatusReportRequest) XXX_Size() int {
	return xxx_messageInfo_StatusReportRequest.Size(m)
}
func (m *StatusReportRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusReportRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StatusReportRequest proto.InternalMessageInfo

func (m *StatusReportRequest) GetToken() *TokenInfos {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *StatusReportRequest) GetStatusSurvey() *SurveyResponse {
	if m != nil {
		return m.StatusSurvey
	}
	return nil
}

func init() {
	proto.RegisterType((*NewStudyRequest)(nil), "inf.study_service_api.NewStudyRequest")
	proto.RegisterType((*SurveyAndContext)(nil), "inf.study_service_api.SurveyAndContext")
	proto.RegisterType((*StudyReferenceReq)(nil), "inf.study_service_api.StudyReferenceReq")
	proto.RegisterType((*SurveyInfoResp)(nil), "inf.study_service_api.SurveyInfoResp")
	proto.RegisterType((*SurveyInfoResp_SurveyInfo)(nil), "inf.study_service_api.SurveyInfoResp.SurveyInfo")
	proto.RegisterType((*AddSurveyReq)(nil), "inf.study_service_api.AddSurveyReq")
	proto.RegisterType((*SubmitResponseReq)(nil), "inf.study_service_api.SubmitResponseReq")
	proto.RegisterType((*EnterStudyRequest)(nil), "inf.study_service_api.EnterStudyRequest")
	proto.RegisterType((*SurveyReferenceRequest)(nil), "inf.study_service_api.SurveyReferenceRequest")
	proto.RegisterType((*StatusReportRequest)(nil), "inf.study_service_api.StatusReportRequest")
}

func init() {
	proto.RegisterFile("study-service-api.proto", fileDescriptor_81f0d8f98f9be15c)
}

var fileDescriptor_81f0d8f98f9be15c = []byte{
	// 724 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xdf, 0x4a, 0x1b, 0x4f,
	0x14, 0x66, 0xcd, 0x2f, 0xd1, 0x9c, 0xf8, 0x53, 0x33, 0xa2, 0x4d, 0x57, 0x0a, 0x12, 0xd1, 0x8a,
	0x90, 0xb5, 0xe8, 0xb5, 0x94, 0x54, 0xd4, 0xfe, 0x91, 0x16, 0x76, 0xa5, 0x85, 0x52, 0x48, 0x37,
	0xc9, 0x49, 0x98, 0x26, 0xd9, 0x59, 0x77, 0x26, 0xb6, 0x52, 0xe8, 0x65, 0x2f, 0x7b, 0xdd, 0xe7,
	0xe8, 0x9b, 0xf4, 0x31, 0xfa, 0x16, 0x65, 0xe7, 0xcc, 0xc6, 0x35, 0x31, 0x51, 0x4b, 0xee, 0x76,
	0xce, 0x9e, 0xef, 0x3b, 0x67, 0xbe, 0xf3, 0x67, 0xe0, 0x81, 0x54, 0xfd, 0xe6, 0x65, 0x45, 0x62,
	0x74, 0xc1, 0x1b, 0x58, 0xf1, 0x43, 0xee, 0x84, 0x91, 0x50, 0x82, 0xad, 0xf0, 0xa0, 0xe5, 0xe8,
	0x9f, 0x35, 0xf3, 0xb3, 0xe6, 0x87, 0xdc, 0x66, 0xed, 0xae, 0xa8, 0xfb, 0xdd, 0x8a, 0xba, 0x0c,
	0x51, 0x92, 0xab, 0x5d, 0xd0, 0x6e, 0xe6, 0x30, 0x2f, 0xfb, 0xd1, 0x05, 0x26, 0xa7, 0x15, 0x3a,
	0x55, 0x22, 0x94, 0xa1, 0x08, 0x24, 0x1a, 0xf3, 0x5a, 0x5b, 0x88, 0x76, 0x17, 0x77, 0xf5, 0xa9,
	0xde, 0x6f, 0xed, 0x62, 0x2f, 0x54, 0x06, 0x53, 0xfe, 0x08, 0x8b, 0xaf, 0xf1, 0xb3, 0x17, 0x73,
	0xba, 0x78, 0xde, 0x47, 0xa9, 0xd8, 0x26, 0x64, 0x95, 0xe8, 0x60, 0x50, 0xb2, 0xd6, 0xad, 0xed,
	0xc2, 0xde, 0xa2, 0x13, 0x27, 0x77, 0x16, 0x5b, 0x5e, 0x04, 0x2d, 0x21, 0x5d, 0xfa, 0xcb, 0xb6,
	0x20, 0xab, 0x53, 0x29, 0xcd, 0x68, 0xb7, 0x25, 0x67, 0x70, 0x07, 0x87, 0xe8, 0xe8, 0x77, 0xf9,
	0x97, 0x05, 0x4b, 0x9e, 0x4e, 0xac, 0x1a, 0x34, 0x0f, 0x45, 0xa0, 0xf0, 0x8b, 0x62, 0x3b, 0x90,
	0xa3, 0x64, 0x4d, 0x10, 0x46, 0x68, 0xba, 0x0d, 0x79, 0xbb, 0xc6, 0x83, 0xed, 0xc3, 0x6c, 0x83,
	0x60, 0x26, 0xd4, 0xc3, 0x51, 0x67, 0xc3, 0xeb, 0x26, 0x9e, 0xec, 0x00, 0x66, 0xc3, 0x08, 0x5b,
	0xbc, 0xdb, 0x2d, 0x65, 0x34, 0x68, 0x23, 0x05, 0xaa, 0x0d, 0x14, 0x32, 0xa1, 0xcc, 0xd1, 0x4d,
	0x30, 0xe5, 0x77, 0x50, 0x34, 0x9a, 0xb4, 0x30, 0xc2, 0xa0, 0x81, 0x2e, 0x9e, 0xdf, 0x55, 0x98,
	0x35, 0xc8, 0x53, 0x29, 0x3b, 0x48, 0xe2, 0xe4, 0xdd, 0x39, 0x6d, 0x78, 0x85, 0x97, 0xe5, 0x3f,
	0x16, 0x2c, 0x50, 0xd0, 0x18, 0x13, 0x07, 0x66, 0xc7, 0x90, 0xe5, 0x31, 0xbe, 0x64, 0xad, 0x67,
	0xb6, 0x0b, 0x7b, 0x4f, 0x9c, 0x1b, 0x9b, 0xc1, 0xb9, 0x8e, 0x4a, 0x1f, 0x09, 0x6e, 0xff, 0xb0,
	0x00, 0xae, 0xac, 0x6c, 0x09, 0x32, 0x1d, 0xa3, 0x6f, 0xde, 0x8d, 0x3f, 0xd9, 0x2e, 0xfc, 0x17,
	0xf8, 0x3d, 0x2c, 0xcd, 0xe8, 0x38, 0x6b, 0x69, 0x15, 0x4f, 0x45, 0xc3, 0xef, 0x72, 0x89, 0xcd,
	0x37, 0xf5, 0x4f, 0xd8, 0x50, 0xae, 0x76, 0x64, 0x07, 0x50, 0x68, 0xa2, 0x6c, 0x44, 0x3c, 0x54,
	0x5c, 0x04, 0xa5, 0xcc, 0xed, 0xb8, 0xb4, 0x7f, 0xf9, 0x1b, 0xcc, 0x57, 0x9b, 0xcd, 0x44, 0xe2,
	0xa9, 0xe8, 0x97, 0x6a, 0x9c, 0xcc, 0x6d, 0x8d, 0x53, 0xfe, 0x69, 0x41, 0xd1, 0xeb, 0xd7, 0x7b,
	0x5c, 0x0d, 0x0a, 0x3c, 0xa5, 0x2c, 0x9e, 0xc2, 0x5c, 0xd2, 0x42, 0xf7, 0x69, 0xaf, 0x01, 0x28,
	0xee, 0xaf, 0xa3, 0x40, 0x61, 0xf4, 0x2f, 0x83, 0x37, 0xb1, 0xbf, 0xbe, 0xc2, 0x6a, 0x12, 0xf4,
	0xaa, 0x73, 0xa7, 0xc5, 0xce, 0x1e, 0x01, 0x98, 0x2b, 0x76, 0x4c, 0x05, 0xf2, 0x6e, 0x9e, 0x2c,
	0x71, 0xf0, 0xef, 0x16, 0x2c, 0x7b, 0xca, 0x57, 0x7d, 0xe9, 0x62, 0x28, 0x22, 0x75, 0xcf, 0xd0,
	0xcf, 0xe1, 0x7f, 0xa9, 0xd1, 0x35, 0x53, 0xe2, 0x99, 0xbb, 0x4b, 0x3b, 0x4f, 0x48, 0xb2, 0xee,
	0xfd, 0xce, 0xc1, 0xa2, 0x96, 0xd6, 0xa3, 0x01, 0xaa, 0x86, 0x9c, 0x55, 0x20, 0x47, 0xb9, 0xb1,
	0x55, 0x87, 0x36, 0xa2, 0x93, 0x6c, 0x44, 0xe7, 0x28, 0xde, 0x88, 0x76, 0x41, 0x07, 0x32, 0x4e,
	0x2e, 0xc0, 0x55, 0x85, 0xd8, 0xf6, 0x98, 0xa1, 0x1c, 0x29, 0xa2, 0x6d, 0xa7, 0xf6, 0x60, 0x55,
	0x4a, 0xde, 0x0e, 0xd0, 0x8c, 0x80, 0x64, 0x55, 0x60, 0x27, 0xa8, 0x86, 0xad, 0xc3, 0x72, 0x4c,
	0xa4, 0xe8, 0x40, 0x71, 0x84, 0x82, 0x55, 0x26, 0xae, 0x8c, 0xe1, 0x4e, 0xb0, 0x1f, 0x4f, 0x74,
	0x4f, 0x6d, 0xe9, 0x0f, 0xc0, 0x68, 0x7e, 0xd2, 0x45, 0x65, 0x3b, 0xe3, 0xe0, 0xa3, 0x95, 0x9f,
	0x78, 0x95, 0xb7, 0xf1, 0x26, 0x4c, 0x4f, 0xe7, 0x58, 0x95, 0x47, 0x86, 0x78, 0x22, 0xef, 0x4b,
	0x58, 0x38, 0x8c, 0xd0, 0x57, 0x98, 0x3c, 0x6c, 0x6c, 0x6b, 0x0c, 0xef, 0xd0, 0xcb, 0x67, 0x8f,
	0xbc, 0x61, 0xec, 0x14, 0x8a, 0x9e, 0x7f, 0x81, 0x44, 0x7d, 0x26, 0xc8, 0xb8, 0x31, 0x86, 0x2e,
	0xbd, 0xec, 0xec, 0x1b, 0x16, 0x13, 0xf3, 0x60, 0xc5, 0xc5, 0x9e, 0x48, 0xf8, 0x8e, 0x23, 0xd1,
	0x23, 0xc6, 0x7b, 0x16, 0xf0, 0x5a, 0xa3, 0xb6, 0x60, 0xf9, 0x04, 0x15, 0x75, 0xfb, 0x60, 0xfb,
	0xcb, 0xf1, 0x5a, 0x0e, 0x3f, 0x6b, 0xf6, 0xe6, 0x9d, 0x1e, 0x9c, 0x67, 0xd9, 0xf7, 0x19, 0x3f,
	0xe4, 0xf5, 0x9c, 0x1e, 0x9a, 0xfd, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3e, 0xa9, 0x7c, 0x90,
	0xcc, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StudyServiceApiClient is the client API for StudyServiceApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StudyServiceApiClient interface {
	Status(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Status, error)
	// Study flow
	EnterStudy(ctx context.Context, in *EnterStudyRequest, opts ...grpc.CallOption) (*AssignedSurveys, error)
	GetAssignedSurveys(ctx context.Context, in *TokenInfos, opts ...grpc.CallOption) (*AssignedSurveys, error)
	GetAssignedSurvey(ctx context.Context, in *SurveyReferenceRequest, opts ...grpc.CallOption) (*SurveyAndContext, error)
	SubmitStatusReport(ctx context.Context, in *StatusReportRequest, opts ...grpc.CallOption) (*AssignedSurveys, error)
	SubmitResponse(ctx context.Context, in *SubmitResponseReq, opts ...grpc.CallOption) (*AssignedSurveys, error)
	// Study management
	CreateNewStudy(ctx context.Context, in *NewStudyRequest, opts ...grpc.CallOption) (*Study, error)
	SaveSurveyToStudy(ctx context.Context, in *AddSurveyReq, opts ...grpc.CallOption) (*Survey, error)
	RemoveSurveyFromStudy(ctx context.Context, in *SurveyReferenceRequest, opts ...grpc.CallOption) (*Status, error)
	GetStudySurveyInfos(ctx context.Context, in *StudyReferenceReq, opts ...grpc.CallOption) (*SurveyInfoResp, error)
}

type studyServiceApiClient struct {
	cc grpc.ClientConnInterface
}

func NewStudyServiceApiClient(cc grpc.ClientConnInterface) StudyServiceApiClient {
	return &studyServiceApiClient{cc}
}

func (c *studyServiceApiClient) Status(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/inf.study_service_api.StudyServiceApi/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studyServiceApiClient) EnterStudy(ctx context.Context, in *EnterStudyRequest, opts ...grpc.CallOption) (*AssignedSurveys, error) {
	out := new(AssignedSurveys)
	err := c.cc.Invoke(ctx, "/inf.study_service_api.StudyServiceApi/EnterStudy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studyServiceApiClient) GetAssignedSurveys(ctx context.Context, in *TokenInfos, opts ...grpc.CallOption) (*AssignedSurveys, error) {
	out := new(AssignedSurveys)
	err := c.cc.Invoke(ctx, "/inf.study_service_api.StudyServiceApi/GetAssignedSurveys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studyServiceApiClient) GetAssignedSurvey(ctx context.Context, in *SurveyReferenceRequest, opts ...grpc.CallOption) (*SurveyAndContext, error) {
	out := new(SurveyAndContext)
	err := c.cc.Invoke(ctx, "/inf.study_service_api.StudyServiceApi/GetAssignedSurvey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studyServiceApiClient) SubmitStatusReport(ctx context.Context, in *StatusReportRequest, opts ...grpc.CallOption) (*AssignedSurveys, error) {
	out := new(AssignedSurveys)
	err := c.cc.Invoke(ctx, "/inf.study_service_api.StudyServiceApi/SubmitStatusReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studyServiceApiClient) SubmitResponse(ctx context.Context, in *SubmitResponseReq, opts ...grpc.CallOption) (*AssignedSurveys, error) {
	out := new(AssignedSurveys)
	err := c.cc.Invoke(ctx, "/inf.study_service_api.StudyServiceApi/SubmitResponse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studyServiceApiClient) CreateNewStudy(ctx context.Context, in *NewStudyRequest, opts ...grpc.CallOption) (*Study, error) {
	out := new(Study)
	err := c.cc.Invoke(ctx, "/inf.study_service_api.StudyServiceApi/CreateNewStudy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studyServiceApiClient) SaveSurveyToStudy(ctx context.Context, in *AddSurveyReq, opts ...grpc.CallOption) (*Survey, error) {
	out := new(Survey)
	err := c.cc.Invoke(ctx, "/inf.study_service_api.StudyServiceApi/SaveSurveyToStudy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studyServiceApiClient) RemoveSurveyFromStudy(ctx context.Context, in *SurveyReferenceRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/inf.study_service_api.StudyServiceApi/RemoveSurveyFromStudy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studyServiceApiClient) GetStudySurveyInfos(ctx context.Context, in *StudyReferenceReq, opts ...grpc.CallOption) (*SurveyInfoResp, error) {
	out := new(SurveyInfoResp)
	err := c.cc.Invoke(ctx, "/inf.study_service_api.StudyServiceApi/GetStudySurveyInfos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StudyServiceApiServer is the server API for StudyServiceApi service.
type StudyServiceApiServer interface {
	Status(context.Context, *empty.Empty) (*Status, error)
	// Study flow
	EnterStudy(context.Context, *EnterStudyRequest) (*AssignedSurveys, error)
	GetAssignedSurveys(context.Context, *TokenInfos) (*AssignedSurveys, error)
	GetAssignedSurvey(context.Context, *SurveyReferenceRequest) (*SurveyAndContext, error)
	SubmitStatusReport(context.Context, *StatusReportRequest) (*AssignedSurveys, error)
	SubmitResponse(context.Context, *SubmitResponseReq) (*AssignedSurveys, error)
	// Study management
	CreateNewStudy(context.Context, *NewStudyRequest) (*Study, error)
	SaveSurveyToStudy(context.Context, *AddSurveyReq) (*Survey, error)
	RemoveSurveyFromStudy(context.Context, *SurveyReferenceRequest) (*Status, error)
	GetStudySurveyInfos(context.Context, *StudyReferenceReq) (*SurveyInfoResp, error)
}

// UnimplementedStudyServiceApiServer can be embedded to have forward compatible implementations.
type UnimplementedStudyServiceApiServer struct {
}

func (*UnimplementedStudyServiceApiServer) Status(ctx context.Context, req *empty.Empty) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (*UnimplementedStudyServiceApiServer) EnterStudy(ctx context.Context, req *EnterStudyRequest) (*AssignedSurveys, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnterStudy not implemented")
}
func (*UnimplementedStudyServiceApiServer) GetAssignedSurveys(ctx context.Context, req *TokenInfos) (*AssignedSurveys, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAssignedSurveys not implemented")
}
func (*UnimplementedStudyServiceApiServer) GetAssignedSurvey(ctx context.Context, req *SurveyReferenceRequest) (*SurveyAndContext, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAssignedSurvey not implemented")
}
func (*UnimplementedStudyServiceApiServer) SubmitStatusReport(ctx context.Context, req *StatusReportRequest) (*AssignedSurveys, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitStatusReport not implemented")
}
func (*UnimplementedStudyServiceApiServer) SubmitResponse(ctx context.Context, req *SubmitResponseReq) (*AssignedSurveys, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitResponse not implemented")
}
func (*UnimplementedStudyServiceApiServer) CreateNewStudy(ctx context.Context, req *NewStudyRequest) (*Study, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewStudy not implemented")
}
func (*UnimplementedStudyServiceApiServer) SaveSurveyToStudy(ctx context.Context, req *AddSurveyReq) (*Survey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveSurveyToStudy not implemented")
}
func (*UnimplementedStudyServiceApiServer) RemoveSurveyFromStudy(ctx context.Context, req *SurveyReferenceRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveSurveyFromStudy not implemented")
}
func (*UnimplementedStudyServiceApiServer) GetStudySurveyInfos(ctx context.Context, req *StudyReferenceReq) (*SurveyInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudySurveyInfos not implemented")
}

func RegisterStudyServiceApiServer(s *grpc.Server, srv StudyServiceApiServer) {
	s.RegisterService(&_StudyServiceApi_serviceDesc, srv)
}

func _StudyServiceApi_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudyServiceApiServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inf.study_service_api.StudyServiceApi/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudyServiceApiServer).Status(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudyServiceApi_EnterStudy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnterStudyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudyServiceApiServer).EnterStudy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inf.study_service_api.StudyServiceApi/EnterStudy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudyServiceApiServer).EnterStudy(ctx, req.(*EnterStudyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudyServiceApi_GetAssignedSurveys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenInfos)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudyServiceApiServer).GetAssignedSurveys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inf.study_service_api.StudyServiceApi/GetAssignedSurveys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudyServiceApiServer).GetAssignedSurveys(ctx, req.(*TokenInfos))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudyServiceApi_GetAssignedSurvey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SurveyReferenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudyServiceApiServer).GetAssignedSurvey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inf.study_service_api.StudyServiceApi/GetAssignedSurvey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudyServiceApiServer).GetAssignedSurvey(ctx, req.(*SurveyReferenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudyServiceApi_SubmitStatusReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudyServiceApiServer).SubmitStatusReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inf.study_service_api.StudyServiceApi/SubmitStatusReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudyServiceApiServer).SubmitStatusReport(ctx, req.(*StatusReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudyServiceApi_SubmitResponse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitResponseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudyServiceApiServer).SubmitResponse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inf.study_service_api.StudyServiceApi/SubmitResponse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudyServiceApiServer).SubmitResponse(ctx, req.(*SubmitResponseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudyServiceApi_CreateNewStudy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewStudyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudyServiceApiServer).CreateNewStudy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inf.study_service_api.StudyServiceApi/CreateNewStudy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudyServiceApiServer).CreateNewStudy(ctx, req.(*NewStudyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudyServiceApi_SaveSurveyToStudy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSurveyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudyServiceApiServer).SaveSurveyToStudy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inf.study_service_api.StudyServiceApi/SaveSurveyToStudy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudyServiceApiServer).SaveSurveyToStudy(ctx, req.(*AddSurveyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudyServiceApi_RemoveSurveyFromStudy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SurveyReferenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudyServiceApiServer).RemoveSurveyFromStudy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inf.study_service_api.StudyServiceApi/RemoveSurveyFromStudy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudyServiceApiServer).RemoveSurveyFromStudy(ctx, req.(*SurveyReferenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudyServiceApi_GetStudySurveyInfos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StudyReferenceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudyServiceApiServer).GetStudySurveyInfos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inf.study_service_api.StudyServiceApi/GetStudySurveyInfos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudyServiceApiServer).GetStudySurveyInfos(ctx, req.(*StudyReferenceReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _StudyServiceApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "inf.study_service_api.StudyServiceApi",
	HandlerType: (*StudyServiceApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _StudyServiceApi_Status_Handler,
		},
		{
			MethodName: "EnterStudy",
			Handler:    _StudyServiceApi_EnterStudy_Handler,
		},
		{
			MethodName: "GetAssignedSurveys",
			Handler:    _StudyServiceApi_GetAssignedSurveys_Handler,
		},
		{
			MethodName: "GetAssignedSurvey",
			Handler:    _StudyServiceApi_GetAssignedSurvey_Handler,
		},
		{
			MethodName: "SubmitStatusReport",
			Handler:    _StudyServiceApi_SubmitStatusReport_Handler,
		},
		{
			MethodName: "SubmitResponse",
			Handler:    _StudyServiceApi_SubmitResponse_Handler,
		},
		{
			MethodName: "CreateNewStudy",
			Handler:    _StudyServiceApi_CreateNewStudy_Handler,
		},
		{
			MethodName: "SaveSurveyToStudy",
			Handler:    _StudyServiceApi_SaveSurveyToStudy_Handler,
		},
		{
			MethodName: "RemoveSurveyFromStudy",
			Handler:    _StudyServiceApi_RemoveSurveyFromStudy_Handler,
		},
		{
			MethodName: "GetStudySurveyInfos",
			Handler:    _StudyServiceApi_GetStudySurveyInfos_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "study-service-api.proto",
}
