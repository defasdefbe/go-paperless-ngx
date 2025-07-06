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

// DocumentTypeService contains methods and other services that help with
// interacting with the paperless-nix API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDocumentTypeService] method instead.
type DocumentTypeService struct {
	Options []option.RequestOption
}

// NewDocumentTypeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDocumentTypeService(opts ...option.RequestOption) (r DocumentTypeService) {
	r = DocumentTypeService{}
	r.Options = opts
	return
}

func (r *DocumentTypeService) New(ctx context.Context, body DocumentTypeNewParams, opts ...option.RequestOption) (res *DocumentType, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/document_types/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *DocumentTypeService) Get(ctx context.Context, id int64, query DocumentTypeGetParams, opts ...option.RequestOption) (res *DocumentType, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/document_types/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

func (r *DocumentTypeService) Update(ctx context.Context, id int64, body DocumentTypeUpdateParams, opts ...option.RequestOption) (res *DocumentType, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/document_types/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

func (r *DocumentTypeService) List(ctx context.Context, query DocumentTypeListParams, opts ...option.RequestOption) (res *DocumentTypeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/document_types/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

func (r *DocumentTypeService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("api/document_types/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type DocumentType struct {
	ID            int64                   `json:"id,required"`
	DocumentCount int64                   `json:"document_count,required"`
	Name          string                  `json:"name,required"`
	Permissions   DocumentTypePermissions `json:"permissions,required"`
	Slug          string                  `json:"slug,required"`
	UserCanChange bool                    `json:"user_can_change,required"`
	IsInsensitive bool                    `json:"is_insensitive"`
	Match         string                  `json:"match"`
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
		Permissions       respjson.Field
		Slug              respjson.Field
		UserCanChange     respjson.Field
		IsInsensitive     respjson.Field
		Match             respjson.Field
		MatchingAlgorithm respjson.Field
		Owner             respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentType) RawJSON() string { return r.JSON.raw }
func (r *DocumentType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentTypePermissions struct {
	Change DocumentTypePermissionsChange `json:"change"`
	View   DocumentTypePermissionsView   `json:"view"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Change      respjson.Field
		View        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentTypePermissions) RawJSON() string { return r.JSON.raw }
func (r *DocumentTypePermissions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentTypePermissionsChange struct {
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
func (r DocumentTypePermissionsChange) RawJSON() string { return r.JSON.raw }
func (r *DocumentTypePermissionsChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentTypePermissionsView struct {
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
func (r DocumentTypePermissionsView) RawJSON() string { return r.JSON.raw }
func (r *DocumentTypePermissionsView) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Name is required.
type DocumentTypeRequestParam struct {
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
	MatchingAlgorithm int64                                  `json:"matching_algorithm,omitzero"`
	SetPermissions    DocumentTypeRequestSetPermissionsParam `json:"set_permissions,omitzero"`
	paramObj
}

func (r DocumentTypeRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow DocumentTypeRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DocumentTypeRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[DocumentTypeRequestParam](
		"matching_algorithm", 0, 1, 2, 3, 4, 5, 6,
	)
}

type DocumentTypeRequestSetPermissionsParam struct {
	Change DocumentTypeRequestSetPermissionsChangeParam `json:"change,omitzero"`
	View   DocumentTypeRequestSetPermissionsViewParam   `json:"view,omitzero"`
	paramObj
}

func (r DocumentTypeRequestSetPermissionsParam) MarshalJSON() (data []byte, err error) {
	type shadow DocumentTypeRequestSetPermissionsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DocumentTypeRequestSetPermissionsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentTypeRequestSetPermissionsChangeParam struct {
	Groups []int64 `json:"groups,omitzero"`
	Users  []int64 `json:"users,omitzero"`
	paramObj
}

func (r DocumentTypeRequestSetPermissionsChangeParam) MarshalJSON() (data []byte, err error) {
	type shadow DocumentTypeRequestSetPermissionsChangeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DocumentTypeRequestSetPermissionsChangeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentTypeRequestSetPermissionsViewParam struct {
	Groups []int64 `json:"groups,omitzero"`
	Users  []int64 `json:"users,omitzero"`
	paramObj
}

func (r DocumentTypeRequestSetPermissionsViewParam) MarshalJSON() (data []byte, err error) {
	type shadow DocumentTypeRequestSetPermissionsViewParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DocumentTypeRequestSetPermissionsViewParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentTypeListResponse struct {
	Count    int64          `json:"count,required"`
	Results  []DocumentType `json:"results,required"`
	All      []any          `json:"all"`
	Next     string         `json:"next,nullable" format:"uri"`
	Previous string         `json:"previous,nullable" format:"uri"`
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
func (r DocumentTypeListResponse) RawJSON() string { return r.JSON.raw }
func (r *DocumentTypeListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentTypeNewParams struct {
	DocumentTypeRequest DocumentTypeRequestParam
	paramObj
}

func (r DocumentTypeNewParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.DocumentTypeRequest)
}
func (r *DocumentTypeNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.DocumentTypeRequest)
}

type DocumentTypeGetParams struct {
	FullPerms param.Opt[bool] `query:"full_perms,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DocumentTypeGetParams]'s query parameters as `url.Values`.
func (r DocumentTypeGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DocumentTypeUpdateParams struct {
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
	MatchingAlgorithm int64                                  `json:"matching_algorithm,omitzero"`
	SetPermissions    DocumentTypeUpdateParamsSetPermissions `json:"set_permissions,omitzero"`
	paramObj
}

func (r DocumentTypeUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow DocumentTypeUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DocumentTypeUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentTypeUpdateParamsSetPermissions struct {
	Change DocumentTypeUpdateParamsSetPermissionsChange `json:"change,omitzero"`
	View   DocumentTypeUpdateParamsSetPermissionsView   `json:"view,omitzero"`
	paramObj
}

func (r DocumentTypeUpdateParamsSetPermissions) MarshalJSON() (data []byte, err error) {
	type shadow DocumentTypeUpdateParamsSetPermissions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DocumentTypeUpdateParamsSetPermissions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentTypeUpdateParamsSetPermissionsChange struct {
	Groups []int64 `json:"groups,omitzero"`
	Users  []int64 `json:"users,omitzero"`
	paramObj
}

func (r DocumentTypeUpdateParamsSetPermissionsChange) MarshalJSON() (data []byte, err error) {
	type shadow DocumentTypeUpdateParamsSetPermissionsChange
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DocumentTypeUpdateParamsSetPermissionsChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentTypeUpdateParamsSetPermissionsView struct {
	Groups []int64 `json:"groups,omitzero"`
	Users  []int64 `json:"users,omitzero"`
	paramObj
}

func (r DocumentTypeUpdateParamsSetPermissionsView) MarshalJSON() (data []byte, err error) {
	type shadow DocumentTypeUpdateParamsSetPermissionsView
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DocumentTypeUpdateParamsSetPermissionsView) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentTypeListParams struct {
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

// URLQuery serializes [DocumentTypeListParams]'s query parameters as `url.Values`.
func (r DocumentTypeListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
