package waybill

import (
	"fmt"
	"strings"
)

func getPath() string {
	return "https://services.rs.ge/WayBillService/WayBillService.asmx"
}

func ChekServiceUser(su, sp string) []byte {
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
