// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"main/app/service/oauth/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/oauth/token/get",
				Handler: GetTokenHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/oauth/token/refresh",
				Handler: RefreshTokenHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/oauth/token/check",
				Handler: CheckTokenHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/oauth/token/read",
				Handler: ReadTokenHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/oauth/token/get/user",
				Handler: GetUserInfoHandler(serverCtx),
			},
		},
	)
}
