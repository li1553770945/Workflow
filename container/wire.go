// +build wireinject

package container

import (
	"HomeWorkGo/infra/conf"
	"HomeWorkGo/infra/database"
	"HomeWorkGo/internal/repo"
	"HomeWorkGo/internal/service/wechat"
	"HomeWorkGo/internal/service/workflow"
	"github.com/google/wire"
)

func GetContainer(path string) *Container {
	panic(wire.Build(

		conf.NewConfig,
		database.NewMySQL,
		repo.NewRepository,

		//service
		wechat.NewWeChat,
		workflow.NewWorkFlow,

		NewContainer,
	))
}
