package middleware

import (
	"github.com/gin-gonic/gin"
)

/*
   @File: casbin_rbac.go
   @Author: khaosles
   @Time: 2023/6/15 00:18
   @Desc:
*/

// CasbinHandler 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		//
		c.Next()
	}
}
