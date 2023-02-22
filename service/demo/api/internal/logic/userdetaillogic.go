package logic

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"

	"go-zero-demo/common/xerr"
	"go-zero-demo/service/demo/api/internal/svc"
	"go-zero-demo/service/demo/api/internal/types"
	"go-zero-demo/service/demo/model"
)

type UserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetailLogic) UserDetail(req *types.UserDetailReq) (resp *types.UserDetailReply, err error) {

	// 用户是否存在
	user, err2 := l.svcCtx.UserModel.FindOne(l.ctx, req.ID)
	userId := l.ctx.Value("userId")

	switch err2 {
	case nil:
		switch userId.(type) {
		case json.Number:
			n, _ := userId.(json.Number).Int64()
			if n != user.Id {
				err = errors.Wrapf(xerr.NewErrMsg("您无权查看其他的用户详情"), "您无权查看其他的用户详情 %d", req.ID)
				return
			}
		default:
			err = errors.Wrapf(xerr.NewErrMsg("您无权查看用户详情"), "您无权查看用户详情 %v", userId)
			return
		}

		resp = &types.UserDetailReply{
			ID:    user.Id,
			Name:  user.Name,
			Email: user.Email,
		}
	case model.ErrNotFound:
		err = errors.Wrapf(xerr.NewErrMsg("用户不存在"), "用户不存在 %d", req.ID)
	default:
		err = errors.Wrapf(xerr.NewErrMsg("查询用户失败"), "查询用户失败 %+v", err)
	}

	return
}
