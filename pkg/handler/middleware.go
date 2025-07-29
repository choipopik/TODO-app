package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authHeader = "Authorization"
	userCtx    = "userId"
	listCtx    = "listId"
)

func (h *Handler) userIdentity(ctx *gin.Context) {
	header := ctx.GetHeader(authHeader)
	if header == "" {
		newErrorResponse(ctx, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(ctx, http.StatusUnauthorized, "invalid auth header")
		return
	}

	//parse token
	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	ctx.Set(userCtx, userId)
}

func getUserId(ctx *gin.Context) (int, error) {
	id, ok := ctx.Get(userCtx)
	if !ok {
		newErrorResponse(ctx, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(ctx, http.StatusInternalServerError, "user id invalid type")
		return 0, errors.New("user id invalid type")
	}

	return idInt, nil
}

func getListId(ctx *gin.Context) (int, error) {
	id, ok := ctx.Get(listCtx)
	if !ok {
		newErrorResponse(ctx, http.StatusInternalServerError, "list id not found")
		return 0, errors.New("list id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(ctx, http.StatusInternalServerError, "list id invalid type")
		return 0, errors.New("list id invalid type")
	}

	return idInt, nil
}
