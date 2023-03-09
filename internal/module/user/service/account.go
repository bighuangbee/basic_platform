package service

import (
	"context"
	pb "github.com/bighuangbee/basic-platform/api/account/v1"
	"github.com/bighuangbee/basic-platform/internal/conf"
	"github.com/bighuangbee/basic-platform/internal/domain"
	commonPb "github.com/bighuangbee/gokit/api/common/v1"
	kitKratos "github.com/bighuangbee/gokit/kratos"
	"github.com/go-kratos/kratos/v2/log"
)

type UserService struct {
	repo   domain.IUserRepo
	logger *log.Helper
	bc     *conf.Bootstrap
}

func NewUserService(repo domain.IUserRepo, logger log.Logger, bc *conf.Bootstrap) *UserService {
	return &UserService{
		repo:   nil,
		logger: log.NewHelper(logger),
		bc:     nil,
	}
}

func (this *UserService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginReply, error) {

	if req.Username == "" {
		//return nil, errors.New("123c")
		return nil, kitKratos.ResponseErr(ctx, commonPb.ErrorInvalidParameter)
	}

	return nil, kitKratos.ResponseErr(ctx, pb.ErrorAccountPwdError)

	return &pb.LoginReply{
		UserId: 10088,
	}, nil
}
