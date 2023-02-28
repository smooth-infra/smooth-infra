package http

import (
	"crypto/tls"
	"fmt"
	"strings"
	"testing"
	"time"

	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
)

func Request(t *testing.T, params map[string]interface{}, expects map[string]interface{}) (err error) {
	tlsConfig := tls.Config{}

	address := fmt.Sprint(params["address"])
	expectedStatusCode := expects["status_code"].(int)
	for _, val := range []string{"https://", "http://"} {
		address = strings.ReplaceAll(address, val, "")
	}
	if val, ok := params["secure"]; ok {
		if val == true {
			address = fmt.Sprintf("https://%s", address)
		} else {
			address = fmt.Sprintf("http://%s", address)
		}
	}

	err = http_helper.HttpGetWithRetryWithCustomValidationE(
		t,
		address,
		&tlsConfig,
		2,
		3*time.Second,
		func(statusCode int, body string) bool {
			return statusCode == expectedStatusCode
		},
	)
	return err
}
