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
	Tin              string         `xml:"TIN"`
	SellerTin        string         `xml:"SELLER_TIN"`
	Name             string         `xml:"NAME"`
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
	WName            string         `xml:"W_NAME"`
	UnitId           int            `xml:"UNIT_ID"`
	Quantity         int            `xml:"QUANTITY"`
	Price            float64        `xml:"PRICE"`
	Amount           float64        `xml:"AMOUNT"`
	BarCode          string         `xml:"BAR_CODE"`
	AId              int            `xml:"A_ID"`
	VATType          int            `xml:"VAT_TYPE"`
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
	IsCorrected      int            `xml:"IS_CORRECTED"`
	SellerST         int            `xml:"SELLER_ST"`
	IsMed            int            `xml:"IS_MED"`
}

type AdjustedWaybill struct {
	Id       int    `xml:"ID"`
	Dt       string `xml:"DT"`
	UserName string `xml:"USER_NAME"`
}

type CorrectedWaybill struct {
	WaybillId     int    `xml:"WAYBILL_ID"`
	WaybillNumber string `xml:"WAYBILL_NUMBER"`
	Dt            string `xml:"DT"`
}

type AkcizCode struct {
	Id          int     `xml:"ID"`
	Title       string  `xml:"TITLE"`
	Measurement string  `xml:"MEASUREMENT"`
	SakonKodi   string  `xml:"SAKON_KODI"`
	AkcisGanakv float64 `xml:"AKCIS_GANAKV"`
	StartDate   string  `xml:"START_DATE"`
	EndDate     string  `xml:"END_DATE"`
}

type BarCodeStruct struct {
	Code     string `xml:"CODE"`
	Name     string `xml:"NAME"`
	UnitId   int    `xml:"UNIT_ID"`
	UnitText string `xml:"UNIT_TXT"`
}

type ErrorCodeStruct struct {
	Id   int    `xml:"ID"`
	Text string `xml:"TEXT"`
	Type int    `xml:"TYPE"`
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
	XMLName xml.Name `xml:"tem:close_waybill"`
	Su      string   `xml:"tem:su"`         // mandatory
	Sp      string   `xml:"tem:sp"`         // mandatory
	Id      int      `xml:"tem:waybill_id"` // mandatory
}

type CloseWaybillResponse struct {
	ResCode int `xml:"Body>close_waybillResponse>close_waybillResult"`
}

type CloseWaybillTransporterRequest struct {
	XMLName       xml.Name `xml:"tem:close_waybill_transporter"`
	Su            string   `xml:"tem:su"`                       // mandatory
	Sp            string   `xml:"tem:sp"`                       // mandatory
	Id            int      `xml:"tem:waybill_id"`               // mandatory
	ReceptionInfo string   `xml:"tem:reception_info,omitempty"` // optional?
	RecieverInfo  string   `xml:"tem:receiver_info,omitempty"`  // optional?
	DeliveryDate  string   `xml:"tem:delivery_date"`            // mandatory, datetime package for date handling
}

type CloseWaybillTransporterResponse struct {
	ResCode int `xml:"Body>close_waybill_transporterResponse>close_waybill_transporterResult"`
}

type CloseWaybillVdRequest struct {
	XMLName      xml.Name `xml:"tem:close_waybill_vd"`
	Su           string   `xml:"tem:su"`            // mandatory
	Sp           string   `xml:"tem:sp"`            // mandatory
	DeliveryDate string   `xml:"tem:delivery_date"` // mandatory, datetime package for date handling
	Id           int      `xml:"tem:waybill_id"`    // mandatory
}

type CloseWaybillVdResponse struct {
	ResCode int `xml:"Body>close_waybill_vdResponse>close_waybill_vdResult"`
}

type ConfirmWaybillRequest struct {
	XMLName xml.Name `xml:"tem:confirm_waybill"`
	Su      string   `xml:"tem:su"`         // mandatory
	Sp      string   `xml:"tem:sp"`         // mandatory
	Id      int      `xml:"tem:waybill_id"` // mandatory
}

type ConfirmWaybillResponse struct {
	ResCode int `xml:"Body>confirm_waybillResponse>confirm_waybillResult"`
}

type CreateServiceUserRequest struct {
	XMLName      xml.Name `xml:"tem:create_service_user"`
	UserName     string   `xml:"tem:user_name"`      // mandatory
	UserPassword string   `xml:"tem:user_password"`  // mandatory
	Ip           string   `xml:"tem:ip,omitempty"`   // optional?
	Name         string   `xml:"tem:name,omitempty"` // optional?
	Su           string   `xml:"tem:su"`             // mandatory
	Sp           string   `xml:"tem:sp"`             // mandatory
}

type CreateServiceUserResponse struct {
	Result bool `xml:"Body>create_service_userResponse>create_service_userResult"`
}

type DelWaybillRequest struct {
	XMLName xml.Name `xml:"tem:del_waybill"`
	Su      string   `xml:"tem:su"`         // mandatory
	Sp      string   `xml:"tem:sp"`         // mandatory
	Id      int      `xml:"tem:waybill_id"` // mandatory
}

type DelWaybillResponse struct {
	ResCode int `xml:"Body>del_waybillResponse>del_waybillResult"`
}

type DeleteBarCodeRequest struct {
	XMLName xml.Name `xml:"tem:delete_bar_code"`
	Su      string   `xml:"tem:su"`       // mandatory
	Sp      string   `xml:"tem:sp"`       // mandatory
	BarCode string   `xml:"tem:bar_code"` // mandatory
}

type DeleteBarCodeResponse struct {
	ResCode int `xml:"Body>delete_bar_codeResponse>delete_bar_codeResult"`
}

type DeleteCarNumbersRequest struct {
	XMLName   xml.Name `xml:"tem:delete_car_numbers"`
	Su        string   `xml:"tem:su"`         // mandatory
	Sp        string   `xml:"tem:sp"`         // mandatory
	CarNumber string   `xml:"tem:car_number"` // mandatory
}

type DeleteCarNumbersResponse struct {
	ResCode int `xml:"Body>delete_car_numbersResponse>delete_car_numbersResult"`
}

type DeleteWaybillTamplateRequest struct {
	XMLName xml.Name `xml:"tem:delete_waybill_tamplate"`
	Su      string   `xml:"tem:su"` // mandatory
	Sp      string   `xml:"tem:sp"` // mandatory
	Id      int      `xml:"tem:id"` // mandatory
}

type DeleteWaybillTamplateResponse struct {
	ResCode int `xml:"Body>delete_waybill_tamplateResponse>delete_waybill_tamplateResult"`
}

type GetAdjustedWaybillRequest struct {
	XMLName xml.Name `xml:"tem:get_adjusted_waybill"`
	Su      string   `xml:"tem:su"` // mandatory
	Sp      string   `xml:"tem:sp"` // mandatory
	Id      int      `xml:"tem:id"` // mandatory
}

type GetAdjustedWaybillResponse struct {
	Waybill Waybill `xml:"Body>get_adjusted_waybillResponse>get_adjusted_waybillResult>WAYBILL"`
}

type GetAdjustedWaybillsRequest struct {
	XMLName xml.Name `xml:"tem:get_adjusted_waybills"`
	Su      string   `xml:"tem:su"`         // mandatory
	Sp      string   `xml:"tem:sp"`         // mandatory
	Id      int      `xml:"tem:waybill_id"` // mandatory
}

type GetAdjustedWaybillsResponse struct {
	AdjustedWaybills []AdjustedWaybill `xml:"Body>get_adjusted_waybillsResponse>get_adjusted_waybillsResult>ADJUSTED_WAYBILLS>ADJUSTED_WAYBILL"`
}

type GetAkcizCodesRequest struct {
	XMLName xml.Name `xml:"tem:get_adjusted_waybills"`
	Su      string   `xml:"tem:su"`               // mandatory
	Sp      string   `xml:"tem:sp"`               // mandatory
	SText   string   `xml:"tem:s_text,omitempty"` // optional
}

type GetAkcizCodesResponse struct {
	AkcizCodes []AkcizCode `xml:"Body>get_akciz_codesResponse>get_akciz_codesResult>AKCIZ_CODES>AKCIZ_CODE"`
}

type GetBarCodesRequest struct {
	XMLName xml.Name `xml:"tem:get_bar_codes"`
	Su      string   `xml:"tem:su"`                 // mandatory
	Sp      string   `xml:"tem:sp"`                 // mandatory
	BarCode string   `xml:"tem:bar_code,omitempty"` // optional
}

type GetBarCodesResponse struct {
	ResCode  int             `xml:"Body>get_bar_codesResponse>get_bar_codesResult"`
	BarCodes []BarCodeStruct `xml:"Body>get_bar_codesResponse>bar_codes>BAR_CODES>BAR_CODE"`
}

type GetBuyerWaybillGoodsListRequest struct {
	XMLName       xml.Name `xml:"tem:get_buyer_waybilll_goods_list"`
	Su            string   `xml:"tem:su"`                        // mandatory
	Sp            string   `xml:"tem:sp"`                        // mandatory
	ITypes        string   `xml:"tem:itypes,omitempty"`          // optional
	SellerTin     string   `xml:"tem:seller_tin,omitempty"`      // optional
	Statuses      string   `xml:"tem:statuses,omitempty"`        // optional
	CarNumber     string   `xml:"tem:car_number,omitempty"`      // optional
	BeginDateS    string   `xml:"tem:begin_date_s,omitempty"`    // optional
	BeginDateE    string   `xml:"tem:begin_date_e,omitempty"`    // optional
	CreateDateS   string   `xml:"tem:create_date_s,omitempty"`   // optional
	CreateDateE   string   `xml:"tem:create_date_e,omitempty"`   // optional
	DriverTin     string   `xml:"tem:driver_tin,omitempty"`      // optional
	DeliveryDateS string   `xml:"tem:delivery_date_s,omitempty"` // optional
	DeliveryDateE string   `xml:"tem:delivery_date_e,omitempty"` // optional
	FullAmount    float64  `xml:"tem:full_amount,omitempty"`     // optional
	WaybillNumber string   `xml:"tem:waybill_number,omitempty"`  // optional
	CloseDateS    string   `xml:"tem:close_date_s,omitempty"`    // optional
	CloseDateE    string   `xml:"tem:close_date_e,omitempty"`    // optional
	SUserIds      string   `xml:"tem:s_user_ids,omitempty"`      // optional
	Comment       string   `xml:"tem:comment,omitempty"`         // optional
	IsConfirmed   int      `xml:"tem:is_confirmed,omitempty"`    // optional
}

type GetBuyerWaybillGoodsListResponse struct {
	WaybillList []Waybill `xml:"Body>get_buyer_waybilll_goods_listResponse>get_buyer_waybilll_goods_listResult>WAYBILL_LIST>WAYBILL"`
}

type GetBuyerWaybillsRequest struct {
	XMLName       xml.Name `xml:"tem:get_buyer_waybills"`
	Su            string   `xml:"tem:su"`                        // mandatory
	Sp            string   `xml:"tem:sp"`                        // mandatory
	ITypes        string   `xml:"tem:itypes,omitempty"`          // optional
	SellerTin     string   `xml:"tem:seller_tin,omitempty"`      // optional
	Statuses      string   `xml:"tem:statuses,omitempty"`        // optional
	CarNumber     string   `xml:"tem:car_number,omitempty"`      // optional
	BeginDateS    string   `xml:"tem:begin_date_s,omitempty"`    // optional
	BeginDateE    string   `xml:"tem:begin_date_e,omitempty"`    // optional
	CreateDateS   string   `xml:"tem:create_date_s,omitempty"`   // optional
	CreateDateE   string   `xml:"tem:create_date_e,omitempty"`   // optional
	DriverTin     string   `xml:"tem:driver_tin,omitempty"`      // optional
	DeliveryDateS string   `xml:"tem:delivery_date_s,omitempty"` // optional
	DeliveryDateE string   `xml:"tem:delivery_date_e,omitempty"` // optional
	FullAmount    float64  `xml:"tem:full_amount,omitempty"`     // optional
	WaybillNumber string   `xml:"tem:waybill_number,omitempty"`  // optional
	CloseDateS    string   `xml:"tem:close_date_s,omitempty"`    // optional
	CloseDateE    string   `xml:"tem:close_date_e,omitempty"`    // optional
	SUserIds      string   `xml:"tem:s_user_ids,omitempty"`      // optional
	Comment       string   `xml:"tem:comment,omitempty"`         // optional
	IsConfirmed   int      `xml:"tem:is_confirmed,omitempty"`    // optional
}

type GetBuyerWaybillsResponse struct {
	WaybillList []Waybill `xml:"Body>get_buyer_waybillsResponse>get_buyer_waybillsResult>WAYBILL_LIST>WAYBILL"`
}

type GetBuyerWaybillsExRequest struct {
	XMLName       xml.Name `xml:"tem:get_buyer_waybills_ex"`
	Su            string   `xml:"tem:su"`                        // mandatory
	Sp            string   `xml:"tem:sp"`                        // mandatory
	ITypes        string   `xml:"tem:itypes,omitempty"`          // optional
	SellerTin     string   `xml:"tem:seller_tin,omitempty"`      // optional
	Statuses      string   `xml:"tem:statuses,omitempty"`        // optional
	CarNumber     string   `xml:"tem:car_number,omitempty"`      // optional
	BeginDateS    string   `xml:"tem:begin_date_s,omitempty"`    // optional
	BeginDateE    string   `xml:"tem:begin_date_e,omitempty"`    // optional
	CreateDateS   string   `xml:"tem:create_date_s,omitempty"`   // optional
	CreateDateE   string   `xml:"tem:create_date_e,omitempty"`   // optional
	DriverTin     string   `xml:"tem:driver_tin,omitempty"`      // optional
	DeliveryDateS string   `xml:"tem:delivery_date_s,omitempty"` // optional
	DeliveryDateE string   `xml:"tem:delivery_date_e,omitempty"` // optional
	FullAmount    float64  `xml:"tem:full_amount,omitempty"`     // optional
	WaybillNumber string   `xml:"tem:waybill_number,omitempty"`  // optional
	CloseDateS    string   `xml:"tem:close_date_s,omitempty"`    // optional
	CloseDateE    string   `xml:"tem:close_date_e,omitempty"`    // optional
	SUserIds      string   `xml:"tem:s_user_ids,omitempty"`      // optional
	Comment       string   `xml:"tem:comment,omitempty"`         // optional
	IsConfirmed   int      `xml:"tem:is_confirmed,omitempty"`    // optional
}

type GetBuyerWaybillsExResponse struct {
	WaybillList []Waybill `xml:"Body>get_buyer_waybills_exResponse>get_buyer_waybills_exResult>WAYBILL_LIST>WAYBILL"`
}

type GetCWaybillRequest struct {
	XMLName xml.Name `xml:"tem:get_c_waybill"`
	Su      string   `xml:"tem:su"`   // mandatory
	Sp      string   `xml:"tem:sp"`   // mandatory
	SDt     string   `xml:"tem:s_dt"` // mandatory
	EDt     string   `xml:"tem:e_dt"` // mandatory
}

type GetCWaybillResponse struct {
	ResCode  int                `xml:"Body>get_c_waybillResponse>get_c_waybillResult"`
	Waybills []CorrectedWaybill `xml:"Body>get_c_waybillResponse>waybill_tamplates>CORRECTED_WAYBILLS>CORRECTED_WAYBILL"`
}

type GetCarNumbersRequest struct {
	XMLName xml.Name `xml:"tem:get_car_numbers"`
	Su      string   `xml:"tem:su"` // mandatory
	Sp      string   `xml:"tem:sp"` // mandatory
}

type GetCarNumbersResponse struct {
	ResCode    int      `xml:"Body>get_car_numbersResponse>get_car_numbersResult"`
	CarNumbers []string `xml:"Body>get_car_numbersResponse>car_numbers>CAR_NUMBERS>CAR_NUMBER"`
}

type GetErrorCodesRequest struct {
	XMLName xml.Name `xml:"tem:get_error_codes"`
	Su      string   `xml:"tem:su"` // mandatory
	Sp      string   `xml:"tem:sp"` // mandatory
}

type GetErrorCodesResponse struct {
	ErrorCodes []ErrorCodeStruct `xml:"Body>get_error_codesResponse>get_error_codesResult>ERROR_CODES>ERROR_CODE"`
}

type GetNameFromTinRequest struct {
	XMLName xml.Name `xml:"tem:get_name_from_tin"`
	Su      string   `xml:"tem:su"`  // mandatory
	Sp      string   `xml:"tem:sp"`  // mandatory
	Tin     string   `xml:"tem:tin"` // mandatory
}

type GetNameFromTinResponse struct {
	Result string `xml:"Body>get_name_from_tinResponse>get_name_from_tinResult"`
}

type GetPayerTypeFromUnIdRequest struct {
	XMLName xml.Name `xml:"tem:get_payer_type_from_un_id"`
	Su      string   `xml:"tem:su"`    // mandatory
	Sp      string   `xml:"tem:sp"`    // mandatory
	UnId    int      `xml:"tem:un_id"` // mandatory
}

type GetPayerTypeFromUnIdResponse struct {
	Result int `xml:"Body>get_payer_type_from_un_idResponse>get_payer_type_from_un_idResult"`
}

//Depricated?
type GetPrintPDFRequest struct {
	XMLName   xml.Name `xml:"tem:get_print_pdf"`
	Su        string   `xml:"tem:su"`         // mandatory
	Sp        string   `xml:"tem:sp"`         // mandatory
	WaybillId int      `xml:"tem:waybill_id"` // mandatory
}

//Depricated?
type GetPrintPDFResponse struct {
	Result string `xml:"Body>get_print_pdfResponse>get_print_pdfResult"`
}

//Dummy struct
type GetServerTimeRequest struct {
	XMLName xml.Name `xml:"tem:get_server_time"`
}

type GetServerTimeResponse struct {
	Result string `xml:"Body>get_server_timeResponse>get_server_timeResult"`
}
