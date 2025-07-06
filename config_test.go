// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package paperlessnix_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/stainless-sdks/paperless-nix-go"
	"github.com/stainless-sdks/paperless-nix-go/internal/testutil"
	"github.com/stainless-sdks/paperless-nix-go/option"
)

func TestConfigGet(t *testing.T) {
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
	_, err := client.Config.Get(context.TODO())
	if err != nil {
		var apierr *paperlessnix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConfigUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Config.Update(
		context.TODO(),
		0,
		paperlessnix.ConfigUpdateParams{
			BarcodeTagMapping:        map[string]interface{}{},
			UserArgs:                 map[string]interface{}{},
			AppLogo:                  io.Reader(bytes.NewBuffer([]byte("some file contents"))),
			AppTitle:                 paperlessnix.String("app_title"),
			BarcodeAsnPrefix:         paperlessnix.String("barcode_asn_prefix"),
			BarcodeDpi:               paperlessnix.Int(1),
			BarcodeEnableAsn:         paperlessnix.Bool(true),
			BarcodeEnableTag:         paperlessnix.Bool(true),
			BarcodeEnableTiffSupport: paperlessnix.Bool(true),
			BarcodeMaxPages:          paperlessnix.Int(1),
			BarcodeRetainSplitPages:  paperlessnix.Bool(true),
			BarcodeString:            paperlessnix.String("barcode_string"),
			BarcodeUpscale:           paperlessnix.Float(1),
			BarcodesEnabled:          paperlessnix.Bool(true),
			ColorConversionStrategy:  paperlessnix.ConfigUpdateParamsColorConversionStrategyLeaveColorUnchanged,
			Deskew:                   paperlessnix.Bool(true),
			ImageDpi:                 paperlessnix.Int(1),
			Language:                 paperlessnix.String("language"),
			MaxImagePixels:           paperlessnix.Float(0),
			Mode:                     paperlessnix.ConfigUpdateParamsModeSkip,
			OutputType:               paperlessnix.ConfigUpdateParamsOutputTypePdf,
			Pages:                    paperlessnix.Int(1),
			RotatePages:              paperlessnix.Bool(true),
			RotatePagesThreshold:     paperlessnix.Float(0),
			SkipArchiveFile:          paperlessnix.ConfigUpdateParamsSkipArchiveFileNever,
			UnpaperClean:             paperlessnix.ConfigUpdateParamsUnpaperCleanClean,
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

func TestConfigDelete(t *testing.T) {
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
	err := client.Config.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *paperlessnix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConfigPatchWithOptionalParams(t *testing.T) {
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
	_, err := client.Config.Patch(
		context.TODO(),
		0,
		paperlessnix.ConfigPatchParams{
			AppLogo:                  io.Reader(bytes.NewBuffer([]byte("some file contents"))),
			AppTitle:                 paperlessnix.String("app_title"),
			BarcodeAsnPrefix:         paperlessnix.String("barcode_asn_prefix"),
			BarcodeDpi:               paperlessnix.Int(1),
			BarcodeEnableAsn:         paperlessnix.Bool(true),
			BarcodeEnableTag:         paperlessnix.Bool(true),
			BarcodeEnableTiffSupport: paperlessnix.Bool(true),
			BarcodeMaxPages:          paperlessnix.Int(1),
			BarcodeRetainSplitPages:  paperlessnix.Bool(true),
			BarcodeString:            paperlessnix.String("barcode_string"),
			BarcodeTagMapping:        map[string]interface{}{},
			BarcodeUpscale:           paperlessnix.Float(1),
			BarcodesEnabled:          paperlessnix.Bool(true),
			ColorConversionStrategy:  paperlessnix.ConfigPatchParamsColorConversionStrategyLeaveColorUnchanged,
			Deskew:                   paperlessnix.Bool(true),
			ImageDpi:                 paperlessnix.Int(1),
			Language:                 paperlessnix.String("language"),
			MaxImagePixels:           paperlessnix.Float(0),
			Mode:                     paperlessnix.ConfigPatchParamsModeSkip,
			OutputType:               paperlessnix.ConfigPatchParamsOutputTypePdf,
			Pages:                    paperlessnix.Int(1),
			RotatePages:              paperlessnix.Bool(true),
			RotatePagesThreshold:     paperlessnix.Float(0),
			SkipArchiveFile:          paperlessnix.ConfigPatchParamsSkipArchiveFileNever,
			UnpaperClean:             paperlessnix.ConfigPatchParamsUnpaperCleanClean,
			UserArgs:                 map[string]interface{}{},
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

func TestConfigGetByID(t *testing.T) {
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
	_, err := client.Config.GetByID(context.TODO(), 0)
	if err != nil {
		var apierr *paperlessnix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
