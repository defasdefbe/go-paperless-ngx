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

func TestStoragePathNewWithOptionalParams(t *testing.T) {
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
	_, err := client.StoragePaths.New(context.TODO(), paperlessngx.StoragePathNewParams{
		StoragePathRequest: paperlessngx.StoragePathRequestParam{
			Name:              "x",
			Path:              "x",
			IsInsensitive:     paperlessngx.Bool(true),
			Match:             paperlessngx.String("match"),
			MatchingAlgorithm: 0,
			Owner:             paperlessngx.Int(0),
			SetPermissions: paperlessngx.StoragePathRequestSetPermissionsParam{
				Change: paperlessngx.StoragePathRequestSetPermissionsChangeParam{
					Groups: []int64{0},
					Users:  []int64{0},
				},
				View: paperlessngx.StoragePathRequestSetPermissionsViewParam{
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

func TestStoragePathGetWithOptionalParams(t *testing.T) {
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
	_, err := client.StoragePaths.Get(
		context.TODO(),
		0,
		paperlessngx.StoragePathGetParams{
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

func TestStoragePathUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.StoragePaths.Update(
		context.TODO(),
		0,
		paperlessngx.StoragePathUpdateParams{
			IsInsensitive:     paperlessngx.Bool(true),
			Match:             paperlessngx.String("match"),
			MatchingAlgorithm: 0,
			Name:              paperlessngx.String("x"),
			Owner:             paperlessngx.Int(0),
			Path:              paperlessngx.String("x"),
			SetPermissions: paperlessngx.StoragePathUpdateParamsSetPermissions{
				Change: paperlessngx.StoragePathUpdateParamsSetPermissionsChange{
					Groups: []int64{0},
					Users:  []int64{0},
				},
				View: paperlessngx.StoragePathUpdateParamsSetPermissionsView{
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

func TestStoragePathListWithOptionalParams(t *testing.T) {
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
	_, err := client.StoragePaths.List(context.TODO(), paperlessngx.StoragePathListParams{
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
		PathIcontains:   paperlessngx.String("path__icontains"),
		PathIendswith:   paperlessngx.String("path__iendswith"),
		PathIexact:      paperlessngx.String("path__iexact"),
		PathIstartswith: paperlessngx.String("path__istartswith"),
	})
	if err != nil {
		var apierr *paperlessngx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestStoragePathDelete(t *testing.T) {
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
	err := client.StoragePaths.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *paperlessngx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestStoragePathTestWithOptionalParams(t *testing.T) {
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
	_, err := client.StoragePaths.Test(context.TODO(), paperlessngx.StoragePathTestParams{
		StoragePathRequest: paperlessngx.StoragePathRequestParam{
			Name:              "x",
			Path:              "x",
			IsInsensitive:     paperlessngx.Bool(true),
			Match:             paperlessngx.String("match"),
			MatchingAlgorithm: 0,
			Owner:             paperlessngx.Int(0),
			SetPermissions: paperlessngx.StoragePathRequestSetPermissionsParam{
				Change: paperlessngx.StoragePathRequestSetPermissionsChangeParam{
					Groups: []int64{0},
					Users:  []int64{0},
				},
				View: paperlessngx.StoragePathRequestSetPermissionsViewParam{
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
