// +build e2e

package test

import(
	"testing"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)
	

func TestHealthEndPoint(t *testing.T) {
	fmt.Println("Running E2E test for health check endpoint")

	client := resty.New()
	resp, err := client.R().Get("http://localhost:8080/health");

	if err!=nil {
		t.Fail()
	}

	assert.Equal(t, 200, resp.StatusCode())

}