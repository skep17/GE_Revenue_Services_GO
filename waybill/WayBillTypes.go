package waybill

import "encoding/xml"

type ServiceUser struct {
	XMLName xml.Name
	Body    struct {
		XMLName           xml.Name
		ListUsersResponse struct {
			XMLName   xml.Name
			ReqResult bool   `xml:"chek_service_userResult"`
			UniqId    string `xml:"un_id"`
			SerUserId string `xml:"s_user_id"`
		} `xml:"chek_service_userResponse"`
	}
}
