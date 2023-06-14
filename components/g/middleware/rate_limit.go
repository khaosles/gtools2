package middleware

import "github.com/gin-gonic/gin"

/*
   @File: rate_limit.go
   @Author: khaosles
   @Time: 2023/6/15 00:16
   @Desc:
*/

func RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		//
		c.Next()
	}
}
