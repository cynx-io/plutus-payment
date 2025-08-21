package helper

import "fmt"

func FormatXenditExternalId(customerId int32, requestId string) string {
	return fmt.Sprintf("%d-%s", customerId, requestId)
}
