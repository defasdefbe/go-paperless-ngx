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

func TestWorkflowActionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.WorkflowActions.New(context.TODO(), paperlessngx.WorkflowActionNewParams{
		WorkflowActionRequest: paperlessngx.WorkflowActionRequestParam{
			ID:                       paperlessngx.Int(0),
			AssignChangeGroups:       []int64{0},
			AssignChangeUsers:        []int64{0},
			AssignCorrespondent:      paperlessngx.Int(0),
			AssignCustomFields:       []int64{0},
			AssignCustomFieldsValues: map[string]interface{}{},
			AssignDocumentType:       paperlessngx.Int(0),
			AssignOwner:              paperlessngx.Int(0),
			AssignStoragePath:        paperlessngx.Int(0),
			AssignTags:               []int64{0},
			AssignTitle:              paperlessngx.String("assign_title"),
			AssignViewGroups:         []int64{0},
			AssignViewUsers:          []int64{0},
			Email: paperlessngx.WorkflowActionEmailParam{
				Body:            "x",
				Subject:         "x",
				To:              "x",
				ID:              paperlessngx.Int(0),
				IncludeDocument: paperlessngx.Bool(true),
			},
			RemoveAllCorrespondents: paperlessngx.Bool(true),
			RemoveAllCustomFields:   paperlessngx.Bool(true),
			RemoveAllDocumentTypes:  paperlessngx.Bool(true),
			RemoveAllOwners:         paperlessngx.Bool(true),
			RemoveAllPermissions:    paperlessngx.Bool(true),
			RemoveAllStoragePaths:   paperlessngx.Bool(true),
			RemoveAllTags:           paperlessngx.Bool(true),
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
			Webhook: paperlessngx.WorkflowActionWebhookParam{
				URL:             "x",
				ID:              paperlessngx.Int(0),
				AsJson:          paperlessngx.Bool(true),
				Body:            paperlessngx.String("body"),
				Headers:         map[string]interface{}{},
				IncludeDocument: paperlessngx.Bool(true),
				Params:          map[string]interface{}{},
				UseParams:       paperlessngx.Bool(true),
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

func TestWorkflowActionGet(t *testing.T) {
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
	_, err := client.WorkflowActions.Get(context.TODO(), 0)
	if err != nil {
		var apierr *paperlessngx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkflowActionUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.WorkflowActions.Update(
		context.TODO(),
		0,
		paperlessngx.WorkflowActionUpdateParams{
			ID:                       paperlessngx.Int(0),
			AssignChangeGroups:       []int64{0},
			AssignChangeUsers:        []int64{0},
			AssignCorrespondent:      paperlessngx.Int(0),
			AssignCustomFields:       []int64{0},
			AssignCustomFieldsValues: map[string]interface{}{},
			AssignDocumentType:       paperlessngx.Int(0),
			AssignOwner:              paperlessngx.Int(0),
			AssignStoragePath:        paperlessngx.Int(0),
			AssignTags:               []int64{0},
			AssignTitle:              paperlessngx.String("assign_title"),
			AssignViewGroups:         []int64{0},
			AssignViewUsers:          []int64{0},
			Email: paperlessngx.WorkflowActionEmailParam{
				Body:            "x",
				Subject:         "x",
				To:              "x",
				ID:              paperlessngx.Int(0),
				IncludeDocument: paperlessngx.Bool(true),
			},
			RemoveAllCorrespondents: paperlessngx.Bool(true),
			RemoveAllCustomFields:   paperlessngx.Bool(true),
			RemoveAllDocumentTypes:  paperlessngx.Bool(true),
			RemoveAllOwners:         paperlessngx.Bool(true),
			RemoveAllPermissions:    paperlessngx.Bool(true),
			RemoveAllStoragePaths:   paperlessngx.Bool(true),
			RemoveAllTags:           paperlessngx.Bool(true),
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
			Webhook: paperlessngx.WorkflowActionWebhookParam{
				URL:             "x",
				ID:              paperlessngx.Int(0),
				AsJson:          paperlessngx.Bool(true),
				Body:            paperlessngx.String("body"),
				Headers:         map[string]interface{}{},
				IncludeDocument: paperlessngx.Bool(true),
				Params:          map[string]interface{}{},
				UseParams:       paperlessngx.Bool(true),
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

func TestWorkflowActionListWithOptionalParams(t *testing.T) {
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
	_, err := client.WorkflowActions.List(context.TODO(), paperlessngx.WorkflowActionListParams{
		Page:     paperlessngx.Int(0),
		PageSize: paperlessngx.Int(0),
	})
	if err != nil {
		var apierr *paperlessngx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkflowActionDelete(t *testing.T) {
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
	err := client.WorkflowActions.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *paperlessngx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
