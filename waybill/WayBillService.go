package waybill

import (
	"GE_Revenue_Services/soap"
	"fmt"
	"strings"
)

const serviceURL string = "https://services.rs.ge/WayBillService/WayBillService.asmx"

func GetURL() string {
	return serviceURL
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
	rawData := soap.SendRequest(req, &response, GetURL())
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
	rawData := soap.SendRequest(req, &response, GetURL())
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

func CloseWaybillTransporter(req CloseWaybillTransporterRequest) (CloseWaybillTransporterResponse, soap.RawData) {
	var response CloseWaybillTransporterResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
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

func CloseWaybillVd(req CloseWaybillVdRequest) (CloseWaybillVdResponse, soap.RawData) {
	var response CloseWaybillVdResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
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

func ConfirmWaybill(req ConfirmWaybillRequest) (ConfirmWaybillResponse, soap.RawData) {
	var response ConfirmWaybillResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
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

func CreateServiceUser(req CreateServiceUserRequest) (CreateServiceUserResponse, soap.RawData) {
	var response CreateServiceUserResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
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

func DelWaybill(req DelWaybillRequest) (DelWaybillResponse, soap.RawData) {
	var response DelWaybillResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
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

func DeleteBarCode(req DeleteBarCodeRequest) (DeleteBarCodeResponse, soap.RawData) {
	var response DeleteBarCodeResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
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

func DeleteCarNumbers(req DeleteCarNumbersRequest) (DeleteCarNumbersResponse, soap.RawData) {
	var response DeleteCarNumbersResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
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

func DeleteWaybillTamplate(req DeleteWaybillTamplateRequest) (DeleteWaybillTamplateResponse, soap.RawData) {
	var response DeleteWaybillTamplateResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
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

func GetAdjustedWaybill(req GetAdjustedWaybillRequest) (GetAdjustedWaybillResponse, soap.RawData) {
	var response GetAdjustedWaybillResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
}

func getAdjustedWaybillsBody(su, sp string, id int) []byte {
	return []byte(strings.TrimSpace(fmt.Sprintf(`
	<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
	<soap:Body>
		<get_adjusted_waybills xmlns="http://tempuri.org/">
			<su>%s</su>
			<sp>%s</sp>
			<waybill_id>%d</waybill_id>
		</get_adjusted_waybills>
	</soap:Body>
	</soap:Envelope>`, su, sp, id)))
}

func GetAdjustedWaybills(req GetAdjustedWaybillsRequest) (GetAdjustedWaybillsResponse, soap.RawData) {
	var response GetAdjustedWaybillsResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
}

func GetAkcizCodes(req GetAkcizCodesRequest) (GetAkcizCodesResponse, soap.RawData) {
	var response GetAkcizCodesResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
}

func GetBarCodes(req GetBarCodesRequest) (GetBarCodesResponse, soap.RawData) {
	var response GetBarCodesResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
}

func GetBuyerWaybillGoodsList(req GetBuyerWaybillGoodsListRequest) (GetBuyerWaybillGoodsListResponse, soap.RawData) {
	var response GetBuyerWaybillGoodsListResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
}

func GetBuyerWaybills(req GetBuyerWaybillsRequest) (GetBuyerWaybillsResponse, soap.RawData) {
	var response GetBuyerWaybillsResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
}

func GetBuyerWaybillsEx(req GetBuyerWaybillsExRequest) (GetBuyerWaybillsExResponse, soap.RawData) {
	var response GetBuyerWaybillsExResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
}

func GetCWaybill(req GetCWaybillRequest) (GetCWaybillResponse, soap.RawData) {
	var response GetCWaybillResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
}

func GetCarNumbers(req GetCarNumbersRequest) (GetCarNumbersResponse, soap.RawData) {
	var response GetCarNumbersResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
}

func GetErrorCodes(req GetErrorCodesRequest) (GetErrorCodesResponse, soap.RawData) {
	var response GetErrorCodesResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
}

func GetNameFromTin(req GetNameFromTinRequest) (GetNameFromTinResponse, soap.RawData) {
	var response GetNameFromTinResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
}

func GetPayerTypeFromUnId(req GetPayerTypeFromUnIdRequest) (GetPayerTypeFromUnIdResponse, soap.RawData) {
	var response GetPayerTypeFromUnIdResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
}

// Depricated?
func GetPrintPDF(req GetPrintPDFRequest) (GetPrintPDFResponse, soap.RawData) {
	var response GetPrintPDFResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
}

func GetServerTime() (GetServerTimeResponse, soap.RawData) {
	var req GetServerTimeRequest
	var response GetServerTimeResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
}

func GetServiceUsers(req GetServiceUsersRequest) (GetServiceUsersResponse, soap.RawData) {
	var response GetServiceUsersResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
}

func GetTinFromUnId(req GetTinFromUnIdRequest) (GetTinFromUnIdResponse, soap.RawData) {
	var response GetTinFromUnIdResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
}

func GetTransTypes(req GetTransTypesRequest) (GetTransTypesResponse, soap.RawData) {
	var response GetTransTypesResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
}

func GetTransporterWaybills(req GetTransporterWaybillsRequest) (GetTransporterWaybillsResponse, soap.RawData) {
	var response GetTransporterWaybillsResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
}

func GetWaybill(req GetWaybillRequest) (GetWaybillResponse, soap.RawData) {
	var response GetWaybillResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
}

func GetWaybillByNumber(req GetWaybillByNumberRequest) (GetWaybillByNumberResponse, soap.RawData) {
	var response GetWaybillByNumberResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
}

func GetWaybillGoodsList(req GetWaybillGoodsListRequest) (GetWaybillGoodsListResponse, soap.RawData) {
	var response GetWaybillGoodsListResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
}

func GetWaybillTamplate(req GetWaybillTamplateRequest) (GetWaybillTamplateResponse, soap.RawData) {
	var response GetWaybillTamplateResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
}

func GetWaybillTamplates(req GetWaybillTamplatesRequest) (GetWaybillTamplatesResponse, soap.RawData) {
	var response GetWaybillTamplatesResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
}

func GetWaybillTypes(req GetWaybillTypesRequest) (GetWaybillTypesResponse, soap.RawData) {
	var response GetWaybillTypesResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
}

func GetWaybillUnits(req GetWaybillUnitsRequest) (GetWaybillUnitsResponse, soap.RawData) {
	var response GetWaybillUnitsResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
}

func GetWaybills(req GetWaybillsRequest) (GetWaybillsResponse, soap.RawData) {
	var response GetWaybillsResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
}

func GetWaybillsEx(req GetWaybillsExRequest) (GetWaybillsExResponse, soap.RawData) {
	var response GetWaybillsExResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
}

func GetWoodTypes(req GetWoodTypesRequest) (GetWoodTypesResponse, soap.RawData) {
	var response GetWoodTypesResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
}

func IsVATPayer(req IsVATPayerRequest) (IsVATPayerResponse, soap.RawData) {
	var response IsVATPayerResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
}

func IsVATPayerTin(req IsVATPayerTinRequest) (IsVATPayerTinResponse, soap.RawData) {
	var response IsVATPayerTinResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
}

func RefWaybill(req RefWaybillRequest) (RefWaybillResponse, soap.RawData) {
	var response RefWaybillResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
}

func RefWaybillVd(req RefWaybillVdRequest) (RefWaybillVdResponse, soap.RawData) {
	var response RefWaybillVdResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
}

func RejectWaybill(req RejectWaybillRequest) (RejectWaybillResponse, soap.RawData) {
	var response RejectWaybillResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
}

func SaveBarCode(req SaveBarCodeRequest) (SaveBarCodeResponse, soap.RawData) {
	var response SaveBarCodeResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
}

func SaveCarNumbers(req SaveCarNumbersRequest) (SaveCarNumbersResponse, soap.RawData) {
	var response SaveCarNumbersResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
}

func SaveInvoice(req SaveInvoiceRequest) (SaveInvoiceResponse, soap.RawData) {
	var response SaveInvoiceResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
}

func SaveWaybill(req SaveWaybillRequest) (SaveWaybillResponse, soap.RawData) {
	var response SaveWaybillResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
}

func SaveWaybillTamplate(req SaveWaybillTamplateRequest) (SaveWaybillTamplateResponse, soap.RawData) {
	var response SaveWaybillTamplateResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
}

func SaveWaybillTransporter(req SaveWaybillTransporterRequest) (SaveWaybillTransporterResponse, soap.RawData) {
	var response SaveWaybillTransporterResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
}

func SendWaybillVd(req SendWaybillVdRequest) (SendWaybillVdResponse, soap.RawData) {
	var response SendWaybillVdResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
}

func SendWaybill(req SendWaybillRequest) (SendWaybillResponse, soap.RawData) {
	var response SendWaybillResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
}

func SendWaybillTransporter(req SendWaybillTransporterRequest) (SendWaybillTransporterResponse, soap.RawData) {
	var response SendWaybillTransporterResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
}

func UpdateServiceUser(req UpdateServiceUserRequest) (UpdateServiceUserResponse, soap.RawData) {
	var response UpdateServiceUserResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
}

func WhatIsMyIP() (WhatIsMyIPResponse, soap.RawData) {
	var req WhatIsMyIPRequest
	var response WhatIsMyIPResponse
	rawData := soap.SendRequest(req, &response, GetURL())
	return response, rawData
}
