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

// GroupService contains methods and other services that help with interacting with
// the paperless-ngx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGroupService] method instead.
type GroupService struct {
	Options []option.RequestOption
}

// NewGroupService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewGroupService(opts ...option.RequestOption) (r GroupService) {
	r = GroupService{}
	r.Options = opts
	return
}

func (r *GroupService) New(ctx context.Context, body GroupNewParams, opts ...option.RequestOption) (res *Group, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/groups/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *GroupService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *Group, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/groups/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *GroupService) Update(ctx context.Context, id int64, body GroupUpdateParams, opts ...option.RequestOption) (res *Group, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/groups/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

func (r *GroupService) List(ctx context.Context, query GroupListParams, opts ...option.RequestOption) (res *GroupListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/groups/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

func (r *GroupService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("api/groups/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type Group struct {
	ID          int64    `json:"id,required"`
	Name        string   `json:"name,required"`
	Permissions []string `json:"permissions,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Permissions respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Group) RawJSON() string { return r.JSON.raw }
func (r *Group) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Permissions are required.
type GroupRequestParam struct {
	Name        string   `json:"name,required"`
	Permissions []string `json:"permissions,omitzero,required"`
	paramObj
}

func (r GroupRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow GroupRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GroupRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GroupListResponse struct {
	Count    int64   `json:"count,required"`
	Results  []Group `json:"results,required"`
	All      []any   `json:"all"`
	Next     string  `json:"next,nullable" format:"uri"`
	Previous string  `json:"previous,nullable" format:"uri"`
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
func (r GroupListResponse) RawJSON() string { return r.JSON.raw }
func (r *GroupListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GroupNewParams struct {
	GroupRequest GroupRequestParam
	paramObj
}

func (r GroupNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.GroupRequest)
}
func (r *GroupNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.GroupRequest)
}

type GroupUpdateParams struct {
	Name        param.Opt[string] `json:"name,omitzero"`
	Permissions []string          `json:"permissions,omitzero"`
	paramObj
}

func (r GroupUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow GroupUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GroupUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GroupListParams struct {
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
	paramObj
}

// URLQuery serializes [GroupListParams]'s query parameters as `url.Values`.
func (r GroupListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
