package main

import (
	"GE_Revenue_Services/waybill"
	"fmt"
)

func main() {
	println("Hello RS.GE!")

	fmt.Println(waybill.ChekServiceUser("vlkuz:12345678910", "123456").UniqId)
}
