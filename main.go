package main

import (
	"GE_Revenue_Services/waybill"
	"fmt"
)

func main() {

	result, raw := waybill.GetServerTime()

	fmt.Println("Here is the request sent:")
	fmt.Println(raw.Request)
	fmt.Println("Here is the result got:")
	fmt.Println(raw.Response)
	fmt.Println("Here is the parsed xml into a struct:")
	fmt.Println(result)

	result2, raw2 := waybill.WhatIsMyIP()

	fmt.Println("Here is the request sent:")
	fmt.Println(raw2.Request)
	fmt.Println("Here is the result got:")
	fmt.Println(raw2.Response)
	fmt.Println("Here is the parsed xml into a struct:")
	fmt.Println(result2)
}
