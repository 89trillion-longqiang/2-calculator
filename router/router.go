package router

import (
	"calculator/ctrl"
	"github.com/gin-gonic/gin"
)

func SetUpRoute() *gin.Engine  {
	r := gin.Default()
	c1 := r.Group("/C1")

	c1.GET("/calculator", ctrl.GetCalculator)
	return r
}

