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
	shimjson "github.com/defasdefbe/go-paperless-ngx/internal/encoding/json"
	"github.com/defasdefbe/go-paperless-ngx/internal/requestconfig"
	"github.com/defasdefbe/go-paperless-ngx/option"
	"github.com/defasdefbe/go-paperless-ngx/packages/param"
	"github.com/defasdefbe/go-paperless-ngx/packages/respjson"
)

// MailAccountService contains methods and other services that help with
// interacting with the paperless-ngx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMailAccountService] method instead.
type MailAccountService struct {
	Options []option.RequestOption
}

// NewMailAccountService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMailAccountService(opts ...option.RequestOption) (r MailAccountService) {
	r = MailAccountService{}
	r.Options = opts
	return
}

func (r *MailAccountService) New(ctx context.Context, body MailAccountNewParams, opts ...option.RequestOption) (res *MailAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/mail_accounts/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *MailAccountService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *MailAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/mail_accounts/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *MailAccountService) Update(ctx context.Context, id int64, body MailAccountUpdateParams, opts ...option.RequestOption) (res *MailAccount, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/mail_accounts/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

func (r *MailAccountService) List(ctx context.Context, query MailAccountListParams, opts ...option.RequestOption) (res *MailAccountListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/mail_accounts/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

func (r *MailAccountService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("api/mail_accounts/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Manually process the selected mail account for new messages.
func (r *MailAccountService) Process(ctx context.Context, id int64, body MailAccountProcessParams, opts ...option.RequestOption) (res *MailAccountProcessResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/mail_accounts/%v/process/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Test a mail account
func (r *MailAccountService) Test(ctx context.Context, body MailAccountTestParams, opts ...option.RequestOption) (res *MailAccountTestResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/mail_accounts/test/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// - `1` - IMAP
// - `2` - Gmail OAuth
// - `3` - Outlook OAuth
type AccountType int64

const (
	AccountType1 AccountType = 1
	AccountType2 AccountType = 2
	AccountType3 AccountType = 3
)

// - `1` - No encryption
// - `2` - Use SSL
// - `3` - Use STARTTLS
type ImapSecurity int64

const (
	ImapSecurity1 ImapSecurity = 1
	ImapSecurity2 ImapSecurity = 2
	ImapSecurity3 ImapSecurity = 3
)

type MailAccount struct {
	ID            int64  `json:"id,required"`
	ImapServer    string `json:"imap_server,required"`
	Name          string `json:"name,required"`
	Password      string `json:"password,required"`
	UserCanChange bool   `json:"user_can_change,required"`
	Username      string `json:"username,required"`
	// - `1` - IMAP
	// - `2` - Gmail OAuth
	// - `3` - Outlook OAuth
	//
	// Any of 1, 2, 3.
	AccountType int64 `json:"account_type"`
	// The character set to use when communicating with the mail server, such as
	// 'UTF-8' or 'US-ASCII'.
	CharacterSet string `json:"character_set"`
	// The expiration date of the refresh token.
	Expiration time.Time `json:"expiration,nullable" format:"date-time"`
	// This is usually 143 for unencrypted and STARTTLS connections, and 993 for SSL
	// connections.
	ImapPort int64 `json:"imap_port,nullable"`
	// - `1` - No encryption
	// - `2` - Use SSL
	// - `3` - Use STARTTLS
	//
	// Any of 1, 2, 3.
	ImapSecurity int64 `json:"imap_security"`
	IsToken      bool  `json:"is_token"`
	Owner        int64 `json:"owner,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		ImapServer    respjson.Field
		Name          respjson.Field
		Password      respjson.Field
		UserCanChange respjson.Field
		Username      respjson.Field
		AccountType   respjson.Field
		CharacterSet  respjson.Field
		Expiration    respjson.Field
		ImapPort      respjson.Field
		ImapSecurity  respjson.Field
		IsToken       respjson.Field
		Owner         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MailAccount) RawJSON() string { return r.JSON.raw }
func (r *MailAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ImapServer, Name, Password, Username are required.
type MailAccountRequestParam struct {
	ImapServer string `json:"imap_server,required"`
	Name       string `json:"name,required"`
	Password   string `json:"password,required"`
	Username   string `json:"username,required"`
	// The expiration date of the refresh token.
	Expiration param.Opt[time.Time] `json:"expiration,omitzero" format:"date-time"`
	// This is usually 143 for unencrypted and STARTTLS connections, and 993 for SSL
	// connections.
	ImapPort param.Opt[int64] `json:"imap_port,omitzero"`
	Owner    param.Opt[int64] `json:"owner,omitzero"`
	// The character set to use when communicating with the mail server, such as
	// 'UTF-8' or 'US-ASCII'.
	CharacterSet param.Opt[string] `json:"character_set,omitzero"`
	IsToken      param.Opt[bool]   `json:"is_token,omitzero"`
	// - `1` - IMAP
	// - `2` - Gmail OAuth
	// - `3` - Outlook OAuth
	//
	// Any of 1, 2, 3.
	AccountType int64 `json:"account_type,omitzero"`
	// - `1` - No encryption
	// - `2` - Use SSL
	// - `3` - Use STARTTLS
	//
	// Any of 1, 2, 3.
	ImapSecurity   int64                                 `json:"imap_security,omitzero"`
	SetPermissions MailAccountRequestSetPermissionsParam `json:"set_permissions,omitzero"`
	paramObj
}

func (r MailAccountRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow MailAccountRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MailAccountRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MailAccountRequestParam](
		"account_type", 1, 2, 3,
	)
	apijson.RegisterFieldValidator[MailAccountRequestParam](
		"imap_security", 1, 2, 3,
	)
}

type MailAccountRequestSetPermissionsParam struct {
	Change MailAccountRequestSetPermissionsChangeParam `json:"change,omitzero"`
	View   MailAccountRequestSetPermissionsViewParam   `json:"view,omitzero"`
	paramObj
}

func (r MailAccountRequestSetPermissionsParam) MarshalJSON() (data []byte, err error) {
	type shadow MailAccountRequestSetPermissionsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MailAccountRequestSetPermissionsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MailAccountRequestSetPermissionsChangeParam struct {
	Groups []int64 `json:"groups,omitzero"`
	Users  []int64 `json:"users,omitzero"`
	paramObj
}

func (r MailAccountRequestSetPermissionsChangeParam) MarshalJSON() (data []byte, err error) {
	type shadow MailAccountRequestSetPermissionsChangeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MailAccountRequestSetPermissionsChangeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MailAccountRequestSetPermissionsViewParam struct {
	Groups []int64 `json:"groups,omitzero"`
	Users  []int64 `json:"users,omitzero"`
	paramObj
}

func (r MailAccountRequestSetPermissionsViewParam) MarshalJSON() (data []byte, err error) {
	type shadow MailAccountRequestSetPermissionsViewParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MailAccountRequestSetPermissionsViewParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MailAccountListResponse struct {
	Count    int64         `json:"count,required"`
	Results  []MailAccount `json:"results,required"`
	All      []any         `json:"all"`
	Next     string        `json:"next,nullable" format:"uri"`
	Previous string        `json:"previous,nullable" format:"uri"`
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
func (r MailAccountListResponse) RawJSON() string { return r.JSON.raw }
func (r *MailAccountListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MailAccountProcessResponse struct {
	Result string `json:"result"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Result      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MailAccountProcessResponse) RawJSON() string { return r.JSON.raw }
func (r *MailAccountProcessResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MailAccountTestResponse struct {
	Success bool `json:"success,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MailAccountTestResponse) RawJSON() string { return r.JSON.raw }
func (r *MailAccountTestResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MailAccountNewParams struct {
	MailAccountRequest MailAccountRequestParam
	paramObj
}

func (r MailAccountNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.MailAccountRequest)
}
func (r *MailAccountNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.MailAccountRequest)
}

type MailAccountUpdateParams struct {
	// The expiration date of the refresh token.
	Expiration param.Opt[time.Time] `json:"expiration,omitzero" format:"date-time"`
	// This is usually 143 for unencrypted and STARTTLS connections, and 993 for SSL
	// connections.
	ImapPort param.Opt[int64] `json:"imap_port,omitzero"`
	Owner    param.Opt[int64] `json:"owner,omitzero"`
	// The character set to use when communicating with the mail server, such as
	// 'UTF-8' or 'US-ASCII'.
	CharacterSet param.Opt[string] `json:"character_set,omitzero"`
	ImapServer   param.Opt[string] `json:"imap_server,omitzero"`
	IsToken      param.Opt[bool]   `json:"is_token,omitzero"`
	Name         param.Opt[string] `json:"name,omitzero"`
	Password     param.Opt[string] `json:"password,omitzero"`
	Username     param.Opt[string] `json:"username,omitzero"`
	// - `1` - IMAP
	// - `2` - Gmail OAuth
	// - `3` - Outlook OAuth
	//
	// Any of 1, 2, 3.
	AccountType int64 `json:"account_type,omitzero"`
	// - `1` - No encryption
	// - `2` - Use SSL
	// - `3` - Use STARTTLS
	//
	// Any of 1, 2, 3.
	ImapSecurity   int64                                 `json:"imap_security,omitzero"`
	SetPermissions MailAccountUpdateParamsSetPermissions `json:"set_permissions,omitzero"`
	paramObj
}

func (r MailAccountUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow MailAccountUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MailAccountUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MailAccountUpdateParamsSetPermissions struct {
	Change MailAccountUpdateParamsSetPermissionsChange `json:"change,omitzero"`
	View   MailAccountUpdateParamsSetPermissionsView   `json:"view,omitzero"`
	paramObj
}

func (r MailAccountUpdateParamsSetPermissions) MarshalJSON() (data []byte, err error) {
	type shadow MailAccountUpdateParamsSetPermissions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MailAccountUpdateParamsSetPermissions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MailAccountUpdateParamsSetPermissionsChange struct {
	Groups []int64 `json:"groups,omitzero"`
	Users  []int64 `json:"users,omitzero"`
	paramObj
}

func (r MailAccountUpdateParamsSetPermissionsChange) MarshalJSON() (data []byte, err error) {
	type shadow MailAccountUpdateParamsSetPermissionsChange
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MailAccountUpdateParamsSetPermissionsChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MailAccountUpdateParamsSetPermissionsView struct {
	Groups []int64 `json:"groups,omitzero"`
	Users  []int64 `json:"users,omitzero"`
	paramObj
}

func (r MailAccountUpdateParamsSetPermissionsView) MarshalJSON() (data []byte, err error) {
	type shadow MailAccountUpdateParamsSetPermissionsView
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MailAccountUpdateParamsSetPermissionsView) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MailAccountListParams struct {
	// A page number within the paginated result set.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Number of results to return per page.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MailAccountListParams]'s query parameters as `url.Values`.
func (r MailAccountListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MailAccountProcessParams struct {
	MailAccountRequest MailAccountRequestParam
	paramObj
}

func (r MailAccountProcessParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.MailAccountRequest)
}
func (r *MailAccountProcessParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.MailAccountRequest)
}

type MailAccountTestParams struct {
	MailAccountRequest MailAccountRequestParam
	paramObj
}

func (r MailAccountTestParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.MailAccountRequest)
}
func (r *MailAccountTestParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.MailAccountRequest)
}
