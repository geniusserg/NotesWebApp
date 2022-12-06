package app

import (
	"awesomeProject/app/internal/config"
	"awesomeProject/app/pkg/logging"
	httpSwagger "http-swagger"
	"httprouter"
	"net/http"
)

type App struct {
	config *config.Config
	logger *logging.Logger
}

func NewApp(config *config.Config, logger *logging.Logger) (App, error) {
	logger.Println("router init")
	router := httprouter.New()

	logger.Println("Swagger")
	router.Handler(http.MethodGet, "/swagger", http.RedirectHandler("/swagger/index.html", http.StatusMovedPermanently))
	router.Handler(http.MethodGet, "/swagger/*any", httpSwagger.WrapHandler)
	app := App{
		config: config,
		logger: logger,
	}

	return app, nil
}
