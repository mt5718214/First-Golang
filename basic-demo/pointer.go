package main

import "fmt"

func pointer() {
	//1.建立存放資料的變數
	var x int = 6
	fmt.Println("原來的資料", x)
	//2.取的記憶體位置: &變數名稱
	fmt.Println("資料的記憶體位址", &x)
	//3.存放到指標變數, 注意指標資料型態: *資料型態
	var xPtr *int = &x
	//4.反解指標變數: *指標變數名稱
	fmt.Println("反解指標回原來的資料", *xPtr)
}
