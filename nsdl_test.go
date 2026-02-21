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

func TestNsdlParseWithOptionalParams(t *testing.T) {
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
	_, err := client.Nsdl.Parse(context.TODO(), casparser.NsdlParseParams{
		Password: casparser.String("password"),
		PdfFile:  casparser.String("pdf_file"),
		PdfURL:   casparser.String("https://example.com"),
	})
	if err != nil {
		var apierr *casparser.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
