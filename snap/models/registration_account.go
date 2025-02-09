package models

type (
	RegistrationAccountRequest[T any] struct {
		PartnerReferenceNo string     `json:"partnerReferenceNo"`
		CountryCode        string     `json:"countryCode"`
		CustomerID         string     `json:"customerId"`
		DeviceInfo         DeviceInfo `json:"deviceInfo"`
		Email              string     `json:"email"`
		Lang               string     `json:"lang"`
		Locale             string     `json:"locale"`
		Name               string     `json:"name"`
		OnboardingPartner  string     `json:"onboardingPartner"`
		PhoneNo            string     `json:"phoneNo"`
		RedirectURL        string     `json:"redirectUrl"`
		Scopes             string     `json:"scopes"`
		SeamlessData       string     `json:"seamlessData"`
		SeamlessSign       string     `json:"seamlessSign"` // Note: Fixed missing colon
		State              string     `json:"state"`
		MerchantID         string     `json:"merchantId"`
		SubMerchantID      string     `json:"subMerchantId"`
		TerminalType       string     `json:"terminalType"`
		AdditionalInfo     T          `json:"additionalInfo"`
	}
	RegistrationAccountResponse[T any] struct {
		PResponseCode      string `json:"responseCode"`
		ResponseMessage    string `json:"responseMessage"`
		ReferenceNo        string `json:"referenceNo"`
		PartnerReferenceNo string `json:"partnerReferenceNo"`
		AuthCode           string `json:"authCode"`
		ApiKey             string `json:"apiKey"`
		AccountId          string `json:"accountId"`
		State              string `json:"state"`
		AdditionalInfo     T      `json:"additionalInfo"`
	}
)

type (
	RegistrationAccountBindingRequest[T any] struct {
		PartnerReferenceNo string        `json:"partnerReferenceNo"`
		Action             string        `json:"action"`
		AuthCode           string        `json:"authCode"`
		GrantType          string        `json:"grantType"`
		IsBindAndPay       string        `json:"isBindAndPay"`
		Lang               string        `json:"lang"`
		Locale             string        `json:"locale"`
		MerchantID         string        `json:"merchantId"`
		SubMerchantID      string        `json:"subMerchantId"`
		Msisdn             string        `json:"msisdn"`
		Otp                string        `json:"otp"`
		PhoneNo            string        `json:"phoneNo"`
		PlatformType       string        `json:"platformType"`
		RedirectURL        string        `json:"redirectUrl"`
		ReferenceID        string        `json:"referenceId"`
		RefreshToken       string        `json:"refreshToken"`
		SuccessParams      SuccessParams `json:"successParams"`
		TerminalID         string        `json:"terminalId"`
		TokenRequestorID   string        `json:"tokenRequestorId"`
		AdditionalInfo     T             `json:"additionalInfo"`
	}
	RegistrationAccountBindingResponse[T any] struct {
		ResponseCode       string          `json:"responseCode"`
		ResponseMessage    string          `json:"responseMessage"`
		ReferenceNo        string          `json:"referenceNo"`
		PartnerReferenceNo string          `json:"partnerReferenceNo"`
		AccountToken       string          `json:"accountToken"`
		AccessTokenInfo    AccessTokenInfo `json:"accessTokenInfo"`
		LinkId             string          `json:"linkId"`
		NextAction         string          `json:"nextAction"`
		LinkageToken       string          `json:"linkageToken"`
		Params             Params          `json:"params"`
		RedirectUrl        string          `json:"redirectUrl"`
		UserInfo           UserInfo        `json:"userInfo"`
		AdditionalInfo     T               `json:"additionalInfo"`
	}
)

type (
	RegistrationAccountInquiryRequest[T any] struct {
		PartnerReferenceNo string `json:"partnerReferenceNo"`
		AdditionalInfo     T      `json:"additionalInfo"`
	}
	RegistrationAccountInquiryResponse[T any] struct {
		ResponseCode            string `json:"responseCode"`
		ResponseMessage         string `json:"responseMessage"`
		ReferenceNo             string `json:"referenceNo"`
		PartnerReferenceNo      string `json:"partnerReferenceNo"`
		AccountCurrency         string `json:"accountCurrency"`
		AccountName             string `json:"accountName"`
		AccountNo               string `json:"accountNo"`
		AccountTransactionLimit string `json:"accountTransactionLimit"`
		EndDatePeriod           string `json:"endDatePeriod"`
		StartDatePeriod         string `json:"startDatePeriod"`
		AdditionalInfo          T      `json:"additionalInfo"`
	}
)

type (
	RegistrationAccountUnbindingRequest[T any] struct {
		PartnerReferenceNo string `json:"partnerReferenceNo"`
		LinkID             string `json:"linkId"`
		MerchantID         string `json:"merchantId"`
		SubMerchantID      string `json:"subMerchantId"`
		TokenID            string `json:"tokenId"`
		AdditionalInfo     T      `json:"additionalInfo"`
	}
	RegistrationAccountUnbindingResponse[T any] struct {
		ResponseCode       string `json:"responseCode"`
		ResponseMessage    string `json:"responseMessage"`
		ReferenceNo        string `json:"referenceNo"`
		PartnerReferenceNo string `json:"partnerReferenceNo"`
		MerchantId         string `json:"merchantId"`
		SubMerchantId      string `json:"subMerchantId"`
		LinkId             string `json:"linkId"`
		UnlinkResult       string `json:"unlinkResult"`
		AdditionalInfo     T      `json:"additionalInfo"`
	}
)

type (
	GetAuthCodeRequest[T any] struct {
		State        string `query:"state"`
		Scope        string `query:"scope"`
		RedirectURL  string `query:"redirectUrl"`
		SeamlessData string `query:"seamlessData"` //base64 encoded data string object
	}
	GetAuthCodeResponse[T any] struct {
		ResponseCode    string `json:"responseCode"`
		ResponseMessage string `json:"responseMessage"`
		AuthCode        string `json:"authCode"`
		State           string `query:"state"`
	}
)
