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

// ShareLinkService contains methods and other services that help with interacting
// with the paperless-ngx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewShareLinkService] method instead.
type ShareLinkService struct {
	Options []option.RequestOption
}

// NewShareLinkService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewShareLinkService(opts ...option.RequestOption) (r ShareLinkService) {
	r = ShareLinkService{}
	r.Options = opts
	return
}

func (r *ShareLinkService) New(ctx context.Context, body ShareLinkNewParams, opts ...option.RequestOption) (res *ShareLink, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/share_links/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *ShareLinkService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *ShareLink, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/share_links/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *ShareLinkService) Update(ctx context.Context, id int64, body ShareLinkUpdateParams, opts ...option.RequestOption) (res *ShareLink, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/share_links/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

func (r *ShareLinkService) List(ctx context.Context, query ShareLinkListParams, opts ...option.RequestOption) (res *ShareLinkListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/share_links/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

func (r *ShareLinkService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("api/share_links/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// - `archive` - Archive
// - `original` - Original
type FileVersionEnum string

const (
	FileVersionEnumArchive  FileVersionEnum = "archive"
	FileVersionEnumOriginal FileVersionEnum = "original"
)

type ShareLink struct {
	ID         int64     `json:"id,required"`
	Created    time.Time `json:"created,required" format:"date-time"`
	Slug       string    `json:"slug,required"`
	Document   int64     `json:"document"`
	Expiration time.Time `json:"expiration,nullable" format:"date-time"`
	// - `archive` - Archive
	// - `original` - Original
	//
	// Any of "archive", "original".
	FileVersion FileVersionEnum `json:"file_version"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Created     respjson.Field
		Slug        respjson.Field
		Document    respjson.Field
		Expiration  respjson.Field
		FileVersion respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShareLink) RawJSON() string { return r.JSON.raw }
func (r *ShareLink) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkRequestParam struct {
	Expiration param.Opt[time.Time] `json:"expiration,omitzero" format:"date-time"`
	Document   param.Opt[int64]     `json:"document,omitzero"`
	// - `archive` - Archive
	// - `original` - Original
	//
	// Any of "archive", "original".
	FileVersion FileVersionEnum `json:"file_version,omitzero"`
	paramObj
}

func (r ShareLinkRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ShareLinkRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ShareLinkRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkListResponse struct {
	Count    int64       `json:"count,required"`
	Results  []ShareLink `json:"results,required"`
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
func (r ShareLinkListResponse) RawJSON() string { return r.JSON.raw }
func (r *ShareLinkListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkNewParams struct {
	ShareLinkRequest ShareLinkRequestParam
	paramObj
}

func (r ShareLinkNewParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.ShareLinkRequest)
}
func (r *ShareLinkNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.ShareLinkRequest)
}

type ShareLinkUpdateParams struct {
	Expiration param.Opt[time.Time] `json:"expiration,omitzero" format:"date-time"`
	Document   param.Opt[int64]     `json:"document,omitzero"`
	// - `archive` - Archive
	// - `original` - Original
	//
	// Any of "archive", "original".
	FileVersion FileVersionEnum `json:"file_version,omitzero"`
	paramObj
}

func (r ShareLinkUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ShareLinkUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ShareLinkUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ShareLinkListParams struct {
	CreatedDateGt     param.Opt[time.Time] `query:"created__date__gt,omitzero" format:"date" json:"-"`
	CreatedDateGte    param.Opt[time.Time] `query:"created__date__gte,omitzero" format:"date" json:"-"`
	CreatedDateLt     param.Opt[time.Time] `query:"created__date__lt,omitzero" format:"date" json:"-"`
	CreatedDateLte    param.Opt[time.Time] `query:"created__date__lte,omitzero" format:"date" json:"-"`
	CreatedDay        param.Opt[float64]   `query:"created__day,omitzero" json:"-"`
	CreatedGt         param.Opt[time.Time] `query:"created__gt,omitzero" format:"date-time" json:"-"`
	CreatedGte        param.Opt[time.Time] `query:"created__gte,omitzero" format:"date-time" json:"-"`
	CreatedLt         param.Opt[time.Time] `query:"created__lt,omitzero" format:"date-time" json:"-"`
	CreatedLte        param.Opt[time.Time] `query:"created__lte,omitzero" format:"date-time" json:"-"`
	CreatedMonth      param.Opt[float64]   `query:"created__month,omitzero" json:"-"`
	CreatedYear       param.Opt[float64]   `query:"created__year,omitzero" json:"-"`
	ExpirationDateGt  param.Opt[time.Time] `query:"expiration__date__gt,omitzero" format:"date" json:"-"`
	ExpirationDateGte param.Opt[time.Time] `query:"expiration__date__gte,omitzero" format:"date" json:"-"`
	ExpirationDateLt  param.Opt[time.Time] `query:"expiration__date__lt,omitzero" format:"date" json:"-"`
	ExpirationDateLte param.Opt[time.Time] `query:"expiration__date__lte,omitzero" format:"date" json:"-"`
	ExpirationDay     param.Opt[float64]   `query:"expiration__day,omitzero" json:"-"`
	ExpirationGt      param.Opt[time.Time] `query:"expiration__gt,omitzero" format:"date-time" json:"-"`
	ExpirationGte     param.Opt[time.Time] `query:"expiration__gte,omitzero" format:"date-time" json:"-"`
	ExpirationLt      param.Opt[time.Time] `query:"expiration__lt,omitzero" format:"date-time" json:"-"`
	ExpirationLte     param.Opt[time.Time] `query:"expiration__lte,omitzero" format:"date-time" json:"-"`
	ExpirationMonth   param.Opt[float64]   `query:"expiration__month,omitzero" json:"-"`
	ExpirationYear    param.Opt[float64]   `query:"expiration__year,omitzero" json:"-"`
	// Which field to use when ordering the results.
	Ordering param.Opt[string] `query:"ordering,omitzero" json:"-"`
	// A page number within the paginated result set.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Number of results to return per page.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ShareLinkListParams]'s query parameters as `url.Values`.
func (r ShareLinkListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
