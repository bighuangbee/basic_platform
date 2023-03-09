package repo

import (
	"context"
	"github.com/bighuangbee/basic-platform/internal/data"
	"github.com/bighuangbee/basic-platform/internal/domain"
	"github.com/bighuangbee/gokit/storage/kitGorm/pagination"
)

type OperationLogRepo struct {
	data *data.Data
}

func NewOperationLogRepo(data *data.Data) domain.IOperationLogRepo {
	return &OperationLogRepo{
		data: data,
	}
}

func (r *OperationLogRepo) ListOperationLogUser(ctx context.Context, userName string) (user []*domain.UserInfo, err error) {
	return nil, err
}

func (r *OperationLogRepo) Add(ctx context.Context, oplog *domain.OperationLog) error {
	return r.data.DB(ctx).Create(oplog).Error
}

func (r *OperationLogRepo) ListOperationLog(ctx context.Context, query *domain.ListOperationLogRequest) ([]*domain.OperationLog, int32, error) {
	db := r.data.DB(ctx).Model(&domain.OperationLog{})

	if len(query.Status) != 0 {
		db.Where("status in ?", query.Status)
	}
	if query.UserID != 0 {
		db.Where("user_id = ?", query.UserID)
	}
	if query.CorpID != 0 {
		db.Where("corp_id = ?", query.CorpID)
	}
	if query.OperationName != "" {
		db.Where("operation_name Like ?", "%"+query.OperationName+"%")
	}
	if query.OperateStartAt != nil && query.OperateEndAt != nil {
		db = db.Where("timestamp between ? and ?", query.OperateStartAt, query.OperateEndAt)
	}

	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	var rows []*domain.OperationLog
	err = pagination.PageQuery(db, query.Pagination).Find(&rows).Error
	if err != nil {
		return nil, 0, err
	}
	return rows, int32(total), nil
}
