package snap

import "net/http"

const (
	E_DUPLICATE                          = "duplicate"
	E_TRX_NOT_FOUND                      = "Transaction Not Found"
	E_NOT_FOUND                          = "not_found"
	E_UNPROCESSABLE_ENTITY               = "unprocessable_entity"
	E_UNAUTHORIZED                       = "Unauthorized"
	E_BAD_REQUEST                        = "Bad Request"
	E_SERVER_ERROR                       = "Internal Server Error"
	E_EXTERNAL_SERVER_ERROR              = "External Server Error"
	E_INVALID_FORMAT                     = "Invalid Field Format"
	E_MANDATORY                          = "Invalid Mandatory Field"
	E_SIGANTURE_INVALID                  = "Signature is invalid"
	E_INVALID_TOKEN                      = "Invalid Token"
	E_INVALID_ACCESS_TOKEN               = "Access Token Invalid"
	E_INVALID_ROUTING                    = "Invalid Routing"
	E_GRANT_TYPE                         = "grant_type must be set to client_credentials"
	E_INVALID_CUSTOMER_TOKEN             = "Invalid Customer Token"
	E_NEED_REQUEST_OTP                   = "Need To Request OTP"
	E_OTP_SENT_TO_CARDHOLDER             = "OTP Sent To Cardholder"
	E_INVALID_OTP                        = "OTP/Verification Code Invalid"
	E_PARSE_SIGNATURE                    = "x509: failed to parse public key (use ParsePKIXPublicKey instead for this key format)"
	E_INVALID_PRODUCT                    = "invalid product id"
	E_BILL_NOT_FOUND                     = "Invalid Bill"
	E_FORBIDDEN                          = "Forbidden"
	E_TERJADI_KESALAHAN                  = "Terjadi Kesalahan"
	E_GENERAL_ERROR                      = "General Error"
	E_INVALID_VA                         = "Invalid Virtual Account"
	E_VA_NOT_FOUND                       = "Invalid Bill / Virtual Account Not Found"
	E_STATE_NOT_WHITELISTED              = "user unable to topup because of state"
	E_PARSING_VALUE                      = "value parsing error"
	E_INVALID_AMOUNT                     = "Invalid Amount"
	E_MONTHLY_IN_LIMIT_EXCEEDED          = "request amount exceed monthly in limit"
	E_BALANCE_LIMIT_REACHED              = "balance limit reached"
	E_TRX_LIMIT_REACHED                  = "Exceeds Transaction Amount Limit"
	E_ACCOUNT_LIMIT_REACHED              = "Account Limit Exceed"
	E_TRX_EXPIRED                        = "Transaction Expired"
	E_CONFLICT                           = "Conflict"
	E_PAID_BILL                          = "Paid Bill"
	E_BILL_HAS_BEEN_PAID                 = "bill has been paid"
	E_PARSING                            = "Parsing Error"
	E_ROWS_ACC_TRX_NOT_UPDATED           = "rows on account_transactions is not updated"
	E_CLOSING_VALIDATION_FAILED          = "unable to close account"
	E_ACCOUNT_TEMPORARILY_BLOCKED        = "Account Temporarily Blocked"
	E_ACCOUNT_FROZEN                     = "Account Frozen/Abnormal"
	E_INVALID_PIN                        = "Invalid PIN"
	E_FAILED_BINDING                     = "Failed Binding"
	E_INSUFFICIENT_FUNDS                 = "Insufficient Funds"
	E_LESS_THAN_MINIMUM_AMOUNT           = "Less Than Minimum Amount"
	E_AVAILABLE_LEDGER_BALANCE_MISMATCH  = "Available balance does not match ledger balance"
	E_INACTIVE_CARD_ACCOUNT_CUSTOMER     = "Inactive Card/Account/Customer"
	E_INACTIVE_ACCOUNT                   = "Inactive Account"
	E_DORMANT_ACCOUNT                    = "The account is dormant"
	E_INVALID_ACCOUNT_INFO               = "Account Information Invalid"
	E_DO_HONOR                           = "Do Not Honor"
	E_WALLET_NOT_FOUND                   = "Wallet Not Found"
	E_TRANSACTION_NOT_PERMITTED          = "Transaction No Permitted"
	E_TRX_AMOUNT_LIMIT                   = "Exceeds Amount Limit"
	E_INVALID_TRX_STATUS                 = "Invalid Transaction Status"
	E_INVALID_CARD_ACCOUNT_CUSTOMER      = "Invalid Card/Account/Customer [info]/Virtual Account"
	E_BILLING_NOT_FOUND                  = "Billing Not Found"
	E_NO_ROWS_AFFECTED                   = "No Rows Affected"
	E_INCONSISTENT_REQUEST               = "Inconsistent Request"
	E_ACCOUNT_ALREADY_EXIST              = "Account Already Exist"
	E_REQUEST_TYPE_DOES_NOT_MATCH        = "Request Type does not match"
	E_INVALID_TRANSACTION_CODE           = "Transaction code not recognized"
	E_INVALID_SUB_CATEGORY               = "Invalid subcategory"
	E_INVALID_GENERAL_TRANSACTION_ACTION = "General transaction action not recognized"
)

var ErrorToHTTPCode = map[string]int{
	E_DUPLICATE:                          http.StatusConflict,            // Conflict
	E_TRX_NOT_FOUND:                      http.StatusNotFound,            // Not Found
	E_NOT_FOUND:                          http.StatusNotFound,            // Not Found
	E_UNPROCESSABLE_ENTITY:               http.StatusUnprocessableEntity, // Unprocessable Entity
	E_UNAUTHORIZED:                       http.StatusUnauthorized,        // Unauthorized
	E_BAD_REQUEST:                        http.StatusBadRequest,          // Bad Request
	E_SERVER_ERROR:                       http.StatusInternalServerError, // Internal Server Error
	E_EXTERNAL_SERVER_ERROR:              http.StatusBadGateway,          // Bad Gateway (External Service Issue)
	E_INVALID_FORMAT:                     http.StatusBadRequest,          // Bad Request (Invalid Field Format)
	E_MANDATORY:                          http.StatusBadRequest,          // Bad Request (Missing Mandatory Field)
	E_SIGANTURE_INVALID:                  http.StatusUnauthorized,        // Unauthorized (Invalid Signature)
	E_INVALID_TOKEN:                      http.StatusUnauthorized,        // Unauthorized
	E_INVALID_ACCESS_TOKEN:               http.StatusUnauthorized,        // Unauthorized
	E_INVALID_ROUTING:                    http.StatusBadRequest,          // Bad Request
	E_GRANT_TYPE:                         http.StatusBadRequest,          // Bad Request
	E_INVALID_CUSTOMER_TOKEN:             http.StatusUnauthorized,        // Unauthorized
	E_NEED_REQUEST_OTP:                   http.StatusForbidden,           // Forbidden (Action Required)
	E_OTP_SENT_TO_CARDHOLDER:             http.StatusOK,                  // OK (Informational)
	E_INVALID_OTP:                        http.StatusBadRequest,          // Bad Request (Invalid OTP)
	E_PARSE_SIGNATURE:                    http.StatusBadRequest,          // Bad Request (Invalid Key Format)
	E_INVALID_PRODUCT:                    http.StatusBadRequest,          // Bad Request
	E_BILL_NOT_FOUND:                     http.StatusNotFound,            // Not Found
	E_FORBIDDEN:                          http.StatusForbidden,           // Forbidden
	E_TERJADI_KESALAHAN:                  http.StatusInternalServerError, // Internal Server Error
	E_GENERAL_ERROR:                      http.StatusInternalServerError, // Internal Server Error
	E_INVALID_VA:                         http.StatusBadRequest,          // Bad Request
	E_VA_NOT_FOUND:                       http.StatusNotFound,            // Not Found
	E_STATE_NOT_WHITELISTED:              http.StatusForbidden,           // Forbidden
	E_PARSING_VALUE:                      http.StatusBadRequest,          // Bad Request (Parsing Error)
	E_INVALID_AMOUNT:                     http.StatusBadRequest,          // Bad Request
	E_MONTHLY_IN_LIMIT_EXCEEDED:          http.StatusForbidden,           // Forbidden (Limit Exceeded)
	E_BALANCE_LIMIT_REACHED:              http.StatusForbidden,           // Forbidden (Balance Limit)
	E_TRX_LIMIT_REACHED:                  http.StatusForbidden,           // Forbidden (Transaction Limit)
	E_ACCOUNT_LIMIT_REACHED:              http.StatusForbidden,           // Forbidden (Account Limit)
	E_TRX_EXPIRED:                        http.StatusBadRequest,          // Bad Request (Expired Transaction)
	E_CONFLICT:                           http.StatusConflict,            // Conflict
	E_PAID_BILL:                          http.StatusConflict,            // Conflict (Bill Already Paid)
	E_BILL_HAS_BEEN_PAID:                 http.StatusConflict,            // Conflict
	E_PARSING:                            http.StatusBadRequest,          // Bad Request (Parsing Error)
	E_ROWS_ACC_TRX_NOT_UPDATED:           http.StatusInternalServerError, // Internal Server Error
	E_CLOSING_VALIDATION_FAILED:          http.StatusBadRequest,          // Bad Request
	E_ACCOUNT_TEMPORARILY_BLOCKED:        http.StatusForbidden,           // Forbidden
	E_ACCOUNT_FROZEN:                     http.StatusForbidden,           // Forbidden
	E_INVALID_PIN:                        http.StatusUnauthorized,        // Unauthorized
	E_FAILED_BINDING:                     http.StatusBadRequest,          // Bad Request
	E_INSUFFICIENT_FUNDS:                 http.StatusPaymentRequired,     // Payment Required
	E_LESS_THAN_MINIMUM_AMOUNT:           http.StatusBadRequest,          // Bad Request
	E_AVAILABLE_LEDGER_BALANCE_MISMATCH:  http.StatusInternalServerError, // Internal Server Error
	E_INACTIVE_CARD_ACCOUNT_CUSTOMER:     http.StatusForbidden,           // Forbidden
	E_INACTIVE_ACCOUNT:                   http.StatusForbidden,           // Forbidden
	E_DORMANT_ACCOUNT:                    http.StatusForbidden,           // Forbidden
	E_INVALID_ACCOUNT_INFO:               http.StatusBadRequest,          // Bad Request
	E_DO_HONOR:                           http.StatusForbidden,           // Forbidden
	E_WALLET_NOT_FOUND:                   http.StatusNotFound,            // Not Found
	E_TRANSACTION_NOT_PERMITTED:          http.StatusForbidden,           // Forbidden
	E_TRX_AMOUNT_LIMIT:                   http.StatusForbidden,           // Forbidden
	E_INVALID_TRX_STATUS:                 http.StatusBadRequest,          // Bad Request
	E_INVALID_CARD_ACCOUNT_CUSTOMER:      http.StatusBadRequest,          // Bad Request
	E_BILLING_NOT_FOUND:                  http.StatusNotFound,            // Not Found
	E_NO_ROWS_AFFECTED:                   http.StatusInternalServerError, // Internal Server Error
	E_INCONSISTENT_REQUEST:               http.StatusBadRequest,          // Bad Request
	E_ACCOUNT_ALREADY_EXIST:              http.StatusConflict,            // Conflict
	E_REQUEST_TYPE_DOES_NOT_MATCH:        http.StatusBadRequest,          // Bad Request
	E_INVALID_TRANSACTION_CODE:           http.StatusBadRequest,          // Bad Request
	E_INVALID_SUB_CATEGORY:               http.StatusBadRequest,          // Bad Request
	E_INVALID_GENERAL_TRANSACTION_ACTION: http.StatusBadRequest,          // Bad Request
}
