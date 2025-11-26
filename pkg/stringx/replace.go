package stringx

import (
	"fmt"
	"strings"
)

func ReplaceVariables(content string, variables map[string]string) string {
	result := content
	for key, value := range variables {
		result = strings.ReplaceAll(result, fmt.Sprintf("${%s}", key), value)
	}
	return result
}
