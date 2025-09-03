// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package paperlessngx_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/defasdefbe/go-paperless-ngx"
	"github.com/defasdefbe/go-paperless-ngx/internal/testutil"
	"github.com/defasdefbe/go-paperless-ngx/option"
)

func TestSearchGetWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := paperlessngx.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Search.Get(context.TODO(), paperlessngx.SearchGetParams{
		Query:  "query",
		DBOnly: paperlessngx.Bool(true),
	})
	if err != nil {
		var apierr *paperlessngx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSearchListTagsWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := paperlessngx.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Search.ListTags(context.TODO(), paperlessngx.SearchListTagsParams{
		Limit: paperlessngx.Int(0),
		Term:  paperlessngx.String("term"),
	})
	if err != nil {
		var apierr *paperlessngx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
