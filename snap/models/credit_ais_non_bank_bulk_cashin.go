package models

type (
	EmoneyBulkCashinPaymentRequest[T any] struct {
		PartnerBulkID   string       `json:"partnerBulkId"`
		TransactionDate string       `json:"transactionDate"`
		Currency        string       `json:"currency"`
		BulkObject      []BulkObject `json:"bulkObject"`
		FeeType         string       `json:"feeType"`
		AdditionalInfo  T            `json:"additionalInfo"`
	}
	EmoneyBulkCashinPaymentResponse[T any] struct {
		ResponseCode    string `json:"responseCode"`
		ResponseMessage string `json:"responseMessage"`
		BulkId          string `json:"bulkId"`
		PartnerBulkId   string `json:"partnerBulkId"`
		AdditionalInfo  T      `json:"additionalInfo"`
	}
)

type (
	EmoneyBulkCashinNotifyRequest[T any] struct {
		BulkID         string       `json:"bulkId"`
		PartnerBulkID  string       `json:"partnerBulkId"`
		BulkObject     []BulkObject `json:"bulkObject"`
		AdditionalInfo T            `json:"additionalInfo"`
	}
	EmoneyBulkCashinNotifyResponse[T any] struct {
		ResponseCode    string `json:"responseCode"`
		ResponseMessage string `json:"responseMessage"`
		BulkId          string `json:"bulkId"`
		PartnerBulkId   string `json:"partnerBulkId"`
		AdditionalInfo  T      `json:"additionalInfo"`
	}
)
