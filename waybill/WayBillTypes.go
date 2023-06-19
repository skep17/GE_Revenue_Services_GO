package waybill

import (
	"encoding/xml"
)

type chekServiceUserEnv struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		Data CheckServiceUserResponse `xml:"chek_service_userResponse"`
	}
}

type CheckServiceUserRequest struct {
	Su string // mandatory
	Sp string // mandatory
}

type CheckServiceUserResponse struct {
	ReqResult bool `xml:"chek_service_userResult"`
	UniqId    int  `xml:"un_id"`
	SerUserId int  `xml:"s_user_id"`
}

type closeWaybillEnv struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		Data CloseWaybillResponse `xml:"close_waybillResponse"`
	}
}

type CloseWaybillRequest struct {
	Su string // mandatory
	Sp string // mandatory
	Id int    // mandatory
}

type CloseWaybillResponse struct {
	ResCode int `xml:"close_waybillResult"`
}

type closeWaybillTransporterEnv struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		Data CloseWaybillTransporterResponse `xml:"close_waybill_transporterResponse"`
	}
}

type CloseWaybillTransporterRequest struct {
	Su            string // mandatory
	Sp            string // mandatory
	Id            int    // mandatory
	ReceptionInfo string // optional?
	RecieverInfo  string // optional?
	DeliveryDate  string // mandatory, datetime package for date handling
}

type CloseWaybillTransporterResponse struct {
	ResCode int `xml:"close_waybill_transporterResult"`
}

type CloseWaybillVdEnv struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		Data CloseWaybillVdResponse `xml:"close_waybill_vdResponse"`
	}
}

type CloseWaybillVdRequest struct {
	Su           string // mandatory
	Sp           string // mandatory
	DeliveryDate string // mandatory, datetime package for date handling
	Id           int    // mandatory
}

type CloseWaybillVdResponse struct {
	ResCode int `xml:"close_waybill_vdResult"`
}

type ConfirmWaybillEnv struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		Data ConfirmWaybillResponse `xml:"confirm_waybillResponse"`
	}
}

type ConfirmWaybillRequest struct {
	Su string // mandatory
	Sp string // mandatory
	Id int    // mandatory
}

type ConfirmWaybillResponse struct {
	ResCode int `xml:"confirm_waybillResult"`
}

type CreateServiceUserEnv struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		Data CreateServiceUserResponse `xml:"create_service_userResponse"`
	}
}

type CreateServiceUserRequest struct {
	UserName     string // mandatory
	UserPassword string // mandatory
	Ip           string // optional?
	Name         string // optional?
	Su           string // mandatory
	Sp           string // mandatory
}

type CreateServiceUserResponse struct {
	result bool `xml:"create_service_userResult"`
}

type DelWaybillEnv struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		Data DelWaybillResponse `xml:"del_waybillResponse"`
	}
}

type DelWaybillRequest struct {
	Su string // mandatory
	Sp string // mandatory
	Id int    // mandatory
}

type DelWaybillResponse struct {
	ResCode int `xml:"del_waybillResult"`
}

type DeleteBarCodeEnv struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		Data DeleteBarCodeResponse `xml:"delete_bar_codeResponse"`
	}
}

type DeleteBarCodeRequest struct {
	Su      string // mandatory
	Sp      string // mandatory
	BarCode string // mandatory
}

type DeleteBarCodeResponse struct {
	ResCode int `xml:"delete_bar_codeResult"`
}

type DeleteCarNumbersEnv struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		Data DeleteCarNumbersResponse `xml:"delete_car_numbersResponse"`
	}
}

type DeleteCarNumbersRequest struct {
	Su        string // mandatory
	Sp        string // mandatory
	CarNumber string // mandatory
}

type DeleteCarNumbersResponse struct {
	ResCode int `xml:"delete_car_numbersResult"`
}

type DeleteWaybillTamplateEnv struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		Data DeleteWaybillTamplateResponse `xml:"delete_waybill_tamplateResponse"`
	}
}

type DeleteWaybillTamplateRequest struct {
	Su string // mandatory
	Sp string // mandatory
	Id int    // mandatory
}

type DeleteWaybillTamplateResponse struct {
	ResCode int `xml:"delete_waybill_tamplateResult"`
}
