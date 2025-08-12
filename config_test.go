// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package paperlessngx_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/defasdefbe/go-paperless-ngx"
	"github.com/defasdefbe/go-paperless-ngx/internal/testutil"
	"github.com/defasdefbe/go-paperless-ngx/option"
)

func TestConfigGet(t *testing.T) {
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
	_, err := client.Config.Get(context.TODO())
	if err != nil {
		var apierr *paperlessngx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConfigUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Config.Update(
		context.TODO(),
		0,
		paperlessngx.ConfigUpdateParams{
			BarcodeTagMapping:        map[string]interface{}{},
			UserArgs:                 map[string]interface{}{},
			AppLogo:                  io.Reader(bytes.NewBuffer([]byte("some file contents"))),
			AppTitle:                 paperlessngx.String("app_title"),
			BarcodeAsnPrefix:         paperlessngx.String("barcode_asn_prefix"),
			BarcodeDpi:               paperlessngx.Int(1),
			BarcodeEnableAsn:         paperlessngx.Bool(true),
			BarcodeEnableTag:         paperlessngx.Bool(true),
			BarcodeEnableTiffSupport: paperlessngx.Bool(true),
			BarcodeMaxPages:          paperlessngx.Int(1),
			BarcodeRetainSplitPages:  paperlessngx.Bool(true),
			BarcodeString:            paperlessngx.String("barcode_string"),
			BarcodeUpscale:           paperlessngx.Float(1),
			BarcodesEnabled:          paperlessngx.Bool(true),
			ColorConversionStrategy:  paperlessngx.ConfigUpdateParamsColorConversionStrategyLeaveColorUnchanged,
			Deskew:                   paperlessngx.Bool(true),
			ImageDpi:                 paperlessngx.Int(1),
			Language:                 paperlessngx.String("language"),
			MaxImagePixels:           paperlessngx.Float(0),
			Mode:                     paperlessngx.ConfigUpdateParamsModeSkip,
			OutputType:               paperlessngx.ConfigUpdateParamsOutputTypePdf,
			Pages:                    paperlessngx.Int(1),
			RotatePages:              paperlessngx.Bool(true),
			RotatePagesThreshold:     paperlessngx.Float(0),
			SkipArchiveFile:          paperlessngx.ConfigUpdateParamsSkipArchiveFileNever,
			UnpaperClean:             paperlessngx.ConfigUpdateParamsUnpaperCleanClean,
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

func TestConfigDelete(t *testing.T) {
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
	err := client.Config.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *paperlessngx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConfigPatchWithOptionalParams(t *testing.T) {
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
	_, err := client.Config.Patch(
		context.TODO(),
		0,
		paperlessngx.ConfigPatchParams{
			AppLogo:                  io.Reader(bytes.NewBuffer([]byte("some file contents"))),
			AppTitle:                 paperlessngx.String("app_title"),
			BarcodeAsnPrefix:         paperlessngx.String("barcode_asn_prefix"),
			BarcodeDpi:               paperlessngx.Int(1),
			BarcodeEnableAsn:         paperlessngx.Bool(true),
			BarcodeEnableTag:         paperlessngx.Bool(true),
			BarcodeEnableTiffSupport: paperlessngx.Bool(true),
			BarcodeMaxPages:          paperlessngx.Int(1),
			BarcodeRetainSplitPages:  paperlessngx.Bool(true),
			BarcodeString:            paperlessngx.String("barcode_string"),
			BarcodeTagMapping:        map[string]interface{}{},
			BarcodeUpscale:           paperlessngx.Float(1),
			BarcodesEnabled:          paperlessngx.Bool(true),
			ColorConversionStrategy:  paperlessngx.ConfigPatchParamsColorConversionStrategyLeaveColorUnchanged,
			Deskew:                   paperlessngx.Bool(true),
			ImageDpi:                 paperlessngx.Int(1),
			Language:                 paperlessngx.String("language"),
			MaxImagePixels:           paperlessngx.Float(0),
			Mode:                     paperlessngx.ConfigPatchParamsModeSkip,
			OutputType:               paperlessngx.ConfigPatchParamsOutputTypePdf,
			Pages:                    paperlessngx.Int(1),
			RotatePages:              paperlessngx.Bool(true),
			RotatePagesThreshold:     paperlessngx.Float(0),
			SkipArchiveFile:          paperlessngx.ConfigPatchParamsSkipArchiveFileNever,
			UnpaperClean:             paperlessngx.ConfigPatchParamsUnpaperCleanClean,
			UserArgs:                 map[string]interface{}{},
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

func TestConfigGetByID(t *testing.T) {
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
	_, err := client.Config.GetByID(context.TODO(), 0)
	if err != nil {
		var apierr *paperlessngx.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
