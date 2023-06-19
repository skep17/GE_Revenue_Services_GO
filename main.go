package main

import (
	"GE_Revenue_Services/waybill"
	"fmt"
)

func main() {
	var req waybill.DelWaybillRequest
	req.Id = 784467146
	req.Su = "vlkuz:12345678910"
	req.Sp = "123456"
	fmt.Println(waybill.DelWaybill(req).ResCode)
	//fmt.Println(time.Date(2023, 6, 20, 2, 0, 0, 0, time.Now().Local().Location()))
}
