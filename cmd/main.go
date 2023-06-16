package main

import (
	"os"
	"text/template"
	"time"

	gpath "github.com/khaosles/gtools2/g/path"
)

/*
   @File: main.go
   @Author: khaosles
   @Time: 2023/6/13 21:17
   @Desc:
*/

func main() {

	filepath := gpath.Join(gpath.RootPath(), "resource", "template", "service.go.internal.tmpl")
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		panic(err)
	}

	data := struct {
		Name      string
		NameUpper string
		Date      string
		Project   string
	}{
		Name:      "metaData",
		NameUpper: "MetaData",
		Date:      time.Now().Format(time.DateTime),
		Project:   "data-search/app",
	}

	p := gpath.Join(gpath.RootPath(), "app/service/internal/meta_data_service.go")
	gpath.MkParentDir(p)
	fp, _ := os.Create(p)

	err = tmpl.Execute(fp, data)
	if err != nil {
		panic(err)
	}
}
