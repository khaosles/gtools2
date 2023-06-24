package main

import (
	"github.com/khaosles/gtools2/components/tmpl"
	gpath "github.com/khaosles/gtools2/utils/path"
)

/*
   @File: main.go
   @Author: khaosles
   @Time: 2023/6/13 21:17
   @Desc:
*/

func main() {
	project := "data-oversea/app"
	root := gpath.Join(gpath.RootPath(), "app")
	models := map[string][]string{
		//"log": {
		//	"LogApi",
		//	"LogTask",
		//},
		//"task": {
		//	"Order",
		//	"TmpSubProduct",
		//	"Product",
		//	"Task",
		//	"SubProduct",
		//},
		"system": {
			"sys_meta_data",
			"SysConfig",
		},
	}

	for pkg, models := range models {
		for _, model := range models {
			tmpl.Run(project, model, root, pkg)
		}
	}
}
