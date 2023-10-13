package utils

import (
	"fmt"
	"strings"
)

func ExtractBearerValueFromHeader(headerValue string) (string, error) {
	onlyToken := strings.Split(headerValue, " ")
	if len(onlyToken) != 2 {
		return "", fmt.Errorf("failed to extract token from header, header is invalid")
	}

	if onlyToken[0] != "Bearer" {
		return "", fmt.Errorf("failed to extract token from header, header is invalid")
	}

	if onlyToken[1] == "" {
		return "", fmt.Errorf("failed to extract token from header, header is invalid")
	}

	return onlyToken[1], nil
}

func ValidateAuthorizationHeader(headers map[string]string) error {
	if len(headers) == 0 {
		return fmt.Errorf("failed to obtain token from header, header is empty")
	}

	if _, err := FindStrInMapKeys(headers, "Authorization"); err != nil {
		return fmt.Errorf("failed to validate Authorization header, "+
			"header's key is invalid. Error: %s", err.Error())
	}

	if err := KeyInMapHasValueSet(headers, "Authorization"); err != nil {
		return fmt.Errorf("failed to validate Authorization header, "+
			"header's value is empty. Error: %s", err.Error())
	}

	return nil
}

func ValidateBearerTokenSchema(headerContent string) error {
	if headerContent == "" {
		return fmt.Errorf("failed to validate bearer token schema, headerContent is empty")
	}

	if !strings.Contains(headerContent, "Bearer") {
		return fmt.Errorf("failed to validate bearer token schema, headerContent is invalid")
	}

	bearerSchema := strings.Split(headerContent, " ")
	if len(bearerSchema) != 2 {
		return fmt.Errorf("failed to validate bearer token schema, malformed headerContent")
	}

	bearerValue := bearerSchema[1]
	if bearerValue == "" {
		return fmt.Errorf("failed to validate bearer token schema, the 'bearerValue' is empty")
	}

	return nil
}
