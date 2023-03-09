package app

import (
	"context"
	pb "github.com/bighuangbee/basic-platform/api/account/v1"
	"github.com/bighuangbee/basic-platform/internal/module/user/service"
	"github.com/go-kratos/kratos/v2/log"
)

type UserApp struct {
	pb.UnimplementedAccountServer
	svc    *service.UserService
	logHelper *log.Helper
}

func NewUserApp(svc *service.UserService, logHelper *log.Helper) pb.AccountServer {

	return &UserApp{
		svc:    svc,
		logHelper: logHelper,
	}
}

func (s *UserApp) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginReply, error) {
	return s.svc.Login(ctx, req)
}

func (s *UserApp) Test(ctx context.Context, req *pb.LoginReq) (*pb.LoginReply, error) {
	s.logHelper.Debug("112233")
	return &pb.LoginReply{UserId: 10086}, nil
}
