package main

import (
	_ "gin-rest-api-ioc/dao"
	service "gin-rest-api-ioc/service"
	"github.com/gin-gonic/gin"
)




func main() {
	r := gin.Default()

	userService:=service.New();

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	});

	r.GET("/user/:id", func(c *gin.Context) {
		var user,err=userService.GetUserById(c.Param("id"));
		if err==nil {
			c.JSON(200,user);
		}else{
			c.JSON(400,err);
		}
	});

	r.POST("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	});


	r.Run() // listen and serve on 0.0.0.0:8080
}
