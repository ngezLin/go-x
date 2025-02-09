package models

type (
	TransferInterbankRequest[T any] struct {
		SourceAccountNumber      string  `json:"source_account_number"`
		DestinationAccountNumber string  `json:"destination_account_number"`
		Amount                   float64 `json:"amount"`
		TransactionType          string  `json:"transaction_type"`
		Reference                string  `json:"reference"`
		AdditionalInfo           T       `json:"additionalInfo"`
	}
	TransferInterbankResponse[T any] struct {
		ResponseCode         string           `json:"responseCode"`
		ResponseMessage      string           `json:"responseMessage"`
		ReferenceNo          string           `json:"referenceNo"`
		PartnerReferenceNo   string           `json:"partnerReferenceNo"`
		Amount               Amount           `json:"amount"`
		BeneficiaryAccountNo string           `json:"beneficiaryAccountNo"`
		BeneficiaryBankCode  string           `json:"beneficiaryBankCode"`
		SourceAccountNo      string           `json:"sourceAccountNo"`
		TraceNo              string           `json:"traceNo"`
		OriginatorInfos      []OriginatorInfo `json:"originatorInfos"`
		AdditionalInfo       T                `json:"additionalInfo"`
	}
)

type (
	TransferRequestPaymentRequest[T any] struct {
		PartnerReferenceNo     string `json:"partnerReferenceNo"`
		BankCode               string `json:"bankCode"`
		BeneficiaryAccountNo   string `json:"beneficiaryAccountNo"`
		BeneficiaryAccountName string `json:"beneficiaryAccountName"`
		Remark                 string `json:"remark"`
		ExpiredDatetime        string `json:"expiredDatetime"`
		SourceAccountNo        string `json:"sourceAccountNo"`
		SourceAccountName      string `json:"sourceAccountName"`
		Currency               string `json:"currency"`
		Amount                 Amount `json:"amount"`
		FeeType                string `json:"feeType"`
		AdditionalInfo         T      `json:"additionalInfo"`
	}
	TransferRequestPaymentResponse[T any] struct {
		ResponseCode       string `json:"responseCode"`
		ResponseMessage    string `json:"responseMessage"`
		ReferenceNo        string `json:"referenceNo"`
		PartnerReferenceNo string `json:"partnerReferenceNo"`
		AdditionalInfo     T      `json:"additionalInfo"`
	}
)

type (
	TransferInterbankBulkRequest[T any] struct {
		PartnerBulkID     string       `json:"partnerBulkId"`
		Currency          string       `json:"currency"`
		CustomerReference string       `json:"customerReference"`
		FeeType           string       `json:"feeType"`
		Remark            string       `json:"remark"`
		SourceAccountNo   string       `json:"sourceAccountNo"`
		TransactionDate   string       `json:"transactionDate"`
		BulkObject        []BulkObject `json:"bulkObject"`
		AdditionalInfo    T            `json:"additionalInfo"`
	}
	BulkObject struct {
		PartnerReferenceNo     string           `json:"partnerReferenceNo"`
		BankCode               string           `json:"bankCode"`
		BeneficiaryAccountNo   string           `json:"beneficiaryAccountNo"`
		BeneficiaryAccountName string           `json:"beneficiaryAccountName"`
		Amount                 Amount           `json:"amount"`
		OriginatorInfos        []OriginatorInfo `json:"originatorInfos"`
	}
	TransferInterbankBulkResponse[T any] struct {
		ResponseCode    string `json:"responseCode"`
		ResponseMessage string `json:"responseMessage"`
		BulkId          string `json:"bulkId"`
		PartnerBulkId   string `json:"partnerBulkId"` // Note: Removed the space after the key
		AdditionalInfo  T      `json:"additionalInfo"`
	}
)

type (
	TransferInterbankBulkNotifyRequest[T any] struct {
		BulkID         string       `json:"bulkId"`
		PartnerBulkID  string       `json:"partnerBulkId"`
		BulkObject     []BulkObject `json:"bulkObject"`
		AdditionalInfo T            `json:"additionalInfo"`
	}
	TransferInterbankBulkNotifyResponse[T any] struct {
		ResponseCode    string `json:"responseCode"`
		ResponseMessage string `json:"responseMessage"`
		BulkId          string `json:"bulkId"`
		PartnerBulkId   string `json:"partnerBulkId"`
	}
)
