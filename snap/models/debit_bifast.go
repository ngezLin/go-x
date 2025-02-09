package models

import "time"

type (
	DebitFastEmandateRequest[T any] struct {
		PartnerReferenceNo string    `json:"partnerReferenceNo"`
		BankCode           string    `json:"bankCode"`
		SourceAccountNo    string    `json:"sourceAccountNo"`
		SourceAccountName  string    `json:"sourceAccountName"`
		MaxAmount          Amount    `json:"maxAmount"`
		BillerId           string    `json:"billerId"`
		BillerName         string    `json:"billerName"`
		CustomerId         string    `json:"customerId"`
		ExpiredDatetime    time.Time `json:"expiredDatetime"`
		AdditionalInfo     T         `json:"additionalInfo"`
	}
	DebitFastEmandateResponse[T any] struct {
		ResponseCode       string `json:"responseCode"`
		ResponseMessage    string `json:"responseMessage"`
		ReferenceNo        string `json:"referenceNo"`
		PartnerReferenceNo string `json:"partnerReferenceNo"`
		EMandateReffId     string `json:"eMandateReffId"`
		AdditionalInfo     T      `json:"additionalInfo"`
	}
)

type (
	DebitFastPaymentRequest[T any] struct {
		PartnerReferenceNo     string    `json:"partnerReferenceNo"`
		Currency               string    `json:"currency"`
		CustomerReference      string    `json:"customerReference"`
		FeeType                string    `json:"feeType"`
		Remark                 string    `json:"remark"`
		BeneficiaryAccountNo   string    `json:"beneficiaryAccountNo"`
		BeneficiaryAccountName string    `json:"beneficiaryAccountName"`
		TransactionDate        time.Time `json:"transactionDate"`
		BankCode               string    `json:"bankCode"`
		SourceAccountNo        string    `json:"sourceAccountNo"`
		SourceAccountName      string    `json:"sourceAccountName"`
		Amount                 Amount    `json:"amount"`
		EMandateReffId         string    `json:"eMandateReffId"`
		AdditionalInfo         T         `json:"additionalInfo"`
	}
	DebitFastPaymentResponse[T any] struct {
		ResponseCode       string `json:"responseCode"`
		ResponseMessage    string `json:"responseMessage"`
		ReferenceNo        string `json:"referenceNo"`
		PartnerReferenceNo string `json:"partnerReferenceNo"`
		AdditionalInfo     T      `json:"additionalInfo"`
	}
)

type (
	DebitFastNotifyRequest[T any] struct {
		OriginalReferenceNo        string `json:"originalReferenceNo"`
		OriginalPartnerReferenceNo string `json:"originalPartnerReferenceNo"`
		OriginalExternalId         string `json:"originalExternalId"`
		TransactionStatus          string `json:"transactionStatus"`
		TransactionStatusDesc      string `json:"transactionStatusDesc"`
		EMandateReffId             string `json:"eMandateReffId"`
		SourceAccountNo            string `json:"sourceAccountNo"`
		SourceAccountName          string `json:"sourceAccountName"`
		Amount                     Amount `json:"amount"`
		TraceNo                    string `json:"traceNo"`
		AdditionalInfo             T      `json:"additionalInfo"`
	}
	DebitFastNotifyResponse[T any] struct {
		ResponseCode    string `json:"responseCode"`
		ResponseMessage string `json:"responseMessage"`
		AdditionalInfo  T      `json:"additionalInfo"`
	}
)
