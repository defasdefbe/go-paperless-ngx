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

func TestStoragePathNewWithOptionalParams(t *testing.T) {
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
	_, err := client.StoragePaths.New(context.TODO(), paperlessnix.StoragePathNewParams{
		StoragePathRequest: paperlessnix.StoragePathRequestParam{
			Name:              "x",
			Path:              "x",
			IsInsensitive:     paperlessnix.Bool(true),
			Match:             paperlessnix.String("match"),
			MatchingAlgorithm: 0,
			Owner:             paperlessnix.Int(0),
			SetPermissions: paperlessnix.StoragePathRequestSetPermissionsParam{
				Change: paperlessnix.StoragePathRequestSetPermissionsChangeParam{
					Groups: []int64{0},
					Users:  []int64{0},
				},
				View: paperlessnix.StoragePathRequestSetPermissionsViewParam{
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

func TestStoragePathGetWithOptionalParams(t *testing.T) {
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
	_, err := client.StoragePaths.Get(
		context.TODO(),
		0,
		paperlessnix.StoragePathGetParams{
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

func TestStoragePathUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.StoragePaths.Update(
		context.TODO(),
		0,
		paperlessnix.StoragePathUpdateParams{
			IsInsensitive:     paperlessnix.Bool(true),
			Match:             paperlessnix.String("match"),
			MatchingAlgorithm: 0,
			Name:              paperlessnix.String("x"),
			Owner:             paperlessnix.Int(0),
			Path:              paperlessnix.String("x"),
			SetPermissions: paperlessnix.StoragePathUpdateParamsSetPermissions{
				Change: paperlessnix.StoragePathUpdateParamsSetPermissionsChange{
					Groups: []int64{0},
					Users:  []int64{0},
				},
				View: paperlessnix.StoragePathUpdateParamsSetPermissionsView{
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

func TestStoragePathListWithOptionalParams(t *testing.T) {
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
	_, err := client.StoragePaths.List(context.TODO(), paperlessnix.StoragePathListParams{
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
		PathIcontains:   paperlessnix.String("path__icontains"),
		PathIendswith:   paperlessnix.String("path__iendswith"),
		PathIexact:      paperlessnix.String("path__iexact"),
		PathIstartswith: paperlessnix.String("path__istartswith"),
	})
	if err != nil {
		var apierr *paperlessnix.Error
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
	client := paperlessnix.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
		option.WithAPIKey("My API Key"),
	)
	err := client.StoragePaths.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *paperlessnix.Error
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
	client := paperlessnix.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.StoragePaths.Test(context.TODO(), paperlessnix.StoragePathTestParams{
		StoragePathRequest: paperlessnix.StoragePathRequestParam{
			Name:              "x",
			Path:              "x",
			IsInsensitive:     paperlessnix.Bool(true),
			Match:             paperlessnix.String("match"),
			MatchingAlgorithm: 0,
			Owner:             paperlessnix.Int(0),
			SetPermissions: paperlessnix.StoragePathRequestSetPermissionsParam{
				Change: paperlessnix.StoragePathRequestSetPermissionsChangeParam{
					Groups: []int64{0},
					Users:  []int64{0},
				},
				View: paperlessnix.StoragePathRequestSetPermissionsViewParam{
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
