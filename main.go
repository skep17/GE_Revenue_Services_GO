package main

import (
	"GE_Revenue_Services/waybill"
	"fmt"
)

func main() {
	var req waybill.CheckServiceUserRequest
	req.Su = "vlkuz:12345678910"
	req.Sp = "123456"

	result, raw := waybill.ChekServiceUser(req)

	fmt.Println("Here is the request sent:")
	fmt.Println(raw.Request)
	fmt.Println("Here is the result got:")
	fmt.Println(raw.Response)
	fmt.Println("Here is the parsed xml into a struct:")
	fmt.Println(result)
}
