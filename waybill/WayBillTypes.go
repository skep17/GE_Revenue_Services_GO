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
	Error        int     `xml:"ERROR"`
	Id           int     `xml:"ID"`
	WName        string  `xml:"W_NAME"`
	UnitId       int     `xml:"UNIT_ID"`
	Quantity     int     `xml:"QUANTITY"`
	QuantityExt  string  `xml:"QUANTITY_EXT"`
	Price        float64 `xml:"PRICE"`
	Amount       float64 `xml:"AMOUNT"`
	BarCode      string  `xml:"BAR_CODE"`
	UnitTxt      string  `xml:"UNIT_TXT"`
	AId          int     `xml:"A_ID"`
	VATType      int     `xml:"VAT_TYPE"`
	Status       int     `xml:"STATUS"`
	GoodType     int     `xml:"GOOD_TYPE"`
	WoodLabel    string  `xml:"WOOD_LABEL"`
	WId          int     `xml:"W_ID"`
	WType        int     `xml:"W_TYPE"`
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
	PageId            string             `xml:"PAGE_ID,omitempty"`
	SubWaybills       []SubWaybill       `xml:"SUB_WAYBILLS>SUB_WAYBILL,omitempty"`
	GoodsList         []Goods            `xml:"GOODS_LIST>GOODS,omitempty"`
	GoodsCount        int                `xml:"GOODS_COUNT,omitempty"`
	CorrectedWaybills []CorrectedWaybill `xml:"CorrectedWaybills>CorrectedWaybil,omitempty"`
	WoodDocsList      []WoodDocument     `xml:"WOOD_DOCS_LIST>WOODDOCUMENT,omitempty"`
	Id                int                `xml:"ID"`
	Type              int                `xml:"TYPE,omitempty"`
	CreateDate        string             `xml:"CREATE_DATE"`
	BuyerTin          string             `xml:"BUYER_TIN,omitempty"`
	ChekBuyerTin      int                `xml:"CHEK_BUYER_TIN,omitempty"`
	BuyerName         string             `xml:"BUYER_NAME,omitempty"`
	StartAddress      string             `xml:"START_ADDRESS,omitempty"`
	EndAddress        string             `xml:"END_ADDRESS,omitempty"`
	DriverTin         string             `xml:"DRIVER_TIN,omitempty"`
	ChekDriverTin     int                `xml:"CHEK_DRIVER_TIN,omitempty"`
	DriverName        string             `xml:"DRIVER_NAME,omitempty"`
	TransportCoast    float64            `xml:"TRANSPORT_COAST,omitempty"`
	ReceptionInfo     string             `xml:"RECEPTION_INFO,omitempty"`
	ReceiverInfo      string             `xml:"RECEIVER_INFO,omitempty"`
	DeliveryDate      string             `xml:"DELIVERY_DATE,omitempty"`
	Status            int                `xml:"STATUS"`
	SelerUnId         int                `xml:"SELER_UN_ID,omitempty"`
	ActivateDate      string             `xml:"ACTIVATE_DATE,omitempty"`
	ParId             int                `xml:"PAR_ID,omitempty"`
	FullAmount        float64            `xml:"FULL_AMOUNT"`
	FullAmountTxt     string             `xml:"FULL_AMOUNT_TXT,omitempty"`
	CarNumber         string             `xml:"CAR_NUMBER,omitempty"`
	WaybillNumber     string             `xml:"WAYBILL_NUMBER"`
	CloseDate         string             `xml:"CLOSE_DATE,omitempty"`
	SUserId           int                `xml:"S_USER_ID,omitempty"`
	BeginDate         string             `xml:"BEGIN_DATE,omitempty"`
	TranCostPayer     string             `xml:"TRAN_COST_PAYER,omitempty"`
	TransId           int                `xml:"TRANS_ID,omitempty"`
	TransTxt          string             `xml:"TRANS_TXT,omitempty"`
	Comment           string             `xml:"COMMENT,omitempty"`
	ReceiverView      int                `xml:"RECEIVER_VIEW,omitempty"`
	IsConfirmed       int                `xml:"IS_CONFIRMED,omitempty"`
	ConfirmationDate  string             `xml:"CONFIRMATION_DATE,omitempty"`
	Tin               string             `xml:"TIN,omitempty"`
	SellerTin         string             `xml:"SELLER_TIN,omitempty"`
	Name              string             `xml:"NAME,omitempty"`
	SellerName        string             `xml:"SELLER_NAME,omitempty"`
	BuyerUnId         int                `xml:"BUYER_UN_ID,omitempty"`
	InvoiceId         int                `xml:"INVOICE_ID,omitempty"`
	Category          int                `xml:"CATEGORY,omitempty"`
	OriginType        int                `xml:"ORIGIN_TYPE,omitempty"`
	OriginText        string             `xml:"ORIGIN_TEXT,omitempty"`
	CancelDate        string             `xml:"CANCEL_DATE,omitempty"`
	SellerStatus      int                `xml:"SELLER_STATUS,omitempty"`
	BuyerStatus       int                `xml:"BUYER_STATUS,omitempty"`
	BuyerSt           int                `xml:"BUYER_ST,omitempty"`
	BuyerSUserId      int                `xml:"BUYER_S_USER_ID,omitempty"`
	CorrectionDate    string             `xml:"CORRECTION_DATE,omitempty"`
	TotalQuantity     int                `xml:"TOTAL_QUANTITY,omitempty"`
	TransporterTin    string             `xml:"TRANSPORTER_TIN,omitempty"`
	TransporterUnId   int                `xml:"TRANSPORTER_UN_ID,omitempty"`
	TransporterName   string             `xml:"TRANSPORTER_NAME,omitempty"`
	TransSubuserId    int                `xml:"TRANS_SUBUSER_ID,omitempty"`
	WoodDoc           int                `xml:"WOOD_DOC,omitempty"`
	WoodDocN          string             `xml:"WOOD_DOC_N,omitempty"`
	WoodDocDate       string             `xml:"WOOD_DOC_DATE,omitempty"`
	WoodLabels        string             `xml:"WOOD_LABELS,omitempty"`
	WName             string             `xml:"W_NAME,omitempty"`
	UnitId            int                `xml:"UNIT_ID,omitempty"`
	Quantity          int                `xml:"QUANTITY,omitempty"`
	Price             float64            `xml:"PRICE,omitempty"`
	Amount            float64            `xml:"AMOUNT,omitempty"`
	BarCode           string             `xml:"BAR_CODE,omitempty"`
	AId               int                `xml:"A_ID,omitempty"`
	VATType           int                `xml:"VAT_TYPE,omitempty"`
	ConfEmpId         int                `xml:"CONF_EMP_ID,omitempty"`
	ConfEmpDate       string             `xml:"CONF_EMP_DATE,omitempty"`
	CustEmpId         int                `xml:"CUST_EMP_ID,omitempty"`
	CustDate          string             `xml:"CUST_DATE,omitempty"`
	CustId            int                `xml:"CUST_ID,omitempty"`
	CustStatus        int                `xml:"CUST_STATUS,omitempty"`
	CustName          string             `xml:"CUST_NAME,omitempty"`
	LastChangeDt      string             `xml:"LAST_CHANGE_DT,omitempty"`
	VerifiedEmpId     int                `xml:"VERIFIED_EMP_ID,omitempty"`
	VerifiedEmpDate   string             `xml:"VERIFIED_EMP_DATE,omitempty"`
	IsCorrected       int                `xml:"IS_CORRECTED,omitempty"`
	SellerST          int                `xml:"SELLER_ST,omitempty"`
	IsMed             int                `xml:"IS_MED"`
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

type ServiceUserStruct struct {
	Id       int    `xml:"ID"`
	UserName string `xml:"USER_NAME"`
	UnId     int    `xml:"UN_ID"`
	IP       string `xml:"IP"`
	Name     string `xml:"NAME"`
}

type TypeStruct struct {
	Id          int    `xml:"ID"`
	Name        string `xml:"NAME"`
	Description string `xml:"DESCRIPTION"`
}

type WaybillTamplate struct {
	Id         int     `xml:"ID"`
	Name       string  `xml:"NAME"`
	Type       string  `xml:"TYPE"`
	BuyerTin   string  `xml:"BUYER_TIN"`
	BuyerName  string  `xml:"BUYER_NAME"`
	FullAmount float64 `xml:"FULL_AMOUNT"`
	CarNumber  string  `xml:"CAR_NUMBER"`
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
	ReceiverInfo  string   `xml:"tem:receiver_info,omitempty"`  // optional?
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

type GetServiceUsersRequest struct {
	XMLName xml.Name `xml:"tem:get_service_users"`
	Su      string   `xml:"tem:su"` // mandatory
	Sp      string   `xml:"tem:sp"` // mandatory
}

type GetServiceUsersResponse struct {
	ServiceUsers []ServiceUserStruct `xml:"Body>get_service_usersResponse>get_service_usersResult>ServiceUsers>ServiceUser"`
}

type GetTinFromUnIdRequest struct {
	XMLName xml.Name `xml:"tem:get_tin_from_un_id"`
	Su      string   `xml:"tem:su"`    // mandatory
	Sp      string   `xml:"tem:sp"`    // mandatory
	UnId    int      `xml:"tem:un_id"` // mandatory
}

type GetTinFromUnIdResponse struct {
	Tin  string `xml:"Body>get_tin_from_un_idResponse>get_tin_from_un_idResult"`
	Name string `xml:"Body>get_tin_from_un_idResponse>name"`
}

type GetTransTypesRequest struct {
	XMLName xml.Name `xml:"tem:get_trans_types"`
	Su      string   `xml:"tem:su"` // mandatory
	Sp      string   `xml:"tem:sp"` // mandatory
}

type GetTransTypesResponse struct {
	TransportTypes []TypeStruct `xml:"Body>get_trans_typesResponse>get_trans_typesResult>TRANSPORT_TYPES>TRANSPORT_TYPE"`
}

type GetTransporterWaybillsRequest struct {
	XMLName       xml.Name `xml:"tem:get_transporter_waybills"`
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

type GetTransporterWaybillsResponse struct {
	WaybillList []Waybill `xml:"Body>get_transporter_waybillsResponse>get_transporter_waybillsResult>WAYBILL_LIST>WAYBILL"`
}

type GetWaybillRequest struct {
	XMLName   xml.Name `xml:"tem:get_waybill"`
	Su        string   `xml:"tem:su"`         // mandatory
	Sp        string   `xml:"tem:sp"`         // mandatory
	WaybillId int      `xml:"tem:waybill_id"` // mandatory
}

type GetWaybillResponse struct {
	Waybill Waybill `xml:"Body>get_waybillResponse>get_waybillResult>WAYBILL"`
}

type GetWaybillByNumberRequest struct {
	XMLName    xml.Name `xml:"tem:get_waybill_by_number"`
	Su         string   `xml:"tem:su"`             // mandatory
	Sp         string   `xml:"tem:sp"`             // mandatory
	WaybillNum string   `xml:"tem:waybill_number"` // mandatory
}

type GetWaybillByNumberResponse struct {
	Waybill Waybill `xml:"Body>get_waybill_by_numberResponse>get_waybill_by_numberResult>WAYBILL"`
}

type GetWaybillGoodsListRequest struct {
	XMLName       xml.Name `xml:"tem:get_waybill_goods_list"`
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
}

type GetWaybillGoodsListResponse struct {
	WaybillList []Waybill `xml:"Body>get_waybill_goods_listResponse>get_waybill_goods_listResult>WAYBILL_LIST>WAYBILL"`
}

type GetWaybillTamplateRequest struct {
	XMLName xml.Name `xml:"tem:get_waybill_tamplate"`
	Su      string   `xml:"tem:su"` // mandatory
	Sp      string   `xml:"tem:sp"` // mandatory
	Id      int      `xml:"tem:id"` // mandatory
}

type GetWaybillTamplateResponse struct {
	ResCode int     `xml:"Body>get_waybill_tamplateResponse>get_waybill_tamplateResult"`
	Waybill Waybill `xml:"Body>get_waybill_tamplateResponse>waybill_tamplate>WAYBILL"`
}

type GetWaybillTamplatesRequest struct {
	XMLName xml.Name `xml:"tem:get_waybill_tamplates"`
	Su      string   `xml:"tem:su"` // mandatory
	Sp      string   `xml:"tem:sp"` // mandatory
}

type GetWaybillTamplatesResponse struct {
	ResCode          int               `xml:"Body>get_waybill_tamplatesResponse>get_waybill_tamplatesResult"`
	WaybillTamplates []WaybillTamplate `xml:"Body>get_waybill_tamplateResponse>waybill_tamplates>WAYBILL_TAMPLATES>WAYBILL_TAMPLATE"`
}

type GetWaybillTypesRequest struct {
	XMLName xml.Name `xml:"tem:get_waybill_types"`
	Su      string   `xml:"tem:su"` // mandatory
	Sp      string   `xml:"tem:sp"` // mandatory
}

type GetWaybillTypesResponse struct {
	WaybillTypes []TypeStruct `xml:"Body>get_waybill_typesResponse>get_waybill_typesResult>WAYBILL_TYPES>WAYBILL_TYPE"`
}

type GetWaybillUnitsRequest struct {
	XMLName xml.Name `xml:"tem:get_waybill_units"`
	Su      string   `xml:"tem:su"` // mandatory
	Sp      string   `xml:"tem:sp"` // mandatory
}

type GetWaybillUnitsResponse struct {
	WaybillUnits []TypeStruct `xml:"Body>get_waybill_unitResponse>get_waybill_unitResult>WAYBILL_UNITS>WAYBILL_UNIT"`
}

type GetWaybillsRequest struct {
	XMLName       xml.Name `xml:"tem:get_waybills"`
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
}

type GetWaybillsResponse struct {
	WaybillList []Waybill `xml:"Body>get_waybillsResponse>get_waybillsResult>WAYBILL_LIST>WAYBILL"`
}

type GetWaybillsExRequest struct {
	XMLName       xml.Name `xml:"tem:get_waybills_ex"`
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

type GetWaybillsExResponse struct {
	WaybillList []Waybill `xml:"Body>get_waybills_exResponse>get_waybills_exResult>WAYBILL_LIST>WAYBILL"`
}

type GetWoodTypesRequest struct {
	XMLName xml.Name `xml:"tem:get_wood_types"`
	Su      string   `xml:"tem:su"` // mandatory
	Sp      string   `xml:"tem:sp"` // mandatory
}

type GetWoodTypesResponse struct {
	WoodTypes []TypeStruct `xml:"Body>get_wood_typesResponse>get_wood_typesResult>WOOD_TYPES>WOOD_TYPE"`
}

type IsVATPayerRequest struct {
	XMLName xml.Name `xml:"tem:is_vat_payer"`
	Su      string   `xml:"tem:su"`    // mandatory
	Sp      string   `xml:"tem:sp"`    // mandatory
	UnId    int      `xml:"tem:un_id"` // mandatory
}

type IsVATPayerResponse struct {
	Result bool `xml:"Body>is_vat_payerResponse>is_vat_payerResult"`
}

type IsVATPayerTinRequest struct {
	XMLName xml.Name `xml:"tem:is_vat_payer_tin"`
	Su      string   `xml:"tem:su"`    // mandatory
	Sp      string   `xml:"tem:sp"`    // mandatory
	UnId    int      `xml:"tem:un_id"` // mandatory
}

type IsVATPayerTinResponse struct {
	Result bool `xml:"Body>is_vat_payer_tinResponse>is_vat_payer_tinResult"`
}

type RefWaybillRequest struct {
	XMLName   xml.Name `xml:"tem:ref_waybill"`
	Su        string   `xml:"tem:su"`         // mandatory
	Sp        string   `xml:"tem:sp"`         // mandatory
	WaybillId int      `xml:"tem:waybill_id"` // mandatory
}

type RefWaybillResponse struct {
	ResCode int `xml:"Body>ref_waybillResponse>ref_waybillResult"`
}

type RefWaybillVdRequest struct {
	XMLName   xml.Name `xml:"tem:ref_waybill_vd"`
	Su        string   `xml:"tem:su"`         // mandatory
	Sp        string   `xml:"tem:sp"`         // mandatory
	WaybillId int      `xml:"tem:waybill_id"` // mandatory
	Comment   string   `xml:"tem:comment"`    // optional
}

type RefWaybillVdResponse struct {
	ResCode int `xml:"Body>ref_waybill_vdResponse>ref_waybill_vdResult"`
}

type RejectWaybillRequest struct {
	XMLName   xml.Name `xml:"tem:reject_waybill"`
	Su        string   `xml:"tem:su"`         // mandatory
	Sp        string   `xml:"tem:sp"`         // mandatory
	WaybillId int      `xml:"tem:waybill_id"` // mandatory
}

type RejectWaybillResponse struct {
	Result bool `xml:"Body>reject_waybillResponse>reject_waybillResult"`
}

type SaveBarCodeRequest struct {
	XMLName   xml.Name `xml:"tem:save_bar_code"`
	Su        string   `xml:"tem:su"`         // mandatory
	Sp        string   `xml:"tem:sp"`         // mandatory
	BarCode   string   `xml:"tem:bar_code"`   // mandatory
	GoodsName string   `xml:"tem:goods_name"` // optional
	UnitId    int      `xml:"tem:unit_id"`    // optional
	UnitTxt   string   `xml:"tem:unit_txt"`   // optional
	AId       int      `xml:"tem:a_id"`       // mandatory
}

type SaveBarCodeResponse struct {
	ResCode int `xml:"Body>save_bar_codeResponse>save_bar_codeResult"`
}

type SaveCarNumbersRequest struct {
	XMLName    xml.Name `xml:"tem:save_car_numbers"`
	Su         string   `xml:"tem:su"`         // mandatory
	Sp         string   `xml:"tem:sp"`         // mandatory
	CarNumbers string   `xml:"tem:car_number"` // mandatory
}

type SaveCarNumbersResponse struct {
	ResCode int `xml:"Body>save_car_numbersResponse>save_car_numbersResult"`
}

type SaveInvoiceRequest struct {
	XMLName   xml.Name `xml:"tem:save_invoice"`
	Su        string   `xml:"tem:su"`         // mandatory
	Sp        string   `xml:"tem:sp"`         // mandatory
	WaybillId int      `xml:"tem:waybill_id"` // mandatory
	InInvId   int      `xml:"tem:in_inv_id"`  // mandatory
}

type SaveInvoiceResponse struct {
	ResCode  int `xml:"Body>save_invoiceResponse>save_invoiceResult"`
	OutInvId int `xml:"Body>save_invoiceResponse>out_inv_id"`
}

type SaveWaybillRequest struct {
	XMLName xml.Name `xml:"tem:save_waybill"`
	Su      string   `xml:"tem:su"`              // mandatory
	Sp      string   `xml:"tem:sp"`              // mandatory
	Waybill Waybill  `xml:"tem:waybill>WAYBILL"` // mandatory
}

type SaveWaybillResponse struct {
	Result struct {
		Status    int     `xml:"RESULT>STATUS"`
		Id        int     `xml:"RESULT>ID"`
		GoodsList []Goods `xml:"RESULT>GOODS_LIST>GOODS"`
	} `xml:"Body>save_waybillResponse>save_waybillResult"`
}

type SaveWaybillTamplateRequest struct {
	XMLName xml.Name `xml:"tem:save_waybill_tamplate"`
	Su      string   `xml:"tem:su"`              // mandatory
	Sp      string   `xml:"tem:sp"`              // mandatory
	VName   string   `xml:"tem:v_name"`          // optional?
	Waybill Waybill  `xml:"tem:waybill>WAYBILL"` // mandatory
}

type SaveWaybillTamplateResponse struct {
	ResCode int `xml:"Body>save_waybill_tamplateResponse>save_waybill_tamplateResult"`
}

type SaveWaybillTransporterRequest struct {
	XMLName        xml.Name `xml:"tem:save_waybill_transporter"`
	Su             string   `xml:"tem:su"`              // mandatory
	Sp             string   `xml:"tem:sp"`              // mandatory
	WaybillId      int      `xml:"tem:waybill_id"`      // mandatory
	CarNumber      string   `xml:"tem:car_number"`      // optional?
	DriverTin      string   `xml:"tem:driver_tin"`      // optional?
	CheckDriverTin int      `xml:"tem:chek_driver_tin"` // mandatory
	DriverName     string   `xml:"tem:driver_name"`     // optional?
	TransId        int      `xml:"tem:trans_id"`        // mandatory
	TransTxt       string   `xml:"tem:trans_txt"`       // optional?
	ReceptionInfo  string   `xml:"tem:reception_info"`  // optional
	ReceiverInfo   string   `xml:"tem:receiver_info"`   // optional
}

type SaveWaybillTransporterResponse struct {
	ResCode int `xml:"Body>save_waybill_transporterResponse>save_waybill_transporterResult"`
}

type SendWaybillVdRequest struct {
	XMLName   xml.Name `xml:"tem:send_waybil_vd"`
	Su        string   `xml:"tem:su"`         // mandatory
	Sp        string   `xml:"tem:sp"`         // mandatory
	BeginDate string   `xml:"tem:begin_date"` // mandatory
	WaybillId int      `xml:"tem:waybill_id"` // mandatory
}

type SendWaybillVdResponse struct {
	Result string `xml:"Body>send_waybil_vdResponse>send_waybil_vdResult"`
}

type SendWaybillRequest struct {
	XMLName   xml.Name `xml:"tem:send_waybil_"`
	Su        string   `xml:"tem:su"`         // mandatory
	Sp        string   `xml:"tem:sp"`         // mandatory
	WaybillId int      `xml:"tem:waybill_id"` // mandatory
}

type SendWaybillResponse struct {
	Result string `xml:"Body>send_waybil_Response>send_waybil_Result"`
}

type SendWaybillTransporterRequest struct {
	XMLName   xml.Name `xml:"tem:send_waybil_transporter"`
	Su        string   `xml:"tem:su"`         // mandatory
	Sp        string   `xml:"tem:sp"`         // mandatory
	WaybillId int      `xml:"tem:waybill_id"` // mandatory
	BeginDate string   `xml:"tem:begin_date"` // mandatory
}

type SendWaybillTransporterResponse struct {
	ResCode    string `xml:"Body>send_waybil_transporterResponse>send_waybil_transporterResult"`
	WaybillNum string `xml:"Body>send_waybil_transporterResponse>waybill_number"`
}

type UpdateServiceUserRequest struct {
	XMLName      xml.Name `xml:"tem:update_service_user"`
	UserName     string   `xml:"tem:user_name"`     // optional?
	UserPassword string   `xml:"tem:user_password"` // optional?
	IP           string   `xml:"tem:ip"`            // optional?
	Name         string   `xml:"tem:name"`          // optional?
	Su           string   `xml:"tem:su"`            // mandatory
	Sp           string   `xml:"tem:sp"`            // mandatory
}

type UpdateServiceUserResponse struct {
	Result bool `xml:"Body>update_service_userResponse>update_service_userResult"`
}

//Dummy struct
type WhatIsMyIPRequest struct {
	XMLName xml.Name `xml:"tem:what_is_my_ip"`
}

type WhatIsMyIPResponse struct {
	Result string `xml:"Body>what_is_my_ipResponse>what_is_my_ipResult"`
}
