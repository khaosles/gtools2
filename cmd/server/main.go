package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/x/beamx"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocarina/gocsv"
	"github.com/khaosles/gtools2/core/db/pgsql"
	gserver "github.com/khaosles/gtools2/g/server"
)

/*
   @File: main.go
   @Author: khaosles
   @Time: 2023/6/3 17:30
   @Desc:
*/

var ll = 2000

type Htp struct {
	gserver.BaseModel
	//Point      string `json:"point" gorm:"type:geometry"`
	Data       string `json:"data" gorm:"type:numeric[]"`
	ReportTime time.Time
	Version    int
}

type Row struct {
	Data       string `csv:"data"`
	Gid        int    `csv:"gid"`
	Version    string `csv:"version"`
	ReportTime string `csv:"report_time"`
}

func InsertIntoDB(rows []*Row) int64 {
	var htps []*Htp
	for _, row := range rows {
		var htp Htp
		//htp.Point = row.Point
		version, _ := strconv.Atoi(row.Version)
		htp.Version = version
		reportTime, _ := time.Parse(time.DateTime, row.ReportTime)
		htp.ReportTime = reportTime
		htp.Data = row.Data
		htps = append(htps, &htp)
	}
	// 执行插入操作
	affected := pgsql.DB.CreateInBatches(&htps, ll).RowsAffected
	return affected
}

func main() {
	se := time.Now()
	path := "/Users/yherguot/doraemon/data/gfs/stofs_2d_glo.t00z.fields.htp.csv"
	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.LazyQuotes = true
		r.Comma = ';'
		return r // Allows use dot as delimiter and use quotes in CSV
	})
	file, _ := os.Open(path)
	var rows []*Row
	err := gocsv.Unmarshal(file, &rows)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(len(rows))
	var d [][]*Row
	for i := 0; i < len(rows); i += ll {
		end := i + ll
		if end > len(rows) {
			end = len(rows)
		}
		d = append(d, rows[i:end])
		//InsertIntoDB(rows[i:end])
	}

	p, s := beam.NewPipelineWithRoot()
	input := beam.CreateList(s, d)

	//// 使用 ParDo 操作执行入库操作
	_ = beam.ParDo(s, func(row []*Row) int64 {
		return InsertIntoDB(row)
	}, input)
	//
	//// 错误处理
	//errors := beam.ParDo(s, func(err int64) int64 {
	//	return err
	//}, inserted)
	//
	//// 输出错误日志
	//beam.ParDo0(s, func(errMsg int64) {
	//	log.Println(errMsg)
	//}, errors)
	//

	//// 运行 Apache Beam 流程
	if err := beamx.Run(context.Background(), p); err != nil {
		log.Fatalf("Pipeline failed: %v", err)
	}
	//
	fmt.Println("数据插入成功")
	println("耗时 =>", time.Now().Sub(se).Seconds())

}
