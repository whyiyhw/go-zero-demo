package logic

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"

	"go-zero-demo/common/xerr"
	"go-zero-demo/service/demo/api/internal/svc"
	"go-zero-demo/service/demo/api/internal/types"
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
	table := l.svcCtx.UserModel.User
	user, err := table.WithContext(l.ctx).Where(table.ID.Eq(req.ID)).First()

	if err != nil {
		err = errors.Wrapf(xerr.NewErrCodeMsg(xerr.DBError, "查询用户失败"), "查询用户失败 %v", err)
		return
	}

	userId := l.ctx.Value("userId")

	if user.ID == 0 {
		err = errors.Wrapf(xerr.NewErrMsg("用户不存在"), "用户不存在 %d", req.ID)
		return
	}
	switch userId.(type) {
	case json.Number:
		n, _ := userId.(json.Number).Int64()
		if user.ID != n {
			err = errors.Wrapf(xerr.NewErrMsg("您无权查看其他的用户详情"), "您无权查看其他的用户详情 %d", req.ID)
			return
		}
	default:
		err = errors.Wrapf(xerr.NewErrMsg("您无权查看用户详情"), "您无权查看用户详情 %v", userId)
		return
	}

	resp = &types.UserDetailReply{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
	return
}
