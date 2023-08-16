package gerr

import "fmt"

/*
   @author: khaosles
   @date: 2023/2/28 15:12
   @description:
*/

var (
	NotImplementedException = Exception("NotImplemented")
)

func NotImplment(method string) error {
	return fmt.Errorf("NotImplemented-> func:%s", method)
}
