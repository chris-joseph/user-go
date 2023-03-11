package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/smallcase/go-be-template/config"
	"github.com/smallcase/go-be-template/internal/api/binotto"
	"github.com/smallcase/go-be-template/pkg/log"
	"github.com/smallcase/go-be-template/pkg/store"
)

func initializeRouter(router *gin.Engine, conf *config.Config, stores Stores) error {
	log.App.Info().Msgf("Lights out and away we go on `%d` port", conf.Port)

	router.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "/healthcheck")
	})
	router.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "The key to winning a race is to finish first",
			"data":    gin.H{},
		})
	})

	// API docs will be available in all the DGNs except for production
	if conf.Environment != config.ProductionEnvironment {
		router.Static("/docs", "./docs")
	}

	binottoGroup := router.Group("/binotto")
	initializeBinottoRouter(binottoGroup, conf, stores.BinottoStore)

	return router.Run(fmt.Sprintf(":%d", conf.Port))
}

func initializeBinottoRouter(router *gin.RouterGroup, conf *config.Config, binottoStore store.BinottoStore) {
	binottoAPI := binotto.New(binottoStore, conf)

	router.GET("/:id", binottoAPI.GetByID())
	router.POST("", binottoAPI.Create())
	router.PATCH("/:id", binottoAPI.UpdateVenue())
}
