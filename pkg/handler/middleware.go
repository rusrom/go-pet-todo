package handler

import (
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
