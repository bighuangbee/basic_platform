package protocol

import (
	pbBasic "github.com/bighuangbee/basic-platform/api/basic/v1"
	pbAccount "github.com/bighuangbee/basic-platform/api/user/v1"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type PbServer struct {
	Account pbAccount.UserServer
	OpLog pbBasic.OperationLogServer
}

func (s *PbServer) RegisterHTTP(srv *http.Server) {
	pbAccount.RegisterUserHTTPServer(srv, s.Account)
	pbBasic.RegisterOperationLogHTTPServer(srv, s.OpLog)
}

func (s *PbServer) RegisterRPC(srv *grpc.Server) {
	pbAccount.RegisterUserServer(srv, s.Account)
	pbBasic.RegisterOperationLogServer(srv, s.OpLog)
}
