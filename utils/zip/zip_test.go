package gzip

import (
	"testing"
)

/*
   @File: zip_test.go
   @Author: khaosles
   @Time: 2023/6/4 00:40
   @Desc:
*/

func TestCompress(t *testing.T) {
	path := "/Users/yherguot/doraemon/data/gfs/product/gfs-atmos/2023/07/09/18/numerical/temperature_2m_20230709180000.bin"
	Compress("/Users/yherguot/doraemon/data/gfs/product/gfs-atmos/2023/07/09/18/numerical/temperature_2m_20230709180000.zip", path)

	//err := Compress("/Users/yherguot/doraemon/data/test/estofs.t06z.fields.htp.csv.zip", "/Users/yherguot/doraemon/data/test/estofs.t06z.fields.htp.csv")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//Decompress("/Users/yherguot/doraemon/data/test/estofs.t06z.fields.htp.csv.zip", "/Users/yherguot/doraemon/data/test/")
}
