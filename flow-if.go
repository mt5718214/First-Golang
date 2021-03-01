package main
import "fmt"
func main(){
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
}