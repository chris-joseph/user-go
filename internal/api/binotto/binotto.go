package binotto

import (
	"github.com/gin-gonic/gin"
	"github.com/smallcase/go-be-template/config"
	"github.com/smallcase/go-be-template/pkg/store"
)

type BinottoAPI struct {
	BinottoStore store.BinottoStore
	Conf         *config.Config
}

func New(binottoStore store.BinottoStore, conf *config.Config) BinottoAPI {
	return BinottoAPI{
		BinottoStore: binottoStore,
		Conf:         conf,
	}
}

func (api *BinottoAPI) GetByID() func(*gin.Context) {
	return func(g *gin.Context) {}
}

func (api *BinottoAPI) Create() func(*gin.Context) {
	return func(g *gin.Context) {}
}

func (api *BinottoAPI) UpdateVenue() func(*gin.Context) {
	return func(g *gin.Context) {}
}
