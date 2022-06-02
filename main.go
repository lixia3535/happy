package main

import (
	"fmt"
	"os"
	"strings"
)

/*func main(){

	// 加载配置文件
	fmt.Print("加载配置文件，初始化DB")
	Egn := InitGin()
	Egn.Run(":8081")

}
*/

var ArgsMap = map[string]string {
	"abc": "abc",
}

func main() {
	var InputKey, InputValue string
	for i := 1; i < len(os.Args); i++ {
		InputArgs := os.Args[i]
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Args Error",r)
				fmt.Println("Args Error, No \"=\" found")
				fmt.Println("Program Exit!")
			}
		}()
		// Equal取值时可能由于用户输入没有"=",导致panic,需要defer recover
		Equal  := strings.Index(InputArgs,"=")
		if Equal < 0 {
			fmt.Println("PS")
			return
		}
		fmt.Println("Equal: ",Equal)
		fmt.Println("异常情况不执行")
		InputKey = InputArgs[:Equal]
		InputValue = InputArgs[Equal+1:]
		if ArgsFunc := ArgsMap[InputKey]; ArgsFunc != "" {
			fmt.Println(ArgsMap[InputKey])
			fmt.Println(InputValue)
		}
	}
	fmt.Println("任何情况均执行")
}