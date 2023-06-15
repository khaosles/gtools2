package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khaosles/gtools2/components/g"
	glog "github.com/khaosles/gtools2/core/log"
	"github.com/unrolled/secure"
)

/*
   @File: ssl_redirect.go
   @Author: khaosles
   @Time: 2023/6/15 00:27
   @Desc:
*/

func SSLRedirect() gin.HandlerFunc {
	return func(c *gin.Context) {
		middleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "localhost:443",
		})
		err := middleware.Process(c.Writer, c.Request)
		if err != nil {
			glog.Error(err)
			c.JSON(http.StatusOK, g.NewJsonResult().CatchErr(err))
			return
		}
		// 继续往下处理
		c.Next()
	}
}
