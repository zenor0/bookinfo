// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.12.4
// source: proto/reviews/review.proto

package reviews

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Empty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_proto_reviews_review_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reviews_review_proto_msgTypes[0]
	if x != nil {
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
	return file_proto_reviews_review_proto_rawDescGZIP(), []int{0}
}

type Review struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	BookId        uint32                 `protobuf:"varint,2,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	UserId        uint32                 `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Username      string                 `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	Content       string                 `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	Rating        float64                `protobuf:"fixed64,6,opt,name=rating,proto3" json:"rating,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string                 `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Review) Reset() {
	*x = Review{}
	mi := &file_proto_reviews_review_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Review) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Review) ProtoMessage() {}

func (x *Review) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reviews_review_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Review.ProtoReflect.Descriptor instead.
func (*Review) Descriptor() ([]byte, []int) {
	return file_proto_reviews_review_proto_rawDescGZIP(), []int{1}
}

func (x *Review) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Review) GetBookId() uint32 {
	if x != nil {
		return x.BookId
	}
	return 0
}

func (x *Review) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Review) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Review) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Review) GetRating() float64 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *Review) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Review) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type ReviewResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Review        *Review                `protobuf:"bytes,1,opt,name=review,proto3" json:"review,omitempty"`
	Rating        float64                `protobuf:"fixed64,2,opt,name=rating,proto3" json:"rating,omitempty"`
	RatingHtml    string                 `protobuf:"bytes,3,opt,name=rating_html,json=ratingHtml,proto3" json:"rating_html,omitempty"`
	Color         string                 `protobuf:"bytes,4,opt,name=color,proto3" json:"color,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReviewResponse) Reset() {
	*x = ReviewResponse{}
	mi := &file_proto_reviews_review_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReviewResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReviewResponse) ProtoMessage() {}

func (x *ReviewResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reviews_review_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReviewResponse.ProtoReflect.Descriptor instead.
func (*ReviewResponse) Descriptor() ([]byte, []int) {
	return file_proto_reviews_review_proto_rawDescGZIP(), []int{2}
}

func (x *ReviewResponse) GetReview() *Review {
	if x != nil {
		return x.Review
	}
	return nil
}

func (x *ReviewResponse) GetRating() float64 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *ReviewResponse) GetRatingHtml() string {
	if x != nil {
		return x.RatingHtml
	}
	return ""
}

func (x *ReviewResponse) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

type CreateReviewRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Review        *Review                `protobuf:"bytes,1,opt,name=review,proto3" json:"review,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateReviewRequest) Reset() {
	*x = CreateReviewRequest{}
	mi := &file_proto_reviews_review_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateReviewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReviewRequest) ProtoMessage() {}

func (x *CreateReviewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reviews_review_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReviewRequest.ProtoReflect.Descriptor instead.
func (*CreateReviewRequest) Descriptor() ([]byte, []int) {
	return file_proto_reviews_review_proto_rawDescGZIP(), []int{3}
}

func (x *CreateReviewRequest) GetReview() *Review {
	if x != nil {
		return x.Review
	}
	return nil
}

type GetReviewRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetReviewRequest) Reset() {
	*x = GetReviewRequest{}
	mi := &file_proto_reviews_review_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetReviewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReviewRequest) ProtoMessage() {}

func (x *GetReviewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reviews_review_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReviewRequest.ProtoReflect.Descriptor instead.
func (*GetReviewRequest) Descriptor() ([]byte, []int) {
	return file_proto_reviews_review_proto_rawDescGZIP(), []int{4}
}

func (x *GetReviewRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetBookReviewsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	BookId        uint32                 `protobuf:"varint,1,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	Page          int32                  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	PageSize      int32                  `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBookReviewsRequest) Reset() {
	*x = GetBookReviewsRequest{}
	mi := &file_proto_reviews_review_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBookReviewsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookReviewsRequest) ProtoMessage() {}

func (x *GetBookReviewsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reviews_review_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookReviewsRequest.ProtoReflect.Descriptor instead.
func (*GetBookReviewsRequest) Descriptor() ([]byte, []int) {
	return file_proto_reviews_review_proto_rawDescGZIP(), []int{5}
}

func (x *GetBookReviewsRequest) GetBookId() uint32 {
	if x != nil {
		return x.BookId
	}
	return 0
}

func (x *GetBookReviewsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetBookReviewsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type GetBookReviewsResponse struct {
	state         protoimpl.MessageState       `protogen:"open.v1"`
	Data          []*ReviewResponse            `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Meta          *GetBookReviewsResponse_Meta `protobuf:"bytes,2,opt,name=meta,proto3" json:"meta,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBookReviewsResponse) Reset() {
	*x = GetBookReviewsResponse{}
	mi := &file_proto_reviews_review_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBookReviewsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookReviewsResponse) ProtoMessage() {}

func (x *GetBookReviewsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reviews_review_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookReviewsResponse.ProtoReflect.Descriptor instead.
func (*GetBookReviewsResponse) Descriptor() ([]byte, []int) {
	return file_proto_reviews_review_proto_rawDescGZIP(), []int{6}
}

func (x *GetBookReviewsResponse) GetData() []*ReviewResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GetBookReviewsResponse) GetMeta() *GetBookReviewsResponse_Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

type GetUserReviewsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        uint32                 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserReviewsRequest) Reset() {
	*x = GetUserReviewsRequest{}
	mi := &file_proto_reviews_review_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserReviewsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserReviewsRequest) ProtoMessage() {}

func (x *GetUserReviewsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reviews_review_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserReviewsRequest.ProtoReflect.Descriptor instead.
func (*GetUserReviewsRequest) Descriptor() ([]byte, []int) {
	return file_proto_reviews_review_proto_rawDescGZIP(), []int{7}
}

func (x *GetUserReviewsRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetUserReviewsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Reviews       []*ReviewResponse      `protobuf:"bytes,1,rep,name=reviews,proto3" json:"reviews,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserReviewsResponse) Reset() {
	*x = GetUserReviewsResponse{}
	mi := &file_proto_reviews_review_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserReviewsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserReviewsResponse) ProtoMessage() {}

func (x *GetUserReviewsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reviews_review_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserReviewsResponse.ProtoReflect.Descriptor instead.
func (*GetUserReviewsResponse) Descriptor() ([]byte, []int) {
	return file_proto_reviews_review_proto_rawDescGZIP(), []int{8}
}

func (x *GetUserReviewsResponse) GetReviews() []*ReviewResponse {
	if x != nil {
		return x.Reviews
	}
	return nil
}

type UpdateReviewRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Review        *Review                `protobuf:"bytes,1,opt,name=review,proto3" json:"review,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateReviewRequest) Reset() {
	*x = UpdateReviewRequest{}
	mi := &file_proto_reviews_review_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateReviewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateReviewRequest) ProtoMessage() {}

func (x *UpdateReviewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reviews_review_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateReviewRequest.ProtoReflect.Descriptor instead.
func (*UpdateReviewRequest) Descriptor() ([]byte, []int) {
	return file_proto_reviews_review_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateReviewRequest) GetReview() *Review {
	if x != nil {
		return x.Review
	}
	return nil
}

type DeleteReviewRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteReviewRequest) Reset() {
	*x = DeleteReviewRequest{}
	mi := &file_proto_reviews_review_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteReviewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteReviewRequest) ProtoMessage() {}

func (x *DeleteReviewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reviews_review_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteReviewRequest.ProtoReflect.Descriptor instead.
func (*DeleteReviewRequest) Descriptor() ([]byte, []int) {
	return file_proto_reviews_review_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteReviewRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetBookReviewsResponse_Meta struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Page          int32                  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize      int32                  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Total         int32                  `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBookReviewsResponse_Meta) Reset() {
	*x = GetBookReviewsResponse_Meta{}
	mi := &file_proto_reviews_review_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBookReviewsResponse_Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookReviewsResponse_Meta) ProtoMessage() {}

func (x *GetBookReviewsResponse_Meta) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reviews_review_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookReviewsResponse_Meta.ProtoReflect.Descriptor instead.
func (*GetBookReviewsResponse_Meta) Descriptor() ([]byte, []int) {
	return file_proto_reviews_review_proto_rawDescGZIP(), []int{6, 0}
}

func (x *GetBookReviewsResponse_Meta) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetBookReviewsResponse_Meta) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GetBookReviewsResponse_Meta) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_proto_reviews_review_proto protoreflect.FileDescriptor

var file_proto_reviews_review_proto_rawDesc = string([]byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2f,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x73, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0xd6,
	0x01, 0x0a, 0x06, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x6f, 0x6f,
	0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b,
	0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x88, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x73, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x06, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x72,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x68, 0x74, 0x6d, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x48, 0x74, 0x6d, 0x6c, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c,
	0x6f, 0x72, 0x22, 0x3e, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x06, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x73, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x06, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x22, 0x22, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x61, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f,
	0x6b, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0xce, 0x01, 0x0a, 0x16, 0x47, 0x65,
	0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x52, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x38, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x24, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f,
	0x6b, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x1a, 0x4d, 0x0a, 0x04, 0x4d,
	0x65, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x30, 0x0a, 0x15, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x4b, 0x0a, 0x16,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x73, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x07, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x22, 0x3e, 0x0a, 0x13, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x27, 0x0a, 0x06, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x52, 0x06, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x22, 0x25, 0x0a, 0x13, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64,
	0x32, 0xaa, 0x03, 0x0a, 0x0d, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x12, 0x1c, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0f, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x12, 0x37, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x19,
	0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x73, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x51, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x12, 0x1e, 0x2e, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x12,
	0x1e, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1f, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3d, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x12, 0x1c, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f,
	0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12,
	0x3c, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12,
	0x1c, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x2a, 0x5a,
	0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x65, 0x6e, 0x6f,
	0x72, 0x30, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x66, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var (
	file_proto_reviews_review_proto_rawDescOnce sync.Once
	file_proto_reviews_review_proto_rawDescData []byte
)

func file_proto_reviews_review_proto_rawDescGZIP() []byte {
	file_proto_reviews_review_proto_rawDescOnce.Do(func() {
		file_proto_reviews_review_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_reviews_review_proto_rawDesc), len(file_proto_reviews_review_proto_rawDesc)))
	})
	return file_proto_reviews_review_proto_rawDescData
}

var file_proto_reviews_review_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_proto_reviews_review_proto_goTypes = []any{
	(*Empty)(nil),                       // 0: reviews.Empty
	(*Review)(nil),                      // 1: reviews.Review
	(*ReviewResponse)(nil),              // 2: reviews.ReviewResponse
	(*CreateReviewRequest)(nil),         // 3: reviews.CreateReviewRequest
	(*GetReviewRequest)(nil),            // 4: reviews.GetReviewRequest
	(*GetBookReviewsRequest)(nil),       // 5: reviews.GetBookReviewsRequest
	(*GetBookReviewsResponse)(nil),      // 6: reviews.GetBookReviewsResponse
	(*GetUserReviewsRequest)(nil),       // 7: reviews.GetUserReviewsRequest
	(*GetUserReviewsResponse)(nil),      // 8: reviews.GetUserReviewsResponse
	(*UpdateReviewRequest)(nil),         // 9: reviews.UpdateReviewRequest
	(*DeleteReviewRequest)(nil),         // 10: reviews.DeleteReviewRequest
	(*GetBookReviewsResponse_Meta)(nil), // 11: reviews.GetBookReviewsResponse.Meta
}
var file_proto_reviews_review_proto_depIdxs = []int32{
	1,  // 0: reviews.ReviewResponse.review:type_name -> reviews.Review
	1,  // 1: reviews.CreateReviewRequest.review:type_name -> reviews.Review
	2,  // 2: reviews.GetBookReviewsResponse.data:type_name -> reviews.ReviewResponse
	11, // 3: reviews.GetBookReviewsResponse.meta:type_name -> reviews.GetBookReviewsResponse.Meta
	2,  // 4: reviews.GetUserReviewsResponse.reviews:type_name -> reviews.ReviewResponse
	1,  // 5: reviews.UpdateReviewRequest.review:type_name -> reviews.Review
	3,  // 6: reviews.ReviewService.CreateReview:input_type -> reviews.CreateReviewRequest
	4,  // 7: reviews.ReviewService.GetReview:input_type -> reviews.GetReviewRequest
	5,  // 8: reviews.ReviewService.GetBookReviews:input_type -> reviews.GetBookReviewsRequest
	7,  // 9: reviews.ReviewService.GetUserReviews:input_type -> reviews.GetUserReviewsRequest
	9,  // 10: reviews.ReviewService.UpdateReview:input_type -> reviews.UpdateReviewRequest
	10, // 11: reviews.ReviewService.DeleteReview:input_type -> reviews.DeleteReviewRequest
	1,  // 12: reviews.ReviewService.CreateReview:output_type -> reviews.Review
	1,  // 13: reviews.ReviewService.GetReview:output_type -> reviews.Review
	6,  // 14: reviews.ReviewService.GetBookReviews:output_type -> reviews.GetBookReviewsResponse
	8,  // 15: reviews.ReviewService.GetUserReviews:output_type -> reviews.GetUserReviewsResponse
	1,  // 16: reviews.ReviewService.UpdateReview:output_type -> reviews.Review
	0,  // 17: reviews.ReviewService.DeleteReview:output_type -> reviews.Empty
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_proto_reviews_review_proto_init() }
func file_proto_reviews_review_proto_init() {
	if File_proto_reviews_review_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_reviews_review_proto_rawDesc), len(file_proto_reviews_review_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_reviews_review_proto_goTypes,
		DependencyIndexes: file_proto_reviews_review_proto_depIdxs,
		MessageInfos:      file_proto_reviews_review_proto_msgTypes,
	}.Build()
	File_proto_reviews_review_proto = out.File
	file_proto_reviews_review_proto_goTypes = nil
	file_proto_reviews_review_proto_depIdxs = nil
}
