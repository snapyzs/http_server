package hanlder

import (
	"github.com/gin-gonic/gin"
	"http_server/model"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var user model.Users
	if err := c.BindJSON(&user); err != nil {
		newErrResp(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := user.Validate(); err != nil {
		newErrResp(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.service.Auth.CreateUser(user)
	if err != nil {
		newErrResp(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var user model.Users
	if err := c.BindJSON(&user); err != nil {
		newErrResp(c, http.StatusBadRequest, err.Error())
		return
	}
	token, err := h.service.Auth.GenerateToken(user.Email, user.Password)
	if err != nil {
		newErrResp(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
