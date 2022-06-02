package utils

import (
	"encoding/base64"
	"fmt"
	"log"
)

func DecryptBase64(pwd string) string{
	pwdString ,err := base64.StdEncoding.DecodeString(pwd)
	if err != nil{
		log.Fatal(err.Error())
		fmt.Println("解析base64报错，",err.Error())
	}
	return string(pwdString)

}
