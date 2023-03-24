package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/brpaz/echozap"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/sondevelopgit/DnF-Project/config"
	"github.com/sondevelopgit/DnF-Project/util"
	"go.uber.org/zap"
)

func init() {
	if config.ConfigInstance().Debug {
		_ = util.LoggerInstance().Level().Enabled(zap.DebugLevel)
	} else {
		_ = util.LoggerInstance().Level().Enabled(zap.ErrorLevel)
	}
}

func main() {
	app := echo.New()

	// 기본 서버 설정
	app.HideBanner = true
	app.HidePort = true
	app.Debug = config.ConfigInstance().Debug

	// 기본 미들웨어 설정
	app.Use(echozap.ZapLogger(util.LoggerInstance()))
	app.Use(echoMiddleware.Recover())
	app.Use(echoMiddleware.RequestID())
	app.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins:  config.ConfigInstance().AllowOrigins,
		MaxAge:        int(time.Hour.Seconds() * 12),
		ExposeHeaders: []string{"*"},
	}))
	app.Use(echoMiddleware.BodyLimit("8M"))

	app.GET("/", func(c echo.Context) error {
		fmt.Println("default")
		return c.String(http.StatusOK, "Hello, World!")
	})

	app.Logger.Fatal(app.Start(":" + config.ConfigInstance().ServerPort))
}
