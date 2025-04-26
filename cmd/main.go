package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"urlshortener/internal/config"
	"urlshortener/internal/service"
	"urlshortener/pkg/logger"
	"urlshortener/pkg/postgres"
	//"urlshortener/pkg/postgres"
)

func main() {
	ctx := context.Background()
	ctx, _ = logger.New(ctx)
	cfg, err := config.New()
	if err != nil {
		logger.GetLoggerFromCtx(ctx).Fatal(ctx, "config.New error", zap.Error(err))
	}
	pool, err := postgres.New(ctx, cfg.Postgres)
	if err != nil {
		logger.GetLoggerFromCtx(ctx).Fatal(ctx, "postgres.New error", zap.Error(err))
	}

	router := gin.Default()
	serv := service.NewService(router, pool)
	serv.RunService()

	fmt.Println("service is running...")
	port := fmt.Sprintf(":%s", cfg.RESTPort)
	router.Run(port)

}
