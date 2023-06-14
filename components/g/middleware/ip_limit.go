package middleware

import "github.com/gin-gonic/gin"

/*
   @File: ip_limit.go
   @Author: khaosles
   @Time: 2023/6/15 00:20
   @Desc:
*/

func IpLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		//
		c.Next()
	}
}
