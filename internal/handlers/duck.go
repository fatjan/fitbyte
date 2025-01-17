package handlers

import (
	"github.com/fatjan/fitbyte/internal/usecases/duck"
	"github.com/gin-gonic/gin"
)

type DuckHandler interface {
	GetDucks(c *gin.Context)
	GetDuckByID(c *gin.Context)
}

type duckHandler struct {
	duckUsecase duck.Usecase
}

func NewDuckHandler(duckUsecase duck.Usecase) DuckHandler {
	return &duckHandler{
		duckUsecase: duckUsecase,
	}
}

func (h *duckHandler) GetDucks(c *gin.Context) {
	ctx := c.Request.Context()
	ducks, err := h.duckUsecase.GetDucks(ctx)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, ducks)
}

func (h *duckHandler) GetDuckByID(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	duck, err := h.duckUsecase.GetDuckByID(ctx, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, duck)
}
