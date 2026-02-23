// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package casparser_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/CASParser/cas-parser-go"
	"github.com/CASParser/cas-parser-go/internal/testutil"
	"github.com/CASParser/cas-parser-go/option"
)

func TestInboundEmailNewWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.InboundEmail.New(context.TODO(), casparser.InboundEmailNewParams{
		CallbackURL:    "https://api.yourapp.com/webhooks/cas-email",
		Alias:          casparser.String("john-portfolio"),
		AllowedSources: []string{"cdsl", "nsdl"},
		Metadata: map[string]string{
			"plan":   "premium",
			"source": "onboarding",
		},
		Reference: casparser.String("user_12345"),
	})
	if err != nil {
		var apierr *casparser.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInboundEmailGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.InboundEmail.Get(context.TODO(), "ie_a1b2c3d4e5f6")
	if err != nil {
		var apierr *casparser.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInboundEmailListWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.InboundEmail.List(context.TODO(), casparser.InboundEmailListParams{
		Limit:  casparser.Int(1),
		Offset: casparser.Int(0),
		Status: casparser.InboundEmailListParamsStatusActive,
	})
	if err != nil {
		var apierr *casparser.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInboundEmailDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.InboundEmail.Delete(context.TODO(), "inbound_email_id")
	if err != nil {
		var apierr *casparser.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
