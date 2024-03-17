package services

import (
	"context"
	"strconv"

	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
)

type HelloWorldService struct {
	BaseService
}

func (this *HelloWorldService) SayHello(ctx context.Context, req *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	// 校验用户权限
	_, _, err := this.ValidateAdminAndUser(ctx, true)
	if err != nil {
		return nil, err
	}

	// 返回信息
	return &pb.SayHelloResponse{
		Result: "Id: " + strconv.FormatInt(req.Id, 10) + ", Name: " + req.Name,
	}, nil
}
