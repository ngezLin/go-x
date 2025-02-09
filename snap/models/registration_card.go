package models

import "time"

type (
	RegistrationCardBindRequest[T any] struct {
		PartnerReferenceNo string `json:"partnerReferenceNo"`
		AccountName        string `json:"accountName"`
		CardData           string `json:"cardData"`
		CustIDMerchant     string `json:"custIdMerchant"`
		IsBindAndPay       string `json:"isBindAndPay"`
		MerchantID         string `json:"merchantId"`
		TerminalID         string `json:"terminalId"`
		JourneyID          string `json:"journeyId"`
		SubMerchantID      string `json:"subMerchantId"`
		ExternalStoreID    string `json:"externalStoreId"`
		Limit              string `json:"limit"`
		MerchantLogoURL    string `json:"merchantLogoUrl"`
		PhoneNo            string `json:"phoneNo"`
		SendOtpFlag        string `json:"sendOtpFlag"`
		Type               string `json:"type"`
		AdditionalInfo     T      `json:"additionalInfo"`
	}
	RegistrationCardBindResponse[T any] struct {
		ResponseCode       string `json:"responseCode"`
		ResponseMessage    string `json:"responseMessage"`
		ReferenceNo        string `json:"referenceNo"`
		PartnerReferenceNo string `json:"partnerReferenceNo"`
		BankCardToken      string `json:"bankCardToken"`
		ChargeToken        string `json:"chargeToken"`
		RandomString       string `json:"randomString"`
		TokenExpiryTime    string `json:"tokenExpiryTime"`
		AdditionalInfo     T      `json:"additionalInfo"`
	}
)

type (
	RegistrationCardInquirySetLimitRequest[T any] struct {
		PartnerReferenceNo string `json:"partnerReferenceNo"`
		BankAccountNo      string `json:"bankAccountNo"`
		BankCardNo         string `json:"bankCardNo"`
		Limit              string `json:"limit"`
		BankCardToken      string `json:"bankCardToken"`
		Otp                string `json:"otp"`
		AdditionalInfo     T      `json:"additionalInfo"`
	}
	RegistrationCardInquirySetLimitResponse[T any] struct {
		ResponseCode       string `json:"responseCode"`
		ResponseMessage    string `json:"responseMessage"`
		ReferenceNo        string `json:"referenceNo"`
		PartnerReferenceNo string `json:"partnerReferenceNo"`
		AdditionalInfo     T      `json:"additionalInfo"`
	}
)

type (
	Account struct {
		AccountData AccountData `json:"accountData"`
	}
	AccountData struct {
		AccountID      string `json:"accountId"`
		CreatedDate    string `json:"createdDate"`
		CredentialNo   string `json:"credentialNo"`
		CredentialType string `json:"credentialType"`
		MaxLimit       string `json:"maxLimit"`
		Status         string `json:"status"`
	}
	RegistrationCardInquiryResponse[T any] struct {
		ResponseCode    string    `json:"responseCode"`
		ResponseMessage string    `json:"responseMessage"`
		AccountList     []Account `json:"accountList"`
		AdditionalInfo  T         `json:"additionalInfo"`
	}
)

type (
	RegistrationCardUnbindRequest[T any] struct {
		PartnerReferenceNo string `json:"partnerReferenceNo"`
		Token              string `json:"token"`
		BankCardNo         string `json:"bankCardNo"`
		Type               string `json:"type"`
		Part               string `json:"part"`
		MerchantID         string `json:"merchantId"`
		SubMerchantID      string `json:"subMerchantId"`
		TerminalID         string `json:"terminalId"`
		TokenRequestorID   string `json:"tokenRequestorId"`
		JourneyID          string `json:"journeyId"`
		TransactionDate    string `json:"transactionDate"`
		AdditionalInfo     T      `json:"additionalInfo"`
	}
	RegistrationCardUnbindResponse[T any] struct {
		ResponseCode       string    `json:"responseCode"`
		ResponseMessage    string    `json:"responseMessage"`
		ReferenceNo        string    `json:"referenceNo"`
		PartnerReferenceNo string    `json:"partnerReferenceNo"`
		CustomerId         string    `json:"customerId"`
		UnsubscribeDate    time.Time `json:"unsubscribeDate"`
		AdditionalInfo     T         `json:"additionalInfo"`
	}
)
