package response

import "github.com/cynx-io/cynx-core/src/response"

const (
	// Expected Error
	CodeSuccess             response.Code = "00"
	codeValidationError     response.Code = "VE"
	codeUnauthorized        response.Code = "UA"
	codeNotAllowed          response.Code = "NA"
	codeNotFound            response.Code = "NF"
	codeInvalidCredentials  response.Code = "IC"
	codeInsufficientBalance response.Code = "IB"

	// Internal Errors
	codeInternalError response.Code = "I-IE"

	// External Errors
	codeXenditError response.Code = "E-XN"

	// Microservice Errors
	codeHermesError response.Code = "M-HR"
	codeAnankeError response.Code = "M-AN"

	// Database Errors
	codeDatabaseCustomerError         response.Code = "D-CR"
	codeDatabaseInvoiceError          response.Code = "D-IV"
	codeDatabaseBalanceError          response.Code = "D-BL"
	codeDatabaseProductPriceListError response.Code = "D-PL"
	codeDatabaseTokenPriceListError   response.Code = "D-TL"
	codeDatabaseTokenInvoiceError     response.Code = "D-TI"
)

var responseCodeNames = map[response.Code]string{
	// Expected Error
	CodeSuccess:             "Success",
	codeValidationError:     "Validation Error",
	codeUnauthorized:        "Not Authorized",
	codeNotAllowed:          "Not Allowed",
	codeNotFound:            "Not Found",
	codeInvalidCredentials:  "Invalid Credentials",
	codeInsufficientBalance: "Insufficient Balance",

	// Internal
	codeInternalError: "Internal Error",

	// External Errors
	codeXenditError: "Xendit Error",

	// Microservice Errors
	codeHermesError: "Hermes Error",
	codeAnankeError: "Ananke Error",

	// Database Errors
	codeDatabaseCustomerError:         "Database Customer Error",
	codeDatabaseInvoiceError:          "Database Invoice Error",
	codeDatabaseBalanceError:          "Database Balance Error",
	codeDatabaseProductPriceListError: "Database Product Price List Error",
	codeDatabaseTokenPriceListError:   "Database Token Price List Error",
	codeDatabaseTokenInvoiceError:     "Database Token Invoice Error",
}
