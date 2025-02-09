package models

import "time"

type DeviceInfo struct {
	OS          string `json:"os"`
	OSVersion   string `json:"osVersion"`
	Model       string `json:"model"`
	Manufacture string `json:"manufacture"`
}

type QParams struct {
	Action string `json:"action"`
}

type SuccessParams struct {
	AccountID string `json:"accountId"`
}

type AccessTokenInfo struct {
	AccessToken  string    `json:"accessToken"`
	ExpiresIn    time.Time `json:"expiresIn"`
	RefreshToken string    `json:"refreshToken"`
	ReExpiresIn  time.Time `json:"reExpiresIn"`
	TokenStatus  string    `json:"tokenStatus"`
}

type Params struct {
	Action             string `json:"action"`
	PinWebViewUrl      string `json:"pinWebViewUrl"`
	RedirectToDeeplink string `json:"redirectToDeeplink"`
}

type UserInfo struct {
	PublicUserId string `json:"publicUserId"`
}

// Amount represents the amount and currency.
type Amount struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
}

type AccountInfo struct {
	BalanceType              string `json:"balanceType"`
	Amount                   Amount `json:"amount"`
	FloatAmount              Amount `json:"floatAmount"`
	HoldAmount               Amount `json:"holdAmount"`
	AvailableBalance         Amount `json:"availableBalance"`
	LedgerBalance            Amount `json:"ledgerBalance"`
	CurrentMultilateralLimit Amount `json:"currentMultilateralLimit"`
	RegistrationStatusCode   string `json:"registrationStatusCode"`
	Status                   string `json:"status"`
}

type SourceOfFund struct {
	Source string `json:"source"`
	Amount Amount `json:"amount"`
}

// BalanceEntry represents an entry in the balance array.
type BalanceEntry struct {
	Amount          Amount `json:"amount"`
	StartingBalance Amount `json:"startingBalance"`
	EndingBalance   Amount `json:"endingBalance"`
}

// TotalEntries represents the total number of credit or debit entries.
type TotalEntries struct {
	NumberOfEntries string `json:"numberOfEntries"`
	Amount          Amount `json:"amount"`
}
type DetailBalance struct {
	StartAmount []Amount `json:"startAmount"`
	EndAmount   []Amount `json:"endAmount"`
}

// DetailInfo represents additional information about the transaction.
type DetailInfo struct {
	Page string `json:"page"`
}

// OriginatorInfo represents the originator information in the JSON.
type OriginatorInfo struct {
	OriginatorCustomerNo   string `json:"originatorCustomerNo"`
	OriginatorCustomerName string `json:"originatorCustomerName"`
	OriginatorBankCode     string `json:"originatorBankCode"`
}

type Lang struct {
	English   string `json:"english"`
	Indonesia string `json:"indonesia"`
}

// BillDetail represents the details of a bill.
type BillDetail struct {
	BillCode        string                 `json:"billCode"`
	BillNo          string                 `json:"billNo"`
	BillName        string                 `json:"billName"`
	BillShortName   string                 `json:"billShortName"`
	BillDescription Lang                   `json:"billDescription"`
	BillSubCompany  string                 `json:"billSubCompany"`
	BillAmount      Amount                 `json:"billAmount"`
	BillAmountLabel string                 `json:"billAmountLabel"`
	BillAmountValue string                 `json:"billAmountValue"`
	AdditionalInfo  map[string]interface{} `json:"additionalInfo"` // Use map for dynamic fields
}

// MerchantInfo represents the information about a merchant.
type MerchantInfo struct {
	MerchantPAN  string `json:"merchantPAN"`
	AcquirerName string `json:"acquirerName"`
}

// UserResource represents the user resource information.
type UserResource struct {
	ResourceType string `json:"resourceType"`
	Value        string `json:"value"`
}

// UrlParam represents the URL parameters.
type UrlParam struct {
	Url        string `json:"url"`
	Type       string `json:"type"`
	IsDeeplink string `json:"isDeeplink"`
}

// PayOptionDetail represents the details of payment options.
type PayOptionDetail[T any] struct {
	PayMethod      string `json:"payMethod"`
	PayOption      string `json:"payOption"`
	TransAmount    Amount `json:"transAmount"`
	FeeAmount      Amount `json:"feeAmount"`
	CardToken      string `json:"cardToken"`
	MerchantToken  string `json:"merchantToken"`
	AdditionalInfo T      `json:"additionalInfo"`
}

// Refund represents the refund details in the refund history.
type Refund struct {
	RefundNo        string    `json:"refundNo"`
	PartnerRefundNo string    `json:"partnerRefundNo"`
	RefundAmount    Amount    `json:"refundAmount"`
	RefundStatus    string    `json:"refundStatus"`
	RefundDate      time.Time `json:"refundDate"`
	Reason          string    `json:"reason"`
}

// Items represents the product details.
type Items struct {
	ProductId   string `json:"productId"`
	ProductName string `json:"productName"`
	Qty         string `json:"qty"`
	Desc        string `json:"desc"`
}

// ScannerInfo represents the scanner device information.
type ScannerInfo struct {
	DeviceId      string `json:"deviceId"`
	DeviceVersion string `json:"deviceVersion"`
	DeviceModel   string `json:"deviceModel"`
	DeviceIp      string `json:"deviceIp"`
}

type Item struct {
	GoodsId  string `json:"goodsId"`
	Price    Amount `json:"price"`
	Category string `json:"category"`
	Unit     string `json:"unit"`
	Quantity string `json:"quantity"`
}
