// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package paperlessnix_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/paperless-nix-go"
	"github.com/stainless-sdks/paperless-nix-go/internal/testutil"
	"github.com/stainless-sdks/paperless-nix-go/option"
)

func TestUsage(t *testing.T) {
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
	bulkEditObject, err := client.BulkEditObjects.New(context.TODO(), paperlessnix.BulkEditObjectNewParams{
		ObjectType: paperlessnix.BulkEditObjectNewParamsObjectTypeTags,
		Objects:    []int64{0},
		Operation:  paperlessnix.BulkEditObjectNewParamsOperationSetPermissions,
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", bulkEditObject.Result)
}
