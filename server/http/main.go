package main

import (
	"TechnicalTestKStyleHub/app/constants"
	"TechnicalTestKStyleHub/domain"
	"TechnicalTestKStyleHub/domain/presenters"
	"TechnicalTestKStyleHub/server"
	memberHandlers "TechnicalTestKStyleHub/src/member/delivery/http/handlers"
	memberRoutes "TechnicalTestKStyleHub/src/member/delivery/http/routes"
	productHandlers "TechnicalTestKStyleHub/src/product/delivery/http/handlers"
	productRoutes "TechnicalTestKStyleHub/src/product/delivery/http/routes"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
)

func main() {

	// Load Configuration
	config, err := domain.LoadConfiguration()
	if err != nil {
		log.Fatal("Error while load configuration, ", err.Error())
	}
	defer config.Sql.Connection.Close()

	// init setup fiber
	app := initEcho()

	// init router
	initRouter(app, config)

	// Listening Http
	if err = app.Start(fmt.Sprintf(":%s", os.Getenv(constants.EnvironmentAppRestPort))); err != nil {
		log.Fatal("Error while listening http protocol, ", err.Error())
	}
}

func initEcho() *echo.Echo {
	// Create a new Echo instance
	e := echo.New()

	// Validator
	validate := validator.New()

	e.Validator = &presenters.CustomValidator{Validator: validate}

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           `{"time":"${time_rfc3339}","method":"${method}","uri":"${uri}","status":${status},"latency_human":"${latency_human}"}` + "\n",
		CustomTimeFormat: "2006-01-02T15:04:05.000Z07:00",
	}))

	log.Println("origin: ", os.Getenv(constants.ConstantOriginCORS))

	return e
}

// @Title          Technical Test KStyle Hub API Backend
// @Version        1.0
// @Description    Technical Test KStyle Hub Backend API docs
// @Contact.name   Fawzi Linggo
// @Contact.email  fawzilinggo@google.com
// @Host           localhost
// @Schemes		   http
// @SecurityDefinitions.apiKey JWT
// @in Cookie
// @name jwt
// @BasePath       /
func initRouter(e *echo.Echo, config domain.Config) {
	// Router for health check
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, "OK")
	})

	// register handler
	registerHandler(e, server.InitContract(config))
}

func registerHandler(app *echo.Echo, contract server.Contract) {

	// member
	memberHandler := memberHandlers.NewMemberHandler(contract.IMemberUseCase)
	memberRoutes.NewMemberRoute(app, memberHandler).RegisterRoutes()

	// product
	productHandler := productHandlers.NewProductHandler(contract.IProductUseCase)
	productRoutes.NewProductRoute(app, productHandler).RegisterRoutes()

}