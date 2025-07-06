// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package paperlessngx

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/defasdefbe/go-paperless-ngx/internal/apijson"
	"github.com/defasdefbe/go-paperless-ngx/internal/apiquery"
	"github.com/defasdefbe/go-paperless-ngx/internal/requestconfig"
	"github.com/defasdefbe/go-paperless-ngx/option"
	"github.com/defasdefbe/go-paperless-ngx/packages/param"
	"github.com/defasdefbe/go-paperless-ngx/packages/respjson"
)

// UserService contains methods and other services that help with interacting with
// the paperless-ngx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	Options []option.RequestOption
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r UserService) {
	r = UserService{}
	r.Options = opts
	return
}

func (r *UserService) New(ctx context.Context, body UserNewParams, opts ...option.RequestOption) (res *User, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/users/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *UserService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *User, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/users/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *UserService) Update(ctx context.Context, id int64, body UserUpdateParams, opts ...option.RequestOption) (res *User, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/users/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

func (r *UserService) List(ctx context.Context, query UserListParams, opts ...option.RequestOption) (res *UserListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/users/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

func (r *UserService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("api/users/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

func (r *UserService) DeactivateTotp(ctx context.Context, id int64, opts ...option.RequestOption) (res *bool, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/users/%v/deactivate_totp/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type User struct {
	ID                   int64    `json:"id,required"`
	InheritedPermissions []string `json:"inherited_permissions,required"`
	IsMfaEnabled         bool     `json:"is_mfa_enabled,required"`
	// Required. 150 characters or fewer. Letters, digits and @/./+/-/\_ only.
	Username   string    `json:"username,required"`
	DateJoined time.Time `json:"date_joined" format:"date-time"`
	Email      string    `json:"email" format:"email"`
	FirstName  string    `json:"first_name"`
	// The groups this user belongs to. A user will get all permissions granted to each
	// of their groups.
	Groups []int64 `json:"groups"`
	// Designates whether this user should be treated as active. Unselect this instead
	// of deleting accounts.
	IsActive bool `json:"is_active"`
	// Designates whether the user can log into this admin site.
	IsStaff bool `json:"is_staff"`
	// Designates that this user has all permissions without explicitly assigning them.
	IsSuperuser     bool     `json:"is_superuser"`
	LastName        string   `json:"last_name"`
	Password        string   `json:"password"`
	UserPermissions []string `json:"user_permissions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		InheritedPermissions respjson.Field
		IsMfaEnabled         respjson.Field
		Username             respjson.Field
		DateJoined           respjson.Field
		Email                respjson.Field
		FirstName            respjson.Field
		Groups               respjson.Field
		IsActive             respjson.Field
		IsStaff              respjson.Field
		IsSuperuser          respjson.Field
		LastName             respjson.Field
		Password             respjson.Field
		UserPermissions      respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r User) RawJSON() string { return r.JSON.raw }
func (r *User) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Username is required.
type UserRequestParam struct {
	// Required. 150 characters or fewer. Letters, digits and @/./+/-/\_ only.
	Username   string               `json:"username,required"`
	DateJoined param.Opt[time.Time] `json:"date_joined,omitzero" format:"date-time"`
	Email      param.Opt[string]    `json:"email,omitzero" format:"email"`
	FirstName  param.Opt[string]    `json:"first_name,omitzero"`
	// Designates whether this user should be treated as active. Unselect this instead
	// of deleting accounts.
	IsActive param.Opt[bool] `json:"is_active,omitzero"`
	// Designates whether the user can log into this admin site.
	IsStaff param.Opt[bool] `json:"is_staff,omitzero"`
	// Designates that this user has all permissions without explicitly assigning them.
	IsSuperuser param.Opt[bool]   `json:"is_superuser,omitzero"`
	LastName    param.Opt[string] `json:"last_name,omitzero"`
	Password    param.Opt[string] `json:"password,omitzero"`
	// The groups this user belongs to. A user will get all permissions granted to each
	// of their groups.
	Groups          []int64  `json:"groups,omitzero"`
	UserPermissions []string `json:"user_permissions,omitzero"`
	paramObj
}

func (r UserRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow UserRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserListResponse struct {
	Count    int64  `json:"count,required"`
	Results  []User `json:"results,required"`
	All      []any  `json:"all"`
	Next     string `json:"next,nullable" format:"uri"`
	Previous string `json:"previous,nullable" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Results     respjson.Field
		All         respjson.Field
		Next        respjson.Field
		Previous    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserListResponse) RawJSON() string { return r.JSON.raw }
func (r *UserListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserNewParams struct {
	UserRequest UserRequestParam
	paramObj
}

func (r UserNewParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.UserRequest)
}
func (r *UserNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.UserRequest)
}

type UserUpdateParams struct {
	DateJoined param.Opt[time.Time] `json:"date_joined,omitzero" format:"date-time"`
	Email      param.Opt[string]    `json:"email,omitzero" format:"email"`
	FirstName  param.Opt[string]    `json:"first_name,omitzero"`
	// Designates whether this user should be treated as active. Unselect this instead
	// of deleting accounts.
	IsActive param.Opt[bool] `json:"is_active,omitzero"`
	// Designates whether the user can log into this admin site.
	IsStaff param.Opt[bool] `json:"is_staff,omitzero"`
	// Designates that this user has all permissions without explicitly assigning them.
	IsSuperuser param.Opt[bool]   `json:"is_superuser,omitzero"`
	LastName    param.Opt[string] `json:"last_name,omitzero"`
	Password    param.Opt[string] `json:"password,omitzero"`
	// Required. 150 characters or fewer. Letters, digits and @/./+/-/\_ only.
	Username param.Opt[string] `json:"username,omitzero"`
	// The groups this user belongs to. A user will get all permissions granted to each
	// of their groups.
	Groups          []int64  `json:"groups,omitzero"`
	UserPermissions []string `json:"user_permissions,omitzero"`
	paramObj
}

func (r UserUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow UserUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserListParams struct {
	// Which field to use when ordering the results.
	Ordering param.Opt[string] `query:"ordering,omitzero" json:"-"`
	// A page number within the paginated result set.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Number of results to return per page.
	PageSize            param.Opt[int64]  `query:"page_size,omitzero" json:"-"`
	UsernameIcontains   param.Opt[string] `query:"username__icontains,omitzero" json:"-"`
	UsernameIendswith   param.Opt[string] `query:"username__iendswith,omitzero" json:"-"`
	UsernameIexact      param.Opt[string] `query:"username__iexact,omitzero" json:"-"`
	UsernameIstartswith param.Opt[string] `query:"username__istartswith,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UserListParams]'s query parameters as `url.Values`.
func (r UserListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
