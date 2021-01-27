package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func test(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "pong",
	})
}

func main() {

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "Hello, Welcome to Gin Framework",
		})
	})

	r.GET("/ping/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"id":  id,
			"msg": "pong, using Params",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		id := c.Query("id")
		role := c.DefaultQuery("role", "0")
		c.JSON(200, gin.H{
			"id":   id,
			"msg":  "pong, using query params",
			"role": role,
		})
	})

	// r.POST("/user", func(c *gin.Context) {
	// 	name := c.PostForm("name")
	// 	role := c.PostForm("role")
	// 	c.JSON(200, gin.H{
	// 		"msg":  "Hello, this is an example for POST with FORM-DATA",
	// 		"name": name,
	// 		"role": role,
	// 	})
	// })

	type user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	r.POST("/user", func(c *gin.Context) {
		var user user
		err := c.Bind(&user)
		if err != nil {
			fmt.Println("Telah terjadi error")
		}
		c.JSON(200, gin.H{
			"msg": "Hello, this is an example for POST with Raw Json",
			"data": map[string]interface{}{
				"username":   user.Username,
				"password":   user.Password,
				"pengalaman": []string{"BE", "FE", "QA"},
			},
		})
	})

	// Grouping
	v1 := r.Group("/v1")
	// this is a scope
	{
		v1.GET("/ping", test)

		v1.GET("/ping/:id/:role", func(c *gin.Context) {
			id := c.Param("id")
			role := c.Param("role")
			c.JSON(200, gin.H{
				"id":   id,
				"msg":  "pong",
				"role": role,
			})
		})
	}

	r.Run() // listen and serve on 0.0.0.0:8080

}
