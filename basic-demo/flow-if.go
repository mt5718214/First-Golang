package main

import (
	"fmt"
	"math/rand"
	"time"
)

func flow() {
	//基本語法
	if true {
		fmt.Println("Go")
	} else {
		fmt.Println("Not Go")
	}
	//簡易情境: ATM
	var money int
	fmt.Println("請問想領多少錢? ")
	fmt.Scanln(&money)
	if money < 100 {
		fmt.Println("Too Few")
	} else if money <= 30000 {
		fmt.Println("OK")
	} else {
		fmt.Println("Too Much")
	}
	fmt.Println("執行完畢")

	// 在條件判斷句宣告變數, 其作用域只在該條件邏輯區內
	if x := computedValue(); x > 50 {
		fmt.Println(x)
		fmt.Println("x is greater than 50")
	} else {
		fmt.Println("x is less than 50")
	}
	// fmt.Println(x) <-- 會報錯

	// goto
	myFunc()

	// for loop
	for index := 0; index < 10; index++ {
		index += index
	}

	// use range to read value of map or slice
	slice := []int{1, 2, 3}
	for k, v := range slice {
		fmt.Println("slice's key:", k)
		fmt.Println("slice's value:", v)
	}

	variadic(1, 2, 3)

}

// goto語句, 用goto跳轉到必須在當前函式內定義的標籤（標籤名稱(label)是區分大小寫的的。）
func myFunc() {
	i := 0
Here: //這行的第一個詞，以冒號結束作為標籤
	println(i)
	i++
	if i > 10 {
		return
	}
	goto Here //跳轉到 Here 去
}

// 可變參數函式, 可接受不定數量的參數, 以下例子接受不定數量的int類型參數
func variadic(arg ...int) {
	for _, n := range arg {
		fmt.Printf("And the number is: %d\n", n)
	}
}

func computedValue() int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return r1.Intn(100)
}
