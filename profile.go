// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package paperlessngx

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/paperless-nix-go/internal/apijson"
	"github.com/stainless-sdks/paperless-nix-go/internal/requestconfig"
	"github.com/stainless-sdks/paperless-nix-go/option"
	"github.com/stainless-sdks/paperless-nix-go/packages/param"
	"github.com/stainless-sdks/paperless-nix-go/packages/respjson"
)

// ProfileService contains methods and other services that help with interacting
// with the paperless-ngx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProfileService] method instead.
type ProfileService struct {
	Options []option.RequestOption
	Totp    ProfileTotpService
}

// NewProfileService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProfileService(opts ...option.RequestOption) (r ProfileService) {
	r = ProfileService{}
	r.Options = opts
	r.Totp = NewProfileTotpService(opts...)
	return
}

// User profile view, only available when logged in
func (r *ProfileService) Get(ctx context.Context, opts ...option.RequestOption) (res *Profile, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/profile/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// User profile view, only available when logged in
func (r *ProfileService) Update(ctx context.Context, body ProfileUpdateParams, opts ...option.RequestOption) (res *Profile, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/profile/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Disconnects a social account provider from the user account
func (r *ProfileService) DisconnectSocialAccount(ctx context.Context, body ProfileDisconnectSocialAccountParams, opts ...option.RequestOption) (res *int64, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/profile/disconnect_social_account/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Generates (or re-generates) an auth token, requires a logged in user unlike the
// default DRF endpoint
func (r *ProfileService) GenerateAuthToken(ctx context.Context, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/profile/generate_auth_token/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// List of social account providers
func (r *ProfileService) ListSocialAccountProviders(ctx context.Context, opts ...option.RequestOption) (res *ProfileListSocialAccountProvidersResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/profile/social_account_providers/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Profile struct {
	AuthToken         string                 `json:"auth_token,required"`
	HasUsablePassword bool                   `json:"has_usable_password,required"`
	IsMfaEnabled      bool                   `json:"is_mfa_enabled,required"`
	SocialAccounts    []ProfileSocialAccount `json:"social_accounts,required"`
	Email             string                 `json:"email" format:"email"`
	FirstName         string                 `json:"first_name"`
	LastName          string                 `json:"last_name"`
	Password          string                 `json:"password"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AuthToken         respjson.Field
		HasUsablePassword respjson.Field
		IsMfaEnabled      respjson.Field
		SocialAccounts    respjson.Field
		Email             respjson.Field
		FirstName         respjson.Field
		LastName          respjson.Field
		Password          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Profile) RawJSON() string { return r.JSON.raw }
func (r *Profile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileSocialAccount struct {
	ID       int64  `json:"id,required"`
	Name     string `json:"name,required"`
	Provider string `json:"provider,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Provider    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileSocialAccount) RawJSON() string { return r.JSON.raw }
func (r *ProfileSocialAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileListSocialAccountProvidersResponse map[string]any

type ProfileUpdateParams struct {
	Email     param.Opt[string] `json:"email,omitzero" format:"email"`
	FirstName param.Opt[string] `json:"first_name,omitzero"`
	LastName  param.Opt[string] `json:"last_name,omitzero"`
	Password  param.Opt[string] `json:"password,omitzero"`
	paramObj
}

func (r ProfileUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ProfileUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileDisconnectSocialAccountParams struct {
	ID int64 `json:"id,required"`
	paramObj
}

func (r ProfileDisconnectSocialAccountParams) MarshalJSON() (data []byte, err error) {
	type shadow ProfileDisconnectSocialAccountParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileDisconnectSocialAccountParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
