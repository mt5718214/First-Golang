package main //可執行程式必須使用 main 封包
import "fmt" //載入內建的fmt封包,用來做基本輸出輸入
func data() { //建立main函式,程式的進入點
	/*
		fmt.Println(3)  //int
		fmt.Println(3.1415)  //float64
		fmt.Println("測試")  //string
		fmt.Println(true)  //boolean
		fmt.Println('a')  //rune
	*/

	//變數使用
	var x int //宣告變數
	x = 4
	fmt.Println(x)
	x = 10
	fmt.Println(x)

	var f float64 = 3.1415
	fmt.Println(f)

	var s string = "哈囉"
	fmt.Println(s)

	var test bool = true //布林值宣告 bool
	fmt.Println(test)

	var c rune = 'b'
	fmt.Println(c)
}
