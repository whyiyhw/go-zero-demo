package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-demo/service/demo/api/internal/svc"
	"go-zero-demo/service/demo/api/internal/types"
	"reflect"
	"testing"
)

func TestUserRegisterLogic_UserRegister(t *testing.T) {
	conf.MustLoad(*configFile, &c)
	type fields struct {
		Logger logx.Logger
		ctx    context.Context
		svcCtx *svc.ServiceContext
	}
	type args struct {
		req *types.UserRegisterReq
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *types.UserRegisterReply
		wantErr  bool
	}{
		{
			name: "success",
			fields: fields{
				Logger: logx.WithContext(context.Background()),
				ctx:    context.Background(),
				svcCtx: svc.NewServiceContext(c),
			},
			args: args{
				req: &types.UserRegisterReq{
					Email:    "demo@163.com",
					Name:     "demo",
					Password: "demo@163.com",
				},
			},
			wantResp: &types.UserRegisterReply{
				Message: "注册成功，去登录吧~",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewUserRegisterLogic(tt.fields.ctx, tt.fields.svcCtx)
			gotResp, err := l.UserRegister(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRegister() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("UserRegister() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
