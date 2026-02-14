// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package casparser_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/cas-parser-go"
	"github.com/stainless-sdks/cas-parser-go/internal/testutil"
	"github.com/stainless-sdks/cas-parser-go/option"
)

func TestCdslFetchRequestOtp(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := casparser.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Cdsl.Fetch.RequestOtp(context.TODO(), casparser.CdslFetchRequestOtpParams{
		BoID: "1234567890123456",
		Dob:  "1990-01-15",
		Pan:  "ABCDE1234F",
	})
	if err != nil {
		var apierr *casparser.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCdslFetchVerifyOtpWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := casparser.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Cdsl.Fetch.VerifyOtp(
		context.TODO(),
		"session_id",
		casparser.CdslFetchVerifyOtpParams{
			Otp:        "123456",
			NumPeriods: casparser.Int(6),
		},
	)
	if err != nil {
		var apierr *casparser.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
