package main

import (
	"github.com/khaosles/gtools2/components/g"
	"github.com/khaosles/gtools2/components/g/pgsql"
)

/*
   @File: main.go
   @Author: khaosles
   @Time: 2023/6/3 17:30
   @Desc:
*/

func main() {
	pgsql.DB.AutoMigrate(&g.Model{})
	mapper := g.MapperImpl[g.Model]{DB: pgsql.DB}
	mapper.Insert(
		&g.Model{},
	)
	//a, _ := dao.SelectByID("3dc2eca4389a4ba9953e6493088cccee")

	//dao.UpdateSelective(a, map[string]any{
	//	"remarks":   "1",
	//	"create_by": "1111",
	//})
	//c := orm.NewConditions().
	//	AndEqual("id", "30fc064892e44ce88b6c448989e59a30").
	//	AndIsNull("remarks").Order("id").Select("id, create_time")
	//dao.SelectByCondition(c)
	_, _ = mapper.SelectByCondition(g.NewConditions().AndEqualTo("update_by", "").Joins("LEFT Join df_order on df_order.id=model.id"))
	//dao.DeleteHardByCondition(c)
	//dao.DeleteByCondition(struct {
	//	Remark string `json:"remark"`
	//}{""})
	//spri := "MIICdwIBADANBgkqhkiG9w0BAQEFAASCAmEwggJdAgEAAoGBAOAsGajbJJqD1QFY/FlnlXXIQOVIZm7Pt/UMEp9JW38bCV7v0qRvCZCf85/JGCpdYQ+ql6ykank8fs00HmRBM5uOlhboNhI86PyBqGU850WQjmOtCCgbEOCOWM4/D1favEu/p8mVqxJ7gYC4gEcgJICyMxjDBU9B4D6begtLaUpZAgMBAAECgYBVoGK4veQwZRTSq/PQDqHnWHN5YPtHbm5c2pyuXS3m0iP1MHPsPUGRDZfYO87QN9TgUBAZcL/+yR3CMhs9vi4AkOMahgvirviXDtYBrT3nIHRQpZxqEw5EYak8OBXHoIfvSaz90iMCgquMbaZ675g/XoPv32u2/w3lyRrq4G8oUQJBAPzJZ0gIsw6iFCy4+1MzPEqH5xEmx+3q4gG7tp/Y3cTVdrDa+YqOtJA/9T5bUT2KUAYXXb2Fez4xs1pdq/gsNI0CQQDjBZUDyhXg8P6R74VeVz9WX3ypKfoR9n98WOH7C/p/Hc0ylwbDm91AnbR+W883zsE0s1g9c+ZaVQCaeRiiz4f9AkEAitihJxrIJwh1Zl8whHGG8zUUgQI5HIBAJU2SsNfwb7YEHH4aRLW/jd/jd5220MOQ0tewwHF50R6BcegzlfvJ3QJAdJh6Vw7kO7oqVNNagQB4VCkIgm0/tRgPk9KmhWQ6jCzHJbNxUudrM/OLLtaCT5xNmH5/1FgBN+WuQKfvIjdKFQJBAKY06Sc4IGZErUQJKFdVAz/NTPgL6Ed4cNzIpTJfJgbX1PCkiVKL2o+aVPFgojyRVglK/t8ZisNlhr/obJEWIRo="
	//spub := "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDgLBmo2ySag9UBWPxZZ5V1yEDlSGZuz7f1DBKfSVt/Gwle79KkbwmQn/OfyRgqXWEPqpespGp5PH7NNB5kQTObjpYW6DYSPOj8gahlPOdFkI5jrQgoGxDgjljOPw9X2rxLv6fJlasSe4GAuIBHICSAsjMYwwVPQeA+m3oLS2lKWQIDAQAB"
	//
	//println("公钥")
	//println(spub)
	//println("私钥")
	//println(spri)
	//data := "123456234567890-09876543"
	//pub, _ := grsa.PublicKeyFromX509PKCS8(spub)
	//pri, _ := grsa.PrivateKeyFromX509PKCS8(spri)
	//encrypt, _ := grsa.Encrypt(data, pub)
	//println("密文")
	//println(encrypt)
	//encrypt = `Y9qINbj9wCMqoXZj2LjQSo5+uG0UPuM2SXwjVwI8/LpezSe9HM92XLV8uERuYat1epwAXT0BKiSlh0qxVK9MnTnhilqCT5nBdxTJYMo3VO6m5lfwrKWsdbSm5tiB5LdDfevCmydH0A9vJ2rOhJAqKCbsraeCn9vAMbf+57rcAYY=`
	//
	//decrypt, _ := grsa.Decrypt(encrypt, pri)
	//println("解密")
	//println(decrypt)
}
