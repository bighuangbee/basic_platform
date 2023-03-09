package app

import (
	"context"
	pb "github.com/bighuangbee/basic-platform/api/basic/v1"
	"github.com/bighuangbee/basic-platform/internal/domain"
	"github.com/bighuangbee/basic-platform/internal/module/operationLog/service"
	pbCommon "github.com/bighuangbee/gokit/api/common/v1"
	pagination2 "github.com/bighuangbee/gokit/storage/kitGorm/pagination"
	"github.com/bighuangbee/gokit/tools/coper"
	"github.com/go-kratos/kratos/v2/log"
)

type OperationLogApp struct {
	pb.UnimplementedOperationLogServer
	svc     *service.OperationLogService
	logger *log.Helper
}

func NewOperationLogApp(svc *service.OperationLogService, logger log.Logger) pb.OperationLogServer {
	return &OperationLogApp{
		svc:     svc,
		logger: log.NewHelper(logger),
	}
}

func (s *OperationLogApp) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddReply, error) {

	var oplog domain.OperationLog
	if err := coper.CopyFromPBMessage(&oplog, req.Log); err != nil {
		s.logger.Errorw("module", "opertionLog", "Add", err)
		return nil, pbCommon.ErrorInvalidParameter("")
	}
	err := s.svc.Add(ctx, &oplog)
	return &pb.AddReply{}, err
}

func (s *OperationLogApp) ListOperationLog(ctx context.Context, req *pb.ListOperationLogRequest) (*pb.ListOperationLogReply, error) {
	pagination, err := pagination2.New(req.Page)
	if err != nil {
		return nil, pbCommon.ErrorInvalidParameter(err.Error())
	}

	query := &domain.ListOperationLogRequest{}
	if err = coper.CopyFromPBMessage(query, req); err != nil {
		s.logger.WithContext(ctx).Error("Failed to copy to the domain.ListOperationLogRequest", err)
		return nil, pbCommon.ErrorInvalidParameter("")
	}
	query.Pagination = pagination

	res, total, err := s.svc.ListOperationLog(ctx, query)
	if err != nil {
		return nil, err
	}
	reply := &pb.ListOperationLogReply{
		Total: total,
		Items: make([]*pb.Log, len(res)),
	}
	for k, v := range res {
		detail := &pb.Log{}
		if err = coper.CopyToPBMessage(detail, &v); err != nil {
			s.logger.WithContext(ctx).Error("Failed to copy to the cloud.Log", err)
			return nil, pbCommon.ErrorInvalidParameter("")
		}
		reply.Items[k] = detail
	}

	return reply, err
}

func (s *OperationLogApp) ListOperationLogUser(ctx context.Context, req *pb.ListOperationLogUserRequest) (*pb.ListOperationLogUserReply, error) {
	apps, err := s.svc.ListOperationLogUser(ctx, req.UserName)
	if err != nil {
		return nil, err
	}

	reply := &pb.ListOperationLogUserReply{
		Items: make([]*pb.ListOperationLogUserReply_User, len(apps)),
	}

	for k, v := range apps {
		detail := &pb.ListOperationLogUserReply_User{}
		if err = coper.CopyToPBMessage(detail, v); err != nil {
			s.logger.Error("Failed to copy to the cloud.ListOperationLogUserReply_User,", err)

			return nil, pbCommon.ErrorInternalError("")
		}
		reply.Items[k] = detail
	}

	return reply, nil
}
