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

func TestWorkflowTriggerNewWithOptionalParams(t *testing.T) {
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
	_, err := client.WorkflowTriggers.New(context.TODO(), paperlessngx.WorkflowTriggerNewParams{
		WorkflowTriggerRequest: paperlessngx.WorkflowTriggerRequestParam{
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

func TestWorkflowTriggerGet(t *testing.T) {
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
	_, err := client.WorkflowTriggers.Get(context.TODO(), 0)
	if err != nil {
		var apierr *paperlessngx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkflowTriggerUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.WorkflowTriggers.Update(
		context.TODO(),
		0,
		paperlessngx.WorkflowTriggerUpdateParams{
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
			Type:                          1,
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

func TestWorkflowTriggerListWithOptionalParams(t *testing.T) {
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
	_, err := client.WorkflowTriggers.List(context.TODO(), paperlessngx.WorkflowTriggerListParams{
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

func TestWorkflowTriggerDelete(t *testing.T) {
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
	err := client.WorkflowTriggers.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *paperlessngx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
