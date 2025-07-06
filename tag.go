// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package paperlessnix

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/paperless-nix-go/internal/apijson"
	"github.com/stainless-sdks/paperless-nix-go/internal/apiquery"
	"github.com/stainless-sdks/paperless-nix-go/internal/requestconfig"
	"github.com/stainless-sdks/paperless-nix-go/option"
	"github.com/stainless-sdks/paperless-nix-go/packages/param"
	"github.com/stainless-sdks/paperless-nix-go/packages/respjson"
)

// TagService contains methods and other services that help with interacting with
// the paperless-nix API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTagService] method instead.
type TagService struct {
	Options []option.RequestOption
}

// NewTagService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTagService(opts ...option.RequestOption) (r TagService) {
	r = TagService{}
	r.Options = opts
	return
}

func (r *TagService) New(ctx context.Context, body TagNewParams, opts ...option.RequestOption) (res *Tag, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/tags/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *TagService) Get(ctx context.Context, id int64, query TagGetParams, opts ...option.RequestOption) (res *Tag, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/tags/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

func (r *TagService) Update(ctx context.Context, id int64, body TagUpdateParams, opts ...option.RequestOption) (res *Tag, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/tags/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

func (r *TagService) List(ctx context.Context, query TagListParams, opts ...option.RequestOption) (res *TagListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/tags/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

func (r *TagService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("api/tags/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type Tag struct {
	ID            int64  `json:"id,required"`
	DocumentCount int64  `json:"document_count,required"`
	Name          string `json:"name,required"`
	Slug          string `json:"slug,required"`
	TextColor     string `json:"text_color,required"`
	UserCanChange bool   `json:"user_can_change,required"`
	Color         string `json:"color"`
	// Marks this tag as an inbox tag: All newly consumed documents will be tagged with
	// inbox tags.
	IsInboxTag    bool   `json:"is_inbox_tag"`
	IsInsensitive bool   `json:"is_insensitive"`
	Match         string `json:"match"`
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
		ID                respjson.Field
		DocumentCount     respjson.Field
		Name              respjson.Field
		Slug              respjson.Field
		TextColor         respjson.Field
		UserCanChange     respjson.Field
		Color             respjson.Field
		IsInboxTag        respjson.Field
		IsInsensitive     respjson.Field
		Match             respjson.Field
		MatchingAlgorithm respjson.Field
		Owner             respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Tag) RawJSON() string { return r.JSON.raw }
func (r *Tag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Name is required.
type TagRequestParam struct {
	Name  string            `json:"name,required"`
	Owner param.Opt[int64]  `json:"owner,omitzero"`
	Color param.Opt[string] `json:"color,omitzero"`
	// Marks this tag as an inbox tag: All newly consumed documents will be tagged with
	// inbox tags.
	IsInboxTag    param.Opt[bool]   `json:"is_inbox_tag,omitzero"`
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
	MatchingAlgorithm int64                         `json:"matching_algorithm,omitzero"`
	SetPermissions    TagRequestSetPermissionsParam `json:"set_permissions,omitzero"`
	paramObj
}

func (r TagRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow TagRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TagRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[TagRequestParam](
		"matching_algorithm", 0, 1, 2, 3, 4, 5, 6,
	)
}

type TagRequestSetPermissionsParam struct {
	Change TagRequestSetPermissionsChangeParam `json:"change,omitzero"`
	View   TagRequestSetPermissionsViewParam   `json:"view,omitzero"`
	paramObj
}

func (r TagRequestSetPermissionsParam) MarshalJSON() (data []byte, err error) {
	type shadow TagRequestSetPermissionsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TagRequestSetPermissionsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagRequestSetPermissionsChangeParam struct {
	Groups []int64 `json:"groups,omitzero"`
	Users  []int64 `json:"users,omitzero"`
	paramObj
}

func (r TagRequestSetPermissionsChangeParam) MarshalJSON() (data []byte, err error) {
	type shadow TagRequestSetPermissionsChangeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TagRequestSetPermissionsChangeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagRequestSetPermissionsViewParam struct {
	Groups []int64 `json:"groups,omitzero"`
	Users  []int64 `json:"users,omitzero"`
	paramObj
}

func (r TagRequestSetPermissionsViewParam) MarshalJSON() (data []byte, err error) {
	type shadow TagRequestSetPermissionsViewParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TagRequestSetPermissionsViewParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagListResponse struct {
	Count    int64  `json:"count,required"`
	Results  []Tag  `json:"results,required"`
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
func (r TagListResponse) RawJSON() string { return r.JSON.raw }
func (r *TagListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagNewParams struct {
	TagRequest TagRequestParam
	paramObj
}

func (r TagNewParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.TagRequest)
}
func (r *TagNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.TagRequest)
}

type TagGetParams struct {
	FullPerms param.Opt[bool] `query:"full_perms,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TagGetParams]'s query parameters as `url.Values`.
func (r TagGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TagUpdateParams struct {
	Owner param.Opt[int64]  `json:"owner,omitzero"`
	Color param.Opt[string] `json:"color,omitzero"`
	// Marks this tag as an inbox tag: All newly consumed documents will be tagged with
	// inbox tags.
	IsInboxTag    param.Opt[bool]   `json:"is_inbox_tag,omitzero"`
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
	MatchingAlgorithm int64                         `json:"matching_algorithm,omitzero"`
	SetPermissions    TagUpdateParamsSetPermissions `json:"set_permissions,omitzero"`
	paramObj
}

func (r TagUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow TagUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TagUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagUpdateParamsSetPermissions struct {
	Change TagUpdateParamsSetPermissionsChange `json:"change,omitzero"`
	View   TagUpdateParamsSetPermissionsView   `json:"view,omitzero"`
	paramObj
}

func (r TagUpdateParamsSetPermissions) MarshalJSON() (data []byte, err error) {
	type shadow TagUpdateParamsSetPermissions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TagUpdateParamsSetPermissions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagUpdateParamsSetPermissionsChange struct {
	Groups []int64 `json:"groups,omitzero"`
	Users  []int64 `json:"users,omitzero"`
	paramObj
}

func (r TagUpdateParamsSetPermissionsChange) MarshalJSON() (data []byte, err error) {
	type shadow TagUpdateParamsSetPermissionsChange
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TagUpdateParamsSetPermissionsChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagUpdateParamsSetPermissionsView struct {
	Groups []int64 `json:"groups,omitzero"`
	Users  []int64 `json:"users,omitzero"`
	paramObj
}

func (r TagUpdateParamsSetPermissionsView) MarshalJSON() (data []byte, err error) {
	type shadow TagUpdateParamsSetPermissionsView
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TagUpdateParamsSetPermissionsView) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagListParams struct {
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

// URLQuery serializes [TagListParams]'s query parameters as `url.Values`.
func (r TagListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
