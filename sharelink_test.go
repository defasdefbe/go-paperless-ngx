// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package paperlessngx_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/defasdefbe/go-paperless-ngx"
	"github.com/defasdefbe/go-paperless-ngx/internal/testutil"
	"github.com/defasdefbe/go-paperless-ngx/option"
)

func TestShareLinkNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ShareLinks.New(context.TODO(), paperlessngx.ShareLinkNewParams{
		ShareLinkRequest: paperlessngx.ShareLinkRequestParam{
			Document:    paperlessngx.Int(0),
			Expiration:  paperlessngx.Time(time.Now()),
			FileVersion: paperlessngx.FileVersionEnumArchive,
		},
	})
	if err != nil {
		var apierr *paperlessngx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestShareLinkGet(t *testing.T) {
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
	_, err := client.ShareLinks.Get(context.TODO(), 0)
	if err != nil {
		var apierr *paperlessngx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestShareLinkUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ShareLinks.Update(
		context.TODO(),
		0,
		paperlessngx.ShareLinkUpdateParams{
			Document:    paperlessngx.Int(0),
			Expiration:  paperlessngx.Time(time.Now()),
			FileVersion: paperlessngx.FileVersionEnumArchive,
		},
	)
	if err != nil {
		var apierr *paperlessngx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestShareLinkListWithOptionalParams(t *testing.T) {
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
	_, err := client.ShareLinks.List(context.TODO(), paperlessngx.ShareLinkListParams{
		CreatedDateGt:     paperlessngx.Time(time.Now()),
		CreatedDateGte:    paperlessngx.Time(time.Now()),
		CreatedDateLt:     paperlessngx.Time(time.Now()),
		CreatedDateLte:    paperlessngx.Time(time.Now()),
		CreatedDay:        paperlessngx.Float(0),
		CreatedGt:         paperlessngx.Time(time.Now()),
		CreatedGte:        paperlessngx.Time(time.Now()),
		CreatedLt:         paperlessngx.Time(time.Now()),
		CreatedLte:        paperlessngx.Time(time.Now()),
		CreatedMonth:      paperlessngx.Float(0),
		CreatedYear:       paperlessngx.Float(0),
		ExpirationDateGt:  paperlessngx.Time(time.Now()),
		ExpirationDateGte: paperlessngx.Time(time.Now()),
		ExpirationDateLt:  paperlessngx.Time(time.Now()),
		ExpirationDateLte: paperlessngx.Time(time.Now()),
		ExpirationDay:     paperlessngx.Float(0),
		ExpirationGt:      paperlessngx.Time(time.Now()),
		ExpirationGte:     paperlessngx.Time(time.Now()),
		ExpirationLt:      paperlessngx.Time(time.Now()),
		ExpirationLte:     paperlessngx.Time(time.Now()),
		ExpirationMonth:   paperlessngx.Float(0),
		ExpirationYear:    paperlessngx.Float(0),
		Ordering:          paperlessngx.String("ordering"),
		Page:              paperlessngx.Int(0),
		PageSize:          paperlessngx.Int(0),
	})
	if err != nil {
		var apierr *paperlessngx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestShareLinkDelete(t *testing.T) {
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
	err := client.ShareLinks.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *paperlessngx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
