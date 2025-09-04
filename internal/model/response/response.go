package response

import (
	"github.com/cynx-io/cynx-core/src/response"
)

func setResponse[Resp response.Generic](resp Resp, code response.Code) {
	resp.GetBase().Code = code.String()
	resp.GetBase().Desc = responseCodeNames[code]
}

func Success[Resp response.Generic](resp Resp) {
	setResponse(resp, CodeSuccess)
}

func ErrorValidation[Resp response.Generic](resp Resp) {
	setResponse(resp, codeValidationError)
}

func ErrorUnauthorized[Resp response.Generic](resp Resp) {
	setResponse(resp, codeUnauthorized)
}

func ErrorNotAllowed[Resp response.Generic](resp Resp) {
	setResponse(resp, codeNotAllowed)
}

func ErrorNotFound[Resp response.Generic](resp Resp) {
	setResponse(resp, codeNotFound)
}

func ErrorInvalidCredentials[Resp response.Generic](resp Resp) {
	setResponse(resp, codeInvalidCredentials)
}

func ErrorInternal[Resp response.Generic](resp Resp) {
	setResponse(resp, codeInternalError)
}

func ErrorHermes[Resp response.Generic](resp Resp) {
	setResponse(resp, codeHermesError)
}

func ErrorDatabaseCustomer[Resp response.Generic](resp Resp) {
	setResponse(resp, codeDatabaseCustomerError)
}

func ErrorDatabaseInvoice[Resp response.Generic](resp Resp) {
	setResponse(resp, codeDatabaseInvoiceError)
}

func ErrorXendit[Resp response.Generic](resp Resp) {
	setResponse(resp, codeXenditError)
}

func ErrorAnanke[Resp response.Generic](resp Resp) {
	setResponse(resp, codeAnankeError)
}

func ErrorDatabaseBalance[Resp response.Generic](resp Resp) {
	setResponse(resp, codeDatabaseBalanceError)
}

func ErrorDatabaseProductPriceList[Resp response.Generic](resp Resp) {
	setResponse(resp, codeDatabaseProductPriceListError)
}

func ErrorDatabaseTokenPriceList[Resp response.Generic](resp Resp) {
	setResponse(resp, codeDatabaseTokenPriceListError)
}

func ErrorInsufficientBalance[Resp response.Generic](resp Resp) {
	setResponse(resp, codeInsufficientBalance)
}

func ErrorDatabaseTokenInvoice[Resp response.Generic](resp Resp) {
	setResponse(resp, codeDatabaseTokenInvoiceError)
}
