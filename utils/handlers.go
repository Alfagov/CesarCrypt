package utils

import "fmt"

func GenerateStoreBackendName(handler string, vault string) string {
	return fmt.Sprintf("%s-%s", handler, vault)
}
