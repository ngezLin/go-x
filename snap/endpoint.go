package snap

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
	SERVICE_CODE_TRANSFER_VA_DELETE_VA                = "31"
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

	// Debit [AIS Non Bank] - Debit Payment Service Codes
	SERVICE_CODE_DEBIT_PAYMENT_HOST_TO_HOST = "54"
	SERVICE_CODE_DEBIT_STATUS               = "55"
	SERVICE_CODE_DEBIT_NOTIFY               = "56"
	SERVICE_CODE_DEBIT_CANCEL               = "57"
	SERVICE_CODE_DEBIT_REFUND               = "58"

	// Debit QR-CPM Service Codes
	SERVICE_CODE_QR_CPM_GENERATE = "59"
	SERVICE_CODE_QR_CPM_PAYMENT  = "60"
	SERVICE_CODE_QR_CPM_QUERY    = "61"
	SERVICE_CODE_QR_CPM_CANCEL   = "62"
	SERVICE_CODE_QR_CPM_NOTIFY   = "79"
	SERVICE_CODE_QR_CPM_REFUND   = "80"

	// Debit Auth Payment Service Codes
	SERVICE_CODE_DEBIT_AUTH_PAYMENT        = "63"
	SERVICE_CODE_DEBIT_AUTH_QUERY          = "64"
	SERVICE_CODE_DEBIT_AUTH_CAPTURE        = "65"
	SERVICE_CODE_DEBIT_AUTH_CAPTURE_INQURY = "66"
	SERVICE_CODE_DEBIT_AUTH_VOID           = "67"
	SERVICE_CODE_DEBIT_AUTH_VOID_INQURY    = "68"
	SERVICE_CODE_DEBIT_AUTH_REFUND         = "69"
	// Debit BI-FAST Service Codes
	SERVICE_CODE_DEBIT_BI_FAST_EMANDATE = "70"
	SERVICE_CODE_DEBIT_BI_FAST_PAYMENT  = "71"
	SERVICE_CODE_DEBIT_BI_FAST_NOTIFY   = "72"
)

// Endpoint Constants
const (
	// Authentication Endpoints
	ENDPOINT_ACCESS_TOKEN_B2B   = "/access-token/b2b"
	ENDPOINT_ACCESS_TOKEN_B2B2C = "/access-token/b2b2c"

	// Registration Endpoints
	ENDPOINT_REGISTRATION_CARD_BIND                  = "/registration-card-bind"
	ENDPOINT_REGISTRATION_CARD_INQUIRY               = "/registration-card-inquiry"
	ENDPOINT_REGISTRATION_CARD_INQUIRY_CUST_MERCHANT = "/registration-card-inquiry/custIdMerchant"
	ENDPOINT_OTP_VERIFICATION                        = "/otp-verification"
	ENDPOINT_REGISTRATION_CARD_UNBIND                = "/registration-card-unbind"
	ENDPOINT_REGISTRATION_ACCOUNT_CREATION           = "/registration-account-creation"
	ENDPOINT_REGISTRATION_ACCOUNT_BINDING            = "/registration-account-binding"
	ENDPOINT_REGISTRATION_ACCOUNT_INQUIRY            = "/registration-account-inquiry"
	ENDPOINT_REGISTRATION_ACCOUNT_UNBINDING          = "/registration-account-unbinding"
	ENDPOINT_GET_AUTH_CODE                           = "/get-auth-code"
	ENDPOINT_OTP                                     = "/otp"

	// Account Information Endpoints
	ENDPOINT_BALANCE_INQUIRY = "/balance-inquiry"

	// Account Transaction Information Endpoints
	ENDPOINT_TRANSACTION_HISTORY_LIST   = "/transaction-history-list"
	ENDPOINT_TRANSACTION_HISTORY_DETAIL = "/transaction-history-detail"
	ENDPOINT_BANK_STATEMENT             = "/bank-statement"

	// Credit [AIS Bank] - Account Inquiry Endpoints
	ENDPOINT_ACCOUNT_INQUIRY_INTERNAL = "/account-inquiry-internal"
	ENDPOINT_ACCOUNT_INQUIRY_EXTERNAL = "/account-inquiry-external"

	// Credit [AIS Bank] - Intrabank Endpoints
	ENDPOINT_TRANSFER_INTRABANK = "/transfer-intrabank"

	// Credit [AIS Bank] - Interbank Endpoints
	ENDPOINT_TRANSFER_INTERBANK             = "/transfer-interbank"
	ENDPOINT_TRANSFER_INTERBANK_BULK        = "/transfer-interbank-bulk"
	ENDPOINT_TRANSFER_INTERBANK_BULK_NOTIFY = "/transfer-interbank-bulk/notify"

	// Credit [AIS Bank] - Transfer Endpoints
	ENDPOINT_TRANSFER_REQUEST_FOR_PAYMENT = "/transfer-request-for-payment"
	ENDPOINT_TRANSFER_RTGS                = "/transfer-rtgs"
	ENDPOINT_TRANSFER_RTGS_NOTIFY         = "/transfer-rtgs/notify"
	ENDPOINT_TRANSFER_SKN                 = "/transfer-skn"
	ENDPOINT_TRANSFER_SKN_NOTIFY          = "/transfer-skn/notify"

	// Credit [AIS Bank] - Payment VA Related Endpoints
	ENDPOINT_TRANSFER_VA_INQUIRY                  = "/transfer-va/inquiry"
	ENDPOINT_TRANSFER_VA_PAYMENT                  = "/transfer-va/payment"
	ENDPOINT_TRANSFER_VA_STATUS                   = "/transfer-va/status"
	ENDPOINT_TRANSFER_VA_CREATE_VA                = "/transfer-va/create-va"
	ENDPOINT_TRANSFER_VA_UPDATE_VA                = "/transfer-va/update-va"
	ENDPOINT_TRANSFER_VA_UPDATE_STATUS            = "/transfer-va/update-status"
	ENDPOINT_TRANSFER_VA_INQUIRY_VA               = "/transfer-va/inquiry-va"
	ENDPOINT_TRANSFER_VA_DELETE_VA                = "/transfer-va/delete-va"
	ENDPOINT_TRANSFER_VA_INQUIRY_INTRABANK        = "/transfer-va/inquiry-intrabank"
	ENDPOINT_TRANSFER_VA_PAYMENT_INTRABANK        = "/transfer-va/payment-intrabank"
	ENDPOINT_TRANSFER_VA_NOTIFY_PAYMENT_INTRABANK = "/transfer-va/notify-payment-intrabank"
	ENDPOINT_TRANSFER_VA_REPORT                   = "/transfer-va/report"

	// Transfer Status Endpoint
	ENDPOINT_TRANSFER_STATUS = "/transfer/status"

	// Credit [AIS Non Bank] - Customer Topup Endpoints
	ENDPOINT_EMONEY_ACCOUNT_INQUIRY = "/emoney/account-inquiry"
	ENDPOINT_EMONEY_TOPUP           = "/emoney/topup"
	ENDPOINT_EMONEY_TOPUP_STATUS    = "/emoney/topup-status"

	// Credit [AIS Non Bank] - Bulk Cashin Endpoints
	ENDPOINT_EMONEY_BULK_CASHIN_PAYMENT = "/emoney/bulk-cashin-payment"
	ENDPOINT_EMONEY_BULK_CASHIN_NOTIFY  = "/emoney/bulk-cashin-notify"

	// Credit [AIS Non Bank] - Transfer to Bank Endpoints
	ENDPOINT_EMONEY_BANK_ACCOUNT_INQUIRY = "/emoney/bank-account-inquiry"
	ENDPOINT_EMONEY_TRANSFER_BANK        = "/emoney/transfer-bank"

	// Credit [AIS Non Bank] - Transfer to OTC Endpoints
	ENDPOINT_EMONEY_OTC_CASHOUT = "/emoney/otc-cashout"
	ENDPOINT_EMONEY_OTC_STATUS  = "/emoney/otc-status"
	ENDPOINT_OTC_CASHOUT_CANCEL = "/otc/cashout/cancel"

	// QR-MPM Endpoints
	ENDPOINT_QR_MPM_GENERATE = "/qr/qr-mpm-generate"
	ENDPOINT_QR_MPM_DECODE   = "/qr/qr-mpm-decode"
	ENDPOINT_QR_APPLY_OTT    = "/qr/apply-ott"
	ENDPOINT_QR_MPM_PAYMENT  = "/qr/qr-mpm-payment"
	ENDPOINT_QR_MPM_QUERY    = "/qr/qr-mpm-query"
	ENDPOINT_QR_MPM_NOTIFY   = "/qr/qr-mpm-notify"
	ENDPOINT_QR_MPM_CANCEL   = "/qr/qr-mpm-cancel"
	ENDPOINT_QR_MPM_REFUND   = "/qr/qr-mpm-refund"
	ENDPOINT_QR_MPM_STATUS   = "/qr/qr-mpm-status"

	// Debit Payment Endpoints
	ENDPOINT_DEBIT_PAYMENT_HOST_TO_HOST = "/debit/payment-host-to-host"
	ENDPOINT_DEBIT_STATUS               = "/debit/status"
	ENDPOINT_DEBIT_NOTIFY               = "/debit/notify"
	ENDPOINT_DEBIT_CANCEL               = "/debit/cancel"
	ENDPOINT_DEBIT_REFUND               = "/debit/refund"

	// Debit QR-CPM Endpoints
	ENDPOINT_QR_CPM_GENERATE = "/qr/qr-cpm-generate"
	ENDPOINT_QR_CPM_PAYMENT  = "/qr/qr-cpm-payment"
	ENDPOINT_QR_CPM_QUERY    = "/qr/qr-cpm-query"
	ENDPOINT_QR_CPM_CANCEL   = "/qr/qr-cpm-cancel"
	ENDPOINT_QR_CPM_NOTIFY   = "/qr/qr-cpm-notify"
	ENDPOINT_QR_CPM_REFUND   = "/qr/qr-cpm-refund"

	// Debit Auth Payment Endpoints
	ENDPOINT_DEBIT_AUTH_PAYMENT        = "/auth/payment"
	ENDPOINT_DEBIT_AUTH_QUERY          = "/auth/query"
	ENDPOINT_DEBIT_AUTH_CAPTURE        = "/auth/capture"
	ENDPOINT_DEBIT_AUTH_CAPTURE_INQURY = "/auth/capture-query"
	ENDPOINT_DEBIT_AUTH_VOID           = "/auth/void"
	ENDPOINT_DEBIT_AUTH_VOID_INQURY    = "/auth/void-query"
	ENDPOINT_DEBIT_AUTH_REFUND         = "/auth/refund"

	// Debit BI-FAST Endpoints
	ENDPOINT_DEBIT_BI_FAST_EMANDATE = "/debit/fast-emandate"
	ENDPOINT_DEBIT_BI_FAST_PAYMENT  = "/debit/fast-payment"
	ENDPOINT_DEBIT_BI_FAST_NOTIFY   = "/debit/fast-notify"
)
