// Code generated by goctl. DO NOT EDIT.
// Source: collection.proto

package cat

import (
	"context"

	"github.com/xh-polaris/meowchat-collection-rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddCatReq      = pb.AddCatReq
	AddCatResp     = pb.AddCatResp
	Cat            = pb.Cat
	DeleteCatReq   = pb.DeleteCatReq
	DeleteCatResp  = pb.DeleteCatResp
	GetCatReq      = pb.GetCatReq
	GetCatResp     = pb.GetCatResp
	GetManyCatReq  = pb.GetManyCatReq
	GetManyCatResp = pb.GetManyCatResp
	UpdateCatReq   = pb.UpdateCatReq
	UpdateCatResp  = pb.UpdateCatResp

	Cat interface {
		GetManyCat(ctx context.Context, in *GetManyCatReq, opts ...grpc.CallOption) (*GetManyCatResp, error)
		GetCat(ctx context.Context, in *GetCatReq, opts ...grpc.CallOption) (*GetCatResp, error)
		AddCat(ctx context.Context, in *AddCatReq, opts ...grpc.CallOption) (*AddCatResp, error)
		UpdateCat(ctx context.Context, in *UpdateCatReq, opts ...grpc.CallOption) (*UpdateCatResp, error)
		DeleteCat(ctx context.Context, in *DeleteCatReq, opts ...grpc.CallOption) (*DeleteCatResp, error)
	}

	defaultCat struct {
		cli zrpc.Client
	}
)

func NewCat(cli zrpc.Client) Cat {
	return &defaultCat{
		cli: cli,
	}
}

func (m *defaultCat) GetManyCat(ctx context.Context, in *GetManyCatReq, opts ...grpc.CallOption) (*GetManyCatResp, error) {
	client := pb.NewCatClient(m.cli.Conn())
	return client.GetManyCat(ctx, in, opts...)
}

func (m *defaultCat) GetCat(ctx context.Context, in *GetCatReq, opts ...grpc.CallOption) (*GetCatResp, error) {
	client := pb.NewCatClient(m.cli.Conn())
	return client.GetCat(ctx, in, opts...)
}

func (m *defaultCat) AddCat(ctx context.Context, in *AddCatReq, opts ...grpc.CallOption) (*AddCatResp, error) {
	client := pb.NewCatClient(m.cli.Conn())
	return client.AddCat(ctx, in, opts...)
}

func (m *defaultCat) UpdateCat(ctx context.Context, in *UpdateCatReq, opts ...grpc.CallOption) (*UpdateCatResp, error) {
	client := pb.NewCatClient(m.cli.Conn())
	return client.UpdateCat(ctx, in, opts...)
}

func (m *defaultCat) DeleteCat(ctx context.Context, in *DeleteCatReq, opts ...grpc.CallOption) (*DeleteCatResp, error) {
	client := pb.NewCatClient(m.cli.Conn())
	return client.DeleteCat(ctx, in, opts...)
}
