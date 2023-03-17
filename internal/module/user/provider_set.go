package user

import (
	"github.com/bighuangbee/basic-platform/internal/module/user/app"
	"github.com/bighuangbee/basic-platform/internal/module/user/repo"
	"github.com/bighuangbee/basic-platform/internal/module/user/service"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(app.NewUserApp, service.NewUserAccess, service.NewUserService, repo.NewUserRepo)
