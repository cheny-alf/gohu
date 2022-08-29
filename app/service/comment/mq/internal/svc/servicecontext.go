package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"main/app/service/comment/mq/internal/config"
	"main/app/service/comment/rpc/crud/crud"
)

type ServiceContext struct {
	Config config.Config

	CrudRpcClient crud.Crud
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		CrudRpcClient: crud.NewCrud(zrpc.MustNewClient(c.CrudRpcClientConf)),
	}
}
