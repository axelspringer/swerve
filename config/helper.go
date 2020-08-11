package config

import (
	"os"
	"strings"
)

func getPrefixedOSEnv(s string) string {
	if e := strings.TrimSpace(
		os.Getenv(envVarPrefix + s)); len(e) > 0 {
		return e
	}

	return ""
}

func hasDuplicates(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr) && j != i; j++ {
			if arr[i] == arr[j] {
				return true
			}
		}
	}
	return false
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
