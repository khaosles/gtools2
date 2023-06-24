package assert

import (
	"strings"
	"time"

	"github.com/khaosles/gtools2/components/g/result"
	glog "github.com/khaosles/gtools2/core/log"
	gerr "github.com/khaosles/gtools2/utils/err"
	gpath "github.com/khaosles/gtools2/utils/path"
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

func IsZero[T gsafe.Numeric](num T, msg string) {
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

func IsEmpty[T any](arr []T, msg string) {
	if len(arr) == 0 {
		panic(result.NewAssertError(msg))
	}
}

func IsEmptyPointer(obj any, msg string) {
	if obj == nil {
		panic(result.NewAssertError(msg))
	}
}

func IsNotEmptyPointer(obj any, msg string) {
	if obj != nil {
		panic(result.NewAssertError(msg))
	}
}

func IsEmptyTime(t time.Time, msg string) {
	if t.IsZero() {
		panic(result.NewAssertError(msg))
	}
}

func IsFile(file string, msg string) {
	if gpath.IsFile(file) {
		panic(result.NewAssertError(msg))
	}
}

func IsNotFile(file string, msg string) {
	if !gpath.IsFile(file) {
		panic(result.NewAssertError(msg))
	}
}

func IsNotImplemented(point any, msg string) {
	if point == nil {
		panic(gerr.NotImplementedException.New(msg))
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
