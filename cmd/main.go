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

	filepath := gpath.Join(gpath.RootPath(), "resource", "template", "mapper.go.internal.tmpl")
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
		Name:      "user",
		NameUpper: "User",
		Date:      time.Now().Format(time.DateTime),
		Project:   "github.com/khaosles/gtools2/components/example",
	}

	fp, _ := os.Create(gpath.Join(gpath.RootPath(), "components/example/mapper/internal/user_mapper1.go"))

	err = tmpl.Execute(fp, data)
	if err != nil {
		panic(err)
	}
}
