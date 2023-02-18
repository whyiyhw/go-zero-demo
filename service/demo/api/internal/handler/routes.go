// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"go-zero-demo/service/demo/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AccessLog},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/api/user/register",
					Handler: UserRegisterHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/api/user/login",
					Handler: UserLoginHandler(serverCtx),
				},
			}...,
		),
	)
}