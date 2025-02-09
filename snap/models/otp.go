package models

type (
	OTPVerificationRequest[T any] struct {
		OriginalPartnerReferenceNo string `json:"originalPartnerReferenceNo"`
		OriginalReferenceNo        string `json:"originalReferenceNo"`
		Action                     string `json:"action"`
		MerchantID                 string `json:"merchantId"`
		Otp                        string `json:"otp"`
		ChargeToken                string `json:"chargeToken"`
		Type                       string `json:"type"`
		AdditionalInfo             T      `json:"additionalInfo"`
	}

	OTPVerificationResponse[T any] struct {
		ResponseCode               string  `json:"responseCode"`
		ResponseMessage            string  `json:"responseMessage"`
		OriginalReferenceNo        string  `json:"originalReferenceNo"`
		OriginalPartnerReferenceNo string  `json:"originalPartnerReferenceNo"`
		AccountNo                  string  `json:"accountNo"`
		BankCardToken              string  `json:"bankCardToken"`
		CardPan                    string  `json:"cardPan"`
		CustomerID                 string  `json:"customerId"`
		Email                      string  `json:"email"`
		ExpiredDatetime            string  `json:"expiredDatetime"`
		ExpiryDate                 string  `json:"expiryDate"`
		IdentificationNo           string  `json:"identificationNo"`
		LinkageToken               string  `json:"linkageToken"`
		PhoneNo                    string  `json:"phoneNo"`
		QParamsURL                 string  `json:"qParamsURL"`
		QParams                    QParams `json:"qParams"`
		SendOtpFlag                string  `json:"sendOtpFlag"`
		SubscribeDatetime          string  `json:"subscribeDatetime"`
		TokenExpiryTime            string  `json:"tokenExpiryTime"`
		TransactionTimestamp       string  `json:"transactionTimestamp"`
		AdditionalInfo             T       `json:"additionalInfo"`
	}
)

type (
	OTPRequest[T any] struct {
		PartnerReferenceNo string `json:"partnerReferenceNo"`
		JourneyID          string `json:"journeyId"`
		MerchantID         string `json:"merchantId"`
		SubMerchant        string `json:"subMerchant"`
		ExternalStoreID    string `json:"externalStoreId"`
		TrxDateTime        string `json:"trxDateTime"`
		BankCardToken      string `json:"bankCardToken"`
		OtpTrxCode         string `json:"otpTrxCode"`
		OtpReasonCode      string `json:"otpReasonCode"`
		OtpReasonMessage   string `json:"otpReasonMessage"`
		AdditionalInfo     T      `json:"additionalInfo"`
	}
	OTPResponse[T any] struct {
		ResponseCode       string `json:"responseCode"`
		ResponseMessage    string `json:"responseMessage"`
		ReferenceNo        string `json:"referenceNo"`
		PartnerReferenceNo string `json:"partnerReferenceNo"`
		ChargeToken        string `json:"chargeToken"`
		AdditionalInfo     T      `json:"additionalInfo"`
	}
)
