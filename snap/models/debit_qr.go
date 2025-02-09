package models

import "time"

type (
	QrCpmGenerateRequest[T any] struct {
		PartnerReferenceNo string    `json:"partnerReferenceNo"`
		UserAccessToken    string    `json:"userAccessToken"`
		MerchantId         string    `json:"merchantId"`
		SubMerchantId      string    `json:"subMerchantId"`
		PartnerTrxDate     time.Time `json:"partnerTrxDate"`
		AdditionalInfo     T         `json:"additionalInfo"`
	}
	QrCpmGenerateResponse[T any] struct {
		ResponseCode       string    `json:"responseCode"`
		ResponseMessage    string    `json:"responseMessage"`
		ReferenceNo        string    `json:"referenceNo"`
		PartnerReferenceNo string    `json:"partnerReferenceNo"`
		QrContent          string    `json:"qrContent"`
		QrUrl              string    `json:"qrUrl"`
		ExpiryTime         time.Time `json:"expiryTime"`
		AdditionalInfo     T         `json:"additionalInfo"`
	}
)

type (
	QrCpmPaymentRequest[T any] struct {
		PartnerReferenceNo string      `json:"partnerReferenceNo"`
		QrContent          string      `json:"qrContent"`
		Amount             Amount      `json:"amount"`
		FeeAmount          Amount      `json:"feeAmount"`
		MerchantId         string      `json:"merchantId"`
		SubMerchantId      string      `json:"subMerchantId"`
		Title              string      `json:"title"`
		ExpiryTime         time.Time   `json:"expiryTime"`
		Items              Items       `json:"items"`
		ExternalStoreId    string      `json:"externalStoreId"`
		MerchantName       string      `json:"merchantName"`
		MerchantLocation   string      `json:"merchantLocation"`
		AcquirerName       string      `json:"acquirerName"`
		TerminalId         string      `json:"terminalId"`
		ScannerInfo        ScannerInfo `json:"scannerInfo"`
		AdditionalInfo     T           `json:"additionalInfo"`
	}
	QrCpmPaymentResponse[T any] struct {
		ResponseCode       string    `json:"responseCode"`
		ResponseMessage    string    `json:"responseMessage"`
		ReferenceNo        string    `json:"referenceNo"`
		PartnerReferenceNo string    `json:"partnerReferenceNo"`
		TransactionDate    time.Time `json:"transactionDate"`
		AdditionalInfo     T         `json:"additionalInfo"`
	}
)

type (
	QrCpmQueryRequest[T any] struct {
		OriginalReferenceNo        string `json:"originalReferenceNo"`
		OriginalPartnerReferenceNo string `json:"originalPartnerReferenceNo"`
		OriginalExternalId         string `json:"originalExternalId"`
		MerchantId                 string `json:"merchantId"`
		SubMerchantId              string `json:"subMerchantId"`
		ExternalStoreId            string `json:"externalStoreId"`
		AdditionalInfo             T      `json:"additionalInfo"`
	}
	QrCpmQueryResponse[T any] struct {
		ResponseCode               string    `json:"responseCode"`
		ResponseMessage            string    `json:"responseMessage"`
		OriginalReferenceNo        string    `json:"originalReferenceNo"`
		OriginalPartnerReferenceNo string    `json:"originalPartnerReferenceNo"`
		OriginalExternalId         string    `json:"originalExternalId"`
		Title                      string    `json:"title"`
		LatestTransactionStatus    string    `json:"latestTransactionStatus"`
		TransactionStatusDesc      string    `json:"transactionStatusDesc"`
		PaidTime                   time.Time `json:"paidTime"`
		AdditionalInfo             T         `json:"additionalInfo"`
	}
)

type (
	QrCpmCancelRequest[T any] struct {
		OriginalPartnerReferenceNo string `json:"originalPartnerReferenceNo"`
		OriginalReferenceNo        string `json:"originalReferenceNo"`
		OriginalExternalId         string `json:"originalExternalId"`
		MerchantId                 string `json:"merchantId"`
		SubMerchantId              string `json:"subMerchantId"`
		ExternalStoreId            string `json:"externalStoreId"`
		Amount                     Amount `json:"amount"`
		Reason                     string `json:"reason"`
		AdditionalInfo             T      `json:"additionalInfo"`
	}
	QrCpmCancelResponse[T any] struct {
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
	QrCpmNotifyRequest[T any] struct {
		OriginalPartnerReferenceNo string `json:"originalPartnerReferenceNo"`
		OriginalReferenceNo        string `json:"originalReferenceNo"`
		MerchantId                 string `json:"merchantId"`
		SubMerchantId              string `json:"subMerchantId"`
		ExternalStoreId            string `json:"externalStoreId"`
		Amount                     Amount `json:"amount"`
		LatestTransactionStatus    string `json:"latestTransactionStatus"`
		TransactionStatusDesc      string `json:"transactionStatusDesc"`
		CustomerNumber             string `json:"customerNumber"`
		AccountType                string `json:"accountType"`
		DestinationNumber          string `json:"destinationNumber"`
		DestinationAccountName     string `json:"destinationAccountName"`
		SessionId                  string `json:"sessionId"`
		BankCode                   string `json:"bankCode"`
		AdditionalInfo             T      `json:"additionalInfo"`
	}
	QrCpmNotifyResponse[T any] struct {
		ResponseCode    string `json:"responseCode"`
		ResponseMessage string `json:"responseMessage"`
		AdditionalInfo  T      `json:"additionalInfo"`
	}
)

type (
	QrCpmRefundRequest[T any] struct {
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
	QrCpmRefundResponse[T any] struct {
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
