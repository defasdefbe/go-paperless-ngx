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

func TestMailRuleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.MailRules.New(context.TODO(), paperlessngx.MailRuleNewParams{
		MailRuleRequest: paperlessngx.MailRuleRequestParam{
			Account:                         0,
			Name:                            "x",
			Action:                          1,
			ActionParameter:                 paperlessngx.String("x"),
			AssignCorrespondent:             paperlessngx.Int(0),
			AssignCorrespondentFrom:         1,
			AssignDocumentType:              paperlessngx.Int(0),
			AssignOwnerFromRule:             paperlessngx.Bool(true),
			AssignTags:                      []int64{0},
			AssignTitleFrom:                 1,
			AttachmentType:                  1,
			ConsumptionScope:                1,
			Enabled:                         paperlessngx.Bool(true),
			FilterAttachmentFilenameExclude: paperlessngx.String("filter_attachment_filename_exclude"),
			FilterAttachmentFilenameInclude: paperlessngx.String("filter_attachment_filename_include"),
			FilterBody:                      paperlessngx.String("filter_body"),
			FilterFrom:                      paperlessngx.String("filter_from"),
			FilterSubject:                   paperlessngx.String("filter_subject"),
			FilterTo:                        paperlessngx.String("filter_to"),
			Folder:                          paperlessngx.String("x"),
			MaximumAge:                      paperlessngx.Int(0),
			Order:                           paperlessngx.Int(0),
			Owner:                           paperlessngx.Int(0),
			PdfLayout:                       0,
			SetPermissions: paperlessngx.MailRuleRequestSetPermissionsParam{
				Change: paperlessngx.MailRuleRequestSetPermissionsChangeParam{
					Groups: []int64{0},
					Users:  []int64{0},
				},
				View: paperlessngx.MailRuleRequestSetPermissionsViewParam{
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

func TestMailRuleGet(t *testing.T) {
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
	_, err := client.MailRules.Get(context.TODO(), 0)
	if err != nil {
		var apierr *paperlessngx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMailRuleUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.MailRules.Update(
		context.TODO(),
		0,
		paperlessngx.MailRuleUpdateParams{
			Account:                         paperlessngx.Int(0),
			Action:                          1,
			ActionParameter:                 paperlessngx.String("x"),
			AssignCorrespondent:             paperlessngx.Int(0),
			AssignCorrespondentFrom:         1,
			AssignDocumentType:              paperlessngx.Int(0),
			AssignOwnerFromRule:             paperlessngx.Bool(true),
			AssignTags:                      []int64{0},
			AssignTitleFrom:                 1,
			AttachmentType:                  1,
			ConsumptionScope:                1,
			Enabled:                         paperlessngx.Bool(true),
			FilterAttachmentFilenameExclude: paperlessngx.String("filter_attachment_filename_exclude"),
			FilterAttachmentFilenameInclude: paperlessngx.String("filter_attachment_filename_include"),
			FilterBody:                      paperlessngx.String("filter_body"),
			FilterFrom:                      paperlessngx.String("filter_from"),
			FilterSubject:                   paperlessngx.String("filter_subject"),
			FilterTo:                        paperlessngx.String("filter_to"),
			Folder:                          paperlessngx.String("x"),
			MaximumAge:                      paperlessngx.Int(0),
			Name:                            paperlessngx.String("x"),
			Order:                           paperlessngx.Int(0),
			Owner:                           paperlessngx.Int(0),
			PdfLayout:                       0,
			SetPermissions: paperlessngx.MailRuleUpdateParamsSetPermissions{
				Change: paperlessngx.MailRuleUpdateParamsSetPermissionsChange{
					Groups: []int64{0},
					Users:  []int64{0},
				},
				View: paperlessngx.MailRuleUpdateParamsSetPermissionsView{
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

func TestMailRuleListWithOptionalParams(t *testing.T) {
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
	_, err := client.MailRules.List(context.TODO(), paperlessngx.MailRuleListParams{
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

func TestMailRuleDelete(t *testing.T) {
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
	err := client.MailRules.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *paperlessngx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
