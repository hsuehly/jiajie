package main

import (
	"go_test/common"
	"log"
)

func main() {
	//f,err := os.Open("./1.txt")
	//if err != nil {
	//	log.Println(err)
	//}

	//http://alist.yangtuyun.cn/d/v2/header_1.png?sign=ddbIpKU1c7hij9vDOnB6ZK5_c-QFLIfVrH3uboJWExI=:0
	//v2/header_1.png
	//ddbIpKU1c7hij9vDOnB6ZK5_c-QFLIfVrH3uboJWExI=:0
	//rawPath := "v2/header_1.png"
	//s := "ddbIpKU1c7hij9vDOnB6ZK5_c-QFLIfVrH3uboJWExI=:0"
	// s := ddbIpKU1c7hij9vDOnB6ZK5_c-QFLIfVrH3uboJWExI=:0
	//err := sign.Verify(rawPath, strings.TrimSuffix(s, "/"))
	//if err != nil {
	//	log.Println(err, "err")
	//}
	//http://alist.yangtuyun.cn/d/v2/header_1.png?sign=HqHkGbfmmbWD53JNeieglR3wvNPs_nXB7sdOJenpq2Q=:1679053686
	// 6vr6PcfkBVzJ4Q330OldD15VtaeF8ASB4byGLSx8NV8=:1679053714
	var sing = common.Sign("/hsuehly.jpg", "/v2/23-03-20-23-09-10/")
	log.Println(sing)
	//hL5wcbzmGGjdLOEmB9gt6ms4fe9n_ur9ngP0aKioWs4=:1679400340
	//wuqL0w63FvIJrpYRx6hAQjRvEKCf-1yj9-QkRykFuBM=:1679400385
	//7IyF_GJWUlIKmrRt_vzD7tGfue7GctpW00Va3b8JPRc=:1679400397
	//
	//http://alist.yangtuyun.cn/d/v2/header_1.png?sign=CvGZ2hC-2Uc3kSB1fedG_RzQcXkirv8aKvCEfPqUOz8=:1679045627
	//ddbIpKU1c7hij9vDOnB6ZK5_c-QFLIfVrH3uboJWExI=:0
	//rawPath := "v2/23-03-20-23-09-10/hsuehly.jpg"
	//// 1小时 http://alist.yangtuyun.cn/d/v2/23-03-20-23-09-10/hsuehly.jpg?sign=7sfx4v3QwZNx4J0mSGXgC86mPexy_ER5OEf30ikAeU0=:1679399242
	//s := "7sfx4v3QwZNx4J0mSGXgC86mPexy_ER5OEf30ikAeU0=:1679399242"
	//log.Println(strings.TrimSuffix(s, "/"), "sss")
	//err := sign.Verify(rawPath, strings.TrimSuffix(s, "/"))
	//if err != nil {
	//	log.Println(err, "err")
	//}
	//http://alist.yangtuyun.cn/d/v2/23-03-20-23-09-10/hsuehly.jpg?sign=BtSPgs8yS2jYDxS5PBiBJVmJvpQq9K1DlO5D6y28GBo=:1679400171

}
