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

func TestInboxCheckConnectionStatus(t *testing.T) {
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
	_, err := client.Inbox.CheckConnectionStatus(context.TODO(), casparser.InboxCheckConnectionStatusParams{
		XInboxToken: "x-inbox-token",
	})
	if err != nil {
		var apierr *casparser.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInboxConnectEmailWithOptionalParams(t *testing.T) {
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
	_, err := client.Inbox.ConnectEmail(context.TODO(), casparser.InboxConnectEmailParams{
		RedirectUri: "https://yourapp.com/oauth-callback",
		State:       casparser.String("abc123"),
	})
	if err != nil {
		var apierr *casparser.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInboxDisconnectEmail(t *testing.T) {
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
	_, err := client.Inbox.DisconnectEmail(context.TODO(), casparser.InboxDisconnectEmailParams{
		XInboxToken: "x-inbox-token",
	})
	if err != nil {
		var apierr *casparser.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInboxListCasFilesWithOptionalParams(t *testing.T) {
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
	_, err := client.Inbox.ListCasFiles(context.TODO(), casparser.InboxListCasFilesParams{
		XInboxToken: "x-inbox-token",
		CasTypes:    []string{"cdsl", "nsdl"},
		EndDate:     casparser.Time(time.Now()),
		StartDate:   casparser.Time(time.Now()),
	})
	if err != nil {
		var apierr *casparser.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
