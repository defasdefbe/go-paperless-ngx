// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package paperlessngx

import (
	"context"
	"net/http"

	"github.com/defasdefbe/go-paperless-ngx/internal/apijson"
	"github.com/defasdefbe/go-paperless-ngx/internal/requestconfig"
	"github.com/defasdefbe/go-paperless-ngx/option"
	"github.com/defasdefbe/go-paperless-ngx/packages/param"
)

// ProfileTotpService contains methods and other services that help with
// interacting with the paperless-ngx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProfileTotpService] method instead.
type ProfileTotpService struct {
	Options []option.RequestOption
}

// NewProfileTotpService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProfileTotpService(opts ...option.RequestOption) (r ProfileTotpService) {
	r = ProfileTotpService{}
	r.Options = opts
	return
}

// Deactivates the TOTP authenticator
func (r *ProfileTotpService) Deactivate(ctx context.Context, opts ...option.RequestOption) (res *bool, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/profile/totp/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Generates a new TOTP secret and returns the URL and SVG
func (r *ProfileTotpService) Generate(ctx context.Context, opts ...option.RequestOption) (res *ProfileTotpGenerateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/profile/totp/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Validates a TOTP code and activates the TOTP authenticator
func (r *ProfileTotpService) Validate(ctx context.Context, body ProfileTotpValidateParams, opts ...option.RequestOption) (res *ProfileTotpValidateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/profile/totp/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ProfileTotpGenerateResponse map[string]any

type ProfileTotpValidateResponse map[string]any

type ProfileTotpValidateParams struct {
	Code   string `json:"code,required"`
	Secret string `json:"secret,required"`
	paramObj
}

func (r ProfileTotpValidateParams) MarshalJSON() (data []byte, err error) {
	type shadow ProfileTotpValidateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileTotpValidateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
