package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authHeader = "Authorization"
	userCtx    = "userId"
)

func (h *TodoHandler) userIdentity(c *gin.Context) {
	authHeader := c.GetHeader(authHeader)
	if authHeader == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty authorization header")
		return
	}

	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid authorization header")
		return
	}

	userId, err := h.services.UserAuthorization.ParseJWT(authHeaderParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
	}

	c.Set(userCtx, userId)
}

func getAuthUserId(ctx *gin.Context) (int, error) {
	val, ok := ctx.Get(userCtx)
	if !ok {
		message := "user id not found in context"
		newErrorResponse(ctx, http.StatusUnauthorized, message)
		return 0, errors.New(message)
	}

	idInt, ok := val.(int)
	if !ok {
		message := "auth user id has invalid type"
		newErrorResponse(ctx, http.StatusInternalServerError, message)
		return 0, errors.New(message)
	}

	return idInt, nil
}
