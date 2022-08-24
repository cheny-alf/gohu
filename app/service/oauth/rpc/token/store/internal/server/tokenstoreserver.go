// Code generated by goctl. DO NOT EDIT!
// Source: token_store.proto

package server

import (
	"context"

	"main/app/service/oauth/rpc/token/store/internal/logic"
	"main/app/service/oauth/rpc/token/store/internal/svc"
	"main/app/service/oauth/rpc/token/store/pb"
)

type TokenStoreServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedTokenStoreServer
}

func NewTokenStoreServer(svcCtx *svc.ServiceContext) *TokenStoreServer {
	return &TokenStoreServer{
		svcCtx: svcCtx,
	}
}

func (s *TokenStoreServer) StoreToken(ctx context.Context, in *pb.StoreTokenReq) (*pb.StoreTokenRes, error) {
	l := logic.NewStoreTokenLogic(ctx, s.svcCtx)
	return l.StoreToken(in)
}

func (s *TokenStoreServer) CheckToken(ctx context.Context, in *pb.CheckTokenReq) (*pb.CheckTokenRes, error) {
	l := logic.NewCheckTokenLogic(ctx, s.svcCtx)
	return l.CheckToken(in)
}

func (s *TokenStoreServer) GetToken(ctx context.Context, in *pb.GetTokenReq) (*pb.GetTokenRes, error) {
	l := logic.NewGetTokenLogic(ctx, s.svcCtx)
	return l.GetToken(in)
}

func (s *TokenStoreServer) RemoveToken(ctx context.Context, in *pb.RemoveTokenReq) (*pb.RemoveTokenRes, error) {
	l := logic.NewRemoveTokenLogic(ctx, s.svcCtx)
	return l.RemoveToken(in)
}