package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type RestHandler interface {
	SignUp(c *gin.Context)
}

func GetLoggedInUserID(c *gin.Context) (string, error) {
	uID, isExisted := c.Get("x-user-id")
	if !isExisted {
		return "", errors.New("Missing x-user-id (set at middleware) in Gin Context")
	}

	return uID.(string), nil
}
