// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/bighuangbee/basic-platform/internal/conf"
	"github.com/bighuangbee/basic-platform/internal/data"
	"github.com/bighuangbee/basic-platform/internal/module/operationLog"
	"github.com/bighuangbee/basic-platform/internal/module/user"
	"github.com/bighuangbee/basic-platform/internal/pkg/middleware"
	"github.com/bighuangbee/basic-platform/internal/protocol"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

func autoWireApp(*conf.Bootstrap, log.Logger, *log.Helper, *middleware.OpLog) (*kratos.App, func(), error) {
	panic(wire.Build(data.ProviderSet, protocol.ProviderSet, newApp,
		user.ProviderSet,
		operationLog.ProviderSet,
	))
}
