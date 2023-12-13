package logic

import (
	"context"
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-demo/service/demo/api/internal/config"
	"go-zero-demo/service/demo/api/internal/svc"
	"go-zero-demo/service/demo/api/internal/types"
	"reflect"
	"testing"
)

var configFile = flag.String("f", "../../etc/demo-api-test.yaml", "the config file")
var c config.Config

func TestUserLoginLogic_UserLogin(t *testing.T) {
	conf.MustLoad(*configFile, &c)
	type fields struct {
		Logger logx.Logger
		ctx    context.Context
		svcCtx *svc.ServiceContext
	}
	type args struct {
		req *types.UserLoginReq
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *types.UserLoginReply
		wantErr  bool
	}{
		{
			name: "fail",
			fields: fields{
				Logger: logx.WithContext(context.Background()),
				ctx:    context.Background(),
				svcCtx: svc.NewServiceContext(c),
			},
			args: args{
				req: &types.UserLoginReq{
					Email:    "",
					Password: "",
				},
			},
			wantResp: nil,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewUserLoginLogic(tt.fields.ctx, tt.fields.svcCtx)
			gotResp, err := l.UserLogin(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserLogin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("UserLogin() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
