package gresult

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

func (r Result) Success(data any) {
	r.ctx.JSON(http.StatusOK, JsonResult{}.Yes(data))
	return
}

func (r Result) Set() {

}

type Option func(any) *JsonResult

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

func (jsonResult JsonResult) SetCode(code int) JsonResult {
	jsonResult.Code = code
	return jsonResult
}

func (jsonResult JsonResult) SetMsg(Msg string) JsonResult {
	jsonResult.Msg = Msg
	return jsonResult
}

func (jsonResult JsonResult) SetSuccess(success bool) JsonResult {
	jsonResult.Success = success
	return jsonResult
}

func (jsonResult JsonResult) SetData(data interface{}) JsonResult {
	jsonResult.Data = data
	return jsonResult
}

// Yes is run successful
func (jsonResult JsonResult) Yes(data interface{}) JsonResult {
	return jsonResult.SetCode(SUCCESS.Code).SetMsg(SUCCESS.Msg).SetSuccess(true).SetData(data)
}

// No is run failed
func (jsonResult JsonResult) No(err ErrorCode) JsonResult {
	return jsonResult.SetCode(err.Code).SetMsg(err.Msg)
}

// CatchErr 异常捕获
func (jsonResult JsonResult) CatchErr(err any) JsonResult {
	return jsonResult.No(CATCH_ERROR.SetMsg(fmt.Sprintf("%v", err)))
}
