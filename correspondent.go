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

// CorrespondentService contains methods and other services that help with
// interacting with the paperless-ngx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCorrespondentService] method instead.
type CorrespondentService struct {
	Options []option.RequestOption
}

// NewCorrespondentService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCorrespondentService(opts ...option.RequestOption) (r CorrespondentService) {
	r = CorrespondentService{}
	r.Options = opts
	return
}

func (r *CorrespondentService) New(ctx context.Context, body CorrespondentNewParams, opts ...option.RequestOption) (res *Correspondent, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/correspondents/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *CorrespondentService) Get(ctx context.Context, id int64, query CorrespondentGetParams, opts ...option.RequestOption) (res *Correspondent, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/correspondents/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

func (r *CorrespondentService) Update(ctx context.Context, id int64, body CorrespondentUpdateParams, opts ...option.RequestOption) (res *Correspondent, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/correspondents/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

func (r *CorrespondentService) List(ctx context.Context, query CorrespondentListParams, opts ...option.RequestOption) (res *CorrespondentListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/correspondents/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

func (r *CorrespondentService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("api/correspondents/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type Correspondent struct {
	ID                 int64                    `json:"id,required"`
	DocumentCount      int64                    `json:"document_count,required"`
	LastCorrespondence time.Time                `json:"last_correspondence,required" format:"date"`
	Name               string                   `json:"name,required"`
	Permissions        CorrespondentPermissions `json:"permissions,required"`
	Slug               string                   `json:"slug,required"`
	UserCanChange      bool                     `json:"user_can_change,required"`
	IsInsensitive      bool                     `json:"is_insensitive"`
	Match              string                   `json:"match"`
	// - `0` - None
	// - `1` - Any word
	// - `2` - All words
	// - `3` - Exact match
	// - `4` - Regular expression
	// - `5` - Fuzzy word
	// - `6` - Automatic
	//
	// Any of 0, 1, 2, 3, 4, 5, 6.
	MatchingAlgorithm int64 `json:"matching_algorithm"`
	Owner             int64 `json:"owner,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		DocumentCount      respjson.Field
		LastCorrespondence respjson.Field
		Name               respjson.Field
		Permissions        respjson.Field
		Slug               respjson.Field
		UserCanChange      respjson.Field
		IsInsensitive      respjson.Field
		Match              respjson.Field
		MatchingAlgorithm  respjson.Field
		Owner              respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Correspondent) RawJSON() string { return r.JSON.raw }
func (r *Correspondent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CorrespondentPermissions struct {
	Change CorrespondentPermissionsChange `json:"change"`
	View   CorrespondentPermissionsView   `json:"view"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Change      respjson.Field
		View        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CorrespondentPermissions) RawJSON() string { return r.JSON.raw }
func (r *CorrespondentPermissions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CorrespondentPermissionsChange struct {
	Groups []int64 `json:"groups"`
	Users  []int64 `json:"users"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Groups      respjson.Field
		Users       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CorrespondentPermissionsChange) RawJSON() string { return r.JSON.raw }
func (r *CorrespondentPermissionsChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CorrespondentPermissionsView struct {
	Groups []int64 `json:"groups"`
	Users  []int64 `json:"users"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Groups      respjson.Field
		Users       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CorrespondentPermissionsView) RawJSON() string { return r.JSON.raw }
func (r *CorrespondentPermissionsView) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Name is required.
type CorrespondentRequestParam struct {
	Name          string            `json:"name,required"`
	Owner         param.Opt[int64]  `json:"owner,omitzero"`
	IsInsensitive param.Opt[bool]   `json:"is_insensitive,omitzero"`
	Match         param.Opt[string] `json:"match,omitzero"`
	// - `0` - None
	// - `1` - Any word
	// - `2` - All words
	// - `3` - Exact match
	// - `4` - Regular expression
	// - `5` - Fuzzy word
	// - `6` - Automatic
	//
	// Any of 0, 1, 2, 3, 4, 5, 6.
	MatchingAlgorithm int64                                   `json:"matching_algorithm,omitzero"`
	SetPermissions    CorrespondentRequestSetPermissionsParam `json:"set_permissions,omitzero"`
	paramObj
}

func (r CorrespondentRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CorrespondentRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CorrespondentRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CorrespondentRequestParam](
		"matching_algorithm", 0, 1, 2, 3, 4, 5, 6,
	)
}

type CorrespondentRequestSetPermissionsParam struct {
	Change CorrespondentRequestSetPermissionsChangeParam `json:"change,omitzero"`
	View   CorrespondentRequestSetPermissionsViewParam   `json:"view,omitzero"`
	paramObj
}

func (r CorrespondentRequestSetPermissionsParam) MarshalJSON() (data []byte, err error) {
	type shadow CorrespondentRequestSetPermissionsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CorrespondentRequestSetPermissionsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CorrespondentRequestSetPermissionsChangeParam struct {
	Groups []int64 `json:"groups,omitzero"`
	Users  []int64 `json:"users,omitzero"`
	paramObj
}

func (r CorrespondentRequestSetPermissionsChangeParam) MarshalJSON() (data []byte, err error) {
	type shadow CorrespondentRequestSetPermissionsChangeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CorrespondentRequestSetPermissionsChangeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CorrespondentRequestSetPermissionsViewParam struct {
	Groups []int64 `json:"groups,omitzero"`
	Users  []int64 `json:"users,omitzero"`
	paramObj
}

func (r CorrespondentRequestSetPermissionsViewParam) MarshalJSON() (data []byte, err error) {
	type shadow CorrespondentRequestSetPermissionsViewParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CorrespondentRequestSetPermissionsViewParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// - `0` - None
// - `1` - Any word
// - `2` - All words
// - `3` - Exact match
// - `4` - Regular expression
// - `5` - Fuzzy word
// - `6` - Automatic
type MatchingAlgorithm int64

const (
	MatchingAlgorithm0 MatchingAlgorithm = 0
	MatchingAlgorithm1 MatchingAlgorithm = 1
	MatchingAlgorithm2 MatchingAlgorithm = 2
	MatchingAlgorithm3 MatchingAlgorithm = 3
	MatchingAlgorithm4 MatchingAlgorithm = 4
	MatchingAlgorithm5 MatchingAlgorithm = 5
	MatchingAlgorithm6 MatchingAlgorithm = 6
)

type CorrespondentListResponse struct {
	Count    int64           `json:"count,required"`
	Results  []Correspondent `json:"results,required"`
	All      []any           `json:"all"`
	Next     string          `json:"next,nullable" format:"uri"`
	Previous string          `json:"previous,nullable" format:"uri"`
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
func (r CorrespondentListResponse) RawJSON() string { return r.JSON.raw }
func (r *CorrespondentListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CorrespondentNewParams struct {
	CorrespondentRequest CorrespondentRequestParam
	paramObj
}

func (r CorrespondentNewParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.CorrespondentRequest)
}
func (r *CorrespondentNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CorrespondentRequest)
}

type CorrespondentGetParams struct {
	FullPerms param.Opt[bool] `query:"full_perms,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CorrespondentGetParams]'s query parameters as `url.Values`.
func (r CorrespondentGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CorrespondentUpdateParams struct {
	Owner         param.Opt[int64]  `json:"owner,omitzero"`
	IsInsensitive param.Opt[bool]   `json:"is_insensitive,omitzero"`
	Match         param.Opt[string] `json:"match,omitzero"`
	Name          param.Opt[string] `json:"name,omitzero"`
	// - `0` - None
	// - `1` - Any word
	// - `2` - All words
	// - `3` - Exact match
	// - `4` - Regular expression
	// - `5` - Fuzzy word
	// - `6` - Automatic
	//
	// Any of 0, 1, 2, 3, 4, 5, 6.
	MatchingAlgorithm int64                                   `json:"matching_algorithm,omitzero"`
	SetPermissions    CorrespondentUpdateParamsSetPermissions `json:"set_permissions,omitzero"`
	paramObj
}

func (r CorrespondentUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow CorrespondentUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CorrespondentUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CorrespondentUpdateParamsSetPermissions struct {
	Change CorrespondentUpdateParamsSetPermissionsChange `json:"change,omitzero"`
	View   CorrespondentUpdateParamsSetPermissionsView   `json:"view,omitzero"`
	paramObj
}

func (r CorrespondentUpdateParamsSetPermissions) MarshalJSON() (data []byte, err error) {
	type shadow CorrespondentUpdateParamsSetPermissions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CorrespondentUpdateParamsSetPermissions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CorrespondentUpdateParamsSetPermissionsChange struct {
	Groups []int64 `json:"groups,omitzero"`
	Users  []int64 `json:"users,omitzero"`
	paramObj
}

func (r CorrespondentUpdateParamsSetPermissionsChange) MarshalJSON() (data []byte, err error) {
	type shadow CorrespondentUpdateParamsSetPermissionsChange
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CorrespondentUpdateParamsSetPermissionsChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CorrespondentUpdateParamsSetPermissionsView struct {
	Groups []int64 `json:"groups,omitzero"`
	Users  []int64 `json:"users,omitzero"`
	paramObj
}

func (r CorrespondentUpdateParamsSetPermissionsView) MarshalJSON() (data []byte, err error) {
	type shadow CorrespondentUpdateParamsSetPermissionsView
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CorrespondentUpdateParamsSetPermissionsView) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CorrespondentListParams struct {
	ID              param.Opt[int64]  `query:"id,omitzero" json:"-"`
	FullPerms       param.Opt[bool]   `query:"full_perms,omitzero" json:"-"`
	NameIcontains   param.Opt[string] `query:"name__icontains,omitzero" json:"-"`
	NameIendswith   param.Opt[string] `query:"name__iendswith,omitzero" json:"-"`
	NameIexact      param.Opt[string] `query:"name__iexact,omitzero" json:"-"`
	NameIstartswith param.Opt[string] `query:"name__istartswith,omitzero" json:"-"`
	// Which field to use when ordering the results.
	Ordering param.Opt[string] `query:"ordering,omitzero" json:"-"`
	// A page number within the paginated result set.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Number of results to return per page.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Multiple values may be separated by commas.
	IDIn []int64 `query:"id__in,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CorrespondentListParams]'s query parameters as
// `url.Values`.
func (r CorrespondentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
