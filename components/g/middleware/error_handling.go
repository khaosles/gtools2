package middleware

import (
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/khaosles/gtools2/components/g"
	glog "github.com/khaosles/gtools2/core/log"
)

/*
   @File: error_handling.go
   @Author: khaosles
   @Time: 2023/3/6 10:13
   @Desc: 全局异常捕获
*/

func ErrorHandling() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Abort()
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					glog.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					_ = c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}
				glog.Error("[Recovery from panic]",
					zap.Any("error", err),
					zap.String("request", string(httpRequest)),
					zap.String("stack", string(debug.Stack())),
				)
				c.JSON(http.StatusOK, g.NewJsonResult().CatchErr(err))
			}
		}()
		c.Next()
	}
}
