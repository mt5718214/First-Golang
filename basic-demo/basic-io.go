package main

import "fmt"

func basic() {
	//基本輸出: fmt.Println(data, data, ...)
	fmt.Println(3)

	//基本輸出: fmt.Scanln(&data, &data, ...)
	//&data: 取得變數的指標(Pointer)
	var x int
	fmt.Println("輸入一個數字")
	fmt.Scanln(&x)
	fmt.Println(x)
}
