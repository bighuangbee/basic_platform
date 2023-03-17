// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/bighuangbee/basic-platform/internal/conf"
	"github.com/bighuangbee/basic-platform/internal/data"
	app2 "github.com/bighuangbee/basic-platform/internal/module/operationLog/app"
	repo2 "github.com/bighuangbee/basic-platform/internal/module/operationLog/repo"
	service2 "github.com/bighuangbee/basic-platform/internal/module/operationLog/service"
	"github.com/bighuangbee/basic-platform/internal/module/user/app"
	"github.com/bighuangbee/basic-platform/internal/module/user/repo"
	"github.com/bighuangbee/basic-platform/internal/module/user/service"
	"github.com/bighuangbee/basic-platform/internal/pkg/middleware"
	"github.com/bighuangbee/basic-platform/internal/protocol"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "net/http/pprof"
)

// Injectors from wire.go:

func autoWireApp(bootstrap *conf.Bootstrap, logger log.Logger, helper *log.Helper, opLog *middleware.OpLog) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(bootstrap, logger)
	if err != nil {
		return nil, nil, err
	}
	iUserRepo := repo.NewUserRepo(dataData, helper, bootstrap)
	iUserAccess := service.NewUserAccess(dataData)
	userService := service.NewUserService(iUserRepo, logger, iUserAccess)
	userServer := app.NewUserApp(userService, helper)
	iOperationLogRepo := repo2.NewOperationLogRepo(dataData)
	operationLogService := service2.NewOperationLogService(iOperationLogRepo, logger)
	operationLogServer := app2.NewOperationLogApp(operationLogService, logger)
	pbServer := &protocol.PbServer{
		Account: userServer,
		OpLog:   operationLogServer,
	}
	server := protocol.NewHTTPServer(bootstrap, logger, pbServer, dataData, opLog, iUserAccess)
	grpcServer := protocol.NewGRPCServer(bootstrap, logger, pbServer)
	kratosApp := newApp(bootstrap, logger, server, grpcServer, opLog)
	return kratosApp, func() {
		cleanup()
	}, nil
}
