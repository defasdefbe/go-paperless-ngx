// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package paperlessngx

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

// StoragePathService contains methods and other services that help with
// interacting with the paperless-ngx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStoragePathService] method instead.
type StoragePathService struct {
	Options []option.RequestOption
}

// NewStoragePathService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStoragePathService(opts ...option.RequestOption) (r StoragePathService) {
	r = StoragePathService{}
	r.Options = opts
	return
}

func (r *StoragePathService) New(ctx context.Context, body StoragePathNewParams, opts ...option.RequestOption) (res *StoragePath, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/storage_paths/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *StoragePathService) Get(ctx context.Context, id int64, query StoragePathGetParams, opts ...option.RequestOption) (res *StoragePath, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/storage_paths/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

func (r *StoragePathService) Update(ctx context.Context, id int64, body StoragePathUpdateParams, opts ...option.RequestOption) (res *StoragePath, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/storage_paths/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

func (r *StoragePathService) List(ctx context.Context, query StoragePathListParams, opts ...option.RequestOption) (res *StoragePathListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/storage_paths/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// When a storage path is deleted, see if documents using it require a rename/move
func (r *StoragePathService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("api/storage_paths/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Test storage path against a document
func (r *StoragePathService) Test(ctx context.Context, body StoragePathTestParams, opts ...option.RequestOption) (res *StoragePath, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/storage_paths/test/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type StoragePath struct {
	ID            int64  `json:"id,required"`
	DocumentCount int64  `json:"document_count,required"`
	Name          string `json:"name,required"`
	Path          string `json:"path,required"`
	Slug          string `json:"slug,required"`
	UserCanChange bool   `json:"user_can_change,required"`
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
		Path              respjson.Field
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
func (r StoragePath) RawJSON() string { return r.JSON.raw }
func (r *StoragePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Path are required.
type StoragePathRequestParam struct {
	Name          string            `json:"name,required"`
	Path          string            `json:"path,required"`
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
	MatchingAlgorithm int64                                 `json:"matching_algorithm,omitzero"`
	SetPermissions    StoragePathRequestSetPermissionsParam `json:"set_permissions,omitzero"`
	paramObj
}

func (r StoragePathRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow StoragePathRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StoragePathRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[StoragePathRequestParam](
		"matching_algorithm", 0, 1, 2, 3, 4, 5, 6,
	)
}

type StoragePathRequestSetPermissionsParam struct {
	Change StoragePathRequestSetPermissionsChangeParam `json:"change,omitzero"`
	View   StoragePathRequestSetPermissionsViewParam   `json:"view,omitzero"`
	paramObj
}

func (r StoragePathRequestSetPermissionsParam) MarshalJSON() (data []byte, err error) {
	type shadow StoragePathRequestSetPermissionsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StoragePathRequestSetPermissionsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StoragePathRequestSetPermissionsChangeParam struct {
	Groups []int64 `json:"groups,omitzero"`
	Users  []int64 `json:"users,omitzero"`
	paramObj
}

func (r StoragePathRequestSetPermissionsChangeParam) MarshalJSON() (data []byte, err error) {
	type shadow StoragePathRequestSetPermissionsChangeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StoragePathRequestSetPermissionsChangeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StoragePathRequestSetPermissionsViewParam struct {
	Groups []int64 `json:"groups,omitzero"`
	Users  []int64 `json:"users,omitzero"`
	paramObj
}

func (r StoragePathRequestSetPermissionsViewParam) MarshalJSON() (data []byte, err error) {
	type shadow StoragePathRequestSetPermissionsViewParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StoragePathRequestSetPermissionsViewParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StoragePathListResponse struct {
	Count    int64         `json:"count,required"`
	Results  []StoragePath `json:"results,required"`
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
func (r StoragePathListResponse) RawJSON() string { return r.JSON.raw }
func (r *StoragePathListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StoragePathNewParams struct {
	StoragePathRequest StoragePathRequestParam
	paramObj
}

func (r StoragePathNewParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.StoragePathRequest)
}
func (r *StoragePathNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.StoragePathRequest)
}

type StoragePathGetParams struct {
	FullPerms param.Opt[bool] `query:"full_perms,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StoragePathGetParams]'s query parameters as `url.Values`.
func (r StoragePathGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type StoragePathUpdateParams struct {
	Owner         param.Opt[int64]  `json:"owner,omitzero"`
	IsInsensitive param.Opt[bool]   `json:"is_insensitive,omitzero"`
	Match         param.Opt[string] `json:"match,omitzero"`
	Name          param.Opt[string] `json:"name,omitzero"`
	Path          param.Opt[string] `json:"path,omitzero"`
	// - `0` - None
	// - `1` - Any word
	// - `2` - All words
	// - `3` - Exact match
	// - `4` - Regular expression
	// - `5` - Fuzzy word
	// - `6` - Automatic
	//
	// Any of 0, 1, 2, 3, 4, 5, 6.
	MatchingAlgorithm int64                                 `json:"matching_algorithm,omitzero"`
	SetPermissions    StoragePathUpdateParamsSetPermissions `json:"set_permissions,omitzero"`
	paramObj
}

func (r StoragePathUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow StoragePathUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StoragePathUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StoragePathUpdateParamsSetPermissions struct {
	Change StoragePathUpdateParamsSetPermissionsChange `json:"change,omitzero"`
	View   StoragePathUpdateParamsSetPermissionsView   `json:"view,omitzero"`
	paramObj
}

func (r StoragePathUpdateParamsSetPermissions) MarshalJSON() (data []byte, err error) {
	type shadow StoragePathUpdateParamsSetPermissions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StoragePathUpdateParamsSetPermissions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StoragePathUpdateParamsSetPermissionsChange struct {
	Groups []int64 `json:"groups,omitzero"`
	Users  []int64 `json:"users,omitzero"`
	paramObj
}

func (r StoragePathUpdateParamsSetPermissionsChange) MarshalJSON() (data []byte, err error) {
	type shadow StoragePathUpdateParamsSetPermissionsChange
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StoragePathUpdateParamsSetPermissionsChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StoragePathUpdateParamsSetPermissionsView struct {
	Groups []int64 `json:"groups,omitzero"`
	Users  []int64 `json:"users,omitzero"`
	paramObj
}

func (r StoragePathUpdateParamsSetPermissionsView) MarshalJSON() (data []byte, err error) {
	type shadow StoragePathUpdateParamsSetPermissionsView
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StoragePathUpdateParamsSetPermissionsView) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StoragePathListParams struct {
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
	PageSize        param.Opt[int64]  `query:"page_size,omitzero" json:"-"`
	PathIcontains   param.Opt[string] `query:"path__icontains,omitzero" json:"-"`
	PathIendswith   param.Opt[string] `query:"path__iendswith,omitzero" json:"-"`
	PathIexact      param.Opt[string] `query:"path__iexact,omitzero" json:"-"`
	PathIstartswith param.Opt[string] `query:"path__istartswith,omitzero" json:"-"`
	// Multiple values may be separated by commas.
	IDIn []int64 `query:"id__in,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StoragePathListParams]'s query parameters as `url.Values`.
func (r StoragePathListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type StoragePathTestParams struct {
	StoragePathRequest StoragePathRequestParam
	paramObj
}

func (r StoragePathTestParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.StoragePathRequest)
}
func (r *StoragePathTestParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.StoragePathRequest)
}
