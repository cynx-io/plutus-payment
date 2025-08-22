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

func XenditInvoiceStatusToProto(status string) proto.PaymentInvoiceStatus {
	switch status {
	case "PENDING":
		return proto.PaymentInvoiceStatus_PENDING
	case "EXPIRED":
		return proto.PaymentInvoiceStatus_FAILED
	case "PAID":
		return proto.PaymentInvoiceStatus_COMPLETED
	case "FAILED":
		return proto.PaymentInvoiceStatus_FAILED
	}
	return proto.PaymentInvoiceStatus_FAILED
}
