package models

import "time"

type (
	VirtualAccountData struct {
		InquiryStatus         string       `json:"inquiryStatus"`
		InquiryReason         Lang         `json:"inquiryReason"`
		SubCompany            string       `json:"subCompany"`
		VirtualAccountTrxType string       `json:"virtualAccountTrxType"`
		PaymentFlagReason     Lang         `json:"paymentFlagReason"`
		PartnerServiceId      string       `json:"partnerServiceId"`
		CustomerNo            string       `json:"customerNo"`
		VirtualAccountNo      string       `json:"virtualAccountNo"`
		VirtualAccountName    string       `json:"virtualAccountName"`
		VirtualAccountEmail   string       `json:"virtualAccountEmail"`
		VirtualAccountPhone   string       `json:"virtualAccountPhone"`
		SourceAccountNo       string       `json:"sourceAccountNo"`
		SourceAccountType     string       `json:"sourceAccountType"`
		InquiryRequestId      string       `json:"inquiryRequestId"`
		PaymentRequestId      string       `json:"paymentRequestId"`
		PartnerReferenceNo    string       `json:"partnerReferenceNo"`
		ReferenceNo           string       `json:"referenceNo"`
		PaidAmount            Amount       `json:"paidAmount"`
		PaidBills             string       `json:"paidBills"`
		TotalAmount           Amount       `json:"totalAmount"`
		TrxDateTime           time.Time    `json:"trxDateTime"`
		JournalNum            string       `json:"journalNum"`
		PaymentType           int          `json:"paymentType"`
		FlagAdvise            string       `json:"flagAdvise"`
		BillDetails           []BillDetail `json:"billDetails"`
		FreeTexts             []Lang       `json:"freeTexts"`
		FeeAmount             Amount       `json:"feeAmount"`
		ProductName           string       `json:"productName"`
		TrxId                 string       `json:"trxId"`
	}
)

type (
	TransferVaInquiryRequest[T any] struct {
		PartnerServiceID      string `json:"partnerServiceId"`
		CustomerNo            string `json:"customerNo"`
		VirtualAccountNo      string `json:"virtualAccountNo"`
		TxnDateInit           string `json:"txnDateInit"`
		ChannelCode           int    `json:"channelCode"`
		Language              string `json:"language"`
		Amount                Amount `json:"amount"`
		HashedSourceAccountNo string `json:"hashedSourceAccountNo"`
		SourceBankCode        string `json:"sourceBankCode"`
		PassApp               string `json:"passApp"`
		InquiryRequestID      string `json:"inquiryRequestId"`
		AdditionalInfo        T      `json:"additionalInfo"`
	}
	TransferVaInquiryResponse[T any] struct {
		ResponseCode       string             `json:"responseCode"`
		ResponseMessage    string             `json:"responseMessage"`
		VirtualAccountData VirtualAccountData `json:"virtualAccountData"`
		AdditionalInfo     T                  `json:"additionalInfo"`
	}
)
type (
	TransferVaPaymentRequest[T any] struct {
		PartnerServiceID        string       `json:"partnerServiceId"`
		CustomerNo              string       `json:"customerNo"`
		VirtualAccountNo        string       `json:"virtualAccountNo"`
		VirtualAccountName      string       `json:"virtualAccountName"`
		VirtualAccountEmail     string       `json:"virtualAccountEmail"`
		VirtualAccountPhone     string       `json:"virtualAccountPhone"`
		TrxID                   string       `json:"trxId"`
		PaymentRequestID        string       `json:"paymentRequestId"`
		ChannelCode             int          `json:"channelCode"`
		HashedSourceAccountNo   string       `json:"hashedSourceAccountNo"`
		SourceBankCode          string       `json:"sourceBankCode"`
		PaidAmount              Amount       `json:"paidAmount"`
		CumulativePaymentAmount Amount       `json:"cumulativePaymentAmount"`
		PaidBills               string       `json:"paidBills"`
		TotalAmount             Amount       `json:"totalAmount"`
		TrxDateTime             string       `json:"trxDateTime"`
		ReferenceNo             string       `json:"referenceNo"`
		JournalNum              string       `json:"journalNum"`
		PaymentType             int          `json:"paymentType"`
		FlagAdvise              string       `json:"flagAdvise"`
		SubCompany              string       `json:"subCompany"`
		BillDetails             []BillDetail `json:"billDetails"`
		FreeTexts               []Lang       `json:"freeTexts"`
		AdditionalInfo          T            `json:"additionalInfo"`
	}
	TransferVaPaymentResponse[T any] struct {
		ResponseCode       string             `json:"responseCode"`
		ResponseMessage    string             `json:"responseMessage"`
		VirtualAccountData VirtualAccountData `json:"virtualAccountData"`
		AdditionalInfo     T                  `json:"additionalInfo"`
	}
)

type (
	TransferVaStatusRequest[T any] struct {
		PartnerServiceID string `json:"partnerServiceId"`
		CustomerNo       int64  `json:"customerNo"` // Use int64 for large numbers
		VirtualAccountNo string `json:"virtualAccountNo"`
		InquiryRequestID string `json:"inquiryRequestId"`
		PaymentRequestID string `json:"paymentRequestId"`
		AdditionalInfo   T      `json:"additionalInfo"`
	}
	TransferVaStatusResponse[T any] struct {
		ResponseCode       string             `json:"responseCode"`
		ResponseMessage    string             `json:"responseMessage"`
		VirtualAccountData VirtualAccountData `json:"virtualAccountData"`
		AdditionalInfo     T                  `json:"additionalInfo"`
	}
)

type (
	TransferVaCreateVaRequest[T any] struct {
		PartnerServiceID      string       `json:"partnerServiceId"`
		CustomerNo            string       `json:"customerNo"` // Keep as string for large numbers
		VirtualAccountNo      string       `json:"virtualAccountNo"`
		VirtualAccountName    string       `json:"virtualAccountName"`
		VirtualAccountEmail   string       `json:"virtualAccountEmail"`
		VirtualAccountPhone   string       `json:"virtualAccountPhone"`
		TrxID                 string       `json:"trxId"`
		TotalAmount           Amount       `json:"totalAmount"`
		BillDetails           []BillDetail `json:"billDetails"`
		FreeTexts             []Lang       `json:"freeTexts"`
		VirtualAccountTrxType string       `json:"virtualAccountTrxType"`
		FeeAmount             Amount       `json:"feeAmount"`
		ExpiredDate           string       `json:"expiredDate"`
		AdditionalInfo        T            `json:"additionalInfo"`
	}
	TransferVaCreateVaResponse[T any] struct {
		ResponseCode       string             `json:"responseCode"`
		ResponseMessage    string             `json:"responseMessage"`
		VirtualAccountData VirtualAccountData `json:"virtualAccountData"`
		AdditionalInfo     T                  `json:"additionalInfo"`
	}
)

type (
	TransferVaUpdateVaRequest[T any] struct {
		PartnerServiceID      string       `json:"partnerServiceId"`
		CustomerNo            string       `json:"customerNo"` // Keep as string for large numbers
		VirtualAccountNo      string       `json:"virtualAccountNo"`
		VirtualAccountName    string       `json:"virtualAccountName"`
		VirtualAccountEmail   string       `json:"virtualAccountEmail"`
		VirtualAccountPhone   string       `json:"virtualAccountPhone"`
		TrxID                 string       `json:"trxId"`
		TotalAmount           Amount       `json:"totalAmount"`
		BillDetails           []BillDetail `json:"billDetails"`
		FreeTexts             []Lang       `json:"freeTexts"`
		VirtualAccountTrxType string       `json:"virtualAccountTrxType"`
		FeeAmount             Amount       `json:"feeAmount"`
		ExpiredDate           string       `json:"expiredDate"`
		AdditionalInfo        T            `json:"additionalInfo"`
	}
	TransferVaUpdateVaResponse[T any] struct {
		ResponseCode       string             `json:"responseCode"`
		ResponseMessage    string             `json:"responseMessage"`
		VirtualAccountData VirtualAccountData `json:"virtualAccountData"`
		AdditionalInfo     T                  `json:"additionalInfo"`
	}
)

type (
	TransferVaUpdateStatusVaRequest[T any] struct {
		PartnerServiceID string `json:"partnerServiceId"`
		CustomerNo       string `json:"customerNo"` // Keep as string for large numbers
		VirtualAccountNo string `json:"virtualAccountNo"`
		TrxID            string `json:"trxId"`
	}
	TransferVaUpdateStatusVaResponse[T any] struct {
		ResponseCode       string             `json:"responseCode"`
		ResponseMessage    string             `json:"responseMessage"`
		VirtualAccountData VirtualAccountData `json:"virtualAccountData"`
		AdditionalInfo     T                  `json:"additionalInfo"`
	}
)

type (
	TransferVaInquiryVaRequest[T any] struct {
		PartnerServiceID string `json:"partnerServiceId"`
		CustomerNo       string `json:"customerNo"` // Keep as string for large numbers
		VirtualAccountNo string `json:"virtualAccountNo"`
		AdditionalInfo   T      `json:"additionalInfo"`
	}
	TransferVaInquiryVaResponse[T any] struct {
		ResponseCode       string             `json:"responseCode"`
		ResponseMessage    string             `json:"responseMessage"`
		VirtualAccountData VirtualAccountData `json:"virtualAccountData"`
		AdditionalInfo     T                  `json:"additionalInfo"`
	}
)

type (
	TransferVaDeleteVaRequest[T any] struct {
		PartnerServiceID string `json:"partnerServiceId"`
		CustomerNo       string `json:"customerNo"` // Keep as string for large numbers
		VirtualAccountNo string `json:"virtualAccountNo"`
		TrxID            string `json:"trxId"`
		AdditionalInfo   T      `json:"additionalInfo"`
	}
	TransferVaDeleteVaResponse[T any] struct {
		ResponseCode       string             `json:"responseCode"`
		ResponseMessage    string             `json:"responseMessage"`
		VirtualAccountData VirtualAccountData `json:"virtualAccountData"`
		AdditionalInfo     T                  `json:"additionalInfo"`
	}
)

type (
	TransferVaInquiryIntrabankRequest[T any] struct {
		PartnerServiceID   string `json:"partnerServiceId"`
		PartnerReferenceNo string `json:"partnerReferenceNo"`
		CustomerNo         int64  `json:"customerNo"` // Use int64 for large numbers
		VirtualAccountNo   string `json:"virtualAccountNo"`
		TrxDateTime        string `json:"trxDateTime"`
		ChannelCode        int    `json:"channelCode"`
		Language           string `json:"language"`
		Amount             Amount `json:"amount"`
		SourceAccountNo    string `json:"sourceAccountNo"`
		SourceAccountType  string `json:"sourceAccountType"`
		AdditionalInfo     T      `json:"additionalInfo"`
	}
	TransferVaInquiryIntrabankResponse[T any] struct {
		ResponseCode       string             `json:"responseCode"`
		ResponseMessage    string             `json:"responseMessage"`
		VirtualAccountData VirtualAccountData `json:"virtualAccountData"`
		AdditionalInfo     T                  `json:"additionalInfo"`
	}
)

type (
	TransferVaPaymentIntrabankRequest[T any] struct {
		PartnerServiceID        string       `json:"partnerServiceId"`
		CustomerNo              int64        `json:"customerNo"`  // Use int64 for large numbers
		ReferenceNo             int64        `json:"referenceNo"` // Use int64 for large numbers
		VirtualAccountNo        string       `json:"virtualAccountNo"`
		VirtualAccountName      string       `json:"virtualAccountName"`
		VirtualAccountEmail     string       `json:"virtualAccountEmail"`
		VirtualAccountPhone     string       `json:"virtualAccountPhone"`
		SourceAccountNo         string       `json:"sourceAccountNo"`
		SourceAccountType       string       `json:"sourceAccountType"`
		InquiryRequestID        string       `json:"inquiryRequestId"`
		PartnerReferenceNo      string       `json:"partnerReferenceNo"`
		PaidAmount              Amount       `json:"paidAmount"`
		CumulativePaymentAmount Amount       `json:"cumulativePaymentAmount"`
		PaidBills               string       `json:"paidBills"`
		TotalAmount             Amount       `json:"totalAmount"`
		TrxDateTime             string       `json:"trxDateTime"`
		JournalNum              string       `json:"journalNum"`
		PaymentType             int          `json:"paymentType"`
		FlagAdvise              string       `json:"flagAdvise"`
		PaymentStatus           string       `json:"paymentStatus"`
		BillDetails             []BillDetail `json:"billDetails"`
		FreeTexts               []Lang       `json:"freeTexts"`
		FeeAmount               Amount       `json:"feeAmount"`
		AdditionalInfo          T            `json:"additionalInfo"`
	}
	TransferVaPaymentIntrabankResponse[T any] struct {
		ResponseCode       string             `json:"responseCode"`
		ResponseMessage    string             `json:"responseMessage"`
		VirtualAccountData VirtualAccountData `json:"virtualAccountData"`
		AdditionalInfo     T                  `json:"additionalInfo"`
	}
)

type (
	TransferVaNotifyIntrabankRequest[T any] struct {
		PartnerServiceID   string `json:"partnerServiceId"`
		CustomerNo         int64  `json:"customerNo"` // Use int64 for large numbers
		VirtualAccountNo   string `json:"virtualAccountNo"`
		InquiryRequestID   string `json:"inquiryRequestId"`
		PaymentRequestID   string `json:"paymentRequestId"`
		PartnerReferenceNo string `json:"partnerReferenceNo"`
		TrxDateTime        string `json:"trxDateTime"`
		PaymentStatus      string `json:"paymentStatus"`
		AdditionalInfo     T      `json:"additionalInfo"`
		PaymentFlagReason  Lang   `json:"paymentFlagReason"`
	}
	TransferVaNotifyIntrabankResponse[T any] struct {
		ResponseCode       string             `json:"responseCode"`
		ResponseMessage    string             `json:"responseMessage"`
		VirtualAccountData VirtualAccountData `json:"virtualAccountData"`
		AdditionalInfo     T                  `json:"additionalInfo"`
	}
)

type (
	TransferVaGetReportRequest[T any] struct {
		PartnerServiceID string `json:"partnerServiceId"`
		StartDate        string `json:"startDate"`
		StartTime        string `json:"startTime"`
		EndDate          string `json:"endDate"`
		EndTime          string `json:"endTime"`
		AdditionalInfo   T      `json:"additionalInfo"`
	}
	TransferVaGetReportResponse[T any] struct {
		ResponseCode       string             `json:"responseCode"`
		ResponseMessage    string             `json:"responseMessage"`
		VirtualAccountData VirtualAccountData `json:"virtualAccountData"`
		AdditionalInfo     T                  `json:"additionalInfo"`
	}
)
