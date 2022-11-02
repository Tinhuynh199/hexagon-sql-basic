package app

import (
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	_ "github.com/go-sql-driver/mysql"

	"hexrestapi/internal/user/adapter/handler"
	"hexrestapi/internal/user/adapter/repository"
	"hexrestapi/internal/user/port"
	"hexrestapi/internal/user/service"
)

type App struct {
	Router *mux.Router
	User port.UserTransport
}

func (a *App) Initialize(config *Config) {
	db, err := sql.Open(config.Sql.Driver, config.Sql.Data_Source)
	if err != nil {
		log.Fatal("Could not connect to database !!!")
	}

	userRepository := repository.NewUserAdapter(db)
	userService := service.NewUserService(db, userRepository)
	userHandler := handler.NewUserHandler(userService)
	
	a.User = userHandler
	a.Router = mux.NewRouter()
	a.setRouters()
}
