package models

import "time"

type (
	DebitPaymentRequest[T, J any] struct {
		PartnerReferenceNo string               `json:"partnerReferenceNo"`
		BankCardToken      string               `json:"bankCardToken"`
		ChargeToken        string               `json:"chargeToken"`
		Otp                string               `json:"otp"`
		MerchantId         string               `json:"merchantId"`
		TerminalId         string               `json:"terminalId"`
		JourneyId          string               `json:"journeyId"`
		SubMerchantId      string               `json:"subMerchantId"`
		Amount             Amount               `json:"amount"`
		UrlParam           []UrlParam           `json:"urlParam"`
		ExternalStoreId    string               `json:"externalStoreId"`
		ValidUpTo          string               `json:"validUpTo"`
		PointOfInitiation  string               `json:"pointOfInitiation"`
		FeeType            string               `json:"feeType"`
		DisabledPayMethods string               `json:"disabledPayMethods"`
		PayOptionDetails   []PayOptionDetail[J] `json:"payOptionDetails"`
		AdditionalInfo     T                    `json:"additionalInfo"`
	}
	DebitPaymentResponse[T any] struct {
		ResponseCode       string `json:"responseCode"`
		ResponseMessage    string `json:"responseMessage"`
		ReferenceNo        string `json:"referenceNo"`
		PartnerReferenceNo string `json:"partnerReferenceNo"`
		ApprovalCode       string `json:"approvalCode"`
		AppRedirectUrl     string `json:"appRedirectUrl"`
		WebRedirectUrl     string `json:"webRedirectUrl"`
		AdditionalInfo     T      `json:"additionalInfo"`
	}
)

type (
	DebitStatusRequest[T any] struct {
		OriginalPartnerReferenceNo string    `json:"originalPartnerReferenceNo"`
		OriginalReferenceNo        string    `json:"originalReferenceNo"`
		OriginalExternalId         string    `json:"originalExternalId"`
		ServiceCode                string    `json:"serviceCode"`
		TransactionDate            time.Time `json:"transactionDate"`
		Amount                     Amount    `json:"amount"`
		MerchantId                 string    `json:"merchantId"`
		SubMerchantId              string    `json:"subMerchantId"`
		ExternalStoreId            string    `json:"externalStoreId"`
		AdditionalInfo             T         `json:"additionalInfo"`
	}
	DebitStatusResponse[T any] struct {
		ResponseCode               string    `json:"responseCode"`
		ResponseMessage            string    `json:"responseMessage"`
		OriginalPartnerReferenceNo string    `json:"originalPartnerReferenceNo"`
		OriginalReferenceNo        string    `json:"originalReferenceNo"`
		ApprovalCode               string    `json:"approvalCode"`
		OriginalExternalId         string    `json:"originalExternalId"`
		ServiceCode                string    `json:"serviceCode"`
		LatestTransactionStatus    string    `json:"latestTransactionStatus"`
		TransactionStatusDesc      string    `json:"transactionStatusDesc"`
		OriginalResponseCode       string    `json:"originalResponseCode"`
		OriginalResponseMessage    string    `json:"originalResponseMessage"`
		SessionId                  string    `json:"sessionId"`
		RequestId                  string    `json:"requestId"`
		RefundHistory              []Refund  `json:"refundHistory"`
		TransAmount                Amount    `json:"transAmount"`
		FeeAmount                  Amount    `json:"feeAmount"`
		PaidTime                   time.Time `json:"paidTime"`
	}
)

type (
	DebitPaymentNotifyRequest[T any] struct {
		OriginalPartnerReferenceNo string    `json:"originalPartnerReferenceNo"`
		OriginalReferenceNo        string    `json:"originalReferenceNo"`
		OriginalExternalId         string    `json:"originalExternalId"`
		MerchantId                 string    `json:"merchantId"`
		SubMerchantId              string    `json:"subMerchantId"`
		Amount                     Amount    `json:"amount"`
		LatestTransactionStatus    string    `json:"latestTransactionStatus"`
		TransactionStatusDesc      string    `json:"transactionStatusDesc"`
		CreatedTime                time.Time `json:"createdTime"`
		FinishedTime               time.Time `json:"finishedTime"`
		ExternalStoreId            string    `json:"externalStoreId"`
		AdditionalInfo             T         `json:"additionalInfo"`
	}
	DebitPaymentNotifyResponse[T any] struct {
		ResponseCode    string `json:"responseCode"`
		ResponseMessage string `json:"responseMessage"`
		ApprovalCode    string `json:"approvalCode"`
	}
)

type (
	DebitCancelRequest[T any] struct {
		OriginalPartnerReferenceNo string `json:"originalPartnerReferenceNo"`
		OriginalReferenceNo        string `json:"originalReferenceNo"`
		ApprovalCode               string `json:"approvalCode"`
		OriginalExternalId         string `json:"originalExternalId"`
		MerchantId                 string `json:"merchantId"`
		SubMerchantId              string `json:"subMerchantId"`
		Reason                     string `json:"reason"`
		ExternalStoreId            string `json:"externalStoreId"`
		Amount                     Amount `json:"amount"`
		AdditionalInfo             T      `json:"additionalInfo"`
	}
	DebitCancelResponse[T any] struct {
		ResponseCode               string    `json:"responseCode"`
		ResponseMessage            string    `json:"responseMessage"`
		OriginalPartnerReferenceNo string    `json:"originalPartnerReferenceNo"`
		OriginalReferenceNo        string    `json:"originalReferenceNo"`
		OriginalExternalId         string    `json:"originalExternalId"`
		CancelTime                 time.Time `json:"cancelTime"`
		TransactionDate            time.Time `json:"transactionDate"`
		AdditionalInfo             T         `json:"additionalInfo"`
	}
)

type (
	DebitRefundRequest[T any] struct {
		MerchantId                 string `json:"merchantId"`
		SubMerchantId              string `json:"subMerchantId"`
		OriginalPartnerReferenceNo string `json:"originalPartnerReferenceNo"`
		OriginalReferenceNo        string `json:"originalReferenceNo"`
		OriginalExternalId         string `json:"originalExternalId"`
		PartnerRefundNo            string `json:"partnerRefundNo"`
		RefundAmount               Amount `json:"refundAmount"`
		ExternalStoreId            string `json:"externalStoreId"`
		Reason                     string `json:"reason"`
		AdditionalInfo             T      `json:"additionalInfo"`
	}
	DebitRefundResponse[T any] struct {
		ResponseCode               string    `json:"responseCode"`
		ResponseMessage            string    `json:"responseMessage"`
		OriginalPartnerReferenceNo string    `json:"originalPartnerReferenceNo"`
		OriginalReferenceNo        string    `json:"originalReferenceNo"`
		OriginalExternalId         string    `json:"originalExternalId"`
		PartnerTrxId               string    `json:"partnerTrxId"`
		RefundNo                   string    `json:"refundNo"`
		PartnerRefundNo            string    `json:"partnerRefundNo"`
		RefundAmount               Amount    `json:"refundAmount"`
		RefundTime                 time.Time `json:"refundTime"`
		AdditionalInfo             T         `json:"additionalInfo"`
	}
)
