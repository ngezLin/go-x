package models

import "time"

type (
	TransferSKNRequest[T any] struct {
		PartnerReferenceNo           string           `json:"partnerReferenceNo"`
		Amount                       Amount           `json:"amount"`
		BeneficiaryAccountName       string           `json:"beneficiaryAccountName"`
		BeneficiaryAccountNo         string           `json:"beneficiaryAccountNo"`
		BeneficiaryAddress           string           `json:"beneficiaryAddress"`
		BeneficiaryBankCode          string           `json:"beneficiaryBankCode"`
		BeneficiaryBankName          string           `json:"beneficiaryBankName"`
		BeneficiaryCustomerResidence string           `json:"beneficiaryCustomerResidence"`
		BeneficiaryCustomerType      string           `json:"beneficiaryCustomerType"`
		BeneficiaryEmail             string           `json:"beneficiaryEmail"`
		Currency                     string           `json:"currency"`
		CustomerReference            string           `json:"customerReference"`
		FeeType                      string           `json:"feeType"`
		Kodepos                      string           `json:"kodepos"`
		ReceiverPhone                string           `json:"receiverPhone"`
		Remark                       string           `json:"remark"`
		SenderCustomerResidence      string           `json:"senderCustomerResidence"`
		SenderCustomerType           string           `json:"senderCustomerType"`
		SenderPhone                  string           `json:"senderPhone"`
		SourceAccountNo              string           `json:"sourceAccountNo"`
		OriginatorInfos              []OriginatorInfo `json:"originatorInfos"`
		TransactionDate              string           `json:"transactionDate"`
		AdditionalInfo               T                `json:"additionalInfo"`
	}
	TransferSKNResponse[T any] struct {
		ResponseCode           string           `json:"responseCode"`
		ResponseMessage        string           `json:"responseMessage"`
		ReferenceNo            string           `json:"referenceNo"`
		PartnerReferenceNo     string           `json:"partnerReferenceNo"`
		Amount                 Amount           `json:"amount"`
		BeneficiaryAccountName string           `json:"beneficiaryAccountName"`
		BeneficiaryAccountNo   string           `json:"beneficiaryAccountNo"`
		BeneficiaryAccountType string           `json:"beneficiaryAccountType"`
		BeneficiaryBankCode    string           `json:"beneficiaryBankCode"`
		Currency               string           `json:"currency"`
		CustomerReference      string           `json:"customerReference"`
		SourceAccountNo        string           `json:"sourceAccountNo"`
		OriginatorInfos        []OriginatorInfo `json:"originatorInfos"`
		TraceNo                string           `json:"traceNo"`
		TransactionDate        time.Time        `json:"transactionDate"`
		TransactionStatus      string           `json:"transactionStatus"`
		TransactionStatusDesc  string           `json:"transactionStatusDesc"`
		AdditionalInfo         T                `json:"additionalInfo"`
	}
)

type (
	TransferSKNNotifyRequest[T any] struct {
		OriginalPartnerReferenceNo string `json:"originalPartnerReferenceNo"`
		OriginalReferenceNo        string `json:"originalReferenceNo"`
		OriginalExternalID         string `json:"originalExternalId"`
		LatestTransactionStatus    string `json:"latestTransactionStatus"`
		Amount                     Amount `json:"amount"`
		BeneficiaryAccountName     string `json:"beneficiaryAccountName"`
		BeneficiaryAccountNo       string `json:"beneficiaryAccountNo"`
		BeneficiaryBankCode        string `json:"beneficiaryBankCode"`
		SourceAccountNo            string `json:"sourceAccountNo"`
		TransactionDate            string `json:"transactionDate"`
		AdditionalInfo             T      `json:"additionalInfo"`
	}
	TransferSKNNotifyResponse struct {
		ResponseCode    string `json:"responseCode"`
		ResponseMessage string `json:"responseMessage"`
	}
)
