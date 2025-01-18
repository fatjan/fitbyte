package handlers

import (
	"github.com/fatjan/fitbyte/internal/config"
	"github.com/fatjan/fitbyte/internal/pkg/jwt_helper"
	authRepository "github.com/fatjan/fitbyte/internal/repositories/auth"
	duckRepo "github.com/fatjan/fitbyte/internal/repositories/duck"
	authUseCase "github.com/fatjan/fitbyte/internal/useCases/auth"
	duckUsecase "github.com/fatjan/fitbyte/internal/usecases/duck"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SetupRouter(cfg *config.Config, db *sqlx.DB, r *gin.Engine) {
	jwtMiddleware := jwt_helper.JWTMiddleware(cfg.JwtKey)

	duckRepository := duckRepo.NewDuckRepository(db)
	duckUsecase := duckUsecase.NewDuckUsecase(duckRepository)
	duckHandler := NewDuckHandler(duckUsecase)

	v1 := r.Group("v1")
	
	duckRouter := v1.Group("ducks")
	duckRouter.GET("/", duckHandler.GetDucks)
	duckRouter.GET("/:id", duckHandler.GetDuckByID)

	authRepository := authRepository.NewAuthRepository(db)
	authUseCase := authUseCase.NewUseCase(authRepository, cfg)
	authHandler := NewAuthHandler(authUseCase)

	authRouter := v1.Group("auth")
	authRouter.POST("/", authHandler.Post)
}
