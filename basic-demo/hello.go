package main //封包 一定要叫做main
import "fmt"

//撰寫程式 > 建置(build) > 執行程式
//建置程式: go build 程式檔案名稱
//執行程式: 輸入./hello
func hello() { //建立主程式 一定要叫做main
	sayHello()
}
func sayHello() {
	fmt.Println("Hello Golang")
}
