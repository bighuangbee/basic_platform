package operationLog

import (
	"github.com/bighuangbee/basic-platform/internal/module/operationLog/app"
	"github.com/bighuangbee/basic-platform/internal/module/operationLog/repo"
	"github.com/bighuangbee/basic-platform/internal/module/operationLog/service"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	app.NewOperationLogApp, service.NewOperationLogService, repo.NewOperationLogRepo,
)
