package models

import "time"

type (
	DebitAuthPaymentRequest[T any] struct {
		PartnerReferenceNo string `json:"partnerReferenceNo"`
		MerchantId         string `json:"merchantId"`
		SubMerchantId      string `json:"subMerchantId"`
		Amount             Amount `json:"amount"`
		FeeType            string `json:"feeType"`
		Mcc                string `json:"mcc"`
		ProductCode        string `json:"productCode"`
		Title              string `json:"title"`
		Items              []Item `json:"items"`
		AdditionalInfo     T      `json:"additionalInfo"`
	}
	DebitAuthPaymentResponse[T any] struct {
		ResponseCode       string    `json:"responseCode"`
		ResponseMessage    string    `json:"responseMessage"`
		ReferenceNo        string    `json:"referenceNo"`
		PartnerReferenceNo string    `json:"partnerReferenceNo"`
		Amount             Amount    `json:"amount"`
		PaidTime           time.Time `json:"paidTime"`
		AdditionalInfo     T         `json:"additionalInfo"`
	}
)

type (
	DebitAuthQueryRequest[T any] struct {
		OriginalPartnerReferenceNo string `json:"originalPartnerReferenceNo"`
		OriginalReferenceNo        string `json:"originalReferenceNo"`
		MerchantId                 string `json:"merchantId"`
		SubMerchantId              string `json:"subMerchantId"`
		ExternalStoreId            string `json:"externalStoreId"`
		AdditionalInfo             T      `json:"additionalInfo"`
	}
	DebitAuthQueryResponse[T any] struct {
		ResponseCode               string    `json:"responseCode"`
		ResponseMessage            string    `json:"responseMessage"`
		OriginalPartnerReferenceNo string    `json:"originalpartnerReferenceNo"` // Note the lowercase 'p'
		OriginalReferenceNo        string    `json:"originalReferenceNo"`
		Amount                     Amount    `json:"amount"`
		PaidTime                   time.Time `json:"paidTime"`
		LatestTransactionStatus    string    `json:"latestTransactionStatus"`
		TransactionStatusDesc      string    `json:"transactionStatusDesc"`
		AdditionalInfo             T         `json:"additionalInfo"`
	}
)

type (
	DebitAuthCaptureRequest[T any] struct {
		OriginalReferenceNo        string `json:"originalReferenceNo"`
		OriginalPartnerReferenceNo string `json:"originalPartnerReferenceNo"`
		MerchantId                 string `json:"merchantId"`
		SubMerchantId              string `json:"subMerchantId"`
		PartnerCaptureNo           string `json:"partnerCaptureNo"`
		CaptureAmount              Amount `json:"captureAmount"`
		Title                      string `json:"title"`
		LastCapture                string `json:"lastCapture"`
		AdditionalInfo             T      `json:"additionalInfo"`
	}
	DebitAuthCaptureResponse[T any] struct {
		ResponseCode               string    `json:"responseCode"`
		ResponseMessage            string    `json:"responseMessage"`
		OriginalReferenceNo        string    `json:"originalReferenceNo"`
		OriginalPartnerReferenceNo string    `json:"originalPartnerReferenceNo"`
		PartnerCaptureNo           string    `json:"partnerCaptureNo"`
		CaptureNo                  string    `json:"captureNo"`
		CaptureAmount              Amount    `json:"captureAmount"`
		CaptureTime                time.Time `json:"captureTime"`
		AdditionalInfo             T         `json:"additionalInfo"`
	}
)

type (
	DebitAuthCaptureQueryRequest[T any] struct {
		OriginalReferenceNo        string `json:"originalReferenceNo"`
		OriginalPartnerReferenceNo string `json:"originalPartnerReferenceNo"`
		MerchantId                 string `json:"merchantId"`
		SubMerchantId              string `json:"subMerchantId"`
		CaptureNo                  string `json:"captureNo"`
		PartnerCaptureNo           string `json:"partnerCaptureNo"`
		AdditionalInfo             T      `json:"additionalInfo"`
	}
	DebitAuthCaptureQueryResponse[T any] struct {
		ResponseCode               string    `json:"responseCode"`
		ResponseMessage            string    `json:"responseMessage"`
		OriginalReferenceNo        string    `json:"originalReferenceNo"`
		OriginalPartnerReferenceNo string    `json:"originalPartnerReferenceNo"`
		CaptureNo                  string    `json:"captureNo"`
		CaptureAmount              Amount    `json:"captureAmount"`
		CaptureTime                time.Time `json:"captureTime"`
		LatestCaptureStatus        string    `json:"latestCaptureStatus"`
		PartnerCaptureNo           string    `json:"partnerCaptureNo"`
		AdditionalInfo             T         `json:"additionalInfo"`
	}
)

type (
	DebitAuthVoidRequest[T any] struct {
		OriginalReferenceNo        string `json:"originalReferenceNo"`
		OriginalPartnerReferenceNo string `json:"originalPartnerReferenceNo"`
		MerchantId                 string `json:"merchantId"`
		SubMerchantId              string `json:"subMerchantId"`
		VoidAmount                 Amount `json:"voidAmount"`
		PartnerVoidNo              string `json:"partnerVoidNo"`
		VoidRemainingAmount        string `json:"voidRemainingAmount"`
		Reason                     string `json:"reason"`
		AdditionalInfo             T      `json:"additionalInfo"`
	}
	DebitAuthVoidResponse[T any] struct {
		ResponseCode               string    `json:"responseCode"`
		ResponseMessage            string    `json:"responseMessage"`
		OriginalReferenceNo        string    `json:"originalReferenceNo"`
		OriginalPartnerReferenceNo string    `json:"originalPartnerReferenceNo"`
		VoidNo                     string    `json:"voidNo"`
		PartnerVoidNo              string    `json:"partnerVoidNo"`
		VoidAmount                 Amount    `json:"voidAmount"`
		VoidTime                   time.Time `json:"voidTime"`
		AdditionalInfo             T         `json:"additionalInfo"`
	}
)
type (
	DebitAuthVoidQueryRequest[T any] struct {
		OriginalReferenceNo        string `json:"originalReferenceNo"`
		OriginalPartnerReferenceNo string `json:"originalPartnerReferenceNo"`
		MerchantId                 string `json:"merchantId"`
		SubMerchantId              string `json:"subMerchantId"`
		VoidNo                     string `json:"voidNo"`
		PartnerVoidNo              string `json:"partnerVoidNo"`
		AdditionalInfo             T      `json:"additionalInfo"`
	}
	DebitAuthVoidQueryResponse[T any] struct {
		ResponseCode               string    `json:"responseCode"`
		ResponseMessage            string    `json:"responseMessage"`
		OriginalReferenceNo        string    `json:"originalReferenceNo"`
		OriginalPartnerReferenceNo string    `json:"originalPartnerReferenceNo"`
		VoidNo                     string    `json:"voidNo"`
		VoidAmount                 Amount    `json:"voidAmount"`
		VoidTime                   time.Time `json:"voidTime"`
		LatestVoidStatus           string    `json:"latestVoidStatus"`
		PartnerVoidNo              string    `json:"partnerVoidNo"`
		AdditionalInfo             T         `json:"additionalInfo"`
	}
)
type (
	DebitAuthRefundRequest[T any] struct {
		OriginalPartnerReferenceNo string `json:"originalPartnerReferenceNo"`
		OriginalReferenceNo        string `json:"originalReferenceNo"`
		PartnerRefundNo            string `json:"partnerRefundNo"`
		MerchantId                 string `json:"merchantId"`
		SubMerchantId              string `json:"subMerchantId"`
		OriginalCaptureNo          string `json:"originalCaptureNo"`
		RefundAmount               Amount `json:"refundAmount"`
		ExternalStoreId            string `json:"externalStoreId"`
		Reason                     string `json:"reason"`
		AdditionalInfo             T      `json:"additionalInfo"`
	}
	DebitAuthRefundResponse[T any] struct {
		ResponseCode               string    `json:"responseCode"`
		ResponseMessage            string    `json:"responseMessage"`
		OriginalCaptureNo          string    `json:"originalCaptureNo"`
		OriginalReferenceNo        string    `json:"originalReferenceNo"`
		OriginalPartnerReferenceNo string    `json:"originalPartnerReferenceNo"`
		PartnerRefundNo            string    `json:"partnerRefundNo"`
		RefundNo                   string    `json:"refundNo"`
		RefundAmount               Amount    `json:"refundAmount"`
		RefundTime                 time.Time `json:"refundTime"`
		AdditionalInfo             T         `json:"additionalInfo"`
	}
)
