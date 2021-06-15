package router

import (
	"calculator/handle"
	"github.com/gin-gonic/gin"
)

func SetUpRoute() *gin.Engine  {
	r := gin.Default()
	c1 := r.Group("/C1")

	c1.GET("/calculator", handle.GetCalculator)
	return r
}

