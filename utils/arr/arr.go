package garr

/*
   @File: arr.go
   @Author: khaosles
   @Time: 2023/7/5 17:52
   @Desc:
*/

func ValueInArray[T comparable](target T, arr []T) bool {
	for _, element := range arr {
		if target == element {
			return true
		}
	}
	return false
}
