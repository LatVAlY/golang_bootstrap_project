package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang_bootstrap_project/environment"
	"golang_bootstrap_project/handler"
	"golang_bootstrap_project/helper"
	"golang_bootstrap_project/repository"
	"golang_bootstrap_project/service"
	"golang_bootstrap_project/utils"
	"net/http"
	"os"
)

func setGeneralDefaults() {
	viper.AutomaticEnv()
	viper.SetDefault(environment.ServicePort, 4000)
	viper.SetDefault(environment.LogLevel, 5)
}

func main() {
	setGeneralDefaults()

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.Level(viper.GetUint32(environment.LogLevel)))

	e := echo.New()

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
	userHandler := handler.UserHandlerImpl{
		UserService: service.UserServiceImpl{
			UserRepository: repository.UserRepoImpl{
				Db:         nil,
				UserHelper: helper.UserHelperImpl{},
			},
		},
		StringUtils: utils.StringUtilsImpl{},
	}

	g := e.Group("/v1")
	g.GET("/users/:userId", userHandler.User)
	defer func(e *echo.Echo) {
		closeErr := e.Close()
		if closeErr != nil {
			log.WithError(closeErr).Errorf("the connection to echo server could not be closed")
		}
	}(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", viper.GetInt(environment.ServicePort))))
}
