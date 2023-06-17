package gstru

import (
	"fmt"
	"testing"
	"time"
)

/*
   @File: stru_test.go
   @Author: khaosles
   @Time: 2023/6/4 01:51
   @Desc:
*/

func TestStructToMap(t *testing.T) {

	a := struct {
		A string `json:"a,omitempty"`
		B int
		C struct {
			D float64 `json:"d,omitempty"`
			E int     `json:"e,omitempty"`
		} `json:"c"`
		F time.Duration `json:"f,omitempty"`
	}{
		A: "aaa", B: 1, C: struct {
			D float64 `json:"d,omitempty"`
			E int     `json:"e,omitempty"`
		}(struct {
			D float64
			E int
		}{D: 2.2, E: 1}), F: time.Duration(3),
	}
	res := StructToMap(&a)
	fmt.Println(res)
}
