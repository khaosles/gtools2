package controller

import (
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/khaosles/gtools2/components/ourjson"

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

var (
	UserController = new(userController)
	once           sync.Once
)

type userController struct {
	userService service.UserService
}

func init() {
	once.Do(func() {
		UserController = &userController{userService: service.NewUserService()}
	})
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

func (ctl userController) FindAll(c *gin.Context) {
	entities := ctl.userService.FindAll()
	g.NewResult(c).Yes(entities)
	go func() {
		println("123")
		time.Sleep(time.Second * 2)
		println("....")
	}()
}

func (ctl userController) FindById(c *gin.Context) {
	id := c.Query("id")
	entity := ctl.userService.FindById(id)
	jsonObject := ourjson.NewJsonObject()
	jsonObject.Put("aaa", entity)
	g.NewResult(c).Yes(jsonObject.Value())
}
