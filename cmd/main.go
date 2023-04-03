package main

import (
	"context"
	"fmt"
	"os"
	"time"

	// @Summary 登录
	// @Description 登录
	// @Produce json
	// @Param body body controllers.LoginParams true "body参数"
	// @Success 200 {string} string "ok" "返回用户信息"
	// @Failure 400 {string} string "err_code：10002 参数错误； err_code：10003 校验错误"
	// @Failure 401 {string} string "err_code：10001 登录失败"
	// @Failure 500 {string} string "err_code：20001 服务错误；err_code：20002 接口错误；err_code：20003 无数据错误；err_code：20004 数据库异常；err_code：20005 缓存异常"
	// @Router /user/person/login [post]
	"github.com/gin-gonic/gin"
	"github.com/realtemirov/api-crud-template/api"
	_ "github.com/realtemirov/api-crud-template/api/docs"
	"github.com/realtemirov/api-crud-template/api/handler"
	"github.com/realtemirov/api-crud-template/config"
	"github.com/realtemirov/api-crud-template/services"
	"github.com/realtemirov/api-crud-template/storage/postgres"
	"github.com/rs/zerolog"
)

func main() {

	// init logger

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// create logger by zerolog

	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
		Level(zerolog.InfoLevel).
		With().
		Timestamp().
		Caller().
		Logger()

	// config
	cnf := config.NewConfig()
	logger.Info().Msg("Getting configs...")

	// context
	ctx := context.Background()
	logger.Info().Msg("Initialize context...")

	// database
	db, err := postgres.NewPostgres(ctx, cnf, logger)

	// checking for error
	if err != nil {
		logger.Printf("Method: main Comment: db.NewPostgres Error: %v", err)
	}

	// close database connection
	defer db.CloseDB()

	// services
	services := services.NewService(cnf, db, logger)
	logger.Info().Msg("Initialize services...")

	// handler
	h := handler.NewHandler(cnf, services, logger)
	logger.Info().Msg("Initialize handlers...")

	// initialize gin
	logger.Info().Msg("Initialize engine...")
	r := gin.Default()

	// routes
	api.Route(r, h)
	logger.Info().Msg("Initialize routes...")

	// engine Run
	logger.Info().Msg("Starting web...")
	r.Run(fmt.Sprintf(":%s", cnf.HostPort))
}
