package main

import (
	"GE_Revenue_Services/waybill"
	"fmt"
)

func main() {
	var req waybill.CheckServiceUserRequest
	req.Su = "vlkuz:12345678910"
	req.Sp = "123456"
	//req.Id = 150621205
	fmt.Println(waybill.ChekServiceUser(req).XMLBody)
	//fmt.Println(time.Date(2023, 6, 20, 2, 0, 0, 0, time.Now().Local().Location()))
}
