package helper

import (
	proto "github.com/cynx-io/plutus-payment/api/proto/gen/plutus"
	"strings"
)

func ToProtoInvoiceStatus(status string) proto.PaymentInvoiceStatus {

	upperString := strings.ToUpper(status)

	switch upperString {
	case "PENDING":
		return proto.PaymentInvoiceStatus_PENDING
	case "FAILED":
		return proto.PaymentInvoiceStatus_FAILED
	case "COMPLETED":
		return proto.PaymentInvoiceStatus_COMPLETED
	}

	return proto.PaymentInvoiceStatus_FAILED
}
