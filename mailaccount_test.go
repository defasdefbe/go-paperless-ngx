// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package paperlessnix_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/paperless-nix-go"
	"github.com/stainless-sdks/paperless-nix-go/internal/testutil"
	"github.com/stainless-sdks/paperless-nix-go/option"
)

func TestMailAccountNewWithOptionalParams(t *testing.T) {
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
	_, err := client.MailAccounts.New(context.TODO(), paperlessnix.MailAccountNewParams{
		MailAccountRequest: paperlessnix.MailAccountRequestParam{
			ImapServer:   "x",
			Name:         "x",
			Password:     "x",
			Username:     "x",
			AccountType:  1,
			CharacterSet: paperlessnix.String("x"),
			Expiration:   paperlessnix.Time(time.Now()),
			ImapPort:     paperlessnix.Int(-9007199254740991),
			ImapSecurity: 1,
			IsToken:      paperlessnix.Bool(true),
			Owner:        paperlessnix.Int(0),
			SetPermissions: paperlessnix.MailAccountRequestSetPermissionsParam{
				Change: paperlessnix.MailAccountRequestSetPermissionsChangeParam{
					Groups: []int64{0},
					Users:  []int64{0},
				},
				View: paperlessnix.MailAccountRequestSetPermissionsViewParam{
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

func TestMailAccountGet(t *testing.T) {
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
	_, err := client.MailAccounts.Get(context.TODO(), 0)
	if err != nil {
		var apierr *paperlessnix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMailAccountUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.MailAccounts.Update(
		context.TODO(),
		0,
		paperlessnix.MailAccountUpdateParams{
			AccountType:  1,
			CharacterSet: paperlessnix.String("x"),
			Expiration:   paperlessnix.Time(time.Now()),
			ImapPort:     paperlessnix.Int(-9007199254740991),
			ImapSecurity: 1,
			ImapServer:   paperlessnix.String("x"),
			IsToken:      paperlessnix.Bool(true),
			Name:         paperlessnix.String("x"),
			Owner:        paperlessnix.Int(0),
			Password:     paperlessnix.String("x"),
			SetPermissions: paperlessnix.MailAccountUpdateParamsSetPermissions{
				Change: paperlessnix.MailAccountUpdateParamsSetPermissionsChange{
					Groups: []int64{0},
					Users:  []int64{0},
				},
				View: paperlessnix.MailAccountUpdateParamsSetPermissionsView{
					Groups: []int64{0},
					Users:  []int64{0},
				},
			},
			Username: paperlessnix.String("x"),
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

func TestMailAccountListWithOptionalParams(t *testing.T) {
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
	_, err := client.MailAccounts.List(context.TODO(), paperlessnix.MailAccountListParams{
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

func TestMailAccountDelete(t *testing.T) {
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
	err := client.MailAccounts.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *paperlessnix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMailAccountProcessWithOptionalParams(t *testing.T) {
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
	_, err := client.MailAccounts.Process(
		context.TODO(),
		0,
		paperlessnix.MailAccountProcessParams{
			MailAccountRequest: paperlessnix.MailAccountRequestParam{
				ImapServer:   "x",
				Name:         "x",
				Password:     "x",
				Username:     "x",
				AccountType:  1,
				CharacterSet: paperlessnix.String("x"),
				Expiration:   paperlessnix.Time(time.Now()),
				ImapPort:     paperlessnix.Int(-9007199254740991),
				ImapSecurity: 1,
				IsToken:      paperlessnix.Bool(true),
				Owner:        paperlessnix.Int(0),
				SetPermissions: paperlessnix.MailAccountRequestSetPermissionsParam{
					Change: paperlessnix.MailAccountRequestSetPermissionsChangeParam{
						Groups: []int64{0},
						Users:  []int64{0},
					},
					View: paperlessnix.MailAccountRequestSetPermissionsViewParam{
						Groups: []int64{0},
						Users:  []int64{0},
					},
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

func TestMailAccountTestWithOptionalParams(t *testing.T) {
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
	_, err := client.MailAccounts.Test(context.TODO(), paperlessnix.MailAccountTestParams{
		MailAccountRequest: paperlessnix.MailAccountRequestParam{
			ImapServer:   "x",
			Name:         "x",
			Password:     "x",
			Username:     "x",
			AccountType:  1,
			CharacterSet: paperlessnix.String("x"),
			Expiration:   paperlessnix.Time(time.Now()),
			ImapPort:     paperlessnix.Int(-9007199254740991),
			ImapSecurity: 1,
			IsToken:      paperlessnix.Bool(true),
			Owner:        paperlessnix.Int(0),
			SetPermissions: paperlessnix.MailAccountRequestSetPermissionsParam{
				Change: paperlessnix.MailAccountRequestSetPermissionsChangeParam{
					Groups: []int64{0},
					Users:  []int64{0},
				},
				View: paperlessnix.MailAccountRequestSetPermissionsViewParam{
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
