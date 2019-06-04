// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: sequence/sequence.proto

package vc_pb_sequence

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Sequence service

type SequenceService interface {
	Create(ctx context.Context, in *ReqCreate, opts ...client.CallOption) (*ResDoc, error)
	Find(ctx context.Context, in *ReqFind, opts ...client.CallOption) (*ResDocs, error)
	FindOne(ctx context.Context, in *ReqFind, opts ...client.CallOption) (*ResDoc, error)
	FindById(ctx context.Context, in *Model, opts ...client.CallOption) (*ResDoc, error)
	FindDocsAndCount(ctx context.Context, in *ReqFind, opts ...client.CallOption) (*ResDocsAndCount, error)
	IncByName(ctx context.Context, in *Model, opts ...client.CallOption) (*ResDoc, error)
}

type sequenceService struct {
	c    client.Client
	name string
}

func NewSequenceService(name string, c client.Client) SequenceService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "vc.pb.sequence"
	}
	return &sequenceService{
		c:    c,
		name: name,
	}
}

func (c *sequenceService) Create(ctx context.Context, in *ReqCreate, opts ...client.CallOption) (*ResDoc, error) {
	req := c.c.NewRequest(c.name, "Sequence.Create", in)
	out := new(ResDoc)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sequenceService) Find(ctx context.Context, in *ReqFind, opts ...client.CallOption) (*ResDocs, error) {
	req := c.c.NewRequest(c.name, "Sequence.Find", in)
	out := new(ResDocs)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sequenceService) FindOne(ctx context.Context, in *ReqFind, opts ...client.CallOption) (*ResDoc, error) {
	req := c.c.NewRequest(c.name, "Sequence.FindOne", in)
	out := new(ResDoc)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sequenceService) FindById(ctx context.Context, in *Model, opts ...client.CallOption) (*ResDoc, error) {
	req := c.c.NewRequest(c.name, "Sequence.FindById", in)
	out := new(ResDoc)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sequenceService) FindDocsAndCount(ctx context.Context, in *ReqFind, opts ...client.CallOption) (*ResDocsAndCount, error) {
	req := c.c.NewRequest(c.name, "Sequence.FindDocsAndCount", in)
	out := new(ResDocsAndCount)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sequenceService) IncByName(ctx context.Context, in *Model, opts ...client.CallOption) (*ResDoc, error) {
	req := c.c.NewRequest(c.name, "Sequence.IncByName", in)
	out := new(ResDoc)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Sequence service

type SequenceHandler interface {
	Create(context.Context, *ReqCreate, *ResDoc) error
	Find(context.Context, *ReqFind, *ResDocs) error
	FindOne(context.Context, *ReqFind, *ResDoc) error
	FindById(context.Context, *Model, *ResDoc) error
	FindDocsAndCount(context.Context, *ReqFind, *ResDocsAndCount) error
	IncByName(context.Context, *Model, *ResDoc) error
}

func RegisterSequenceHandler(s server.Server, hdlr SequenceHandler, opts ...server.HandlerOption) error {
	type sequence interface {
		Create(ctx context.Context, in *ReqCreate, out *ResDoc) error
		Find(ctx context.Context, in *ReqFind, out *ResDocs) error
		FindOne(ctx context.Context, in *ReqFind, out *ResDoc) error
		FindById(ctx context.Context, in *Model, out *ResDoc) error
		FindDocsAndCount(ctx context.Context, in *ReqFind, out *ResDocsAndCount) error
		IncByName(ctx context.Context, in *Model, out *ResDoc) error
	}
	type Sequence struct {
		sequence
	}
	h := &sequenceHandler{hdlr}
	return s.Handle(s.NewHandler(&Sequence{h}, opts...))
}

type sequenceHandler struct {
	SequenceHandler
}

func (h *sequenceHandler) Create(ctx context.Context, in *ReqCreate, out *ResDoc) error {
	return h.SequenceHandler.Create(ctx, in, out)
}

func (h *sequenceHandler) Find(ctx context.Context, in *ReqFind, out *ResDocs) error {
	return h.SequenceHandler.Find(ctx, in, out)
}

func (h *sequenceHandler) FindOne(ctx context.Context, in *ReqFind, out *ResDoc) error {
	return h.SequenceHandler.FindOne(ctx, in, out)
}

func (h *sequenceHandler) FindById(ctx context.Context, in *Model, out *ResDoc) error {
	return h.SequenceHandler.FindById(ctx, in, out)
}

func (h *sequenceHandler) FindDocsAndCount(ctx context.Context, in *ReqFind, out *ResDocsAndCount) error {
	return h.SequenceHandler.FindDocsAndCount(ctx, in, out)
}

func (h *sequenceHandler) IncByName(ctx context.Context, in *Model, out *ResDoc) error {
	return h.SequenceHandler.IncByName(ctx, in, out)
}
