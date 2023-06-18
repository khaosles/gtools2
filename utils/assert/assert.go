package assert

import (
	"strings"

	"github.com/khaosles/gtools2/components/g/result"
	glog "github.com/khaosles/gtools2/core/log"
	gsafe "github.com/khaosles/gtools2/utils/safe"
)

/*
   @File: assert.go
   @Author: khaosles
   @Time: 2023/6/18 12:59
   @Desc:
*/

func IsNull(obj any, msg string) {
	if obj == nil {
		panic(result.NewAssertError(msg))
	}
}

func IsNotNull(obj any, msg string) {
	if obj != nil {
		panic(result.NewAssertError(msg))
	}
}

func IsBlank(s string, msg string) {
	if strings.TrimSpace(s) == "" {
		panic(result.NewAssertError(msg))
	}
}

func IsNotBlank(s string, msg string) {
	if strings.TrimSpace(s) != "" {
		panic(result.NewAssertError(msg))
	}
}

func IsEmpty[T gsafe.Numeric](num T, msg string) {
	if num == 0 {
		panic(result.NewAssertError(msg))
	}
}

func IsTrue(expr bool, msg string) {
	if expr == true {
		panic(result.NewAssertError(msg))
	}
}

func IsFalse(expr bool, msg string) {
	if expr == false {
		panic(result.NewAssertError(msg))
	}
}

func ExecSuccess(err error, msg string) {
	if err != nil {
		glog.Error(err)
		panic(result.NewInternalError(msg))
	}
}

func CheckCount(count int, msg string) {
	if count < 1 {
		panic(result.NewInternalError(msg))
	}
}
