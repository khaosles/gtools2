package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/x/beamx"
	_ "github.com/go-sql-driver/mysql"
)

/*
   @File: main.go
   @Author: khaosles
   @Time: 2023/6/3 17:30
   @Desc:
*/

type Person struct {
	ID   int    `beam:"id"`
	Name string `beam:"name"`
	Age  int    `beam:"age"`
}

func InsertIntoDB(person Person, db *sql.DB) string {
	// 执行插入操作
	db.Exec("INSERT INTO user (id, name, age) VALUES (?, ?, ?)", person.ID, person.Name, person.Age)
	return "123"
}

func main() {
	// 连接数据库
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/sea")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// 设置go get -u github.com/apache/beam/sdks/v2/go/pkg/beam Apache Beam 流程
	p, s := beam.NewPipelineWithRoot()

	// 从数据源读取数据
	input := beam.CreateList(s, []Person{
		{ID: 1, Name: "Alice", Age: 25},
		{ID: 2, Name: "Bob", Age: 30},
		{ID: 3, Name: "Charlie", Age: 35},
	})

	// 使用 ParDo 操作执行入库操作
	inserted := beam.ParDo(s, func(person Person) string {
		return InsertIntoDB(person, db)
	}, input)

	// 错误处理
	errors := beam.ParDo(s, func(err string) string {
		return err
	}, inserted)

	// // 输出错误日志
	beam.ParDo0(s, func(errMsg string) {
		log.Println(errMsg)
	}, errors)

	// 运行 Apache Beam 流程
	if err := beamx.Run(context.Background(), p); err != nil {
		log.Fatalf("Pipeline failed: %v", err)
	}

	fmt.Println("数据插入成功")
}
