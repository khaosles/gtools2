package main

import (
	"os"
	"text/template"
	"time"

	"github.com/iancoleman/strcase"

	gpath "github.com/khaosles/gtools2/g/path"
)

/*
   @File: main.go
   @Author: khaosles
   @Time: 2023/6/13 21:17
   @Desc:
*/

func main() {
	project := "data-search/app"
	modelName := "metaData"
	root := gpath.Join(gpath.RootPath(), "app")
	for _, name := range []string{"service.go.internal.tmpl", "service.go.tmpl", "mapper.go.internal.tmpl", "mapper.go.tmpl"} {
		filepath := gpath.Join(gpath.RootPath(), "resource", "template", name)
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
			Name:      strcase.ToLowerCamel(modelName),
			NameUpper: strcase.ToCamel(modelName),
			Date:      time.Now().Format(time.DateTime),
			Project:   project,
		}
		var path string
		switch name {
		case "service.go.internal.tmpl":
			path = gpath.Join(root, "/service/internal/"+strcase.ToSnake(modelName)+"_service_impl.go")
		case "service.go.tmpl":
			path = gpath.Join(root, "/service/"+strcase.ToSnake(modelName)+"_service.go")
		case "mapper.go.internal.tmpl":
			path = gpath.Join(root, "/mapper/internal/"+strcase.ToSnake(modelName)+"_mapper_impl.go")
		case "mapper.go.tmpl":
			path = gpath.Join(root, "/mapper/"+strcase.ToSnake(modelName)+"_mapper.go")
		}

		gpath.MkParentDir(path)
		fp, _ := os.Create(path)

		err = tmpl.Execute(fp, data)
		if err != nil {
			panic(err)
		}
	}

}
