package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rayhan889/go-template/internal/delivery/http"
	"github.com/rayhan889/go-template/internal/delivery/http/middleware"
	"github.com/rayhan889/go-template/internal/delivery/http/route"
	"github.com/rayhan889/go-template/internal/repository"
	"github.com/rayhan889/go-template/internal/usecase"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB 			*gorm.DB
	App 		*fiber.App
	Log 		*logrus.Logger
	Validator 	*validator.Validate
	Config 		*viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
	// setup repoisitories
	userRepoRepository := repository.NewUserRepository(config.Log)

	// setup producers (if needed)

	// setup use cases
	userUsecase := usecase.NewUserUsecase(config.DB, config.Log, config.Validator, userRepoRepository)

	// setup controllers
	userController := http.NewUserController(userUsecase, config.Log)
	docsController := http.NewDocsController()

	// setup middlewares
	authMiddleware := middleware.NewAuth(userUsecase)

	routeConfig := route.RouteConfig{
		App: config.App,
		UserController: userController,
		DocsController: docsController,
		AuthMiddleware: authMiddleware,
	}

	routeConfig.Setup()
}