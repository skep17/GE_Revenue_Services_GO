package waybill

import (
	"GE_Revenue_Services/soap"
	"fmt"
	"strings"
)

func getURL() string {
	return "https://services.rs.ge/WayBillService/WayBillService.asmx"
}

func chekServiceUserBody(su, sp string) []byte {
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

func ChekServiceUser(req CheckServiceUserRequest) (CheckServiceUserResponse, soap.RawData) {
	var response CheckServiceUserResponse
	rawData := soap.SendRequest(req, &response, getURL())
	return response, rawData
}

func closeWaybillBody(su, sp string, id int) []byte {
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

func CloseWaybill(req CloseWaybillRequest) (CloseWaybillResponse, soap.RawData) {
	var response CloseWaybillResponse
	rawData := soap.SendRequest(req, &response, getURL())
	return response, rawData
}

func closeWaybillTransporterBody(su, sp string, id int, reception_info, reciever_info, delivery string) []byte {
	return []byte(strings.TrimSpace(fmt.Sprintf(`
	<?xml version="1.0" encoding="utf-8"?>
	<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
	<soap:Body>
		<close_waybill_transporter xmlns="http://tempuri.org/">
		<su>%s</su>
		<sp>%s</sp>
		<waybill_id>%d</waybill_id>
		<reception_info>%s</reception_info>
		<receiver_info>%s</receiver_info>
		<delivery_date>%s</delivery_date>
		</close_waybill_transporter>
	</soap:Body>
	</soap:Envelope>`, su, sp, id, reception_info, reciever_info, delivery)))
}

func CloseWaybillTransporter(req CloseWaybillTransporterRequest) CloseWaybillTransporterResponse {
	var result CloseWaybillTransporterResponse
	soap.SendRequest(closeWaybillTransporterBody(req.Su, req.Sp, req.Id, req.ReceptionInfo, req.RecieverInfo, req.DeliveryDate), &result, getURL())
	return result
}

func closeWaybillVdBody(su, sp, delivery string, id int) []byte {
	return []byte(strings.TrimSpace(fmt.Sprintf(`
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:tem="http://tempuri.org/">
	<soapenv:Header/>
	<soapenv:Body>
		<tem:close_waybill_vd>
			<!--Optional:-->
			<tem:su>%s</tem:su>
			<!--Optional:-->
			<tem:sp>%s</tem:sp>
			<tem:delivery_date>%s</tem:delivery_date>
			<tem:waybill_id>%d</tem:waybill_id>
		</tem:close_waybill_vd>
	</soapenv:Body>
	</soapenv:Envelope>`, su, sp, delivery, id)))
}

func CloseWaybillVd(req CloseWaybillVdRequest) CloseWaybillVdResponse {
	var result CloseWaybillVdResponse
	soap.SendRequest(closeWaybillVdBody(req.Su, req.Sp, req.DeliveryDate, req.Id), &result, getURL())
	return result
}

func confirmWaybillBody(su, sp string, id int) []byte {
	return []byte(strings.TrimSpace(fmt.Sprintf(`
	<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
	<soap:Body>
		<confirm_waybill xmlns="http://tempuri.org/">
		<su>%s</su>
		<sp>%s</sp>
		<waybill_id>%d</waybill_id>
		</confirm_waybill>
	</soap:Body>
	</soap:Envelope>`, su, sp, id)))
}

func ConfirmWaybill(req ConfirmWaybillRequest) ConfirmWaybillResponse {
	var result ConfirmWaybillResponse
	soap.SendRequest(confirmWaybillBody(req.Su, req.Sp, req.Id), &result, getURL())
	return result
}

func createServiceUserBody(userName, userPassword, ip, name, su, sp string) []byte {
	return []byte(strings.TrimSpace(fmt.Sprintf(`
	<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
	<soap:Body>
		<create_service_user xmlns="http://tempuri.org/">
		<user_name>%s</user_name>
		<user_password>%s</user_password>
		<ip>%s</ip>
		<name>%s</name>
		<su>%s</su>
		<sp>%s</sp>
		</create_service_user>
	</soap:Body>
	</soap:Envelope>`, userName, userPassword, ip, name, su, sp)))
}

func CreateServiceUser(req CreateServiceUserRequest) CreateServiceUserResponse {
	var result CreateServiceUserResponse
	soap.SendRequest(createServiceUserBody(req.UserName, req.UserPassword, req.Ip, req.Name, req.Su, req.Sp), &result, getURL())
	return result
}

func delWaybillBody(su, sp string, id int) []byte {
	return []byte(strings.TrimSpace(fmt.Sprintf(`
	<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
	<soap:Body>
		<del_waybill xmlns="http://tempuri.org/">
		<su>%s</su>
		<sp>%s</sp>
		<waybill_id>%d</waybill_id>
		</del_waybill>
	</soap:Body>
	</soap:Envelope>`, su, sp, id)))
}

func DelWaybill(req DelWaybillRequest) DelWaybillResponse {
	var result DelWaybillResponse
	soap.SendRequest(delWaybillBody(req.Su, req.Sp, req.Id), &result, getURL())
	return result
}

func deleteBarCodeBody(su, sp, bar string) []byte {
	return []byte(strings.TrimSpace(fmt.Sprintf(`
	<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
	<soap:Body>
		<delete_bar_code xmlns="http://tempuri.org/">
		<su>%s</su>
		<sp>%s</sp>
		<bar_code>%s</bar_code>
		</delete_bar_code>
	</soap:Body>
	</soap:Envelope>`, su, sp, bar)))
}

func DeleteBarCode(req DeleteBarCodeRequest) DeleteBarCodeResponse {
	var result DeleteBarCodeResponse
	soap.SendRequest(deleteBarCodeBody(req.Su, req.Sp, req.BarCode), &result, getURL())
	return result
}

func deleteCarNumbersBody(su, sp, number string) []byte {
	return []byte(strings.TrimSpace(fmt.Sprintf(`
	<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
	<soap:Body>
		<delete_car_numbers xmlns="http://tempuri.org/">
		<su>%s</su>
		<sp>%s</sp>
		<car_number>%s</car_number>
		</delete_car_numbers>
	</soap:Body>
	</soap:Envelope>`, su, sp, number)))
}

func DeleteCarNumbers(req DeleteCarNumbersRequest) DeleteCarNumbersResponse {
	var result DeleteCarNumbersResponse
	soap.SendRequest(deleteCarNumbersBody(req.Su, req.Sp, req.CarNumber), &result, getURL())
	return result
}

func deleteWaybillTamplateBody(su, sp string, id int) []byte {
	return []byte(strings.TrimSpace(fmt.Sprintf(`
	<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
	<soap:Body>
		<del_waybill xmlns="http://tempuri.org/">
		<su>%s</su>
		<sp>%s</sp>
		<waybill_id>%d</waybill_id>
		</del_waybill>
	</soap:Body>
	</soap:Envelope>`, su, sp, id)))
}

func DeleteWaybillTamplate(req DeleteWaybillTamplateRequest) DeleteWaybillTamplateResponse {
	var result DeleteWaybillTamplateResponse
	soap.SendRequest(deleteWaybillTamplateBody(req.Su, req.Sp, req.Id), &result, getURL())
	return result
}

func getAdjustedWaybillBody(su, sp string, id int) []byte {
	return []byte(strings.TrimSpace(fmt.Sprintf(`
	<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
	<soap:Body>
		<get_adjusted_waybill xmlns="http://tempuri.org/">
		<su>%s</su>
		<sp>%s</sp>
		<id>%d</id>
		</get_adjusted_waybill>
	</soap:Body>
	</soap:Envelope>`, su, sp, id)))
}

func GetAdjustedWaybill(req GetAdjustedWaybillRequest) GetAdjustedWaybillResponse {
	var result GetAdjustedWaybillResponse
	soap.SendRequest(getAdjustedWaybillBody(req.Su, req.Sp, req.Id), &result, getURL())
	return result
}
