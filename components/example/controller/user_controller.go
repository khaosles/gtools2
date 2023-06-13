package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/khaosles/gtools2/components/example/model"
	"github.com/khaosles/gtools2/components/example/service"
	"github.com/khaosles/gtools2/components/g"
)

/*
   @File: user_controller.go
   @Author: khaosles
   @Time: 2023/6/13 10:07
   @Desc:
*/

var UserController = new(userController)

type userController struct {
	userService service.UserService
}

func init() {
	UserController = &userController{userService: service.NewUserService()}
}

func (ctl userController) Add(c *gin.Context) {
	var entity model.User
	err := c.ShouldBindJSON(&entity)
	if err != nil {
		g.NewResult(c).No(g.SYSTEM_ERROR)
		return
	}
	ctl.userService.Save(&entity)
	g.NewResult(c).Yes(nil)
}

func (ctl userController) Update(c *gin.Context) {
	var entity model.User
	err := c.ShouldBindJSON(&entity)
	if err != nil {
		g.NewResult(c).No(g.SYSTEM_ERROR)
		return
	}
	ctl.userService.Update(&entity)
	g.NewResult(c).Yes(nil)
}

func (ctl userController) DeleteById(c *gin.Context) {
	id := c.Param("id")
	ctl.userService.DeleteById(id)
	g.NewResult(c).Yes(nil)
}

func (ctl userController) FIndAll(c *gin.Context) {
	entities := ctl.userService.FindAll()
	g.NewResult(c).Yes(entities)
}

func (ctl userController) FindById(c *gin.Context) {
	id := c.Query("id")
	entity := ctl.userService.FindById(id)
	g.NewResult(c).Yes(entity)
}
