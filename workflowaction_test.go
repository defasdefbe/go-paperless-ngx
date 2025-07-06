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

func TestWorkflowActionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.WorkflowActions.New(context.TODO(), paperlessnix.WorkflowActionNewParams{
		WorkflowActionRequest: paperlessnix.WorkflowActionRequestParam{
			ID:                       paperlessnix.Int(0),
			AssignChangeGroups:       []int64{0},
			AssignChangeUsers:        []int64{0},
			AssignCorrespondent:      paperlessnix.Int(0),
			AssignCustomFields:       []int64{0},
			AssignCustomFieldsValues: map[string]interface{}{},
			AssignDocumentType:       paperlessnix.Int(0),
			AssignOwner:              paperlessnix.Int(0),
			AssignStoragePath:        paperlessnix.Int(0),
			AssignTags:               []int64{0},
			AssignTitle:              paperlessnix.String("assign_title"),
			AssignViewGroups:         []int64{0},
			AssignViewUsers:          []int64{0},
			Email: paperlessnix.WorkflowActionEmailParam{
				Body:            "x",
				Subject:         "x",
				To:              "x",
				ID:              paperlessnix.Int(0),
				IncludeDocument: paperlessnix.Bool(true),
			},
			RemoveAllCorrespondents: paperlessnix.Bool(true),
			RemoveAllCustomFields:   paperlessnix.Bool(true),
			RemoveAllDocumentTypes:  paperlessnix.Bool(true),
			RemoveAllOwners:         paperlessnix.Bool(true),
			RemoveAllPermissions:    paperlessnix.Bool(true),
			RemoveAllStoragePaths:   paperlessnix.Bool(true),
			RemoveAllTags:           paperlessnix.Bool(true),
			RemoveChangeGroups:      []int64{0},
			RemoveChangeUsers:       []int64{0},
			RemoveCorrespondents:    []int64{0},
			RemoveCustomFields:      []int64{0},
			RemoveDocumentTypes:     []int64{0},
			RemoveOwners:            []int64{0},
			RemoveStoragePaths:      []int64{0},
			RemoveTags:              []int64{0},
			RemoveViewGroups:        []int64{0},
			RemoveViewUsers:         []int64{0},
			Type:                    1,
			Webhook: paperlessnix.WorkflowActionWebhookParam{
				URL:             "x",
				ID:              paperlessnix.Int(0),
				AsJson:          paperlessnix.Bool(true),
				Body:            paperlessnix.String("body"),
				Headers:         map[string]interface{}{},
				IncludeDocument: paperlessnix.Bool(true),
				Params:          map[string]interface{}{},
				UseParams:       paperlessnix.Bool(true),
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

func TestWorkflowActionGet(t *testing.T) {
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
	_, err := client.WorkflowActions.Get(context.TODO(), 0)
	if err != nil {
		var apierr *paperlessnix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkflowActionUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.WorkflowActions.Update(
		context.TODO(),
		0,
		paperlessnix.WorkflowActionUpdateParams{
			ID:                       paperlessnix.Int(0),
			AssignChangeGroups:       []int64{0},
			AssignChangeUsers:        []int64{0},
			AssignCorrespondent:      paperlessnix.Int(0),
			AssignCustomFields:       []int64{0},
			AssignCustomFieldsValues: map[string]interface{}{},
			AssignDocumentType:       paperlessnix.Int(0),
			AssignOwner:              paperlessnix.Int(0),
			AssignStoragePath:        paperlessnix.Int(0),
			AssignTags:               []int64{0},
			AssignTitle:              paperlessnix.String("assign_title"),
			AssignViewGroups:         []int64{0},
			AssignViewUsers:          []int64{0},
			Email: paperlessnix.WorkflowActionEmailParam{
				Body:            "x",
				Subject:         "x",
				To:              "x",
				ID:              paperlessnix.Int(0),
				IncludeDocument: paperlessnix.Bool(true),
			},
			RemoveAllCorrespondents: paperlessnix.Bool(true),
			RemoveAllCustomFields:   paperlessnix.Bool(true),
			RemoveAllDocumentTypes:  paperlessnix.Bool(true),
			RemoveAllOwners:         paperlessnix.Bool(true),
			RemoveAllPermissions:    paperlessnix.Bool(true),
			RemoveAllStoragePaths:   paperlessnix.Bool(true),
			RemoveAllTags:           paperlessnix.Bool(true),
			RemoveChangeGroups:      []int64{0},
			RemoveChangeUsers:       []int64{0},
			RemoveCorrespondents:    []int64{0},
			RemoveCustomFields:      []int64{0},
			RemoveDocumentTypes:     []int64{0},
			RemoveOwners:            []int64{0},
			RemoveStoragePaths:      []int64{0},
			RemoveTags:              []int64{0},
			RemoveViewGroups:        []int64{0},
			RemoveViewUsers:         []int64{0},
			Type:                    1,
			Webhook: paperlessnix.WorkflowActionWebhookParam{
				URL:             "x",
				ID:              paperlessnix.Int(0),
				AsJson:          paperlessnix.Bool(true),
				Body:            paperlessnix.String("body"),
				Headers:         map[string]interface{}{},
				IncludeDocument: paperlessnix.Bool(true),
				Params:          map[string]interface{}{},
				UseParams:       paperlessnix.Bool(true),
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

func TestWorkflowActionListWithOptionalParams(t *testing.T) {
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
	_, err := client.WorkflowActions.List(context.TODO(), paperlessnix.WorkflowActionListParams{
		Page:     paperlessnix.Int(0),
		PageSize: paperlessnix.Int(0),
	})
	if err != nil {
		var apierr *paperlessnix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkflowActionDelete(t *testing.T) {
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
	err := client.WorkflowActions.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *paperlessnix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
