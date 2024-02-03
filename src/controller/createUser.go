package controller

import (
	restError "github.com/Gomes34/CRUD/src/logger"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	err := restError.NewBadRequest("Request error")

	c.JSON(err.Code, err)
}
