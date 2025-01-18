package handlers

import (
	"net/http"

	"github.com/fatjan/fitbyte/internal/dto"
	"github.com/fatjan/fitbyte/internal/pkg/exceptions"
	"github.com/fatjan/fitbyte/internal/useCases/auth"
	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	Login(ginCtx *gin.Context)
	Register(ginCtx *gin.Context)
}

type authHandler struct {
	authUseCase auth.UseCase
}

func NewAuthHandler(authUseCase auth.UseCase) AuthHandler {
	return &authHandler{authUseCase: authUseCase}
}

func (r *authHandler) Register(ginCtx *gin.Context) {
	var authRequest dto.AuthRequest
	if err := ginCtx.BindJSON(&authRequest); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	authResponse, err := r.authUseCase.Register(&authRequest)
	if err != nil {
		ginCtx.JSON(exceptions.MapToHttpStatusCode(err), gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusCreated, authResponse)
}

func (r *authHandler) Login(ginCtx *gin.Context) {
	var authRequest dto.AuthRequest
	if err := ginCtx.BindJSON(&authRequest); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	authResponse, err := r.authUseCase.Login(&authRequest)
	if err != nil {
		ginCtx.JSON(exceptions.MapToHttpStatusCode(err), gin.H{"error": err.Error()})
		return
	}
	ginCtx.JSON(http.StatusOK, authResponse)
}
