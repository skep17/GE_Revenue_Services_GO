package main

import (
	"GE_Revenue_Services/waybill"
	"encoding/xml"
	"fmt"
	"log"
)

func main() {
	println("Hello RS.GE!")

	var result waybill.ServiceUser

	initServicePath("https://services.rs.ge/WayBillService/WayBillService.asmx")

	err := xml.NewDecoder(sendRequest(waybill.ChekServiceUser("vlkuz:12345678910", "123456")).Body).Decode(&result)
	if err != nil {
		log.Fatal("Error on unmarshaling xml. ", err.Error())
		return
	}

	fmt.Println(result.Body.ListUsersResponse.ReqResult)
	fmt.Println(result.Body.ListUsersResponse.UniqId)
	fmt.Println(result.Body.ListUsersResponse.SerUserId)
}
