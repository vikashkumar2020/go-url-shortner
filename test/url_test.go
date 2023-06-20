// +build e2e

package test

import(
	"testing"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestCreateShortUrlEndPoint(t *testing.T) {
	fmt.Println("Running E2E test for create url endpoint")

	client := resty.New()
	resp, err := client.R().
	SetBody(`{
		"long_url":"https://towardsdev.com/"
	}`).
	Post("http://localhost:8080/api/v1/create");

	if err!=nil {
		t.Fail()
	}

	assert.Equal(t, 201, resp.StatusCode())

}

func TestGetShortUrlEndPoint(t *testing.T) {
	fmt.Println("Running E2E test for create url endpoint")

	client := resty.New()
	resp, err := client.R().Get("http://localhost:8080/api/v1/1");

	if err!=nil {
		t.Fail()
	}

	assert.Equal(t, 200, resp.StatusCode())

}