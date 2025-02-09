package models

import "time"

type (
	TransferStatusRequest[T any] struct {
		OriginalPartnerReferenceNo string `json:"originalPartnerReferenceNo"`
		OriginalReferenceNo        string `json:"originalReferenceNo"`
		OriginalExternalID         string `json:"originalExternalId"`
		ServiceCode                string `json:"serviceCode"`
		TransactionDate            string `json:"transactionDate"`
		Amount                     Amount `json:"amount"`
		AdditionalData             T      `json:"additionalData"`
	}
	TransferStatusResponse[T any] struct {
		ResponseCode               string    `json:"responseCode"`
		ResponseMessage            string    `json:"responseMessage"`
		OriginalReferenceNo        string    `json:"originalReferenceNo"`
		OriginalPartnerReferenceNo string    `json:"originalPartnerReferenceNo"`
		OriginalExternalId         string    `json:"originalExternalId"`
		ServiceCode                string    `json:"serviceCode"`
		TransactionDate            time.Time `json:"transactionDate"`
		Amount                     Amount    `json:"amount"`
		BeneficiaryAccountNo       string    `json:"beneficiaryAccountNo"`
		BeneficiaryBankCode        string    `json:"beneficiaryBankCode"`
		Currency                   string    `json:"currency"`
		PreviousResponseCode       string    `json:"previousResponseCode"`
		ReferenceNumber            string    `json:"referenceNumber"`
		SourceAccountNo            string    `json:"sourceAccountNo"`
		TransactionId              string    `json:"transactionId"`
		LatestTransactionStatus    string    `json:"latestTransactionStatus"`
		TransactionStatusDesc      string    `json:"transactionStatusDesc"`
		AdditionalData             T         `json:"additionalData"`
	}
)
