package service

import (
	"context"
	"fmt"
	pb "github.com/bighuangbee/basic-platform/api/user/v1"
	"github.com/bighuangbee/basic-platform/internal/data"
	"github.com/bighuangbee/basic-platform/internal/domain"
	commonPb "github.com/bighuangbee/gokit/api/common/v1"
	kitKratos "github.com/bighuangbee/gokit/kratos"
	"github.com/bighuangbee/gokit/storage/cache"
	"github.com/bighuangbee/gokit/storage/kitGorm"
	"github.com/bighuangbee/gokit/tools"
	"github.com/bighuangbee/gokit/tools/crypto"
	"github.com/bighuangbee/gokit/userAccess"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type UserService struct {
	repo   domain.IUserRepo
	logger *log.Helper
	userAccess userAccess.IUserAccess
}

const cryptoStr = "B10>=F3Z^#!v"
const loginExpire = time.Minute

func NewUserAccess(data *data.Data)userAccess.IUserAccess{
	return userAccess.New(
		cache.NewCacheRedis(data.Redis("user:token")),
		time.Hour,
	)
}

func NewUserService(repo domain.IUserRepo, logger log.Logger, userAccess userAccess.IUserAccess) *UserService {
	return &UserService{
		repo:   repo,
		logger: log.NewHelper(logger),
		userAccess:     userAccess,
	}
}

func (r *UserService) Login(ctx context.Context, req *domain.User) (*pb.LoginReply, error) {
	user, err := r.repo.Get(ctx, req.Account)
	if err != nil{
		r.logger.Errorw("UserService Get 账号不存在", err)
		return nil, kitKratos.ResponseErr(ctx, pb.ErrorAccountPwdError)
	}

	if user.Password != genderPassword(req.Password, user.Salt){
		r.logger.Errorw("UserService Login", "密码不正确")
		return nil, kitKratos.ResponseErr(ctx, pb.ErrorAccountPwdError)
	}

	token, err := r.userAccess.Issue(userAccess.NewUserClaims(user.UserName, req.Account, 0, fmt.Sprintf("%d",user.Id), loginExpire))
	if err != nil{
		r.logger.Errorw("UserService userAccess.Issue", "token签发失败")
		return nil, kitKratos.ResponseErr(ctx, commonPb.ErrorInternalError)
	}

	return &pb.LoginReply{
		UserId: user.Id,
		Account: user.Account,
		UserName: user.UserName,
		Token: token,
		ExpiredAt: int32(time.Now().Add(loginExpire).Unix()),
	}, nil
}

func genderPassword(password, salt string)string{
	tmp, _ := crypto.AesEncrypt([]byte(password), []byte(cryptoStr))
	return tools.MD5(salt + string(tmp)+ password + cryptoStr)
}

func (r *UserService) Create(ctx context.Context, user *domain.User) (*pb.CreateUserReply, error) {
	user.Status = 1
	user.Salt = tools.RandStr(8)
	user.Password = genderPassword(user.Password, user.Salt)
	if err := r.repo.Create(ctx, user); err != nil{
		if is, k, v := kitGorm.IsUniqueErr(err); is {
			return nil, commonPb.ErrorInvalidParameter(fmt.Sprintf("已存在 %s:%s", k, v))
		}
		return nil, err
	}
	return &pb.CreateUserReply{UserId: user.Id}, nil
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
