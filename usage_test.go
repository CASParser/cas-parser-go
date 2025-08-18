// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package casparser_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/cas-parser-go"
	"github.com/stainless-sdks/cas-parser-go/internal/testutil"
	"github.com/stainless-sdks/cas-parser-go/option"
)

func TestUsage(t *testing.T) {
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
	unifiedResponse, err := client.CasParser.SmartParse(context.TODO(), casparser.CasParserSmartParseParams{
		Password: casparser.String("ABCDF"),
		PdfURL:   casparser.String("https://your-cas-pdf-url-here.com"),
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", unifiedResponse.DematAccounts)
}
