package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func main() {
	CreateConnection()
	fmt.Println("done")


	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/add", func(c *gin.Context) {
		data, err := ioutil.ReadAll(c.Request.Body)
		//addToMongo := CreateUser()
		//var user User
		user := &User{
			Address: Address{},
		}
		//var user map[string] interface{}
		_ = json.Unmarshal(data, user)
		addToMongo := CreateUser(user)
		fmt.Println("err", user.DateOfBirth, err, addToMongo)
		c.JSON(200, gin.H{
			"message": string(data),
		})
	})
	r.Run(":8000")
}

