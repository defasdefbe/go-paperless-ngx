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
	shimjson "github.com/defasdefbe/go-paperless-ngx/internal/encoding/json"
	"github.com/defasdefbe/go-paperless-ngx/internal/requestconfig"
	"github.com/defasdefbe/go-paperless-ngx/option"
	"github.com/defasdefbe/go-paperless-ngx/packages/param"
	"github.com/defasdefbe/go-paperless-ngx/packages/respjson"
)

// CustomFieldService contains methods and other services that help with
// interacting with the paperless-ngx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomFieldService] method instead.
type CustomFieldService struct {
	Options []option.RequestOption
}

// NewCustomFieldService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomFieldService(opts ...option.RequestOption) (r CustomFieldService) {
	r = CustomFieldService{}
	r.Options = opts
	return
}

func (r *CustomFieldService) New(ctx context.Context, body CustomFieldNewParams, opts ...option.RequestOption) (res *CustomField, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/custom_fields/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *CustomFieldService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *CustomField, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/custom_fields/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *CustomFieldService) Update(ctx context.Context, id int64, body CustomFieldUpdateParams, opts ...option.RequestOption) (res *CustomField, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/custom_fields/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

func (r *CustomFieldService) List(ctx context.Context, query CustomFieldListParams, opts ...option.RequestOption) (res *CustomFieldListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/custom_fields/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

func (r *CustomFieldService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("api/custom_fields/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type CustomField struct {
	ID int64 `json:"id,required"`
	// - `string` - string
	// - `url` - url
	// - `date` - date
	// - `boolean` - boolean
	// - `integer` - integer
	// - `float` - float
	// - `monetary` - monetary
	// - `documentlink` - documentlink
	// - `select` - select
	//
	// Any of "string", "url", "date", "boolean", "integer", "float", "monetary",
	// "documentlink", "select".
	DataType      DataTypeEnum `json:"data_type,required"`
	DocumentCount int64        `json:"document_count,required"`
	Name          string       `json:"name,required"`
	// Extra data for the custom field, such as select options
	ExtraData any `json:"extra_data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		DataType      respjson.Field
		DocumentCount respjson.Field
		Name          respjson.Field
		ExtraData     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomField) RawJSON() string { return r.JSON.raw }
func (r *CustomField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DataType, Name are required.
type CustomFieldRequestParam struct {
	// - `string` - string
	// - `url` - url
	// - `date` - date
	// - `boolean` - boolean
	// - `integer` - integer
	// - `float` - float
	// - `monetary` - monetary
	// - `documentlink` - documentlink
	// - `select` - select
	//
	// Any of "string", "url", "date", "boolean", "integer", "float", "monetary",
	// "documentlink", "select".
	DataType DataTypeEnum `json:"data_type,omitzero,required"`
	Name     string       `json:"name,required"`
	// Extra data for the custom field, such as select options
	ExtraData any `json:"extra_data,omitzero"`
	paramObj
}

func (r CustomFieldRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CustomFieldRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomFieldRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// - `string` - string
// - `url` - url
// - `date` - date
// - `boolean` - boolean
// - `integer` - integer
// - `float` - float
// - `monetary` - monetary
// - `documentlink` - documentlink
// - `select` - select
type DataTypeEnum string

const (
	DataTypeEnumString       DataTypeEnum = "string"
	DataTypeEnumURL          DataTypeEnum = "url"
	DataTypeEnumDate         DataTypeEnum = "date"
	DataTypeEnumBoolean      DataTypeEnum = "boolean"
	DataTypeEnumInteger      DataTypeEnum = "integer"
	DataTypeEnumFloat        DataTypeEnum = "float"
	DataTypeEnumMonetary     DataTypeEnum = "monetary"
	DataTypeEnumDocumentlink DataTypeEnum = "documentlink"
	DataTypeEnumSelect       DataTypeEnum = "select"
)

type CustomFieldListResponse struct {
	Count    int64         `json:"count,required"`
	Results  []CustomField `json:"results,required"`
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
func (r CustomFieldListResponse) RawJSON() string { return r.JSON.raw }
func (r *CustomFieldListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomFieldNewParams struct {
	CustomFieldRequest CustomFieldRequestParam
	paramObj
}

func (r CustomFieldNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CustomFieldRequest)
}
func (r *CustomFieldNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CustomFieldRequest)
}

type CustomFieldUpdateParams struct {
	Name param.Opt[string] `json:"name,omitzero"`
	// - `string` - string
	// - `url` - url
	// - `date` - date
	// - `boolean` - boolean
	// - `integer` - integer
	// - `float` - float
	// - `monetary` - monetary
	// - `documentlink` - documentlink
	// - `select` - select
	//
	// Any of "string", "url", "date", "boolean", "integer", "float", "monetary",
	// "documentlink", "select".
	DataType DataTypeEnum `json:"data_type,omitzero"`
	// Extra data for the custom field, such as select options
	ExtraData any `json:"extra_data,omitzero"`
	paramObj
}

func (r CustomFieldUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow CustomFieldUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomFieldUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomFieldListParams struct {
	ID              param.Opt[int64]  `query:"id,omitzero" json:"-"`
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

// URLQuery serializes [CustomFieldListParams]'s query parameters as `url.Values`.
func (r CustomFieldListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
