package response

import "github.com/cynx-io/cynx-core/src/response"

const (
	// Expected Error
	CodeSuccess            response.Code = "00"
	codeValidationError    response.Code = "VE"
	codeUnauthorized       response.Code = "UA"
	codeNotAllowed         response.Code = "NA"
	codeNotFound           response.Code = "NF"
	codeInvalidCredentials response.Code = "IC"

	// Internal Errors
	codeInternalError response.Code = "I-IE"

	// External Errors
	codeXenditError response.Code = "E-XN"

	// Microservice Errors
	codeHermesError response.Code = "M-HR"

	// Database Errors
	codeDatabaseCustomerError response.Code = "D-CR"
	codeDatabaseInvoiceError  response.Code = "D-IV"
)

var responseCodeNames = map[response.Code]string{
	// Expected Error
	CodeSuccess:            "Success",
	codeValidationError:    "Validation Error",
	codeUnauthorized:       "Not Authorized",
	codeNotAllowed:         "Not Allowed",
	codeNotFound:           "Not Found",
	codeInvalidCredentials: "Invalid Credentials",

	// Internal
	codeInternalError: "Internal Error",

	// External Errors
	codeXenditError: "Xendit Error",

	// Microservice Errors
	codeHermesError: "Hermes Error",

	// Database Errors
	codeDatabaseCustomerError: "Database Customer Error",
	codeDatabaseInvoiceError:  "Database Invoice Error",
}
