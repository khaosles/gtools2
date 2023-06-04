package main

import (
	"context"

	"github.com/khaosles/gtools2/components/download"
)

/*
   @File: main.go
   @Author: khaosles
   @Time: 2023/6/3 17:30
   @Desc:
*/

type Instance struct {
	Workers  int    `yaml:"workers" default:"4"`
	PartSize string `yaml:"part-size" default:"100MiB"`
	BufSize  string `yaml:"buf-size" default:"2MiB"`
}

func main() {
	//var d Instance
	//err := gcfg.GetComponentConfiguration("download", &d)
	//if err != nil {
	//	glog.Error(err)
	//}
	//fmt.Printf("%+v\n", d)
	//helper_gen.Uuid()
	//var p config.Mysql
	//err := gcfg.GetComponentConfiguration("mysql", &p)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Printf("%+v\n", p)
	//fmt.Printf("%+v\n", !p.LogZap)
	var i download.Instance
	i.Run()
	download.Get().Download(
		context.Background(),
		"/Users/yherguot/doraemon/data/gfs/stofs_2d_glo.t00z.fields.htp.nc",
		download.HttpReader{"https://nomads.ncep.noaa.gov/pub/data/nccf/com/stofs/prod/stofs_2d_glo.20230603/stofs_2d_glo.t00z.fields.htp.nc"},
		nil,
	)
	//println(gmachine.GetId(""))
	//println(helper_gen.Uuid())
	//println(helper_gen.UuidShort())
	//println(helper_gen.UuidNoSeparator())
	//println(helper_gen.RandString(40))
}
