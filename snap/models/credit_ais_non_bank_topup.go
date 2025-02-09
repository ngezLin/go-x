package models

type (
	EmoneyAccountInquiryRequest[T any] struct {
		PartnerReferenceNo string `json:"partnerReferenceNo"`
		CustomerNumber     string `json:"customerNumber"`
		Amount             Amount `json:"amount"`
		TransactionDate    string `json:"transactionDate"`
		AdditionalInfo     T      `json:"additionalInfo"`
	}
	EmoneyAccountInquiryResponse[T any] struct {
		ResponseCode           string `json:"responseCode"`
		ResponseMessage        string `json:"responseMessage"`
		ReferenceNo            string `json:"referenceNo"`
		PartnerReferenceNo     string `json:"partnerReferenceNo"`
		SessionId              string `json:"sessionId"`
		CustomerNumber         string `json:"customerNumber"`
		CustomerName           string `json:"customerName"`
		CustomerMonthlyInLimit string `json:"customerMonthlyInLimit"`
		MinAmount              Amount `json:"minAmount"`
		MaxAmount              Amount `json:"maxAmount"`
		Amount                 Amount `json:"amount"`
		FeeAmount              Amount `json:"feeAmount"`
		FeeType                string `json:"feeType"`
		AdditionalInfo         T      `json:"additionalInfo"`
	}
)

type (
	EmoneyTopupRequest[T any] struct {
		PartnerReferenceNo string `json:"partnerReferenceNo"`
		CustomerNumber     string `json:"customerNumber"`
		CustomerName       string `json:"customerName"`
		Amount             Amount `json:"amount"`
		FeeAmount          Amount `json:"feeAmount"`
		TransactionDate    string `json:"transactionDate"`
		SessionID          string `json:"sessionId"`
		CategoryID         string `json:"categoryId"`
		Notes              string `json:"notes"`
		AdditionalInfo     T      `json:"additionalInfo"`
	}
	EmoneyTopupResponse[T any] struct {
		ResponseCode       string `json:"responseCode"`
		ResponseMessage    string `json:"responseMessage"`
		ReferenceNo        string `json:"referenceNo"`
		PartnerReferenceNo string `json:"partnerReferenceNo"`
		SessionId          string `json:"sessionId"`
		CustomerNumber     string `json:"customerNumber"`
		ReferenceNumber    string `json:"referenceNumber"`
		Amount             Amount `json:"amount"`
		AdditionalInfo     T      `json:"additionalInfo"`
	}
)

type (
	EmoneyTopupStatusRequest[T any] struct {
		OriginalPartnerReferenceNo string `json:"originalPartnerReferenceNo"`
		OriginalReferenceNo        string `json:"originalReferenceNo"`
		OriginalExternalID         string `json:"originalExternalId"`
		ServiceCode                string `json:"serviceCode"`
		AdditionalInfo             T      `json:"additionalInfo"`
	}
	EmoneyTopupStatusResponse[T any] struct {
		ResponseCode               string `json:"responseCode"`
		ResponseMessage            string `json:"responseMessage"`
		OriginalPartnerReferenceNo string `json:"originalPartnerReferenceNo"`
		OriginalReferenceNo        string `json:"originalReferenceNo"`
		OriginalExternalId         string `json:"originalExternalId"`
		ServiceCode                string `json:"serviceCode"`
		Amount                     Amount `json:"amount"`
		LatestTransactionStatus    string `json:"latestTransactionStatus"`
		TransactionStatusDesc      string `json:"transactionStatusDesc"`
		AdditionalInfo             T      `json:"additionalInfo"`
	}
)
