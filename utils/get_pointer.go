package utils

/*
   @File: get_pointer.go
   @Author: khaosles
   @Time: 2023/6/14 22:12
   @Desc:
*/

func Pointer[T any](in T) (out *T) {
	return &in
}
