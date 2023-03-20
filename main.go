package main

import (
	"go_test/internal/sign"
	"log"
	"os"
	"strings"
)

func main() {
	f,err := os.Open("./1.txt")
	if err != nil {
		log.Println(err)
	}

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
	//var sing = common.Sign("header_1.png", "/v2")
	//log.Println(sing)
	//http://alist.yangtuyun.cn/d/v2/header_1.png?sign=CvGZ2hC-2Uc3kSB1fedG_RzQcXkirv8aKvCEfPqUOz8=:1679045627
	//ddbIpKU1c7hij9vDOnB6ZK5_c-QFLIfVrH3uboJWExI=:0
	rawPath := "v2/header_1.png"
	s := "B3iXLLD3LkJoueIOcwdQa6Eud6_pG5ttV-W7uJUWoHo=:1679053776"
	// LDvd5-L31LjtBOVMDSsdrekAwVPdsiey8FyzbsRDf3E=:1679053714
	//s := ddbIpKU1c7hij9vDOnB6ZK5_c-QFLIfVrH3uboJWExI=:0
	err := sign.Verify(rawPath, strings.TrimSuffix(s, "/"))
	if err != nil {
		log.Println(err, "err")
	}

}
