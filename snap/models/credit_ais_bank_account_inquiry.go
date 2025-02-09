package models

type (
	AccountInquiryInternalRequest[T any] struct {
		PartnerReferenceNo   string `json:"partnerReferenceNo"`
		BeneficiaryAccountNo string `json:"beneficiaryAccountNo"`
		AdditionalInfo       T      `json:"additionalInfo"`
	}
	AccountInquiryInternalResponse[T any] struct {
		ResponseCode             string `json:"responseCode"`
		ResponseMessage          string `json:"responseMessage"`
		ReferenceNo              string `json:"referenceNo"`
		PartnerReferenceNo       string `json:"partnerReferenceNo"`
		BeneficiaryAccountName   string `json:"beneficiaryAccountName"`
		BeneficiaryAccountNo     string `json:"beneficiaryAccountNo"`
		BeneficiaryAccountStatus string `json:"beneficiaryAccountStatus"`
		BeneficiaryAccountType   string `json:"beneficiaryAccountType"`
		Currency                 string `json:"currency"`
		AdditionalInfo           T      `json:"additionalInfo"`
	}
)

type (
	AccountInquiryExternalRequest[T any] struct {
		BeneficiaryBankCode  string `json:"beneficiaryBankCode"`
		BeneficiaryAccountNo string `json:"beneficiaryAccountNo"`
		PartnerReferenceNo   string `json:"partnerReferenceNo"`
		AdditionalInfo       T      `json:"additionalInfo"`
	}
	AccountInquiryExternalResponse[T any] struct {
		ResponseCode           string `json:"responseCode"`
		ResponseMessage        string `json:"responseMessage"`
		ReferenceNo            string `json:"referenceNo"`
		PartnerReferenceNo     string `json:"partnerReferenceNo"`
		BeneficiaryAccountName string `json:"beneficiaryAccountName"`
		BeneficiaryAccountNo   string `json:"beneficiaryAccountNo"`
		BeneficiaryBankCode    string `json:"beneficiaryBankCode"`
		BeneficiaryBankName    string `json:"beneficiaryBankName"`
		Currency               string `json:"currency"`
		AdditionalInfo         T      `json:"additionalInfo"`
	}
)
