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

func TestWorkflowTriggerNewWithOptionalParams(t *testing.T) {
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
	_, err := client.WorkflowTriggers.New(context.TODO(), paperlessnix.WorkflowTriggerNewParams{
		WorkflowTriggerRequest: paperlessnix.WorkflowTriggerRequestParam{
			Type:                          1,
			ID:                            paperlessnix.Int(0),
			FilterFilename:                paperlessnix.String("filter_filename"),
			FilterHasCorrespondent:        paperlessnix.Int(0),
			FilterHasDocumentType:         paperlessnix.Int(0),
			FilterHasTags:                 []int64{0},
			FilterMailrule:                paperlessnix.Int(0),
			FilterPath:                    paperlessnix.String("filter_path"),
			IsInsensitive:                 paperlessnix.Bool(true),
			Match:                         paperlessnix.String("match"),
			MatchingAlgorithm:             0,
			ScheduleDateCustomField:       paperlessnix.Int(0),
			ScheduleDateField:             paperlessnix.ScheduleDateFieldEnumAdded,
			ScheduleIsRecurring:           paperlessnix.Bool(true),
			ScheduleOffsetDays:            paperlessnix.Int(-9007199254740991),
			ScheduleRecurringIntervalDays: paperlessnix.Int(1),
			Sources:                       []int64{1},
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

func TestWorkflowTriggerGet(t *testing.T) {
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
	_, err := client.WorkflowTriggers.Get(context.TODO(), 0)
	if err != nil {
		var apierr *paperlessnix.Error
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
	client := paperlessnix.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.WorkflowTriggers.Update(
		context.TODO(),
		0,
		paperlessnix.WorkflowTriggerUpdateParams{
			ID:                            paperlessnix.Int(0),
			FilterFilename:                paperlessnix.String("filter_filename"),
			FilterHasCorrespondent:        paperlessnix.Int(0),
			FilterHasDocumentType:         paperlessnix.Int(0),
			FilterHasTags:                 []int64{0},
			FilterMailrule:                paperlessnix.Int(0),
			FilterPath:                    paperlessnix.String("filter_path"),
			IsInsensitive:                 paperlessnix.Bool(true),
			Match:                         paperlessnix.String("match"),
			MatchingAlgorithm:             0,
			ScheduleDateCustomField:       paperlessnix.Int(0),
			ScheduleDateField:             paperlessnix.ScheduleDateFieldEnumAdded,
			ScheduleIsRecurring:           paperlessnix.Bool(true),
			ScheduleOffsetDays:            paperlessnix.Int(-9007199254740991),
			ScheduleRecurringIntervalDays: paperlessnix.Int(1),
			Sources:                       []int64{1},
			Type:                          1,
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

func TestWorkflowTriggerListWithOptionalParams(t *testing.T) {
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
	_, err := client.WorkflowTriggers.List(context.TODO(), paperlessnix.WorkflowTriggerListParams{
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

func TestWorkflowTriggerDelete(t *testing.T) {
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
	err := client.WorkflowTriggers.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *paperlessnix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
