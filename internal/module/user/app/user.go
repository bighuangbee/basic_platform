package app

import (
	"context"
	"fmt"
	pb "github.com/bighuangbee/basic-platform/api/user/v1"
	"github.com/bighuangbee/basic-platform/internal/domain"
	"github.com/bighuangbee/basic-platform/internal/module/user/service"
	pbCommon "github.com/bighuangbee/gokit/api/common/v1"
	"github.com/bighuangbee/gokit/tools"
	"github.com/bighuangbee/gokit/tools/coper"
	"github.com/go-kratos/kratos/v2/log"
)

type UserApp struct {
	pb.UnimplementedUserServer
	svc *service.UserService
	log *log.Helper
}

const cryptoStr = "B10>=F3Z^#!v"

func NewUserApp(svc *service.UserService, logHelper *log.Helper) pb.UserServer {
	return &UserApp{
		svc: svc,
		log: logHelper,
	}
}

func (s *UserApp) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginReply, error) {
	user := domain.User{}
	if err := coper.CopyFromPBMessage(&user, req); err != nil{
		s.log.Error(err)
		return nil, pbCommon.ErrorInvalidParameter("")
	}
	return s.svc.Login(ctx, req)
}

func (s *UserApp) Create(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	user := domain.User{}
	if err := coper.CopyFromPBMessage(&user, req); err != nil{
		s.log.Error(err)
		return nil, pbCommon.ErrorInvalidParameter("")
	}
	user.Status = 1
	user.Salt = tools.RandStr(8)
	user.Password = tools.MD5(tools.RandStr(10)+req.Password+user.Salt+cryptoStr)
	return s.svc.Create(ctx, &user)
}

func (s *UserApp) Update(ctx context.Context, req *pb.UpdateUserRequest) (*pb.CreateUserReply, error) {
	//user := pb.UpdateUserRequest{}

	fmt.Println("---req========", req.Id, req.Username, req.Mobile, req.Mobile == nil)
	//fmt.Println("---domain.User", user.Id, user.Username, user.Mobile, user.Mobile == nil)

	return s.svc.Update(ctx, req)
}

func (s *UserApp) Test(ctx context.Context, req *pb.LoginReq) (*pb.LoginReply, error) {
	s.log.Debug("112233")
	return &pb.LoginReply{UserId: 10086}, nil
}
