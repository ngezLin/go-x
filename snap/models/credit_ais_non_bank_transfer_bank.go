package models

import "time"

type (
	EmoneyBankAccountInquiryRequest[T any] struct {
		PartnerReferenceNo       string `json:"partnerReferenceNo"`
		CustomerNumber           string `json:"customerNumber"`
		Amount                   Amount `json:"amount"`
		BeneficiaryAccountNumber string `json:"beneficiaryAccountNumber"`
		AdditionalInfo           T      `json:"additionalInfo"`
	}
	EmoneyBankAccountInquiryResponse[T any] struct {
		ResponseCode             string `json:"responseCode"`
		ResponseMessage          string `json:"responseMessage"`
		ReferenceNo              string `json:"referenceNo"`
		PartnerReferenceNo       string `json:"partnerReferenceNo"`
		AccountType              string `json:"accountType"`
		BeneficiaryAccountNumber string `json:"beneficiaryAccountNumber"`
		BeneficiaryAccountName   string `json:"beneficiaryAccountName"`
		BeneficiaryBankCode      string `json:"beneficiaryBankCode"`
		BeneficiaryBankShortName string `json:"beneficiaryBankShortName"`
		BeneficiaryBankName      string `json:"beneficiaryBankName"`
		Amount                   Amount `json:"amount"`
		SessionId                string `json:"sessionId"`
		AdditionalInfo           T      `json:"additionalInfo"`
	}
)

type (
	EmoneyBankPaymentRequest[T any] struct {
		PartnerReferenceNo       string `json:"partnerReferenceNo"`
		CustomerNumber           string `json:"customerNumber"`
		AccountType              string `json:"accountType"`
		BeneficiaryAccountNumber string `json:"beneficiaryAccountNumber"`
		BeneficiaryBankCode      string `json:"beneficiaryBankCode"`
		Amount                   Amount `json:"amount"`
		SessionID                string `json:"sessionId"`
		AdditionalInfo           T      `json:"additionalInfo"`
	}
	EmoneyBankPaymentResponse[T any] struct {
		ResponseCode       string    `json:"responseCode"`
		ResponseMessage    string    `json:"responseMessage"`
		ReferenceNo        string    `json:"referenceNo"`
		PartnerReferenceNo string    `json:"partnerReferenceNo"`
		TransactionDate    time.Time `json:"transactionDate"`
		ReferenceNumber    string    `json:"referenceNumber"`
		AdditionalInfo     T         `json:"additionalInfo"`
	}
)
