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

func Run1(project, modelName, root, pkg string) {
	for _, name := range []string{"service_impl.go.tmpl",
		"service.go.tmpl", "mapper.go.tmpl", "model.go.tmpl", "controller.go.tmpl",
	} {
		filepath := gpath.Join(gpath.RootPath(), "resource", "template1", name)
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
		case "service_impl.go.tmpl":
			path = gpath.Join(root, "/service/impl/"+strcase.ToSnake(modelName)+"_service_impl.go")
		case "service.go.tmpl":
			path = gpath.Join(root, "/service/"+strcase.ToSnake(modelName)+"_service.go")
		case "mapper.go.tmpl":
			path = gpath.Join(root, "/mapper/"+strcase.ToSnake(modelName)+"_mapper.go")
		case "model.go.tmpl":
			path = gpath.Join(root, "/model/"+strcase.ToSnake(modelName)+".go")
		case "controller.go.tmpl":
			path = gpath.Join(root, "/controller/"+strcase.ToSnake(modelName)+"_controller.go")
		}

		gpath.MkParentDir(path)
		fp, _ := os.Create(path)

		err = tmpl.Execute(fp, data)
		if err != nil {
			panic(err)
		}
	}

}
