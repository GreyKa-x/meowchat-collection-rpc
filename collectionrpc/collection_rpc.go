// Code generated by goctl. DO NOT EDIT!
// Source: collection.proto

package collectionrpc

import (
	"context"
	pb2 "github.com/xh-polaris/meowchat-collection-rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Cat               = pb2.Cat
	CreateCatReq      = pb2.CreateCatReq
	CreateCatResp     = pb2.CreateCatResp
	CreateImageReq    = pb2.CreateImageReq
	CreateImageResp   = pb2.CreateImageResp
	DeleteCatReq      = pb2.DeleteCatReq
	DeleteCatResp     = pb2.DeleteCatResp
	DeleteImageReq    = pb2.DeleteImageReq
	DeleteImageResp   = pb2.DeleteImageResp
	GetImageByCatReq  = pb2.GetImageByCatReq
	GetImageByCatResp = pb2.GetImageByCatResp
	ListCatReq        = pb2.ListCatReq
	ListCatResp       = pb2.ListCatResp
	RetrieveCatReq    = pb2.RetrieveCatReq
	RetrieveCatResp   = pb2.RetrieveCatResp
	SearchCatReq      = pb2.SearchCatReq
	SearchCatResp     = pb2.SearchCatResp
	UpdateCatReq      = pb2.UpdateCatReq
	UpdateCatResp     = pb2.UpdateCatResp

	CollectionRpc interface {
		SearchCat(ctx context.Context, in *SearchCatReq, opts ...grpc.CallOption) (*SearchCatResp, error)
		ListCat(ctx context.Context, in *ListCatReq, opts ...grpc.CallOption) (*ListCatResp, error)
		RetrieveCat(ctx context.Context, in *RetrieveCatReq, opts ...grpc.CallOption) (*RetrieveCatResp, error)
		CreateCat(ctx context.Context, in *CreateCatReq, opts ...grpc.CallOption) (*CreateCatResp, error)
		UpdateCat(ctx context.Context, in *UpdateCatReq, opts ...grpc.CallOption) (*UpdateCatResp, error)
		DeleteCat(ctx context.Context, in *DeleteCatReq, opts ...grpc.CallOption) (*DeleteCatResp, error)
		CreateImage(ctx context.Context, in *CreateImageReq, opts ...grpc.CallOption) (*CreateImageResp, error)
		DeleteImage(ctx context.Context, in *DeleteImageReq, opts ...grpc.CallOption) (*DeleteImageResp, error)
		GetImageByCat(ctx context.Context, in *GetImageByCatReq, opts ...grpc.CallOption) (*GetImageByCatResp, error)
	}

	defaultCollectionRpc struct {
		cli zrpc.Client
	}
)

func NewCollectionRpc(cli zrpc.Client) CollectionRpc {
	return &defaultCollectionRpc{
		cli: cli,
	}
}

func (m *defaultCollectionRpc) SearchCat(ctx context.Context, in *SearchCatReq, opts ...grpc.CallOption) (*SearchCatResp, error) {
	client := pb2.NewCollectionRpcClient(m.cli.Conn())
	return client.SearchCat(ctx, in, opts...)
}

func (m *defaultCollectionRpc) ListCat(ctx context.Context, in *ListCatReq, opts ...grpc.CallOption) (*ListCatResp, error) {
	client := pb2.NewCollectionRpcClient(m.cli.Conn())
	return client.ListCat(ctx, in, opts...)
}

func (m *defaultCollectionRpc) RetrieveCat(ctx context.Context, in *RetrieveCatReq, opts ...grpc.CallOption) (*RetrieveCatResp, error) {
	client := pb2.NewCollectionRpcClient(m.cli.Conn())
	return client.RetrieveCat(ctx, in, opts...)
}

func (m *defaultCollectionRpc) CreateCat(ctx context.Context, in *CreateCatReq, opts ...grpc.CallOption) (*CreateCatResp, error) {
	client := pb2.NewCollectionRpcClient(m.cli.Conn())
	return client.CreateCat(ctx, in, opts...)
}

func (m *defaultCollectionRpc) UpdateCat(ctx context.Context, in *UpdateCatReq, opts ...grpc.CallOption) (*UpdateCatResp, error) {
	client := pb2.NewCollectionRpcClient(m.cli.Conn())
	return client.UpdateCat(ctx, in, opts...)
}

func (m *defaultCollectionRpc) DeleteCat(ctx context.Context, in *DeleteCatReq, opts ...grpc.CallOption) (*DeleteCatResp, error) {
	client := pb2.NewCollectionRpcClient(m.cli.Conn())
	return client.DeleteCat(ctx, in, opts...)
}

func (m *defaultCollectionRpc) CreateImage(ctx context.Context, in *CreateImageReq, opts ...grpc.CallOption) (*CreateImageResp, error) {
	client := pb2.NewCollectionRpcClient(m.cli.Conn())
	return client.CreateImage(ctx, in, opts...)
}

func (m *defaultCollectionRpc) DeleteImage(ctx context.Context, in *DeleteImageReq, opts ...grpc.CallOption) (*DeleteImageResp, error) {
	client := pb2.NewCollectionRpcClient(m.cli.Conn())
	return client.DeleteImage(ctx, in, opts...)
}

func (m *defaultCollectionRpc) GetImageByCat(ctx context.Context, in *GetImageByCatReq, opts ...grpc.CallOption) (*GetImageByCatResp, error) {
	client := pb2.NewCollectionRpcClient(m.cli.Conn())
	return client.GetImageByCat(ctx, in, opts...)
}
