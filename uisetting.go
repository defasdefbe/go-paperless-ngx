// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package paperlessnix

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/paperless-nix-go/internal/apijson"
	"github.com/stainless-sdks/paperless-nix-go/internal/requestconfig"
	"github.com/stainless-sdks/paperless-nix-go/option"
	"github.com/stainless-sdks/paperless-nix-go/packages/param"
	"github.com/stainless-sdks/paperless-nix-go/packages/respjson"
)

// UiSettingService contains methods and other services that help with interacting
// with the paperless-nix API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUiSettingService] method instead.
type UiSettingService struct {
	Options []option.RequestOption
}

// NewUiSettingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUiSettingService(opts ...option.RequestOption) (r UiSettingService) {
	r = UiSettingService{}
	r.Options = opts
	return
}

func (r *UiSettingService) New(ctx context.Context, body UiSettingNewParams, opts ...option.RequestOption) (res *UiSettingsView, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/ui_settings/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *UiSettingService) Get(ctx context.Context, opts ...option.RequestOption) (res *UiSettingsView, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/ui_settings/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type UiSettingsView struct {
	ID       int64          `json:"id,required"`
	Settings map[string]any `json:"settings,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Settings    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UiSettingsView) RawJSON() string { return r.JSON.raw }
func (r *UiSettingsView) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UiSettingNewParams struct {
	Settings map[string]any `json:"settings,omitzero"`
	paramObj
}

func (r UiSettingNewParams) MarshalJSON() (data []byte, err error) {
	type shadow UiSettingNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UiSettingNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
