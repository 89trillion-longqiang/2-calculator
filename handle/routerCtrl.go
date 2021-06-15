package handle

import "github.com/gin-gonic/gin"

func GetCalculator(c *gin.Context){
	uri := c.Request.RequestURI///获取url
	ret := HandleGetCalculator(uri)

	c.JSON(200,gin.H{
		"condition": ret["condition"],
		"result":   ret["result"],
	})
}
