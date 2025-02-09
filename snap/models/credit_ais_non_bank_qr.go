package models

import "time"

type (
	QrMpmGenerateRequest[T any] struct {
		PartnerReferenceNo string    `json:"partnerReferenceNo"`
		Amount             Amount    `json:"amount"`
		FeeAmount          Amount    `json:"feeAmount"`
		MerchantId         string    `json:"merchantId"`
		SubMerchantId      string    `json:"subMerchantId"`
		StoreId            string    `json:"storeId"`
		TerminalId         string    `json:"terminalId"`
		ValidityPeriod     time.Time `json:"validityPeriod"`
		AdditionalInfo     T         `json:"additionalInfo"`
	}
	QrMpmGenerateResponse[T any] struct {
		ResponseCode       string `json:"responseCode"`
		ResponseMessage    string `json:"responseMessage"`
		ReferenceNo        string `json:"referenceNo"`
		PartnerReferenceNo string `json:"partnerReferenceNo"`
		QrContent          string `json:"qrContent"`
		QrUrl              string `json:"qrUrl"`
		QrImage            string `json:"qrImage"`
		RedirectUrl        string `json:"redirectUrl"`
		MerchantName       string `json:"merchantName"`
		StoreId            string `json:"storeId"`
		TerminalId         string `json:"terminalId"`
		AdditionalInfo     T      `json:"additionalInfo"`
	}
)

type (
	QrMpmDecodeRequest[T any] struct {
		PartnerReferenceNo string    `json:"partnerReferenceNo"`
		QrContent          string    `json:"qrContent"`
		Amount             Amount    `json:"amount"`
		MerchantId         string    `json:"merchantId"`
		SubMerchantId      string    `json:"subMerchantId"`
		ScanTime           time.Time `json:"scanTime"`
		AdditionalInfo     T         `json:"additionalInfo"`
	}
	QrMpmDecodeResponse[T any] struct {
		ResponseCode       string         `json:"responseCode"`
		ResponseMessage    string         `json:"responseMessage"`
		ReferenceNo        string         `json:"referenceNo"`
		PartnerReferenceNo string         `json:"partnerReferenceNo"`
		RedirectUrl        string         `json:"redirectUrl"`
		MerchantName       string         `json:"merchantName"`
		MerchantCategory   string         `json:"merchantCategory"`
		MerchantLocation   string         `json:"merchantLocation"`
		MerchantInfos      []MerchantInfo `json:"merchantInfos"`
		TransactionAmount  Amount         `json:"transactionAmount"`
		FeeAmount          Amount         `json:"feeAmount"`
		AdditionalInfo     T              `json:"additionalInfo"`
	}
)

type (
	QrMpmApplyOTTRequest[T any] struct {
		UserResources  []string `json:"userResources"`
		AdditionalInfo T        `json:"additionalInfo"`
	}
	QrMpmApplyOTTResponse[T any] struct {
		ResponseCode    string         `json:"responseCode"`
		ResponseMessage string         `json:"responseMessage"`
		UserResources   []UserResource `json:"userResources"`
		AdditionalInfo  T              `json:"additionalInfo"`
	}
)

type (
	QrMpmPaymentRequestt[T any] struct {
		PartnerReferenceNo string `json:"partnerReferenceNo"`
		MerchantId         string `json:"merchantId"`
		SubMerchantId      string `json:"subMerchantId"`
		Amount             Amount `json:"amount"`
		FeeAmount          Amount `json:"feeAmount"`
		Otp                string `json:"otp"`
		VerificationId     string `json:"verificationId"`
		AdditionalInfo     T      `json:"additionalInfo"`
	}
	QrMpmPaymentResponse[T any] struct {
		ResponseCode       string    `json:"responseCode"`
		ResponseMessage    string    `json:"responseMessage"`
		ReferenceNo        string    `json:"referenceNo"`
		PartnerReferenceNo string    `json:"partnerReferenceNo"`
		TransactionDate    time.Time `json:"transactionDate"`
		Amount             Amount    `json:"amount"`
		FeeAmount          Amount    `json:"feeAmount"`
		VerificationId     string    `json:"verificationId"`
		AdditionalInfo     T         `json:"additionalInfo"`
	}
)

type (
	QrMpmQueryRequest[T any] struct {
		OriginalReferenceNo        string `json:"originalReferenceNo"`
		OriginalPartnerReferenceNo string `json:"originalPartnerReferenceNo"`
		OriginalExternalId         string `json:"originalExternalId"`
		ServiceCode                string `json:"serviceCode"`
		MerchantId                 string `json:"merchantId"`
		SubMerchantId              string `json:"subMerchantId"`
		ExternalStoreId            string `json:"externalStoreId"` // Note: Removed the space in the key
		AdditionalInfo             T      `json:"additionalInfo"`
	}
	QrMpmQueryResponse[T any] struct {
		ResponseCode               string    `json:"responseCode"`
		ResponseMessage            string    `json:"responseMessage"`
		OriginalReferenceNo        string    `json:"originalReferenceNo"`
		OriginalPartnerReferenceNo string    `json:"originalPartnerReferenceNo"`
		OriginalExternalId         string    `json:"originalExternalId"`
		ServiceCode                string    `json:"serviceCode"`
		LatestTransactionStatus    string    `json:"latestTransactionStatus"`
		TransactionStatusDesc      string    `json:"transactionStatusDesc"`
		PaidTime                   time.Time `json:"paidTime"`
		Amount                     Amount    `json:"amount"`
		FeeAmount                  Amount    `json:"feeAmount"`
		TerminalId                 string    `json:"terminalId"`
		AdditionalInfo             T         `json:"additionalInfo"`
	}
)

type (
	QrMpmNotifyRequest[T any] struct {
		OriginalReferenceNo        string `json:"originalReferenceNo"`
		OriginalPartnerReferenceNo string `json:"originalPartnerReferenceNo"`
		LatestTransactionStatus    string `json:"latestTransactionStatus"`
		TransactionStatusDesc      string `json:"transactionStatusDesc"`
		CustomerNumber             string `json:"customerNumber"`
		AccountType                string `json:"accountType"`
		DestinationNumber          string `json:"destinationNumber"`
		DestinationAccountName     string `json:"destinationAccountName"`
		Amount                     Amount `json:"amount"`
		SessionId                  string `json:"sessionId"`
		BankCode                   string `json:"bankCode"`
		ExternalStoreId            string `json:"externalStoreId"`
		AdditionalInfo             T      `json:"additionalInfo"`
	}
	QrMpmNotifyResponse[T any] struct {
		ResponseCode    string `json:"responseCode"`
		ResponseMessage string `json:"responseMessage"`
		AdditionalInfo  T      `json:"additionalInfo"`
	}
)

type (
	QrMpmCancelRequest[T any] struct {
		OriginalPartnerReferenceNo string `json:"originalPartnerReferenceNo"`
		OriginalReferenceNo        string `json:"originalReferenceNo"`
		OriginalExternalId         string `json:"originalExternalId"`
		MerchantId                 string `json:"merchantId"`
		SubMerchantId              string `json:"subMerchantId"`
		ExternalStoreId            string `json:"externalStoreId"`
		Reason                     string `json:"reason"`
		Amount                     Amount `json:"amount"`
		AdditionalInfo             T      `json:"additionalInfo"`
	}
	QrMpmCancelResponse[T any] struct {
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
	QrMpmRefundRequest[T any] struct {
		MerchantId                 string `json:"merchantId"`
		SubMerchantId              string `json:"subMerchantId"`
		ExternalStoreId            string `json:"externalStoreId"`
		OriginalPartnerReferenceNo string `json:"originalPartnerReferenceNo"`
		OriginalReferenceNo        string `json:"originalReferenceNo"`
		OriginalExternalId         string `json:"originalExternalId"`
		PartnerRefundNo            string `json:"partnerRefundNo"`
		RefundAmount               Amount `json:"refundAmount"`
		Reason                     string `json:"reason"`
		AdditionalInfo             T      `json:"additionalInfo"`
	}
	QrMpmRefundResponse[T any] struct {
		ResponseCode               string    `json:"responseCode"`
		ResponseMessage            string    `json:"responseMessage"`
		OriginalPartnerReferenceNo string    `json:"originalPartnerReferenceNo"`
		OriginalReferenceNo        string    `json:"originalReferenceNo"`
		OriginalExternalId         string    `json:"originalExternalId"`
		RefundNo                   string    `json:"refundNo"`
		PartnerRefundNo            string    `json:"partnerRefundNo"`
		RefundAmount               Amount    `json:"refundAmount"`
		RefundTime                 time.Time `json:"refundTime"`
		AdditionalInfo             T         `json:"additionalInfo"`
	}
)

type (
	QrMpmStatusRequest[T any] struct {
		OriginalPartnerReferenceNo string    `json:"originalPartnerReferenceNo"`
		OriginalReferenceNo        string    `json:"originalReferenceNo"`
		OriginalExternalId         string    `json:"originalExternalId"`
		ServiceCode                string    `json:"serviceCode"`
		TransactionDate            time.Time `json:"transactionDate"`
		Amount                     Amount    `json:"amount"`
		AdditionalInfo             T         `json:"additionalInfo"`
	}
	QrMpmStatusResponse[T any] struct {
		ResponseCode               string    `json:"responseCode"`
		ResponseMessage            string    `json:"responseMessage"`
		OriginalReferenceNo        string    `json:"originalReferenceNo"`
		OriginalPartnerReferenceNo string    `json:"originalPartnerReferenceNo"`
		OriginalExternalId         string    `json:"originalExternalId"`
		ServiceCode                string    `json:"serviceCode"`
		LatestTransactionStatus    string    `json:"latestTransactionStatus"`
		TransactionStatusDesc      string    `json:"transactionStatusDesc"`
		OriginalResponseCode       string    `json:"originalResponseCode"`
		OriginalResponseMessage    string    `json:"originalResponseMessage"`
		SessionId                  string    `json:"sessionId"`
		RequestId                  string    `json:"requestId"`
		TransactionDate            time.Time `json:"transactionDate"`
		Amount                     Amount    `json:"amount"`
		AdditionalInfo             T         `json:"additionalInfo"`
	}
)
