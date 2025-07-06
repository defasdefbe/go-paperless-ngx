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

func TestWorkflowNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Workflows.New(context.TODO(), paperlessngx.WorkflowNewParams{
		WorkflowRequest: paperlessngx.WorkflowRequestParam{
			Actions: []paperlessngx.WorkflowActionRequestParam{{
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
			}},
			Name: "x",
			Triggers: []paperlessngx.WorkflowTriggerRequestParam{{
				Type:                          1,
				ID:                            paperlessngx.Int(0),
				FilterFilename:                paperlessngx.String("filter_filename"),
				FilterHasCorrespondent:        paperlessngx.Int(0),
				FilterHasDocumentType:         paperlessngx.Int(0),
				FilterHasTags:                 []int64{0},
				FilterMailrule:                paperlessngx.Int(0),
				FilterPath:                    paperlessngx.String("filter_path"),
				IsInsensitive:                 paperlessngx.Bool(true),
				Match:                         paperlessngx.String("match"),
				MatchingAlgorithm:             0,
				ScheduleDateCustomField:       paperlessngx.Int(0),
				ScheduleDateField:             paperlessngx.ScheduleDateFieldEnumAdded,
				ScheduleIsRecurring:           paperlessngx.Bool(true),
				ScheduleOffsetDays:            paperlessngx.Int(-9007199254740991),
				ScheduleRecurringIntervalDays: paperlessngx.Int(1),
				Sources:                       []int64{1},
			}},
			Enabled: paperlessngx.Bool(true),
			Order:   paperlessngx.Int(0),
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

func TestWorkflowGet(t *testing.T) {
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
	_, err := client.Workflows.Get(context.TODO(), 0)
	if err != nil {
		var apierr *paperlessngx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkflowUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Workflows.Update(
		context.TODO(),
		0,
		paperlessngx.WorkflowUpdateParams{
			Actions: []paperlessngx.WorkflowActionRequestParam{{
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
			}},
			Enabled: paperlessngx.Bool(true),
			Name:    paperlessngx.String("x"),
			Order:   paperlessngx.Int(0),
			Triggers: []paperlessngx.WorkflowTriggerRequestParam{{
				Type:                          1,
				ID:                            paperlessngx.Int(0),
				FilterFilename:                paperlessngx.String("filter_filename"),
				FilterHasCorrespondent:        paperlessngx.Int(0),
				FilterHasDocumentType:         paperlessngx.Int(0),
				FilterHasTags:                 []int64{0},
				FilterMailrule:                paperlessngx.Int(0),
				FilterPath:                    paperlessngx.String("filter_path"),
				IsInsensitive:                 paperlessngx.Bool(true),
				Match:                         paperlessngx.String("match"),
				MatchingAlgorithm:             0,
				ScheduleDateCustomField:       paperlessngx.Int(0),
				ScheduleDateField:             paperlessngx.ScheduleDateFieldEnumAdded,
				ScheduleIsRecurring:           paperlessngx.Bool(true),
				ScheduleOffsetDays:            paperlessngx.Int(-9007199254740991),
				ScheduleRecurringIntervalDays: paperlessngx.Int(1),
				Sources:                       []int64{1},
			}},
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

func TestWorkflowListWithOptionalParams(t *testing.T) {
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
	_, err := client.Workflows.List(context.TODO(), paperlessngx.WorkflowListParams{
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

func TestWorkflowDelete(t *testing.T) {
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
	err := client.Workflows.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *paperlessngx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
