package models

import "time"

type (
	TransferIntrabankRequest[T any] struct {
		PartnerReferenceNo   string           `json:"partnerReferenceNo"`
		Amount               Amount           `json:"amount"`
		BeneficiaryAccountNo string           `json:"beneficiaryAccountNo"`
		BeneficiaryEmail     string           `json:"beneficiaryEmail"`
		Currency             string           `json:"currency"`
		CustomerReference    string           `json:"customerReference"`
		FeeType              string           `json:"feeType"`
		Remark               string           `json:"remark"`
		SourceAccountNo      string           `json:"sourceAccountNo"`
		TransactionDate      string           `json:"transactionDate"`
		OriginatorInfos      []OriginatorInfo `json:"originatorInfos"`
		AdditionalInfo       T                `json:"additionalInfo"`
	}
	TransferIntrabankResponse[T any] struct {
		ResponseCode         string           `json:"responseCode"`
		ResponseMessage      string           `json:"responseMessage"`
		ReferenceNo          string           `json:"referenceNo"`
		PartnerReferenceNo   string           `json:"partnerReferenceNo"`
		Amount               Amount           `json:"amount"`
		BeneficiaryAccountNo string           `json:"beneficiaryAccountNo"`
		Currency             string           `json:"currency"`
		CustomerReference    string           `json:"customerReference"`
		SourceAccount        string           `json:"sourceAccount"`
		TransactionDate      time.Time        `json:"transactionDate"`
		OriginatorInfos      []OriginatorInfo `json:"originatorInfos"`
		AdditionalInfo       T                `json:"additionalInfo"`
	}
)
