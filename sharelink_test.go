// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package paperlessnix_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/paperless-nix-go"
	"github.com/stainless-sdks/paperless-nix-go/internal/testutil"
	"github.com/stainless-sdks/paperless-nix-go/option"
)

func TestShareLinkNewWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := paperlessnix.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.ShareLinks.New(context.TODO(), paperlessnix.ShareLinkNewParams{
		ShareLinkRequest: paperlessnix.ShareLinkRequestParam{
			Document:    paperlessnix.Int(0),
			Expiration:  paperlessnix.Time(time.Now()),
			FileVersion: paperlessnix.FileVersionEnumArchive,
		},
	})
	if err != nil {
		var apierr *paperlessnix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestShareLinkGet(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := paperlessnix.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.ShareLinks.Get(context.TODO(), 0)
	if err != nil {
		var apierr *paperlessnix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestShareLinkUpdateWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := paperlessnix.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.ShareLinks.Update(
		context.TODO(),
		0,
		paperlessnix.ShareLinkUpdateParams{
			Document:    paperlessnix.Int(0),
			Expiration:  paperlessnix.Time(time.Now()),
			FileVersion: paperlessnix.FileVersionEnumArchive,
		},
	)
	if err != nil {
		var apierr *paperlessnix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestShareLinkListWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := paperlessnix.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.ShareLinks.List(context.TODO(), paperlessnix.ShareLinkListParams{
		CreatedDateGt:     paperlessnix.Time(time.Now()),
		CreatedDateGte:    paperlessnix.Time(time.Now()),
		CreatedDateLt:     paperlessnix.Time(time.Now()),
		CreatedDateLte:    paperlessnix.Time(time.Now()),
		CreatedDay:        paperlessnix.Float(0),
		CreatedGt:         paperlessnix.Time(time.Now()),
		CreatedGte:        paperlessnix.Time(time.Now()),
		CreatedLt:         paperlessnix.Time(time.Now()),
		CreatedLte:        paperlessnix.Time(time.Now()),
		CreatedMonth:      paperlessnix.Float(0),
		CreatedYear:       paperlessnix.Float(0),
		ExpirationDateGt:  paperlessnix.Time(time.Now()),
		ExpirationDateGte: paperlessnix.Time(time.Now()),
		ExpirationDateLt:  paperlessnix.Time(time.Now()),
		ExpirationDateLte: paperlessnix.Time(time.Now()),
		ExpirationDay:     paperlessnix.Float(0),
		ExpirationGt:      paperlessnix.Time(time.Now()),
		ExpirationGte:     paperlessnix.Time(time.Now()),
		ExpirationLt:      paperlessnix.Time(time.Now()),
		ExpirationLte:     paperlessnix.Time(time.Now()),
		ExpirationMonth:   paperlessnix.Float(0),
		ExpirationYear:    paperlessnix.Float(0),
		Ordering:          paperlessnix.String("ordering"),
		Page:              paperlessnix.Int(0),
		PageSize:          paperlessnix.Int(0),
	})
	if err != nil {
		var apierr *paperlessnix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestShareLinkDelete(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := paperlessnix.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
		option.WithAPIKey("My API Key"),
	)
	err := client.ShareLinks.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *paperlessnix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
