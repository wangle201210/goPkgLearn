package main

import (
	"github.com/gin-gonic/gin"
)

type data struct {
	Message string `json:"message"`
}
func main() {
	r := gin.Default()
	r.Handle("GET","/used/abc", func(c *gin.Context) {
		c.JSON(123,"abc")
	})
	r.Handle("GET","/used/*abc", func(c *gin.Context) {
		c.JSON(123,"*abc")
	})
	//v1 := r.Group("/api")
	//{
	//	v1.GET("/test", test)
	//	v1.GET("/say", say)
	//}
	r.Run(":3001")
}

func test(c *gin.Context)  {
	d := data{
		Message: "success",
	}
	c.JSON(200,d)
}

func say(c *gin.Context)  {
	str := c.Query("str")
	d := data{
		Message: str,
	}
	c.JSON(200,d)
}
