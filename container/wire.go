// +build wireinject

package container

import (
	"github.com/google/wire"
	"workflow_http/infra/conf"
	"workflow_http/infra/database"
	"workflow_http/infra/wechat_bot"
	"workflow_http/internal/repo"
	"workflow_http/internal/service/wechat"
	"workflow_http/internal/service/workflow"
)

func GetContainer(path string) *Container {
	panic(wire.Build(

		//infra
		conf.NewConfig,
		database.NewMySQL,
		wechat_bot.NewBot,
		//repo
		repo.NewRepository,

		//service
		wechat.NewWeChat,
		workflow.NewWorkFlow,

		NewContainer,
	))
}
