package repo

import (
	"context"
	"github.com/bighuangbee/basic-platform/internal/conf"
	"github.com/bighuangbee/basic-platform/internal/data"
	"github.com/bighuangbee/basic-platform/internal/domain"
	"github.com/go-kratos/kratos/v2/log"
)


func NewUserRepo(data *data.Data, logHelper *log.Helper, bootstrap *conf.Bootstrap) domain.IUserRepo {
	return &UserRepo{
		data:   data,
		logHelper: logHelper,
		bc:     bootstrap,
	}
}

type UserRepo struct {
	data   *data.Data
	logHelper *log.Helper
	bc     *conf.Bootstrap
}

func (this *UserRepo) Login(context.Context, *domain.User)  (*domain.User, error) {
	return &domain.User{}, nil
}

func (this *UserRepo) CreateUser(context.Context, *domain.User)  (*domain.User, error) {
	return &domain.User{}, nil
}
