package handler

import (
	"fmt"
	"net/http"

	"go-zero-demo/common/response"
	"go-zero-demo/common/validator"
	"go-zero-demo/service/demo/api/internal/logic"
	"go-zero-demo/service/demo/api/internal/svc"
	"go-zero-demo/service/demo/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserRegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserRegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamError(r, w, err)
			return
		}
		// validate check
		if err := validator.Validate.StructCtx(r.Context(), req); err != nil {
			errMap := validator.Translate(err, &req)
			for _, errFormat := range errMap {
				response.ParamError(r, w, fmt.Errorf(errFormat))
				return
			}
			response.ParamError(r, w, err)
			return
		}

		l := logic.NewUserRegisterLogic(r.Context(), svcCtx)
		resp, err := l.UserRegister(&req)

		response.Response(r, w, resp, err)
	}
}
