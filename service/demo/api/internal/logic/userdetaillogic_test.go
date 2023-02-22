package logic

import (
	"context"
	"encoding/json"
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"go-zero-demo/service/demo/api/internal/config"
	"go-zero-demo/service/demo/api/internal/svc"
	"go-zero-demo/service/demo/api/internal/types"
	"reflect"
	"testing"
)

var configFile = flag.String("f", "../../etc/demo-api.yaml", "the config file")

func TestUserDetailLogic_UserDetail(t *testing.T) {
	var c config.Config
	conf.MustLoad(*configFile, &c)
	type fields struct {
		ctx    context.Context
		svcCtx *svc.ServiceContext
	}
	type args struct {
		req *types.UserDetailReq
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantResp *types.UserDetailReply
		wantErr  bool
	}{
		{
			name: "success",
			fields: fields{
				ctx:    context.WithValue(context.Background(), "userId", json.Number("2")),
				svcCtx: svc.NewServiceContext(c),
			},
			args: args{
				req: &types.UserDetailReq{
					ID: 2,
				},
			},
			wantResp: &types.UserDetailReply{
				ID:    2,
				Email: "demo@163.com",
				Name:  "demo",
			},
			wantErr: false,
		},
		{
			name: "error-without-permission",
			fields: fields{
				ctx:    context.WithValue(context.Background(), "userId", json.Number("1")),
				svcCtx: svc.NewServiceContext(c),
			},
			args: args{
				req: &types.UserDetailReq{
					ID: 2,
				},
			},
			wantResp: nil,
			wantErr:  true,
		},
		{
			name: "error-without-no-permission",
			fields: fields{
				ctx:    context.WithValue(context.Background(), "userId", 1),
				svcCtx: svc.NewServiceContext(c),
			},
			args: args{
				req: &types.UserDetailReq{
					ID: 2,
				},
			},
			wantResp: nil,
			wantErr:  true,
		},
		{
			name: "test",
			fields: fields{
				ctx:    context.Background(),
				svcCtx: svc.NewServiceContext(c),
			},
			args: args{
				req: &types.UserDetailReq{
					ID: -1,
				},
			},
			wantResp: nil,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			l := NewUserDetailLogic(tt.fields.ctx, tt.fields.svcCtx)

			gotResp, err := l.UserDetail(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("UserDetail() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
