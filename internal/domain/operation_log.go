package domain

import (
	"context"
	"github.com/bighuangbee/gokit/model"
	"github.com/bighuangbee/gokit/storage/kitGorm"
	"github.com/bighuangbee/gokit/storage/kitGorm/pagination"
)

type IOperationLogRepo interface {
	Add(ctx context.Context, oplog *OperationLog) error
	ListOperationLog(ctx context.Context, query *ListOperationLogRequest) ([]*OperationLog, int32, error)
	ListOperationLogUser(ctx context.Context, userName string) (user []*UserInfo, err error)
}

type OperationLog struct {
	model.Id
	// 企业ID
	CorpID int32 `json:"corpId"`
	// 企业名称
	CorpName string `json:"corpName"`
	// 操作人账号ID
	UserID int32 `json:"userId"`
	// 操作人姓名
	UserName string `json:"userName"`
	// 操作名称
	OperationName string `json:"operationName"`
	// 操作的模块
	OperationModule string `json:"operationModule"`
	// 操作的类型 0=未知，1=创建，2=查看，3=编辑，4=删除，5=登陆，6=登出，7=导出，8=导入，9=保存
	OperationType uint8 `json:"operationType"`
	// 自定义详细内容
	Detail string `json:"detail"`
	// 操作时间
	Timestamp *kitGorm.PBTime `json:"timestamp"`
	//数据库创建时间
	CreatedAt kitGorm.PBTime `json:"createdAt"`
	// 状态
	Status int `json:"status"`
	// 请求失败原因
	Reason string `json:"reason"`
}

func (c *OperationLog) TableName() string {
	return "operation_log"
}


type ListOperationLogRequest struct {
	Pagination     pagination.IPagination `json:"pagination,omitempty"`
	UserID         int32                    `json:"userId"`
	CorpID         int32                    `json:"corpId"`
	OperateStartAt *kitGorm.PBTime          `json:"operateStartAt"`
	OperateEndAt   *kitGorm.PBTime          `json:"operateEndAt"`
	AppID          int32                    `json:"appId"`
	OperationName  string                   `json:"operationName"`
	Status         []string                 `json:"status"`
}

type OperationLogUserResp struct {
	Items UserInfo `json:"items"`
}

type UserInfo struct {
	UserID   int32  `json:"userId"`
	UserName string `json:"userName"`
}
