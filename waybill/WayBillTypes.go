package waybill

import "encoding/xml"

type chekServiceUserXML struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		Data CheckServiceUserResponse `xml:"chek_service_userResponse"`
	}
}

type CheckServiceUserResponse struct {
	ReqResult bool `xml:"chek_service_userResult"`
	UniqId    int  `xml:"un_id"`
	SerUserId int  `xml:"s_user_id"`
}

type closeWaybillXML struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		Data CloseWaybillResponse `xml:"close_waybillResponse"`
	}
}

type CloseWaybillResponse struct {
	ResCode int `xml:"close_waybillResult"`
}

type closeWaybillTransporterXML struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		Data CloseWaybillTransporterResponse `xml:"close_waybill_transporterResponse"`
	}
}

type CloseWaybillTransporterResponse struct {
	ResCode int `xml:"close_waybill_transporterResult"`
}
