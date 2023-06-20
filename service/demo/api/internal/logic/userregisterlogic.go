package logic

import (
	"context"
	"gorm.io/gorm"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"

	"go-zero-demo/common/xerr"
	"go-zero-demo/service/demo/api/internal/svc"
	"go-zero-demo/service/demo/api/internal/types"
	"go-zero-demo/service/demo/model"

	"golang.org/x/crypto/bcrypt"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterReq) (resp *types.UserRegisterReply, err error) {
	// 判断用户是否已经注册
	table := l.svcCtx.UserModel.User
	exist, err := table.WithContext(l.ctx).Where(table.Email.Eq(req.Email)).First()

	if !errors.Is(err, gorm.ErrRecordNotFound) && err != nil {
		return nil, errors.Wrapf(xerr.NewErrCodeMsg(xerr.DBError, "查询用户失败"), "查询用户失败 %v", err)
	} else {
		if exist != nil && exist.ID > 0 {
			return nil, errors.Wrapf(xerr.NewErrMsg("用户已经注册"), "用户已经注册 %d", exist.ID)
		}
	}

	// 加密密码
	password := []byte(req.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("密码加密失败"), "密码加密失败 %v", err)
	}

	// 未注册的用户进行注册操作
	if err := table.WithContext(l.ctx).Create(&model.User{
		Email:    req.Email,
		Name:     req.Name,
		Password: string(hashedPassword),
	}); err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("用户注册失败"), "用户注册失败 %v", err)
	}

	return &types.UserRegisterReply{Message: "注册成功，去登录吧~"}, nil
}
