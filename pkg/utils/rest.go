package utils

import "fmt"

func IsHTTPVerbValid(ver string) error {
	switch ver {
	case "GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS", "TRACE", "CONNECT":
		return nil
	default:
		return fmt.Errorf("invalid HTTP verb: %s", ver)
	}
}
