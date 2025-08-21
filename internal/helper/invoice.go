package helper

import (
	proto "github.com/cynx-io/plutus-payment/api/proto/gen/plutus"
	"strings"
)

func ToProtoInvoiceStatus(status string) proto.PaymentInvoiceStatus {

	upperString := strings.ToUpper(status)

	switch upperString {
	case "PENDING":
		return proto.PaymentInvoiceStatus_PAYMENT_INVOICE_STATUS_PENDING
	case "FAILED":
		return proto.PaymentInvoiceStatus_PAYMENT_INVOICE_STATUS_FAILED
	case "COMPLETED":
		return proto.PaymentInvoiceStatus_PAYMENT_INVOICE_STATUS_COMPLETED
	}

	return proto.PaymentInvoiceStatus_PAYMENT_INVOICE_STATUS_FAILED
}
