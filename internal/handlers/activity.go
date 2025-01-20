package handlers

import (
	"net/http"

	"github.com/fatjan/fitbyte/internal/dto"
	"github.com/fatjan/fitbyte/internal/pkg/exceptions"
	internal_validator "github.com/fatjan/fitbyte/internal/pkg/validator"
	"github.com/fatjan/fitbyte/internal/usecases/activity"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ActivityHandler interface {
	Get(ginCtx *gin.Context)
	Post(ginCtx *gin.Context)
}

type activityHandler struct {
	activityUseCase activity.UseCase
}

func NewActivityHandler(activityUseCase activity.UseCase) ActivityHandler {
	return &activityHandler{activityUseCase}
}

func (r *activityHandler) Post(ginCtx *gin.Context) {
	if ginCtx.GetHeader("Content-Type") != "application/json" {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "invalid content type"})
		return
	}

	var activityRequest dto.ActivityRequest
	if err := ginCtx.BindJSON(&activityRequest); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	var validate = validator.New()
	validate.RegisterValidation("iso8601", internal_validator.ValidateISODate)
	if err := validate.Struct(activityRequest); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := ginCtx.GetInt("user_id")
	activityResponse, err := r.activityUseCase.PostActivity(ginCtx.Request.Context(), &activityRequest, userId)
	if err != nil {
		ginCtx.JSON(exceptions.MapToHttpStatusCode(err), gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusCreated, activityResponse)
}

func (r *activityHandler) Get(ginCtx *gin.Context) {
	var activityRequest dto.ActivityQueryParamRequest
	if err := ginCtx.ShouldBindQuery(&activityRequest); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	userId := ginCtx.GetInt("user_id")
	activityResponses, err := r.activityUseCase.GetActivity(ginCtx.Request.Context(), &activityRequest, userId)
	if err != nil {
		ginCtx.JSON(exceptions.MapToHttpStatusCode(err), gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusCreated, activityResponses)
}
