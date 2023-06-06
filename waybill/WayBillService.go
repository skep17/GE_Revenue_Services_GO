package waybill

import (
	"GE_Revenue_Services/soap"
	"fmt"
	"strings"
	"time"
)

func getURL() string {
	return "https://services.rs.ge/WayBillService/WayBillService.asmx"
}

func chekServiceUserReq(su, sp string) []byte {
	return []byte(strings.TrimSpace(fmt.Sprintf(`
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:tem="http://tempuri.org/">
		<soapenv:Header/>
		<soapenv:Body>
			<tem:chek_service_user>
				<!--Optional:-->
				<tem:su>%s</tem:su>
				<!--Optional:-->
				<tem:sp>%s</tem:sp>
			</tem:chek_service_user>
		</soapenv:Body>
 	</soapenv:Envelope>`, su, sp)))
}

func ChekServiceUser(su, sp string) CheckServiceUserResponse {
	var result chekServiceUserXML
	soap.SendRequest(chekServiceUserReq(su, sp), &result, getURL())
	return result.Body.Data
}

func closeWaybillReq(su, sp string, id int) []byte {
	return []byte(strings.TrimSpace(fmt.Sprintf(`
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:tem="http://tempuri.org/">
	<soapenv:Header/>
		<soapenv:Body>
			<tem:close_waybill>
				<!--Optional:-->
				<tem:su>%s</tem:su>
				<!--Optional:-->
				<tem:sp>%s</tem:sp>
				<tem:waybill_id>%d</tem:waybill_id>
			</tem:close_waybill>
		</soapenv:Body>
 	</soapenv:Envelope>`, su, sp, id)))
}

func CloseWaybill(su, sp string, id int) CloseWaybillResponse {
	var result closeWaybillXML
	soap.SendRequest(closeWaybillReq(su, sp, id), &result, getURL())
	return result.Body.Data
}

func closeWaybillTransporterReq(su, sp string, id int, reception_info, reciever_info string, delivery time.Time) []byte {
	return []byte(strings.TrimSpace(fmt.Sprintf(`
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:tem="http://tempuri.org/">
	<soapenv:Header/>
		<soapenv:Body>
			<tem:close_waybill>
				<!--Optional:-->
				<tem:su>%s</tem:su>
				<!--Optional:-->
				<tem:sp>%s</tem:sp>
				<tem:waybill_id>%d</tem:waybill_id>
			</tem:close_waybill>
		</soapenv:Body>
 	</soapenv:Envelope>`, su, sp, id)))
}

func CloseWaybillTransporter(su, sp string, id int, reception_info, reciever_info string, delivery time.Time) CloseWaybillTransporterResponse {
	var result closeWaybillTransporterXML
	soap.SendRequest(closeWaybillTransporterReq(su, sp, id, reception_info, reciever_info, delivery), &result, getURL())
	return result.Body.Data
}
