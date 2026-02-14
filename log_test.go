// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package casparser_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/CASParser/cas-parser-go"
	"github.com/CASParser/cas-parser-go/internal/testutil"
	"github.com/CASParser/cas-parser-go/option"
)

func TestLogNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Logs.New(context.TODO(), casparser.LogNewParams{
		EndTime:   casparser.Time(time.Now()),
		Limit:     casparser.Int(1),
		StartTime: casparser.Time(time.Now()),
	})
	if err != nil {
		var apierr *casparser.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLogGetSummaryWithOptionalParams(t *testing.T) {
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
	_, err := client.Logs.GetSummary(context.TODO(), casparser.LogGetSummaryParams{
		EndTime:   casparser.Time(time.Now()),
		StartTime: casparser.Time(time.Now()),
	})
	if err != nil {
		var apierr *casparser.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
