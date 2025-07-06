// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package paperlessngx

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/defasdefbe/go-paperless-ngx/internal/apijson"
	"github.com/defasdefbe/go-paperless-ngx/internal/apiquery"
	"github.com/defasdefbe/go-paperless-ngx/internal/requestconfig"
	"github.com/defasdefbe/go-paperless-ngx/option"
	"github.com/defasdefbe/go-paperless-ngx/packages/param"
	"github.com/defasdefbe/go-paperless-ngx/packages/respjson"
)

// SavedViewService contains methods and other services that help with interacting
// with the paperless-ngx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSavedViewService] method instead.
type SavedViewService struct {
	Options []option.RequestOption
}

// NewSavedViewService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSavedViewService(opts ...option.RequestOption) (r SavedViewService) {
	r = SavedViewService{}
	r.Options = opts
	return
}

func (r *SavedViewService) New(ctx context.Context, body SavedViewNewParams, opts ...option.RequestOption) (res *SavedView, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/saved_views/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *SavedViewService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *SavedView, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/saved_views/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *SavedViewService) Update(ctx context.Context, id int64, body SavedViewUpdateParams, opts ...option.RequestOption) (res *SavedView, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/saved_views/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

func (r *SavedViewService) List(ctx context.Context, query SavedViewListParams, opts ...option.RequestOption) (res *SavedViewListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/saved_views/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

func (r *SavedViewService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("api/saved_views/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// - `0` - title contains
// - `1` - content contains
// - `2` - ASN is
// - `3` - correspondent is
// - `4` - document type is
// - `5` - is in inbox
// - `6` - has tag
// - `7` - has any tag
// - `8` - created before
// - `9` - created after
// - `10` - created year is
// - `11` - created month is
// - `12` - created day is
// - `13` - added before
// - `14` - added after
// - `15` - modified before
// - `16` - modified after
// - `17` - does not have tag
// - `18` - does not have ASN
// - `19` - title or content contains
// - `20` - fulltext query
// - `21` - more like this
// - `22` - has tags in
// - `23` - ASN greater than
// - `24` - ASN less than
// - `25` - storage path is
// - `26` - has correspondent in
// - `27` - does not have correspondent in
// - `28` - has document type in
// - `29` - does not have document type in
// - `30` - has storage path in
// - `31` - does not have storage path in
// - `32` - owner is
// - `33` - has owner in
// - `34` - does not have owner
// - `35` - does not have owner in
// - `36` - has custom field value
// - `37` - is shared by me
// - `38` - has custom fields
// - `39` - has custom field in
// - `40` - does not have custom field in
// - `41` - does not have custom field
// - `42` - custom fields query
// - `43` - created to
// - `44` - created from
// - `45` - added to
// - `46` - added from
// - `47` - mime type is
type RuleTypeEnum int64

const (
	RuleTypeEnum0  RuleTypeEnum = 0
	RuleTypeEnum1  RuleTypeEnum = 1
	RuleTypeEnum2  RuleTypeEnum = 2
	RuleTypeEnum3  RuleTypeEnum = 3
	RuleTypeEnum4  RuleTypeEnum = 4
	RuleTypeEnum5  RuleTypeEnum = 5
	RuleTypeEnum6  RuleTypeEnum = 6
	RuleTypeEnum7  RuleTypeEnum = 7
	RuleTypeEnum8  RuleTypeEnum = 8
	RuleTypeEnum9  RuleTypeEnum = 9
	RuleTypeEnum10 RuleTypeEnum = 10
	RuleTypeEnum11 RuleTypeEnum = 11
	RuleTypeEnum12 RuleTypeEnum = 12
	RuleTypeEnum13 RuleTypeEnum = 13
	RuleTypeEnum14 RuleTypeEnum = 14
	RuleTypeEnum15 RuleTypeEnum = 15
	RuleTypeEnum16 RuleTypeEnum = 16
	RuleTypeEnum17 RuleTypeEnum = 17
	RuleTypeEnum18 RuleTypeEnum = 18
	RuleTypeEnum19 RuleTypeEnum = 19
	RuleTypeEnum20 RuleTypeEnum = 20
	RuleTypeEnum21 RuleTypeEnum = 21
	RuleTypeEnum22 RuleTypeEnum = 22
	RuleTypeEnum23 RuleTypeEnum = 23
	RuleTypeEnum24 RuleTypeEnum = 24
	RuleTypeEnum25 RuleTypeEnum = 25
	RuleTypeEnum26 RuleTypeEnum = 26
	RuleTypeEnum27 RuleTypeEnum = 27
	RuleTypeEnum28 RuleTypeEnum = 28
	RuleTypeEnum29 RuleTypeEnum = 29
	RuleTypeEnum30 RuleTypeEnum = 30
	RuleTypeEnum31 RuleTypeEnum = 31
	RuleTypeEnum32 RuleTypeEnum = 32
	RuleTypeEnum33 RuleTypeEnum = 33
	RuleTypeEnum34 RuleTypeEnum = 34
	RuleTypeEnum35 RuleTypeEnum = 35
	RuleTypeEnum36 RuleTypeEnum = 36
	RuleTypeEnum37 RuleTypeEnum = 37
	RuleTypeEnum38 RuleTypeEnum = 38
	RuleTypeEnum39 RuleTypeEnum = 39
	RuleTypeEnum40 RuleTypeEnum = 40
	RuleTypeEnum41 RuleTypeEnum = 41
	RuleTypeEnum42 RuleTypeEnum = 42
	RuleTypeEnum43 RuleTypeEnum = 43
	RuleTypeEnum44 RuleTypeEnum = 44
	RuleTypeEnum45 RuleTypeEnum = 45
	RuleTypeEnum46 RuleTypeEnum = 46
	RuleTypeEnum47 RuleTypeEnum = 47
)

type SavedView struct {
	ID              int64                 `json:"id,required"`
	FilterRules     []SavedViewFilterRule `json:"filter_rules,required"`
	Name            string                `json:"name,required"`
	ShowInSidebar   bool                  `json:"show_in_sidebar,required"`
	ShowOnDashboard bool                  `json:"show_on_dashboard,required"`
	UserCanChange   bool                  `json:"user_can_change,required"`
	DisplayFields   any                   `json:"display_fields"`
	// Any of "table", "smallCards", "largeCards", "".
	DisplayMode SavedViewDisplayMode `json:"display_mode,nullable"`
	Owner       int64                `json:"owner,nullable"`
	PageSize    int64                `json:"page_size,nullable"`
	SortField   string               `json:"sort_field,nullable"`
	SortReverse bool                 `json:"sort_reverse"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		FilterRules     respjson.Field
		Name            respjson.Field
		ShowInSidebar   respjson.Field
		ShowOnDashboard respjson.Field
		UserCanChange   respjson.Field
		DisplayFields   respjson.Field
		DisplayMode     respjson.Field
		Owner           respjson.Field
		PageSize        respjson.Field
		SortField       respjson.Field
		SortReverse     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SavedView) RawJSON() string { return r.JSON.raw }
func (r *SavedView) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SavedViewFilterRule struct {
	// - `0` - title contains
	// - `1` - content contains
	// - `2` - ASN is
	// - `3` - correspondent is
	// - `4` - document type is
	// - `5` - is in inbox
	// - `6` - has tag
	// - `7` - has any tag
	// - `8` - created before
	// - `9` - created after
	// - `10` - created year is
	// - `11` - created month is
	// - `12` - created day is
	// - `13` - added before
	// - `14` - added after
	// - `15` - modified before
	// - `16` - modified after
	// - `17` - does not have tag
	// - `18` - does not have ASN
	// - `19` - title or content contains
	// - `20` - fulltext query
	// - `21` - more like this
	// - `22` - has tags in
	// - `23` - ASN greater than
	// - `24` - ASN less than
	// - `25` - storage path is
	// - `26` - has correspondent in
	// - `27` - does not have correspondent in
	// - `28` - has document type in
	// - `29` - does not have document type in
	// - `30` - has storage path in
	// - `31` - does not have storage path in
	// - `32` - owner is
	// - `33` - has owner in
	// - `34` - does not have owner
	// - `35` - does not have owner in
	// - `36` - has custom field value
	// - `37` - is shared by me
	// - `38` - has custom fields
	// - `39` - has custom field in
	// - `40` - does not have custom field in
	// - `41` - does not have custom field
	// - `42` - custom fields query
	// - `43` - created to
	// - `44` - created from
	// - `45` - added to
	// - `46` - added from
	// - `47` - mime type is
	//
	// Any of 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
	// 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
	// 41, 42, 43, 44, 45, 46, 47.
	RuleType int64  `json:"rule_type,required"`
	Value    string `json:"value,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RuleType    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SavedViewFilterRule) RawJSON() string { return r.JSON.raw }
func (r *SavedViewFilterRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SavedViewDisplayMode string

const (
	SavedViewDisplayModeTable      SavedViewDisplayMode = "table"
	SavedViewDisplayModeSmallCards SavedViewDisplayMode = "smallCards"
	SavedViewDisplayModeLargeCards SavedViewDisplayMode = "largeCards"
	SavedViewDisplayModeEmpty      SavedViewDisplayMode = ""
)

// The property RuleType is required.
type SavedViewFilterRuleParam struct {
	// - `0` - title contains
	// - `1` - content contains
	// - `2` - ASN is
	// - `3` - correspondent is
	// - `4` - document type is
	// - `5` - is in inbox
	// - `6` - has tag
	// - `7` - has any tag
	// - `8` - created before
	// - `9` - created after
	// - `10` - created year is
	// - `11` - created month is
	// - `12` - created day is
	// - `13` - added before
	// - `14` - added after
	// - `15` - modified before
	// - `16` - modified after
	// - `17` - does not have tag
	// - `18` - does not have ASN
	// - `19` - title or content contains
	// - `20` - fulltext query
	// - `21` - more like this
	// - `22` - has tags in
	// - `23` - ASN greater than
	// - `24` - ASN less than
	// - `25` - storage path is
	// - `26` - has correspondent in
	// - `27` - does not have correspondent in
	// - `28` - has document type in
	// - `29` - does not have document type in
	// - `30` - has storage path in
	// - `31` - does not have storage path in
	// - `32` - owner is
	// - `33` - has owner in
	// - `34` - does not have owner
	// - `35` - does not have owner in
	// - `36` - has custom field value
	// - `37` - is shared by me
	// - `38` - has custom fields
	// - `39` - has custom field in
	// - `40` - does not have custom field in
	// - `41` - does not have custom field
	// - `42` - custom fields query
	// - `43` - created to
	// - `44` - created from
	// - `45` - added to
	// - `46` - added from
	// - `47` - mime type is
	//
	// Any of 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
	// 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
	// 41, 42, 43, 44, 45, 46, 47.
	RuleType int64             `json:"rule_type,omitzero,required"`
	Value    param.Opt[string] `json:"value,omitzero"`
	paramObj
}

func (r SavedViewFilterRuleParam) MarshalJSON() (data []byte, err error) {
	type shadow SavedViewFilterRuleParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SavedViewFilterRuleParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SavedViewFilterRuleParam](
		"rule_type", 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47,
	)
}

// The properties FilterRules, Name, ShowInSidebar, ShowOnDashboard are required.
type SavedViewRequestParam struct {
	FilterRules     []SavedViewFilterRuleParam `json:"filter_rules,omitzero,required"`
	Name            string                     `json:"name,required"`
	ShowInSidebar   bool                       `json:"show_in_sidebar,required"`
	ShowOnDashboard bool                       `json:"show_on_dashboard,required"`
	Owner           param.Opt[int64]           `json:"owner,omitzero"`
	PageSize        param.Opt[int64]           `json:"page_size,omitzero"`
	SortField       param.Opt[string]          `json:"sort_field,omitzero"`
	SortReverse     param.Opt[bool]            `json:"sort_reverse,omitzero"`
	// Any of "table", "smallCards", "largeCards", "".
	DisplayMode   SavedViewRequestDisplayMode `json:"display_mode,omitzero"`
	DisplayFields any                         `json:"display_fields,omitzero"`
	paramObj
}

func (r SavedViewRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow SavedViewRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SavedViewRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SavedViewRequestDisplayMode string

const (
	SavedViewRequestDisplayModeTable      SavedViewRequestDisplayMode = "table"
	SavedViewRequestDisplayModeSmallCards SavedViewRequestDisplayMode = "smallCards"
	SavedViewRequestDisplayModeLargeCards SavedViewRequestDisplayMode = "largeCards"
	SavedViewRequestDisplayModeEmpty      SavedViewRequestDisplayMode = ""
)

type SavedViewListResponse struct {
	Count    int64       `json:"count,required"`
	Results  []SavedView `json:"results,required"`
	All      []any       `json:"all"`
	Next     string      `json:"next,nullable" format:"uri"`
	Previous string      `json:"previous,nullable" format:"uri"`
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
func (r SavedViewListResponse) RawJSON() string { return r.JSON.raw }
func (r *SavedViewListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SavedViewNewParams struct {
	SavedViewRequest SavedViewRequestParam
	paramObj
}

func (r SavedViewNewParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.SavedViewRequest)
}
func (r *SavedViewNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.SavedViewRequest)
}

type SavedViewUpdateParams struct {
	Owner           param.Opt[int64]  `json:"owner,omitzero"`
	PageSize        param.Opt[int64]  `json:"page_size,omitzero"`
	SortField       param.Opt[string] `json:"sort_field,omitzero"`
	Name            param.Opt[string] `json:"name,omitzero"`
	ShowInSidebar   param.Opt[bool]   `json:"show_in_sidebar,omitzero"`
	ShowOnDashboard param.Opt[bool]   `json:"show_on_dashboard,omitzero"`
	SortReverse     param.Opt[bool]   `json:"sort_reverse,omitzero"`
	// Any of "table", "smallCards", "largeCards", "".
	DisplayMode   SavedViewUpdateParamsDisplayMode `json:"display_mode,omitzero"`
	DisplayFields any                              `json:"display_fields,omitzero"`
	FilterRules   []SavedViewFilterRuleParam       `json:"filter_rules,omitzero"`
	paramObj
}

func (r SavedViewUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SavedViewUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SavedViewUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SavedViewUpdateParamsDisplayMode string

const (
	SavedViewUpdateParamsDisplayModeTable      SavedViewUpdateParamsDisplayMode = "table"
	SavedViewUpdateParamsDisplayModeSmallCards SavedViewUpdateParamsDisplayMode = "smallCards"
	SavedViewUpdateParamsDisplayModeLargeCards SavedViewUpdateParamsDisplayMode = "largeCards"
	SavedViewUpdateParamsDisplayModeEmpty      SavedViewUpdateParamsDisplayMode = ""
)

type SavedViewListParams struct {
	// A page number within the paginated result set.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Number of results to return per page.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SavedViewListParams]'s query parameters as `url.Values`.
func (r SavedViewListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
