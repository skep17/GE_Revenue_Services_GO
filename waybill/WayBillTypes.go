package waybill

import "encoding/xml"

type SubWaybill struct {
	Id            int     `xml:"ID"`
	WaybillNumber string  `xml:"WAYBILL_NUMBER"`
	BuyerTin      string  `xml:"BUYER_TIN"`
	BuyerName     string  `xml:"BUYER_NAME"`
	FullAmount    float64 `xml:"FULL_AMOUNT"`
	Status        int     `xml:"STATUS"`
	TotalQuantity int     `xml:"TOTAL_QUANTITY"`
	Type          int     `xml:"TYPE"`
}

type Goods struct {
	Id           int     `xml:"ID"`
	WName        string  `xml:"W_NAME"`
	UnitId       int     `xml:"UNIT_ID"`
	Quantity     int     `xml:"QUANTITY"`
	Price        float64 `xml:"PRICE"`
	Amount       float64 `xml:"AMOUNT"`
	BarCode      string  `xml:"BAR_CODE"`
	AId          int     `xml:"A_ID"`
	VATType      int     `xml:"VAT_TYPE"`
	Status       int     `xml:"STATUS"`
	WoodLabel    string  `xml:"WOOD_LABEL"`
	WId          int     `xml:"W_ID"`
	LastChangeDt string  `xml:"LAST_CHANGE_DT"`
	QuantityF    int     `xml:"QUANTITY_F"`
}

type WoodDocument struct {
	Id      int    `xml:"ID"`
	DocN    string `xml:"DOC_N"`
	DocDate string `xml:"DOC_DATE"`
	DocDesc string `xml:"DOC_DESC"`
}

type Waybill struct {
	SubWaybills      []SubWaybill   `xml:"SUB_WAYBILLS>SUB_WAYBILL"`
	GoodsList        []Goods        `xml:"GOODS_LIST>GOODS"`
	WoodDocsList     []WoodDocument `xml:"WOOD_DOCS_LIST>WOODDOCUMENT"`
	Id               int            `xml:"ID"`
	Type             int            `xml:"TYPE"`
	CreateDate       string         `xml:"CREATE_DATE"`
	BuyerTin         string         `xml:"BUYER_TIN"`
	ChekBuyerTin     int            `xml:"CHEK_BUYER_TIN"`
	BuyerName        string         `xml:"BUYER_NAME"`
	StartAddress     string         `xml:"START_ADDRESS"`
	EndAddress       string         `xml:"END_ADDRESS"`
	DriverTin        string         `xml:"DRIVER_TIN"`
	ChekDriverTin    int            `xml:"CHEK_DRIVER_TIN"`
	DriverName       string         `xml:"DRIVER_NAME"`
	TransportCoast   float64        `xml:"TRANSPORT_COAST"`
	ReceptionInfo    string         `xml:"RECEPTION_INFO"`
	RecieverInfo     string         `xml:"RECEIVER_INFO"`
	DeliveryDate     string         `xml:"DELIVERY_DATE"`
	Status           int            `xml:"STATUS"`
	SelerUnId        int            `xml:"SELER_UN_ID"`
	ActivateDate     string         `xml:"ACTIVATE_DATE"`
	ParId            int            `xml:"PAR_ID"`
	FullAmount       float64        `xml:"FULL_AMOUNT"`
	FullAmountTxt    string         `xml:"FULL_AMOUNT_TXT"`
	CarNumber        string         `xml:"CAR_NUMBER"`
	WaybillNumber    string         `xml:"WAYBILL_NUMBER"`
	CloseDate        string         `xml:"CLOSE_DATE"`
	SUserId          int            `xml:"S_USER_ID"`
	BeginDate        string         `xml:"BEGIN_DATE"`
	TranCostPayer    string         `xml:"TRAN_COST_PAYER"`
	TransId          int            `xml:"TRANS_ID"`
	TransTxt         string         `xml:"TRANS_TXT"`
	Comment          string         `xml:"COMMENT"`
	ReceiverView     int            `xml:"RECEIVER_VIEW"`
	IsConfirmed      int            `xml:"IS_CONFIRMED"`
	ConfirmationDate string         `xml:"CONFIRMATION_DATE"`
	SellerTin        string         `xml:"SELLER_TIN"`
	SellerName       string         `xml:"SELLER_NAME"`
	BuyerUnId        int            `xml:"BUYER_UN_ID"`
	InvoiceId        int            `xml:"INVOICE_ID"`
	Category         int            `xml:"CATEGORY"`
	OriginType       int            `xml:"ORIGIN_TYPE"`
	OriginText       string         `xml:"ORIGIN_TEXT"`
	CancelDate       string         `xml:"CANCEL_DATE"`
	SellerStatus     int            `xml:"SELLER_STATUS"`
	BuyerStatus      int            `xml:"BUYER_STATUS"`
	BuyerSUserId     int            `xml:"BUYER_S_USER_ID"`
	CorrectionDate   string         `xml:"CORRECTION_DATE"`
	TotalQuantity    int            `xml:"TOTAL_QUANTITY"`
	TransporterTin   string         `xml:"TRANSPORTER_TIN"`
	TransporterUnId  int            `xml:"TRANSPORTER_UN_ID"`
	TransporterName  string         `xml:"TRANSPORTER_NAME"`
	TransSubuserId   int            `xml:"TRANS_SUBUSER_ID"`
	WoodDoc          int            `xml:"WOOD_DOC"`
	WoodDocN         string         `xml:"WOOD_DOC_N"`
	WoodDocDate      string         `xml:"WOOD_DOC_DATE"`
	WoodLabels       string         `xml:"WOOD_LABELS"`
	ConfEmpId        int            `xml:"CONF_EMP_ID"`
	ConfEmpDate      string         `xml:"CONF_EMP_DATE"`
	CustEmpId        int            `xml:"CUST_EMP_ID"`
	CustDate         string         `xml:"CUST_DATE"`
	CustId           int            `xml:"CUST_ID"`
	CustStatus       int            `xml:"CUST_STATUS"`
	CustName         string         `xml:"CUST_NAME"`
	LastChangeDt     string         `xml:"LAST_CHANGE_DT"`
	VerifiedEmpId    int            `xml:"VERIFIED_EMP_ID"`
	VerifiedEmpDate  string         `xml:"VERIFIED_EMP_DATE"`
	IsMed            int            `xml:"IS_MED"`
}

type CheckServiceUserRequest struct {
	XMLName xml.Name `xml:"tem:chek_service_user"`
	Su      string   `xml:"tem:su"` // mandatory
	Sp      string   `xml:"tem:sp"` // mandatory
}

type CheckServiceUserResponse struct {
	ReqResult bool `xml:"Body>chek_service_userResponse>chek_service_userResult"`
	UniqId    int  `xml:"Body>chek_service_userResponse>un_id"`
	SerUserId int  `xml:"Body>chek_service_userResponse>s_user_id"`
}

type CloseWaybillRequest struct {
	Su string // mandatory
	Sp string // mandatory
	Id int    // mandatory
}

type CloseWaybillResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	XMLBody string
	ResCode int `xml:"Body>close_waybillResponse>close_waybillResult"`
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
	XMLName xml.Name `xml:"Envelope"`
	XMLBody string
	ResCode int `xml:"Body>close_waybill_transporterResponse>close_waybill_transporterResult"`
}

type CloseWaybillVdRequest struct {
	Su           string // mandatory
	Sp           string // mandatory
	DeliveryDate string // mandatory, datetime package for date handling
	Id           int    // mandatory
}

type CloseWaybillVdResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	XMLBody string
	ResCode int `xml:"Body>close_waybill_vdResponse>close_waybill_vdResult"`
}

type ConfirmWaybillRequest struct {
	Su string // mandatory
	Sp string // mandatory
	Id int    // mandatory
}

type ConfirmWaybillResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	XMLBody string
	ResCode int `xml:"Body>confirm_waybillResponse>confirm_waybillResult"`
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
	XMLName xml.Name `xml:"Envelope"`
	XMLBody string
	Result  bool `xml:"Body>create_service_userResponse>create_service_userResult"`
}

type DelWaybillRequest struct {
	Su string // mandatory
	Sp string // mandatory
	Id int    // mandatory
}

type DelWaybillResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	XMLBody string
	ResCode int `xml:"Body>del_waybillResponse>del_waybillResult"`
}

type DeleteBarCodeRequest struct {
	Su      string // mandatory
	Sp      string // mandatory
	BarCode string // mandatory
}

type DeleteBarCodeResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	XMLBody string
	ResCode int `xml:"Body>delete_bar_codeResponse>delete_bar_codeResult"`
}

type DeleteCarNumbersRequest struct {
	Su        string // mandatory
	Sp        string // mandatory
	CarNumber string // mandatory
}

type DeleteCarNumbersResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	XMLBody string
	ResCode int `xml:"Body>delete_car_numbersResponse>delete_car_numbersResult"`
}

type DeleteWaybillTamplateRequest struct {
	Su string // mandatory
	Sp string // mandatory
	Id int    // mandatory
}

type DeleteWaybillTamplateResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	XMLBody string
	ResCode int `xml:"Body>delete_waybill_tamplateResponse>delete_waybill_tamplateResult"`
}

type GetAdjustedWaybillRequest struct {
	Su string // mandatory
	Sp string // mandatory
	Id int    // mandatory
}

type GetAdjustedWaybillResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	XMLBody string
	Waybill Waybill `xml:"Body>get_adjusted_waybillResponse>get_adjusted_waybillResult>WAYBILL"`
}
