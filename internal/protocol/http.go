package protocol

import (
	"context"
	v1 "github.com/bighuangbee/basic-platform/api/user/v1"
	"github.com/bighuangbee/basic-platform/internal/conf"
	"github.com/bighuangbee/basic-platform/internal/data"
	"github.com/bighuangbee/basic-platform/internal/pkg/middleware"
	kitKratos "github.com/bighuangbee/gokit/kratos"
	"github.com/bighuangbee/gokit/userAccess"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	"time"
)

func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})
	//whiteList["/api.user.v1.User/Login"] = struct{}{}
	whiteList[v1.OperationUserLogin] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

// NewHTTPServer new a HTTP server.
func NewHTTPServer(bc *conf.Bootstrap, logger log.Logger, server *PbServer, data *data.Data, opLog *middleware.OpLog, access userAccess.IUserAccess) *http.Server {

	c := bc.Server
	srv := http.NewServer(
		http.Address(c.HTTP.Addr),
		http.Timeout(time.Duration(c.HTTP.Timeout)*time.Second),
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(logger),
			validate.Validator(),
			opLog.Save(),
			selector.Server(userAccess.CheckToken(access)).Match(NewWhiteListMatcher()).Build(),
		),

		kitKratos.SuccessEncoder(),
		kitKratos.ErrorEncoder(),
	)
	openAPIhandler := openapiv2.NewHandler()
	srv.HandlePrefix("/q/", openAPIhandler)
	server.RegisterHTTP(srv)

	return srv
}
