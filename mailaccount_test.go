// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package paperlessngx_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/defasdefbe/go-paperless-ngx"
	"github.com/defasdefbe/go-paperless-ngx/internal/testutil"
	"github.com/defasdefbe/go-paperless-ngx/option"
)

func TestMailAccountNewWithOptionalParams(t *testing.T) {
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
	_, err := client.MailAccounts.New(context.TODO(), paperlessngx.MailAccountNewParams{
		MailAccountRequest: paperlessngx.MailAccountRequestParam{
			ImapServer:   "x",
			Name:         "x",
			Password:     "x",
			Username:     "x",
			AccountType:  1,
			CharacterSet: paperlessngx.String("x"),
			Expiration:   paperlessngx.Time(time.Now()),
			ImapPort:     paperlessngx.Int(-9007199254740991),
			ImapSecurity: 1,
			IsToken:      paperlessngx.Bool(true),
			Owner:        paperlessngx.Int(0),
			SetPermissions: paperlessngx.MailAccountRequestSetPermissionsParam{
				Change: paperlessngx.MailAccountRequestSetPermissionsChangeParam{
					Groups: []int64{0},
					Users:  []int64{0},
				},
				View: paperlessngx.MailAccountRequestSetPermissionsViewParam{
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

func TestMailAccountGet(t *testing.T) {
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
	_, err := client.MailAccounts.Get(context.TODO(), 0)
	if err != nil {
		var apierr *paperlessngx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMailAccountUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.MailAccounts.Update(
		context.TODO(),
		0,
		paperlessngx.MailAccountUpdateParams{
			AccountType:  1,
			CharacterSet: paperlessngx.String("x"),
			Expiration:   paperlessngx.Time(time.Now()),
			ImapPort:     paperlessngx.Int(-9007199254740991),
			ImapSecurity: 1,
			ImapServer:   paperlessngx.String("x"),
			IsToken:      paperlessngx.Bool(true),
			Name:         paperlessngx.String("x"),
			Owner:        paperlessngx.Int(0),
			Password:     paperlessngx.String("x"),
			SetPermissions: paperlessngx.MailAccountUpdateParamsSetPermissions{
				Change: paperlessngx.MailAccountUpdateParamsSetPermissionsChange{
					Groups: []int64{0},
					Users:  []int64{0},
				},
				View: paperlessngx.MailAccountUpdateParamsSetPermissionsView{
					Groups: []int64{0},
					Users:  []int64{0},
				},
			},
			Username: paperlessngx.String("x"),
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

func TestMailAccountListWithOptionalParams(t *testing.T) {
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
	_, err := client.MailAccounts.List(context.TODO(), paperlessngx.MailAccountListParams{
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

func TestMailAccountDelete(t *testing.T) {
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
	err := client.MailAccounts.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *paperlessngx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMailAccountProcessWithOptionalParams(t *testing.T) {
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
	_, err := client.MailAccounts.Process(
		context.TODO(),
		0,
		paperlessngx.MailAccountProcessParams{
			MailAccountRequest: paperlessngx.MailAccountRequestParam{
				ImapServer:   "x",
				Name:         "x",
				Password:     "x",
				Username:     "x",
				AccountType:  1,
				CharacterSet: paperlessngx.String("x"),
				Expiration:   paperlessngx.Time(time.Now()),
				ImapPort:     paperlessngx.Int(-9007199254740991),
				ImapSecurity: 1,
				IsToken:      paperlessngx.Bool(true),
				Owner:        paperlessngx.Int(0),
				SetPermissions: paperlessngx.MailAccountRequestSetPermissionsParam{
					Change: paperlessngx.MailAccountRequestSetPermissionsChangeParam{
						Groups: []int64{0},
						Users:  []int64{0},
					},
					View: paperlessngx.MailAccountRequestSetPermissionsViewParam{
						Groups: []int64{0},
						Users:  []int64{0},
					},
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

func TestMailAccountTestWithOptionalParams(t *testing.T) {
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
	_, err := client.MailAccounts.Test(context.TODO(), paperlessngx.MailAccountTestParams{
		MailAccountRequest: paperlessngx.MailAccountRequestParam{
			ImapServer:   "x",
			Name:         "x",
			Password:     "x",
			Username:     "x",
			AccountType:  1,
			CharacterSet: paperlessngx.String("x"),
			Expiration:   paperlessngx.Time(time.Now()),
			ImapPort:     paperlessngx.Int(-9007199254740991),
			ImapSecurity: 1,
			IsToken:      paperlessngx.Bool(true),
			Owner:        paperlessngx.Int(0),
			SetPermissions: paperlessngx.MailAccountRequestSetPermissionsParam{
				Change: paperlessngx.MailAccountRequestSetPermissionsChangeParam{
					Groups: []int64{0},
					Users:  []int64{0},
				},
				View: paperlessngx.MailAccountRequestSetPermissionsViewParam{
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
