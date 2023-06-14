package middleware

import "github.com/gin-gonic/gin"

/*
   @File: error_to_mail.go
   @Author: khaosles
   @Time: 2023/6/15 00:19
   @Desc:
*/

func ErrorToMail() gin.HandlerFunc {
	return func(c *gin.Context) {
		//
		c.Next()
	}
}
