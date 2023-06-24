package g

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khaosles/gtools2/components/g/result"
	"github.com/khaosles/gtools2/utils/assert"
)

/*
   @File: base_controller.go
   @Author: khaosles
   @Time: 2023/6/13 10:07
   @Desc:
*/

type BaseController[T any] struct {
	Service Service[T]
}

func (ctl BaseController[T]) Add(c *gin.Context) {
	var entity T
	err := c.ShouldBindJSON(&entity)
	assert.ExecSuccess(err, "参数错误")
	err = ctl.Service.Save(&entity)
	assert.ExecSuccess(err, "保存失败")
	c.JSON(http.StatusOK, result.NewJsonResult().Yes(nil))
}

func (ctl BaseController[T]) Update(c *gin.Context) {
	var entity T
	err := c.ShouldBindJSON(&entity)
	assert.ExecSuccess(err, "参数错误")
	err = ctl.Service.Update(&entity)
	assert.ExecSuccess(err, "更新失败")
	c.JSON(http.StatusOK, result.NewJsonResult().Yes(nil))
}

func (ctl BaseController[T]) DeleteById(c *gin.Context) {
	id := c.Param("id")
	assert.IsBlank(id, "id不能为空")
	err := ctl.Service.DeleteById(id)
	assert.ExecSuccess(err, "删除失败")
	c.JSON(http.StatusOK, result.NewJsonResult().Yes(nil))
}

func (ctl BaseController[T]) FindAll(c *gin.Context) {
	entities, err := ctl.Service.FindAll()
	assert.ExecSuccess(err, "查询失败")
	c.JSON(http.StatusOK, result.NewJsonResult().Yes(entities))
}

func (ctl BaseController[T]) FindById(c *gin.Context) {
	id := c.Query("id")
	assert.IsBlank(id, "id不能为空")
	entity, err := ctl.Service.FindById(id)
	assert.ExecSuccess(err, "查询失败")
	c.JSON(http.StatusOK, result.NewJsonResult().Yes(entity))
}
