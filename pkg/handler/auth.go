package handler

import (
	"github.com/gin-gonic/gin"
	todo "github.com/rusrom/yt-todo"
	"net/http"
)

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body todo.User true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up [post]
func (h *TodoHandler) signUp(c *gin.Context) {
	var newUser todo.User

	err := c.BindJSON(&newUser)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateUser(newUser)
	if err != nil {
		//newErrorResponse(c, http.StatusInternalServerError, err.Error())
		newErrorResponse(c, http.StatusBadRequest, "such username already exists")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary SignIn
// @Tags auth
// @Description login
// @ID login
// @Accept  json
// @Produce  json
// @Param input body todo.SignInInput true "credentials"
// @Success 200 {string} string "token"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-in [post]
func (h *TodoHandler) signIn(c *gin.Context) {
	var userCredentials todo.SignInInput

	err := c.BindJSON(&userCredentials)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.GenerateJWT(userCredentials)
	if err != nil {
		//newErrorResponse(c, http.StatusInternalServerError, err.Error())
		newErrorResponse(c, http.StatusBadRequest, "not valid credentials")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
