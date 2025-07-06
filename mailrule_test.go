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

func TestMailRuleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.MailRules.New(context.TODO(), paperlessnix.MailRuleNewParams{
		MailRuleRequest: paperlessnix.MailRuleRequestParam{
			Account:                         0,
			Name:                            "x",
			Action:                          1,
			ActionParameter:                 paperlessnix.String("x"),
			AssignCorrespondent:             paperlessnix.Int(0),
			AssignCorrespondentFrom:         1,
			AssignDocumentType:              paperlessnix.Int(0),
			AssignOwnerFromRule:             paperlessnix.Bool(true),
			AssignTags:                      []int64{0},
			AssignTitleFrom:                 1,
			AttachmentType:                  1,
			ConsumptionScope:                1,
			Enabled:                         paperlessnix.Bool(true),
			FilterAttachmentFilenameExclude: paperlessnix.String("filter_attachment_filename_exclude"),
			FilterAttachmentFilenameInclude: paperlessnix.String("filter_attachment_filename_include"),
			FilterBody:                      paperlessnix.String("filter_body"),
			FilterFrom:                      paperlessnix.String("filter_from"),
			FilterSubject:                   paperlessnix.String("filter_subject"),
			FilterTo:                        paperlessnix.String("filter_to"),
			Folder:                          paperlessnix.String("x"),
			MaximumAge:                      paperlessnix.Int(0),
			Order:                           paperlessnix.Int(0),
			Owner:                           paperlessnix.Int(0),
			PdfLayout:                       0,
			SetPermissions: paperlessnix.MailRuleRequestSetPermissionsParam{
				Change: paperlessnix.MailRuleRequestSetPermissionsChangeParam{
					Groups: []int64{0},
					Users:  []int64{0},
				},
				View: paperlessnix.MailRuleRequestSetPermissionsViewParam{
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

func TestMailRuleGet(t *testing.T) {
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
	_, err := client.MailRules.Get(context.TODO(), 0)
	if err != nil {
		var apierr *paperlessnix.Error
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
	client := paperlessnix.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.MailRules.Update(
		context.TODO(),
		0,
		paperlessnix.MailRuleUpdateParams{
			Account:                         paperlessnix.Int(0),
			Action:                          1,
			ActionParameter:                 paperlessnix.String("x"),
			AssignCorrespondent:             paperlessnix.Int(0),
			AssignCorrespondentFrom:         1,
			AssignDocumentType:              paperlessnix.Int(0),
			AssignOwnerFromRule:             paperlessnix.Bool(true),
			AssignTags:                      []int64{0},
			AssignTitleFrom:                 1,
			AttachmentType:                  1,
			ConsumptionScope:                1,
			Enabled:                         paperlessnix.Bool(true),
			FilterAttachmentFilenameExclude: paperlessnix.String("filter_attachment_filename_exclude"),
			FilterAttachmentFilenameInclude: paperlessnix.String("filter_attachment_filename_include"),
			FilterBody:                      paperlessnix.String("filter_body"),
			FilterFrom:                      paperlessnix.String("filter_from"),
			FilterSubject:                   paperlessnix.String("filter_subject"),
			FilterTo:                        paperlessnix.String("filter_to"),
			Folder:                          paperlessnix.String("x"),
			MaximumAge:                      paperlessnix.Int(0),
			Name:                            paperlessnix.String("x"),
			Order:                           paperlessnix.Int(0),
			Owner:                           paperlessnix.Int(0),
			PdfLayout:                       0,
			SetPermissions: paperlessnix.MailRuleUpdateParamsSetPermissions{
				Change: paperlessnix.MailRuleUpdateParamsSetPermissionsChange{
					Groups: []int64{0},
					Users:  []int64{0},
				},
				View: paperlessnix.MailRuleUpdateParamsSetPermissionsView{
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

func TestMailRuleListWithOptionalParams(t *testing.T) {
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
	_, err := client.MailRules.List(context.TODO(), paperlessnix.MailRuleListParams{
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

func TestMailRuleDelete(t *testing.T) {
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
	err := client.MailRules.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *paperlessnix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
