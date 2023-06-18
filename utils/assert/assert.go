package assert

import (
	"strings"

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
		panic(msg)
	}
}

func IsNotNull(obj any, msg string) {
	if obj != nil {
		panic(msg)
	}
}

func IsBlank(s string, msg string) {
	if strings.TrimSpace(s) == "" {
		panic(msg)
	}
}

func IsNotBlank(s string, msg string) {
	if strings.TrimSpace(s) != "" {
		panic(msg)
	}
}

func IsEmpty[T gsafe.Numeric](num T, msg string) {
	if num == 0 {
		panic(msg)
	}
}

func IsTrue(expr bool, msg string) {
	if expr == true {
		panic(msg)
	}
}

func IsFalse(expr bool, msg string) {
	if expr == false {
		panic(msg)
	}
}

func ExecSuccess(err error, msg string) {
	if err != nil {
		panic(msg)
	}
}

func CheckCount(count int, msg string) {
	if count < 1 {
		panic(msg)
	}
}
