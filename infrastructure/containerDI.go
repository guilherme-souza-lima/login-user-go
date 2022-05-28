package infrastructure

import (
	"gorm.io/gorm"
	"loginUserGo/infrastructure/database"
	"loginUserGo/infrastructure/database/db_postgres"
	"loginUserGo/infrastructure/migrations"
	"loginUserGo/user_case/handler"
	"loginUserGo/user_case/repository"
	"loginUserGo/user_case/service"
)

type ContainerDI struct {
	Config         Config
	DB             *gorm.DB
	Migration      migrations.DatabaseMakeMigrations
	UserRepository repository.UserRepository
	UserService    service.UserService
	UserHandler    handler.UserHandler
}

func NewContainerDI(config Config) *ContainerDI {
	container := &ContainerDI{
		Config: config,
	}

	configDB := database.Config{
		Hostname: container.Config.Host,
		Port:     container.Config.Port,
		UserName: container.Config.User,
		Password: container.Config.Password,
		Database: container.Config.Database,
	}

	container.DB = db_postgres.InitGorm(&configDB)
	container.Migration = migrations.NewDatabaseMakeMigrations(container.DB)

	container.build()
	return container
}

func (c *ContainerDI) build() {
	c.UserRepository = repository.NewUserRepository(c.DB)
	c.UserService = service.NewUserService(c.UserRepository)
	c.UserHandler = handler.NewUserHandler(c.UserService)
}

func (c *ContainerDI) ShutDown() {}
