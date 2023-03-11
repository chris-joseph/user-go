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

// swagger:route GET /pets pets listPets
//
// Lists pets filtered by some parameters.
//
// This will show all available pets by default.
// You can get the pets that are out of stock
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Schemes: http, https
//
//	Parameters:
//	  + name: limit
//	    in: query
//	    description: maximum numnber of results to return
//	    required: false
//	    type: integer
//	    format: int32
//
//	Responses:
//	  200: Binotto
func (api *BinottoAPI) GetByID() func(*gin.Context) {
	return func(g *gin.Context) {}
}

func (api *BinottoAPI) Create() func(*gin.Context) {
	return func(g *gin.Context) {}
}

func (api *BinottoAPI) UpdateVenue() func(*gin.Context) {
	return func(g *gin.Context) {}
}
