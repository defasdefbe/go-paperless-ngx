// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package paperlessnix_test

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
	client := paperlessnix.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Correspondents.New(context.TODO(), paperlessnix.CorrespondentNewParams{
		CorrespondentRequest: paperlessnix.CorrespondentRequestParam{
			Name:              "x",
			IsInsensitive:     paperlessnix.Bool(true),
			Match:             paperlessnix.String("match"),
			MatchingAlgorithm: 0,
			Owner:             paperlessnix.Int(0),
			SetPermissions: paperlessnix.CorrespondentRequestSetPermissionsParam{
				Change: paperlessnix.CorrespondentRequestSetPermissionsChangeParam{
					Groups: []int64{0},
					Users:  []int64{0},
				},
				View: paperlessnix.CorrespondentRequestSetPermissionsViewParam{
					Groups: []int64{0},
					Users:  []int64{0},
				},
			},
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

func TestCorrespondentGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Correspondents.Get(
		context.TODO(),
		0,
		paperlessnix.CorrespondentGetParams{
			FullPerms: paperlessnix.Bool(true),
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

func TestCorrespondentUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Correspondents.Update(
		context.TODO(),
		0,
		paperlessnix.CorrespondentUpdateParams{
			IsInsensitive:     paperlessnix.Bool(true),
			Match:             paperlessnix.String("match"),
			MatchingAlgorithm: 0,
			Name:              paperlessnix.String("x"),
			Owner:             paperlessnix.Int(0),
			SetPermissions: paperlessnix.CorrespondentUpdateParamsSetPermissions{
				Change: paperlessnix.CorrespondentUpdateParamsSetPermissionsChange{
					Groups: []int64{0},
					Users:  []int64{0},
				},
				View: paperlessnix.CorrespondentUpdateParamsSetPermissionsView{
					Groups: []int64{0},
					Users:  []int64{0},
				},
			},
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

func TestCorrespondentListWithOptionalParams(t *testing.T) {
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
	_, err := client.Correspondents.List(context.TODO(), paperlessnix.CorrespondentListParams{
		ID:              paperlessnix.Int(0),
		FullPerms:       paperlessnix.Bool(true),
		IDIn:            []int64{0},
		NameIcontains:   paperlessnix.String("name__icontains"),
		NameIendswith:   paperlessnix.String("name__iendswith"),
		NameIexact:      paperlessnix.String("name__iexact"),
		NameIstartswith: paperlessnix.String("name__istartswith"),
		Ordering:        paperlessnix.String("ordering"),
		Page:            paperlessnix.Int(0),
		PageSize:        paperlessnix.Int(0),
	})
	if err != nil {
		var apierr *paperlessnix.Error
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
	client := paperlessnix.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
		option.WithAPIKey("My API Key"),
	)
	err := client.Correspondents.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *paperlessnix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
