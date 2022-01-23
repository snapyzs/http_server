package hanlder

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	Auth    = "Authorization"
	userCtx = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(Auth)
	if header == "" {
		newErrResp(c, http.StatusUnauthorized, "empty header auth")
		return
	}
	headerPairs := strings.Split(header, " ")
	if len(headerPairs) != 2 {
		newErrResp(c, http.StatusUnauthorized, "invalid auth header")
		return
	}
	userId, err := h.service.Auth.ParseToken(headerPairs[1])
	if err != nil {
		newErrResp(c, http.StatusUnauthorized, err.Error())
	}
	c.Set(userCtx, userId)
}
