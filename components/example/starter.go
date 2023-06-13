package main

import (
	"github.com/gin-gonic/gin"
	"github.com/khaosles/gtools2/components/example/controller"
)

/*
   @File: starter.go
   @Author: khaosles
   @Time: 2023/6/13 22:31
   @Desc:
*/

func main() {
	r := gin.Default()
	r.POST("/add", controller.UserController.Add)
	r.POST("/update", controller.UserController.Update)
	r.DELETE("/:id", controller.UserController.Delete)
	r.GET("/getById", controller.UserController.GetById)
	r.GET("/getAll", controller.UserController.GetAll)

	panic(r.Run(":8000"))
	//fmt.Println(service.NewUserService().FindAll())
}
