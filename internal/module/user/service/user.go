package service

import (
	"context"
	"fmt"
	pb "github.com/bighuangbee/basic-platform/api/user/v1"
	"github.com/bighuangbee/basic-platform/internal/conf"
	"github.com/bighuangbee/basic-platform/internal/domain"
	commonPb "github.com/bighuangbee/gokit/api/common/v1"
	kitKratos "github.com/bighuangbee/gokit/kratos"
	"github.com/bighuangbee/gokit/storage/kitGorm"
	"github.com/go-kratos/kratos/v2/log"
)

type UserService struct {
	repo   domain.IUserRepo
	logger *log.Helper
	bc     *conf.Bootstrap
}

func NewUserService(repo domain.IUserRepo, logger log.Logger, bc *conf.Bootstrap) *UserService {
	return &UserService{
		repo:   repo,
		logger: log.NewHelper(logger),
		bc:     nil,
	}
}

func (r *UserService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginReply, error) {

	if req.Username == "" {
		//return nil, errors.New("123c")
		return nil, kitKratos.ResponseErr(ctx, commonPb.ErrorInvalidParameter)
	}

	return nil, kitKratos.ResponseErr(ctx, pb.ErrorAccountPwdError)

	return &pb.LoginReply{
		UserId: 10088,
	}, nil
}

func (r *UserService) Create(ctx context.Context, req *domain.User) (*pb.CreateUserReply, error) {
	if err := r.repo.Create(ctx, req); err != nil{
		return nil, err
	}
	return &pb.CreateUserReply{UserId: req.Id}, nil
}

func (r *UserService) Update(ctx context.Context, req *pb.UpdateUserRequest) (*pb.CreateUserReply, error) {
	if err := r.repo.Update(ctx, req); err != nil{
		if is, k, v := kitGorm.IsUniqueErr(err); is {
			return nil, commonPb.ErrorInvalidParameter(fmt.Sprintf("已存在 %s:%s", k, v))
		}
		return nil, err
	}
	return &pb.CreateUserReply{UserId: req.Id}, nil
}
