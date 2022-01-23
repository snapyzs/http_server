package hanlder

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResp struct {
	Message string `json:"message"`
}

type statusResp struct {
	Status string `json:"status"`
}

func newErrResp(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResp{message})
}
