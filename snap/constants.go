package snap

const X_TIMESTAMP = "X-TIMESTAMP"
const X_SIGNATURE = "X-SIGNATURE"
const X_CLIENT_KEY = "X-CLIENT-KEY"
const X_PARTNER_ID = "X-PARTNER-ID"
const X_EXTERNAL_ID = "X-EXTERNAL-ID"
const X_IP_ADDRESS = "X-IP-ADDRESS"
const X_DEVICE_ID = "X-DEVICE-ID"
const X_LONGITUDE = "X-LONGITUDE"
const X_LATITUDE = "X-LATITUDE"
const X_IDEMPOTENTCY = "X-Idempotency-Key"

// Service Code Constants
const (
	// Authentication Service Codes
	SERVICE_CODE_ACCESS_TOKEN_B2B   = "73"
	SERVICE_CODE_ACCESS_TOKEN_B2B2C = "74"

	// Registration Service Codes
	SERVICE_CODE_REGISTRATION_CARD_BIND                  = "01"
	SERVICE_CODE_REGISTRATION_CARD_INQUIRY               = "02"
	SERVICE_CODE_REGISTRATION_CARD_INQUIRY_CUST_MERCHANT = "03"
	SERVICE_CODE_OTP_VERIFICATION                        = "04"
	SERVICE_CODE_REGISTRATION_CARD_UNBIND                = "05"
	SERVICE_CODE_REGISTRATION_ACCOUNT_CREATION           = "06"
	SERVICE_CODE_REGISTRATION_ACCOUNT_BINDING            = "07"
	SERVICE_CODE_REGISTRATION_ACCOUNT_INQUIRY            = "08"
	SERVICE_CODE_REGISTRATION_ACCOUNT_UNBINDING          = "09"
	SERVICE_CODE_GET_AUTH_CODE                           = "10"
	SERVICE_CODE_OTP                                     = "81"

	// Account Information Service Codes
	SERVICE_CODE_BALANCE_INQUIRY = "11"

	// Account Transaction Information Service Codes
	SERVICE_CODE_TRANSACTION_HISTORY_LIST   = "12"
	SERVICE_CODE_TRANSACTION_HISTORY_DETAIL = "13"
	SERVICE_CODE_BANK_STATEMENT             = "14"

	// Credit [AIS Bank] - Account Inquiry Service Codes
	SERVICE_CODE_ACCOUNT_INQUIRY_INTERNAL = "15"
	SERVICE_CODE_ACCOUNT_INQUIRY_EXTERNAL = "16"

	// Credit [AIS Bank] - Intrabank Service Codes
	SERVICE_CODE_TRANSFER_INTRABANK = "17"

	// Credit [AIS Bank] - Interbank Service Codes
	SERVICE_CODE_TRANSFER_INTERBANK             = "18"
	SERVICE_CODE_TRANSFER_INTERBANK_BULK        = "20"
	SERVICE_CODE_TRANSFER_INTERBANK_BULK_NOTIFY = "21"

	// Credit [AIS Bank] - Transfer Service Codes
	SERVICE_CODE_TRANSFER_REQUEST_FOR_PAYMENT = "19"
	SERVICE_CODE_TRANSFER_RTGS                = "22"
	SERVICE_CODE_TRANSFER_RTGS_NOTIFY         = "76"
	SERVICE_CODE_TRANSFER_SKN                 = "23"
	SERVICE_CODE_TRANSFER_SKN_NOTIFY          = "75"

	// Credit [AIS Bank] - Payment VA Related Service Codes
	SERVICE_CODE_TRANSFER_VA_INQUIRY                  = "24"
	SERVICE_CODE_TRANSFER_VA_PAYMENT                  = "25"
	SERVICE_CODE_TRANSFER_VA_STATUS                   = "26"
	SERVICE_CODE_TRANSFER_VA_CREATE                   = "27"
	SERVICE_CODE_TRANSFER_VA_UPDATE                   = "28"
	SERVICE_CODE_TRANSFER_VA_UPDATE_STATUS            = "29"
	SERVICE_CODE_TRANSFER_VA_INQUIRY_VA               = "30"
	SERVICE_CODE_TRANSFER_VA_DELETE                   = "31"
	SERVICE_CODE_TRANSFER_VA_INQUIRY_INTRABANK        = "32"
	SERVICE_CODE_TRANSFER_VA_PAYMENT_INTRABANK        = "33"
	SERVICE_CODE_TRANSFER_VA_NOTIFY_PAYMENT_INTRABANK = "34"
	SERVICE_CODE_TRANSFER_VA_REPORT                   = "35"

	// Transfer Status Service Code
	SERVICE_CODE_TRANSFER_STATUS = "36"

	// Credit [AIS Non Bank] - Customer Topup Service Codes
	SERVICE_CODE_EMONEY_ACCOUNT_INQUIRY = "37"
	SERVICE_CODE_EMONEY_TOPUP           = "38"
	SERVICE_CODE_EMONEY_TOPUP_STATUS    = "39"

	// Credit [AIS Non Bank] - Bulk Cashin Service Codes
	SERVICE_CODE_EMONEY_BULK_CASHIN_PAYMENT = "40"
	SERVICE_CODE_EMONEY_BULK_CASHIN_NOTIFY  = "41"

	// Credit [AIS Non Bank] - Transfer to Bank Service Codes
	SERVICE_CODE_EMONEY_BANK_ACCOUNT_INQUIRY = "42"
	SERVICE_CODE_EMONEY_TRANSFER_BANK        = "43"

	// Credit [AIS Non Bank] - Transfer to OTC Service Codes
	SERVICE_CODE_EMONEY_OTC_CASHOUT = "44"
	SERVICE_CODE_EMONEY_OTC_STATUS  = "45"
	SERVICE_CODE_OTC_CASHOUT_CANCEL = "46"

	// QR-MPM Service Codes
	SERVICE_CODE_QR_MPM_GENERATE = "47"
	SERVICE_CODE_QR_MPM_DECODE   = "48"
	SERVICE_CODE_QR_APPLY_OTT    = "49"
	SERVICE_CODE_QR_MPM_PAYMENT  = "50"
	SERVICE_CODE_QR_MPM_QUERY    = "51"
	SERVICE_CODE_QR_MPM_NOTIFY   = "52"
	SERVICE_CODE_QR_MPM_CANCEL   = "77"
	SERVICE_CODE_QR_MPM_REFUND   = "78"
	SERVICE_CODE_QR_MPM_STATUS   = "53"
)

// Endpoint Constants
const (
	// Authentication Endpoints
	ENDPOINT_ACCESS_TOKEN_B2B   = "/v1.0/access-token/b2b"
	ENDPOINT_ACCESS_TOKEN_B2B2C = "/v1.0/access-token/b2b2c"

	// Registration Endpoints
	ENDPOINT_REGISTRATION_CARD_BIND                  = "/v1.0/registration-card-bind"
	ENDPOINT_REGISTRATION_CARD_INQUIRY               = "/v1.0/registration-card-inquiry"
	ENDPOINT_REGISTRATION_CARD_INQUIRY_CUST_MERCHANT = "/v1.0/registration-card-inquiry/custIdMerchant"
	ENDPOINT_OTP_VERIFICATION                        = "/v1.0/otp-verification"
	ENDPOINT_REGISTRATION_CARD_UNBIND                = "/v1.0/registration-card-unbind"
	ENDPOINT_REGISTRATION_ACCOUNT_CREATION           = "/v1.0/registration-account-creation"
	ENDPOINT_REGISTRATION_ACCOUNT_BINDING            = "/v1.0/registration-account-binding"
	ENDPOINT_REGISTRATION_ACCOUNT_INQUIRY            = "/v1.0/registration-account-inquiry"
	ENDPOINT_REGISTRATION_ACCOUNT_UNBINDING          = "/v1.0/registration-account-unbinding"
	ENDPOINT_GET_AUTH_CODE                           = "/v1.0/get-auth-code"
	ENDPOINT_OTP                                     = "/v1.0/otp"

	// Account Information Endpoints
	ENDPOINT_BALANCE_INQUIRY = "/v1.0/balance-inquiry"

	// Account Transaction Information Endpoints
	ENDPOINT_TRANSACTION_HISTORY_LIST   = "/v1.0/transaction-history-list"
	ENDPOINT_TRANSACTION_HISTORY_DETAIL = "/v1.0/transaction-history-detail"
	ENDPOINT_BANK_STATEMENT             = "/v1.0/bank-statement"

	// Credit [AIS Bank] - Account Inquiry Endpoints
	ENDPOINT_ACCOUNT_INQUIRY_INTERNAL = "/v1.0/account-inquiry-internal"
	ENDPOINT_ACCOUNT_INQUIRY_EXTERNAL = "/v1.0/account-inquiry-external"

	// Credit [AIS Bank] - Intrabank Endpoints
	ENDPOINT_TRANSFER_INTRABANK = "/v1.0/transfer-intrabank"

	// Credit [AIS Bank] - Interbank Endpoints
	ENDPOINT_TRANSFER_INTERBANK             = "/v1.0/transfer-interbank"
	ENDPOINT_TRANSFER_INTERBANK_BULK        = "/v1.0/transfer-interbank-bulk"
	ENDPOINT_TRANSFER_INTERBANK_BULK_NOTIFY = "/v1.0/transfer-interbank-bulk/notify"

	// Credit [AIS Bank] - Transfer Endpoints
	ENDPOINT_TRANSFER_REQUEST_FOR_PAYMENT = "/v1.0/transfer-request-for-payment"
	ENDPOINT_TRANSFER_RTGS                = "/v1.0/transfer-rtgs"
	ENDPOINT_TRANSFER_RTGS_NOTIFY         = "/v1.0/transfer-rtgs/notify"
	ENDPOINT_TRANSFER_SKN                 = "/v1.0/transfer-skn"
	ENDPOINT_TRANSFER_SKN_NOTIFY          = "/v1.0/transfer-skn/notify"

	// Credit [AIS Bank] - Payment VA Related Endpoints
	ENDPOINT_TRANSFER_VA_INQUIRY                  = "/v1.0/transfer-va/inquiry"
	ENDPOINT_TRANSFER_VA_PAYMENT                  = "/v1.0/transfer-va/payment"
	ENDPOINT_TRANSFER_VA_STATUS                   = "/v1.0/transfer-va/status"
	ENDPOINT_TRANSFER_VA_CREATE                   = "/v1.0/transfer-va/create-va"
	ENDPOINT_TRANSFER_VA_UPDATE                   = "/v1.0/transfer-va/update-va"
	ENDPOINT_TRANSFER_VA_UPDATE_STATUS            = "/v1.0/transfer-va/update-status"
	ENDPOINT_TRANSFER_VA_INQUIRY_VA               = "/v1.0/transfer-va/inquiry-va"
	ENDPOINT_TRANSFER_VA_DELETE                   = "/v1.0/transfer-va/delete-va"
	ENDPOINT_TRANSFER_VA_INQUIRY_INTRABANK        = "/v1.0/transfer-va/inquiry-intrabank"
	ENDPOINT_TRANSFER_VA_PAYMENT_INTRABANK        = "/v1.0/transfer-va/payment-intrabank"
	ENDPOINT_TRANSFER_VA_NOTIFY_PAYMENT_INTRABANK = "/v1.0/transfer-va/notify-payment-intrabank"
	ENDPOINT_TRANSFER_VA_REPORT                   = "/v1.0/transfer-va/report"

	// Transfer Status Endpoint
	ENDPOINT_TRANSFER_STATUS = "/v1.0/transfer/status"

	// Credit [AIS Non Bank] - Customer Topup Endpoints
	ENDPOINT_EMONEY_ACCOUNT_INQUIRY = "/v1.0/emoney/account-inquiry"
	ENDPOINT_EMONEY_TOPUP           = "/v1.0/emoney/topup"
	ENDPOINT_EMONEY_TOPUP_STATUS    = "/v1.0/emoney/topup-status"

	// Credit [AIS Non Bank] - Bulk Cashin Endpoints
	ENDPOINT_EMONEY_BULK_CASHIN_PAYMENT = "/v1.0/emoney/bulk-cashin-payment"
	ENDPOINT_EMONEY_BULK_CASHIN_NOTIFY  = "/v1.0/emoney/bulk-cashin-notify"

	// Credit [AIS Non Bank] - Transfer to Bank Endpoints
	ENDPOINT_EMONEY_BANK_ACCOUNT_INQUIRY = "/v1.0/emoney/bank-account-inquiry"
	ENDPOINT_EMONEY_TRANSFER_BANK        = "/v1.0/emoney/transfer-bank"

	// Credit [AIS Non Bank] - Transfer to OTC Endpoints
	ENDPOINT_EMONEY_OTC_CASHOUT = "/v1.0/emoney/otc-cashout"
	ENDPOINT_EMONEY_OTC_STATUS  = "/v1.0/emoney/otc-status"
	ENDPOINT_OTC_CASHOUT_CANCEL = "/v1.0/otc/cashout/cancel"

	// QR-MPM Endpoints
	ENDPOINT_QR_MPM_GENERATE = "/v1.0/qr/qr-mpm-generate"
	ENDPOINT_QR_MPM_DECODE   = "/v1.0/qr/qr-mpm-decode"
	ENDPOINT_QR_APPLY_OTT    = "/v1.0/qr/apply-ott"
	ENDPOINT_QR_MPM_PAYMENT  = "/v1.0/qr/qr-mpm-payment"
	ENDPOINT_QR_MPM_QUERY    = "/v1.0/qr/qr-mpm-query"
	ENDPOINT_QR_MPM_NOTIFY   = "/v1.0/qr/qr-mpm-notify"
	ENDPOINT_QR_MPM_CANCEL   = "/v1.0/qr/qr-mpm-cancel"
	ENDPOINT_QR_MPM_REFUND   = "/v1.0/qr/qr-mpm-refund"
	ENDPOINT_QR_MPM_STATUS   = "/v1.0/qr/qr-mpm-status"
)
