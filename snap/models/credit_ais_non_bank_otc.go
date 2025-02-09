package models

import "time"

type (
	EmoneyOTCCashOutRequest[T any] struct {
		PartnerReferenceNo string `json:"partnerReferenceNo"`
		CustomerNumber     string `json:"customerNumber"`
		OTP                string `json:"otp"`
		Amount             Amount `json:"amount"`
		AdditionalInfo     T      `json:"additionalInfo"`
	}
	EmoneyOTCCashOutResponse[T any] struct {
		ResponseCode       string    `json:"responseCode"`
		ResponseMessage    string    `json:"responseMessage"`
		ReferenceNo        string    `json:"referenceNo"`
		PartnerReferenceNo string    `json:"partnerReferenceNo"`
		TransactionDate    time.Time `json:"transactionDate"`
		AdditionalInfo     T         `json:"additionalInfo"`
	}
)

type (
	EmoneyOTCStatusRequest[T any] struct {
		OriginalPartnerReferenceNo string    `json:"originalPartnerReferenceNo"`
		OriginalReferenceNo        string    `json:"originalReferenceNo"`
		OriginalExternalId         string    `json:"originalExternalId"`
		ServiceCode                string    `json:"serviceCode"`
		CustomerNumber             string    `json:"customerNumber"`
		TransactionDate            time.Time `json:"transactionDate"`
		Amount                     Amount    `json:"amount"`

		AdditionalInfo T `json:"additionalInfo"`
	}
	EmoneyOTCStatusResponse[T any] struct {
		ResponseCode               string    `json:"responseCode"`
		ResponseMessage            string    `json:"responseMessage"`
		OriginalReferenceNo        string    `json:"originalReferenceNo"`
		OriginalPartnerReferenceNo string    `json:"originalPartnerReferenceNo"`
		OriginalExternalId         string    `json:"originalExternalId"`
		ServiceCode                string    `json:"serviceCode"`
		LatestTransactionStatus    string    `json:"latestTransactionStatus"`
		TransactionStatusDesc      string    `json:"transactionStatusDesc"`
		TransactionDate            time.Time `json:"transactionDate"`
		Amount                     Amount    `json:"amount"`
		AdditionalInfo             T         `json:"additionalInfo"`
	}
)

type (
	EmoneyOTCCancelRequest[T any] struct {
		OriginalReferenceNo        string `json:"originalReferenceNo"`
		OriginalPartnerReferenceNo string `json:"originalPartnerReferenceNo"`
		OriginalExternalId         string `json:"originalExternalId"`
		CustomerNumber             string `json:"customerNumber"`
		Reason                     string `json:"reason"`
		AdditionalInfo             T      `json:"additionalInfo"`
	}
	EmoneyOTCCancelResponse[T any] struct {
		ResponseCode               string    `json:"responseCode"`
		ResponseMessage            string    `json:"responseMessage"`
		OriginalReferenceNo        string    `json:"originalReferenceNo"`
		OriginalPartnerReferenceNo string    `json:"originalPartnerReferenceNo"`
		OriginalExternalId         string    `json:"originalExternalId"`
		CancelTime                 time.Time `json:"cancelTime"`
		TransactionDate            time.Time `json:"transactionDate"`
		AdditionalInfo             T         `json:"additionalInfo"`
	}
)
