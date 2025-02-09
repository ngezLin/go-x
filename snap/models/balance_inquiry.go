package models

type (
	BalanceInquiryRequest[T any] struct {
		PartnerReferenceNo string   `json:"partnerReferenceNo"`
		BankCardToken      string   `json:"bankCardToken"`
		AccountNo          string   `json:"accountNo"`
		BalanceTypes       []string `json:"balanceTypes"` // Slice of strings for balance types
		AdditionalInfo     T        `json:"additionalInfo"`
	}
	BalanceInquiryResponse[T any] struct {
		ResponseCode       string        `json:"responseCode"`
		ResponseMessage    string        `json:"responseMessage"`
		ReferenceNo        string        `json:"referenceNo"`
		PartnerReferenceNo string        `json:"partnerReferenceNo"`
		AccountNo          string        `json:"accountNo"`
		Name               string        `json:"name"`
		AccountInfos       []AccountInfo `json:"accountInfos"`
		AdditionalInfo     T             `json:"additionalInfo"`
	}
)
