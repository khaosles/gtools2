package g

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
   @File: json_result.go
   @Author: khaosles
   @Time: 2023/3/7 21:54
   @Desc:
*/

type Result struct {
	ctx *gin.Context
}

func NewResult(ctx *gin.Context) *Result {
	return &Result{ctx: ctx}
}

func (r Result) Yes(data any) {
	r.ctx.JSON(http.StatusOK, NewJsonResult().Yes(data))
	return
}

func (r Result) No(code ErrorCode) {
	r.ctx.JSON(http.StatusOK, NewJsonResult().No(code))
	return
}

type JsonResult struct {
	// code
	Code int `json:"code" default:"0"`
	// response information
	Msg string `json:"msg" default:""`
	// data
	Data interface{} `json:"data,omitempty" default:"nil"` // 默认无数据时不显示该字段
	// whether success
	Success bool `json:"success" default:"false"`
}

func NewJsonResult() *JsonResult {
	return &JsonResult{}
}

func (jr *JsonResult) SetCode(code int) *JsonResult {
	jr.Code = code
	return jr
}

func (jr *JsonResult) SetMsg(Msg string) *JsonResult {
	jr.Msg = Msg
	return jr
}

func (jr *JsonResult) SetSuccess(success bool) *JsonResult {
	jr.Success = success
	return jr
}

func (jr *JsonResult) SetData(data interface{}) *JsonResult {
	jr.Data = data
	return jr
}

// Yes is run successful
func (jr *JsonResult) Yes(data interface{}) *JsonResult {
	return jr.SetCode(SUCCESS.Code).SetMsg(SUCCESS.Msg).SetSuccess(true).SetData(data)
}

// No is run failed
func (jr *JsonResult) No(err ErrorCode) *JsonResult {
	return jr.SetCode(err.Code).SetMsg(err.Msg)
}

// CatchErr 异常捕获
func (jr *JsonResult) CatchErr(err any) *JsonResult {
	return jr.No(CATCH_ERROR.SetMsg(fmt.Sprintf("%v", err)))
}
