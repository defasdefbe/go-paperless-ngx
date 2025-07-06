// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package paperlessngx_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/paperless-nix-go"
	"github.com/stainless-sdks/paperless-nix-go/internal/testutil"
	"github.com/stainless-sdks/paperless-nix-go/option"
)

func TestCorrespondentNewWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Correspondents.New(context.TODO(), paperlessngx.CorrespondentNewParams{
		CorrespondentRequest: paperlessngx.CorrespondentRequestParam{
			Name:              "x",
			IsInsensitive:     paperlessngx.Bool(true),
			Match:             paperlessngx.String("match"),
			MatchingAlgorithm: 0,
			Owner:             paperlessngx.Int(0),
			SetPermissions: paperlessngx.CorrespondentRequestSetPermissionsParam{
				Change: paperlessngx.CorrespondentRequestSetPermissionsChangeParam{
					Groups: []int64{0},
					Users:  []int64{0},
				},
				View: paperlessngx.CorrespondentRequestSetPermissionsViewParam{
					Groups: []int64{0},
					Users:  []int64{0},
				},
			},
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

func TestCorrespondentGetWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Correspondents.Get(
		context.TODO(),
		0,
		paperlessngx.CorrespondentGetParams{
			FullPerms: paperlessngx.Bool(true),
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

func TestCorrespondentUpdateWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Correspondents.Update(
		context.TODO(),
		0,
		paperlessngx.CorrespondentUpdateParams{
			IsInsensitive:     paperlessngx.Bool(true),
			Match:             paperlessngx.String("match"),
			MatchingAlgorithm: 0,
			Name:              paperlessngx.String("x"),
			Owner:             paperlessngx.Int(0),
			SetPermissions: paperlessngx.CorrespondentUpdateParamsSetPermissions{
				Change: paperlessngx.CorrespondentUpdateParamsSetPermissionsChange{
					Groups: []int64{0},
					Users:  []int64{0},
				},
				View: paperlessngx.CorrespondentUpdateParamsSetPermissionsView{
					Groups: []int64{0},
					Users:  []int64{0},
				},
			},
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

func TestCorrespondentListWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Correspondents.List(context.TODO(), paperlessngx.CorrespondentListParams{
		ID:              paperlessngx.Int(0),
		FullPerms:       paperlessngx.Bool(true),
		IDIn:            []int64{0},
		NameIcontains:   paperlessngx.String("name__icontains"),
		NameIendswith:   paperlessngx.String("name__iendswith"),
		NameIexact:      paperlessngx.String("name__iexact"),
		NameIstartswith: paperlessngx.String("name__istartswith"),
		Ordering:        paperlessngx.String("ordering"),
		Page:            paperlessngx.Int(0),
		PageSize:        paperlessngx.Int(0),
	})
	if err != nil {
		var apierr *paperlessngx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCorrespondentDelete(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	err := client.Correspondents.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *paperlessngx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
