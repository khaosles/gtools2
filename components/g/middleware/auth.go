package middleware

import "github.com/gin-gonic/gin"

/*
   @File: auth.go
   @Author: khaosles
   @Time: 2023/6/15 00:18
   @Desc:
*/

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//
		c.Next()
	}
}
