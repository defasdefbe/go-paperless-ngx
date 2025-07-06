// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package paperlessngx

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/defasdefbe/go-paperless-ngx/internal/apiform"
	"github.com/defasdefbe/go-paperless-ngx/internal/apijson"
	"github.com/defasdefbe/go-paperless-ngx/internal/requestconfig"
	"github.com/defasdefbe/go-paperless-ngx/option"
	"github.com/defasdefbe/go-paperless-ngx/packages/param"
	"github.com/defasdefbe/go-paperless-ngx/packages/respjson"
)

// ConfigService contains methods and other services that help with interacting
// with the paperless-ngx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConfigService] method instead.
type ConfigService struct {
	Options []option.RequestOption
}

// NewConfigService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewConfigService(opts ...option.RequestOption) (r ConfigService) {
	r = ConfigService{}
	r.Options = opts
	return
}

// Get the application configuration
func (r *ConfigService) Get(ctx context.Context, opts ...option.RequestOption) (res *[]ApplicationConfiguration, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/config/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *ConfigService) Update(ctx context.Context, id int64, body ConfigUpdateParams, opts ...option.RequestOption) (res *ApplicationConfiguration, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/config/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

func (r *ConfigService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("api/config/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

func (r *ConfigService) Patch(ctx context.Context, id int64, body ConfigPatchParams, opts ...option.RequestOption) (res *ApplicationConfiguration, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/config/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

func (r *ConfigService) GetByID(ctx context.Context, id int64, opts ...option.RequestOption) (res *ApplicationConfiguration, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/config/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ApplicationConfiguration struct {
	ID                       int64   `json:"id,required"`
	BarcodeTagMapping        any     `json:"barcode_tag_mapping,required"`
	UserArgs                 any     `json:"user_args,required"`
	AppLogo                  string  `json:"app_logo,nullable" format:"uri"`
	AppTitle                 string  `json:"app_title,nullable"`
	BarcodeAsnPrefix         string  `json:"barcode_asn_prefix,nullable"`
	BarcodeDpi               int64   `json:"barcode_dpi,nullable"`
	BarcodeEnableAsn         bool    `json:"barcode_enable_asn,nullable"`
	BarcodeEnableTag         bool    `json:"barcode_enable_tag,nullable"`
	BarcodeEnableTiffSupport bool    `json:"barcode_enable_tiff_support,nullable"`
	BarcodeMaxPages          int64   `json:"barcode_max_pages,nullable"`
	BarcodeRetainSplitPages  bool    `json:"barcode_retain_split_pages,nullable"`
	BarcodeString            string  `json:"barcode_string,nullable"`
	BarcodeUpscale           float64 `json:"barcode_upscale,nullable"`
	BarcodesEnabled          bool    `json:"barcodes_enabled,nullable"`
	// Any of "LeaveColorUnchanged", "RGB", "UseDeviceIndependentColor", "Gray",
	// "CMYK", "".
	ColorConversionStrategy ApplicationConfigurationColorConversionStrategy `json:"color_conversion_strategy,nullable"`
	Deskew                  bool                                            `json:"deskew,nullable"`
	ImageDpi                int64                                           `json:"image_dpi,nullable"`
	Language                string                                          `json:"language,nullable"`
	MaxImagePixels          float64                                         `json:"max_image_pixels,nullable"`
	// Any of "skip", "redo", "force", "skip_noarchive", "".
	Mode ApplicationConfigurationMode `json:"mode,nullable"`
	// Any of "pdf", "pdfa", "pdfa-1", "pdfa-2", "pdfa-3", "".
	OutputType           ApplicationConfigurationOutputType `json:"output_type,nullable"`
	Pages                int64                              `json:"pages,nullable"`
	RotatePages          bool                               `json:"rotate_pages,nullable"`
	RotatePagesThreshold float64                            `json:"rotate_pages_threshold,nullable"`
	// Any of "never", "with_text", "always", "".
	SkipArchiveFile ApplicationConfigurationSkipArchiveFile `json:"skip_archive_file,nullable"`
	// Any of "clean", "clean-final", "none", "".
	UnpaperClean ApplicationConfigurationUnpaperClean `json:"unpaper_clean,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                       respjson.Field
		BarcodeTagMapping        respjson.Field
		UserArgs                 respjson.Field
		AppLogo                  respjson.Field
		AppTitle                 respjson.Field
		BarcodeAsnPrefix         respjson.Field
		BarcodeDpi               respjson.Field
		BarcodeEnableAsn         respjson.Field
		BarcodeEnableTag         respjson.Field
		BarcodeEnableTiffSupport respjson.Field
		BarcodeMaxPages          respjson.Field
		BarcodeRetainSplitPages  respjson.Field
		BarcodeString            respjson.Field
		BarcodeUpscale           respjson.Field
		BarcodesEnabled          respjson.Field
		ColorConversionStrategy  respjson.Field
		Deskew                   respjson.Field
		ImageDpi                 respjson.Field
		Language                 respjson.Field
		MaxImagePixels           respjson.Field
		Mode                     respjson.Field
		OutputType               respjson.Field
		Pages                    respjson.Field
		RotatePages              respjson.Field
		RotatePagesThreshold     respjson.Field
		SkipArchiveFile          respjson.Field
		UnpaperClean             respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ApplicationConfiguration) RawJSON() string { return r.JSON.raw }
func (r *ApplicationConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ApplicationConfigurationColorConversionStrategy string

const (
	ApplicationConfigurationColorConversionStrategyLeaveColorUnchanged       ApplicationConfigurationColorConversionStrategy = "LeaveColorUnchanged"
	ApplicationConfigurationColorConversionStrategyRgb                       ApplicationConfigurationColorConversionStrategy = "RGB"
	ApplicationConfigurationColorConversionStrategyUseDeviceIndependentColor ApplicationConfigurationColorConversionStrategy = "UseDeviceIndependentColor"
	ApplicationConfigurationColorConversionStrategyGray                      ApplicationConfigurationColorConversionStrategy = "Gray"
	ApplicationConfigurationColorConversionStrategyCmyk                      ApplicationConfigurationColorConversionStrategy = "CMYK"
	ApplicationConfigurationColorConversionStrategyEmpty                     ApplicationConfigurationColorConversionStrategy = ""
)

type ApplicationConfigurationMode string

const (
	ApplicationConfigurationModeSkip          ApplicationConfigurationMode = "skip"
	ApplicationConfigurationModeRedo          ApplicationConfigurationMode = "redo"
	ApplicationConfigurationModeForce         ApplicationConfigurationMode = "force"
	ApplicationConfigurationModeSkipNoarchive ApplicationConfigurationMode = "skip_noarchive"
	ApplicationConfigurationModeEmpty         ApplicationConfigurationMode = ""
)

type ApplicationConfigurationOutputType string

const (
	ApplicationConfigurationOutputTypePdf   ApplicationConfigurationOutputType = "pdf"
	ApplicationConfigurationOutputTypePdfa  ApplicationConfigurationOutputType = "pdfa"
	ApplicationConfigurationOutputTypePdfa1 ApplicationConfigurationOutputType = "pdfa-1"
	ApplicationConfigurationOutputTypePdfa2 ApplicationConfigurationOutputType = "pdfa-2"
	ApplicationConfigurationOutputTypePdfa3 ApplicationConfigurationOutputType = "pdfa-3"
	ApplicationConfigurationOutputTypeEmpty ApplicationConfigurationOutputType = ""
)

type ApplicationConfigurationSkipArchiveFile string

const (
	ApplicationConfigurationSkipArchiveFileNever    ApplicationConfigurationSkipArchiveFile = "never"
	ApplicationConfigurationSkipArchiveFileWithText ApplicationConfigurationSkipArchiveFile = "with_text"
	ApplicationConfigurationSkipArchiveFileAlways   ApplicationConfigurationSkipArchiveFile = "always"
	ApplicationConfigurationSkipArchiveFileEmpty    ApplicationConfigurationSkipArchiveFile = ""
)

type ApplicationConfigurationUnpaperClean string

const (
	ApplicationConfigurationUnpaperCleanClean      ApplicationConfigurationUnpaperClean = "clean"
	ApplicationConfigurationUnpaperCleanCleanFinal ApplicationConfigurationUnpaperClean = "clean-final"
	ApplicationConfigurationUnpaperCleanNone       ApplicationConfigurationUnpaperClean = "none"
	ApplicationConfigurationUnpaperCleanEmpty      ApplicationConfigurationUnpaperClean = ""
)

type ConfigUpdateParams struct {
	BarcodeTagMapping        any                `json:"barcode_tag_mapping,omitzero,required"`
	UserArgs                 any                `json:"user_args,omitzero,required"`
	AppTitle                 param.Opt[string]  `json:"app_title,omitzero"`
	BarcodeAsnPrefix         param.Opt[string]  `json:"barcode_asn_prefix,omitzero"`
	BarcodeDpi               param.Opt[int64]   `json:"barcode_dpi,omitzero"`
	BarcodeEnableAsn         param.Opt[bool]    `json:"barcode_enable_asn,omitzero"`
	BarcodeEnableTag         param.Opt[bool]    `json:"barcode_enable_tag,omitzero"`
	BarcodeEnableTiffSupport param.Opt[bool]    `json:"barcode_enable_tiff_support,omitzero"`
	BarcodeMaxPages          param.Opt[int64]   `json:"barcode_max_pages,omitzero"`
	BarcodeRetainSplitPages  param.Opt[bool]    `json:"barcode_retain_split_pages,omitzero"`
	BarcodeString            param.Opt[string]  `json:"barcode_string,omitzero"`
	BarcodeUpscale           param.Opt[float64] `json:"barcode_upscale,omitzero"`
	BarcodesEnabled          param.Opt[bool]    `json:"barcodes_enabled,omitzero"`
	Deskew                   param.Opt[bool]    `json:"deskew,omitzero"`
	ImageDpi                 param.Opt[int64]   `json:"image_dpi,omitzero"`
	Language                 param.Opt[string]  `json:"language,omitzero"`
	MaxImagePixels           param.Opt[float64] `json:"max_image_pixels,omitzero"`
	Pages                    param.Opt[int64]   `json:"pages,omitzero"`
	RotatePages              param.Opt[bool]    `json:"rotate_pages,omitzero"`
	RotatePagesThreshold     param.Opt[float64] `json:"rotate_pages_threshold,omitzero"`
	AppLogo                  io.Reader          `json:"app_logo,omitzero" format:"binary"`
	// Any of "LeaveColorUnchanged", "RGB", "UseDeviceIndependentColor", "Gray",
	// "CMYK", "".
	ColorConversionStrategy ConfigUpdateParamsColorConversionStrategy `json:"color_conversion_strategy,omitzero"`
	// Any of "skip", "redo", "force", "skip_noarchive", "".
	Mode ConfigUpdateParamsMode `json:"mode,omitzero"`
	// Any of "pdf", "pdfa", "pdfa-1", "pdfa-2", "pdfa-3", "".
	OutputType ConfigUpdateParamsOutputType `json:"output_type,omitzero"`
	// Any of "never", "with_text", "always", "".
	SkipArchiveFile ConfigUpdateParamsSkipArchiveFile `json:"skip_archive_file,omitzero"`
	// Any of "clean", "clean-final", "none", "".
	UnpaperClean ConfigUpdateParamsUnpaperClean `json:"unpaper_clean,omitzero"`
	paramObj
}

func (r ConfigUpdateParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

type ConfigUpdateParamsColorConversionStrategy string

const (
	ConfigUpdateParamsColorConversionStrategyLeaveColorUnchanged       ConfigUpdateParamsColorConversionStrategy = "LeaveColorUnchanged"
	ConfigUpdateParamsColorConversionStrategyRgb                       ConfigUpdateParamsColorConversionStrategy = "RGB"
	ConfigUpdateParamsColorConversionStrategyUseDeviceIndependentColor ConfigUpdateParamsColorConversionStrategy = "UseDeviceIndependentColor"
	ConfigUpdateParamsColorConversionStrategyGray                      ConfigUpdateParamsColorConversionStrategy = "Gray"
	ConfigUpdateParamsColorConversionStrategyCmyk                      ConfigUpdateParamsColorConversionStrategy = "CMYK"
	ConfigUpdateParamsColorConversionStrategyEmpty                     ConfigUpdateParamsColorConversionStrategy = ""
)

type ConfigUpdateParamsMode string

const (
	ConfigUpdateParamsModeSkip          ConfigUpdateParamsMode = "skip"
	ConfigUpdateParamsModeRedo          ConfigUpdateParamsMode = "redo"
	ConfigUpdateParamsModeForce         ConfigUpdateParamsMode = "force"
	ConfigUpdateParamsModeSkipNoarchive ConfigUpdateParamsMode = "skip_noarchive"
	ConfigUpdateParamsModeEmpty         ConfigUpdateParamsMode = ""
)

type ConfigUpdateParamsOutputType string

const (
	ConfigUpdateParamsOutputTypePdf   ConfigUpdateParamsOutputType = "pdf"
	ConfigUpdateParamsOutputTypePdfa  ConfigUpdateParamsOutputType = "pdfa"
	ConfigUpdateParamsOutputTypePdfa1 ConfigUpdateParamsOutputType = "pdfa-1"
	ConfigUpdateParamsOutputTypePdfa2 ConfigUpdateParamsOutputType = "pdfa-2"
	ConfigUpdateParamsOutputTypePdfa3 ConfigUpdateParamsOutputType = "pdfa-3"
	ConfigUpdateParamsOutputTypeEmpty ConfigUpdateParamsOutputType = ""
)

type ConfigUpdateParamsSkipArchiveFile string

const (
	ConfigUpdateParamsSkipArchiveFileNever    ConfigUpdateParamsSkipArchiveFile = "never"
	ConfigUpdateParamsSkipArchiveFileWithText ConfigUpdateParamsSkipArchiveFile = "with_text"
	ConfigUpdateParamsSkipArchiveFileAlways   ConfigUpdateParamsSkipArchiveFile = "always"
	ConfigUpdateParamsSkipArchiveFileEmpty    ConfigUpdateParamsSkipArchiveFile = ""
)

type ConfigUpdateParamsUnpaperClean string

const (
	ConfigUpdateParamsUnpaperCleanClean      ConfigUpdateParamsUnpaperClean = "clean"
	ConfigUpdateParamsUnpaperCleanCleanFinal ConfigUpdateParamsUnpaperClean = "clean-final"
	ConfigUpdateParamsUnpaperCleanNone       ConfigUpdateParamsUnpaperClean = "none"
	ConfigUpdateParamsUnpaperCleanEmpty      ConfigUpdateParamsUnpaperClean = ""
)

type ConfigPatchParams struct {
	AppTitle                 param.Opt[string]  `json:"app_title,omitzero"`
	BarcodeAsnPrefix         param.Opt[string]  `json:"barcode_asn_prefix,omitzero"`
	BarcodeDpi               param.Opt[int64]   `json:"barcode_dpi,omitzero"`
	BarcodeEnableAsn         param.Opt[bool]    `json:"barcode_enable_asn,omitzero"`
	BarcodeEnableTag         param.Opt[bool]    `json:"barcode_enable_tag,omitzero"`
	BarcodeEnableTiffSupport param.Opt[bool]    `json:"barcode_enable_tiff_support,omitzero"`
	BarcodeMaxPages          param.Opt[int64]   `json:"barcode_max_pages,omitzero"`
	BarcodeRetainSplitPages  param.Opt[bool]    `json:"barcode_retain_split_pages,omitzero"`
	BarcodeString            param.Opt[string]  `json:"barcode_string,omitzero"`
	BarcodeUpscale           param.Opt[float64] `json:"barcode_upscale,omitzero"`
	BarcodesEnabled          param.Opt[bool]    `json:"barcodes_enabled,omitzero"`
	Deskew                   param.Opt[bool]    `json:"deskew,omitzero"`
	ImageDpi                 param.Opt[int64]   `json:"image_dpi,omitzero"`
	Language                 param.Opt[string]  `json:"language,omitzero"`
	MaxImagePixels           param.Opt[float64] `json:"max_image_pixels,omitzero"`
	Pages                    param.Opt[int64]   `json:"pages,omitzero"`
	RotatePages              param.Opt[bool]    `json:"rotate_pages,omitzero"`
	RotatePagesThreshold     param.Opt[float64] `json:"rotate_pages_threshold,omitzero"`
	AppLogo                  io.Reader          `json:"app_logo,omitzero" format:"binary"`
	// Any of "LeaveColorUnchanged", "RGB", "UseDeviceIndependentColor", "Gray",
	// "CMYK", "".
	ColorConversionStrategy ConfigPatchParamsColorConversionStrategy `json:"color_conversion_strategy,omitzero"`
	// Any of "skip", "redo", "force", "skip_noarchive", "".
	Mode ConfigPatchParamsMode `json:"mode,omitzero"`
	// Any of "pdf", "pdfa", "pdfa-1", "pdfa-2", "pdfa-3", "".
	OutputType ConfigPatchParamsOutputType `json:"output_type,omitzero"`
	// Any of "never", "with_text", "always", "".
	SkipArchiveFile ConfigPatchParamsSkipArchiveFile `json:"skip_archive_file,omitzero"`
	// Any of "clean", "clean-final", "none", "".
	UnpaperClean      ConfigPatchParamsUnpaperClean `json:"unpaper_clean,omitzero"`
	BarcodeTagMapping any                           `json:"barcode_tag_mapping,omitzero"`
	UserArgs          any                           `json:"user_args,omitzero"`
	paramObj
}

func (r ConfigPatchParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

type ConfigPatchParamsColorConversionStrategy string

const (
	ConfigPatchParamsColorConversionStrategyLeaveColorUnchanged       ConfigPatchParamsColorConversionStrategy = "LeaveColorUnchanged"
	ConfigPatchParamsColorConversionStrategyRgb                       ConfigPatchParamsColorConversionStrategy = "RGB"
	ConfigPatchParamsColorConversionStrategyUseDeviceIndependentColor ConfigPatchParamsColorConversionStrategy = "UseDeviceIndependentColor"
	ConfigPatchParamsColorConversionStrategyGray                      ConfigPatchParamsColorConversionStrategy = "Gray"
	ConfigPatchParamsColorConversionStrategyCmyk                      ConfigPatchParamsColorConversionStrategy = "CMYK"
	ConfigPatchParamsColorConversionStrategyEmpty                     ConfigPatchParamsColorConversionStrategy = ""
)

type ConfigPatchParamsMode string

const (
	ConfigPatchParamsModeSkip          ConfigPatchParamsMode = "skip"
	ConfigPatchParamsModeRedo          ConfigPatchParamsMode = "redo"
	ConfigPatchParamsModeForce         ConfigPatchParamsMode = "force"
	ConfigPatchParamsModeSkipNoarchive ConfigPatchParamsMode = "skip_noarchive"
	ConfigPatchParamsModeEmpty         ConfigPatchParamsMode = ""
)

type ConfigPatchParamsOutputType string

const (
	ConfigPatchParamsOutputTypePdf   ConfigPatchParamsOutputType = "pdf"
	ConfigPatchParamsOutputTypePdfa  ConfigPatchParamsOutputType = "pdfa"
	ConfigPatchParamsOutputTypePdfa1 ConfigPatchParamsOutputType = "pdfa-1"
	ConfigPatchParamsOutputTypePdfa2 ConfigPatchParamsOutputType = "pdfa-2"
	ConfigPatchParamsOutputTypePdfa3 ConfigPatchParamsOutputType = "pdfa-3"
	ConfigPatchParamsOutputTypeEmpty ConfigPatchParamsOutputType = ""
)

type ConfigPatchParamsSkipArchiveFile string

const (
	ConfigPatchParamsSkipArchiveFileNever    ConfigPatchParamsSkipArchiveFile = "never"
	ConfigPatchParamsSkipArchiveFileWithText ConfigPatchParamsSkipArchiveFile = "with_text"
	ConfigPatchParamsSkipArchiveFileAlways   ConfigPatchParamsSkipArchiveFile = "always"
	ConfigPatchParamsSkipArchiveFileEmpty    ConfigPatchParamsSkipArchiveFile = ""
)

type ConfigPatchParamsUnpaperClean string

const (
	ConfigPatchParamsUnpaperCleanClean      ConfigPatchParamsUnpaperClean = "clean"
	ConfigPatchParamsUnpaperCleanCleanFinal ConfigPatchParamsUnpaperClean = "clean-final"
	ConfigPatchParamsUnpaperCleanNone       ConfigPatchParamsUnpaperClean = "none"
	ConfigPatchParamsUnpaperCleanEmpty      ConfigPatchParamsUnpaperClean = ""
)
