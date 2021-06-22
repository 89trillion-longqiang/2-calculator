package ctrl

import (
	"calculator/handle"
	"github.com/gin-gonic/gin"
)

func GetCalculator(c *gin.Context){
	uri := c.Request.RequestURI///获取url
	if uri == "" {
		c.JSON(200,gin.H{
			"condition": "parameter error",
		})
		return
	}

	ret := handle.HandleGetCalculator(uri)

	c.JSON(200,gin.H{
		"condition": ret["condition"],
		"result":   ret["result"],
	})
}
