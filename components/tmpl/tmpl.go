package tmpl

import (
	"fmt"
	"os"
	"text/template"
	"time"

	"github.com/iancoleman/strcase"
	gpath "github.com/khaosles/gtools2/utils/path"
)

/*
   @File: main.go
   @Author: khaosles
   @Time: 2023/6/13 21:17
   @Desc:
*/

func Run(project, modelName, root, pkg string) {
	for _, name := range []string{"service.go.internal.tmpl",
		"service.go.tmpl", "mapper.go.internal.tmpl",
		"mapper.go.tmpl", "model.go.tmpl", "controller.go.tmpl",
	} {
		filepath := gpath.Join(gpath.RootPath(), "resource", "template", name)
		fmt.Println(filepath)
		tmpl, err := template.ParseFiles(filepath)
		if err != nil {
			panic(err)
		}
		data := struct {
			Name      string
			NameUpper string
			NameSnake string
			Date      string
			Project   string
			Pkg       string
		}{
			Name:      strcase.ToLowerCamel(modelName),
			NameUpper: strcase.ToCamel(modelName),
			NameSnake: strcase.ToSnake(modelName),
			Date:      time.Now().Format(time.DateTime),
			Project:   project,
			Pkg:       pkg,
		}
		var path string
		switch name {
		case "service.go.internal.tmpl":
			path = gpath.Join(root, "/service/internal/"+strcase.ToSnake(modelName)+"_service_impl.go")
		case "service.go.tmpl":
			path = gpath.Join(root, "/service/"+pkg+"/"+strcase.ToSnake(modelName)+"_service.go")
		case "mapper.go.internal.tmpl":
			path = gpath.Join(root, "/mapper/internal/"+strcase.ToSnake(modelName)+"_mapper_impl.go")
		case "mapper.go.tmpl":
			path = gpath.Join(root, "/mapper/"+pkg+"/"+strcase.ToSnake(modelName)+"_mapper.go")
		case "model.go.tmpl":
			path = gpath.Join(root, "/model/"+pkg+"/"+strcase.ToSnake(modelName)+".go")
		case "controller.go.tmpl":
			path = gpath.Join(root, "/controller/"+pkg+"/"+strcase.ToSnake(modelName)+"_controller.go")
		}

		gpath.MkParentDir(path)
		fp, _ := os.Create(path)

		err = tmpl.Execute(fp, data)
		if err != nil {
			panic(err)
		}
	}

}
