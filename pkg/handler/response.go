package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(ctx *gin.Context, statusCode int, msg string) {
	logrus.Error(msg)
	ctx.AbortWithStatusJSON(statusCode, errorResponse{msg})
}
