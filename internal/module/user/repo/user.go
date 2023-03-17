package repo

import (
	"context"
	"fmt"
	v1 "github.com/bighuangbee/basic-platform/api/user/v1"
	"github.com/bighuangbee/basic-platform/internal/conf"
	"github.com/bighuangbee/basic-platform/internal/data"
	"github.com/bighuangbee/basic-platform/internal/domain"
	"github.com/bighuangbee/gokit/tools"
	"github.com/bighuangbee/gokit/userAccess"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
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

func (this *UserRepo) Get(ctx context.Context, account string)  (user *domain.User, err error) {
	user = &domain.User{}
	err = this.data.DB(ctx).Where("account=?",account).Where("status=1").First(&user).Error
	return
}

func (this *UserRepo) Create(ctx context.Context, data *domain.User) (error) {
	data.CreatedBy = uint64(userAccess.GetUserId(ctx))
	return this.data.DB(ctx).Create(&data).Error
}


func (this *UserRepo) Update(ctx context.Context, data *v1.UpdateUserRequest) (error) {

	updateMap := tools.PbToUpdateMap(data, &domain.User{}, userAccess.GetUserId(ctx))
	fmt.Println("==updateMap",updateMap )
	return this.data.DB(ctx).Transaction(func(tx *gorm.DB) error {
		user := domain.User{}
		if err := tx.Where("id=?", data.Id).First(&user).Error; err != nil{
			return err
		}
		//账号唯一
		return tx.Model(&domain.User{}).Where("id=?", data.Id).Where("account=?",user.Account).Updates(updateMap).Error
	})

}
