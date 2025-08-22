package helper

import "fmt"

func FormatXenditExternalId(customerId int32, requestId string, feature int32) string {
	return fmt.Sprintf("%d-%d-%s", feature, customerId, requestId)
}
