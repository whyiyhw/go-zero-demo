package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/conf"
	"go-zero-demo/service/demo/api/internal/svc"
	"go-zero-demo/service/demo/api/internal/types"
	"reflect"
	"testing"
)

func TestUserDetailLogic_UserDetail(t *testing.T) {
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
				ctx:    context.WithValue(context.Background(), "userId", json.Number("1")),
				svcCtx: svc.NewServiceContext(c),
			},
			args: args{
				req: &types.UserDetailReq{
					ID: 1,
				},
			},
			wantResp: &types.UserDetailReply{
				ID:    1,
				Email: "demo@163.com",
				Name:  "demo",
			},
			wantErr: false,
		},
		{
			name: "error-without-user",
			fields: fields{
				ctx:    context.WithValue(context.Background(), "userId", json.Number("10")),
				svcCtx: svc.NewServiceContext(c),
			},
			args: args{
				req: &types.UserDetailReq{
					ID: 10,
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
