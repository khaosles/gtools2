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
	project := "admin-server/app"
	root := gpath.Join(gpath.RootPath(), "app")
	models := []string{
		//"sys_config",
		//"sys_config_data",
		"log_login",
		"log_operator",
		//"sys_dictionary",
		//"sys_dictionary_data",
		//"sys_user",
		//"sys_role",
		//"sys_permission",
		//"sys_user_role",
		//"sys_role_permission",
	}

	for _, model := range models {
		tmpl.Run(project, model, root, "log")
	}

}
