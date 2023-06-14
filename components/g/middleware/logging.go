package middleware

import "github.com/gin-gonic/gin"

/*
   @File: logging.go
   @Author: khaosles
   @Time: 2023/6/15 00:16
   @Desc:
*/

func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		//
		c.Next()
	}
}
