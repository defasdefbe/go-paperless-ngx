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

// WorkflowService contains methods and other services that help with interacting
// with the paperless-ngx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkflowService] method instead.
type WorkflowService struct {
	Options []option.RequestOption
}

// NewWorkflowService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkflowService(opts ...option.RequestOption) (r WorkflowService) {
	r = WorkflowService{}
	r.Options = opts
	return
}

func (r *WorkflowService) New(ctx context.Context, body WorkflowNewParams, opts ...option.RequestOption) (res *Workflow, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/workflows/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *WorkflowService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *Workflow, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/workflows/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *WorkflowService) Update(ctx context.Context, id int64, body WorkflowUpdateParams, opts ...option.RequestOption) (res *Workflow, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/workflows/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

func (r *WorkflowService) List(ctx context.Context, query WorkflowListParams, opts ...option.RequestOption) (res *WorkflowListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/workflows/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

func (r *WorkflowService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("api/workflows/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type Workflow struct {
	ID       int64             `json:"id,required"`
	Actions  []WorkflowAction  `json:"actions,required"`
	Name     string            `json:"name,required"`
	Triggers []WorkflowTrigger `json:"triggers,required"`
	Enabled  bool              `json:"enabled"`
	Order    int64             `json:"order"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Actions     respjson.Field
		Name        respjson.Field
		Triggers    respjson.Field
		Enabled     respjson.Field
		Order       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Workflow) RawJSON() string { return r.JSON.raw }
func (r *Workflow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Actions, Name, Triggers are required.
type WorkflowRequestParam struct {
	Actions  []WorkflowActionRequestParam  `json:"actions,omitzero,required"`
	Name     string                        `json:"name,required"`
	Triggers []WorkflowTriggerRequestParam `json:"triggers,omitzero,required"`
	Enabled  param.Opt[bool]               `json:"enabled,omitzero"`
	Order    param.Opt[int64]              `json:"order,omitzero"`
	paramObj
}

func (r WorkflowRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowListResponse struct {
	Count    int64      `json:"count,required"`
	Results  []Workflow `json:"results,required"`
	All      []any      `json:"all"`
	Next     string     `json:"next,nullable" format:"uri"`
	Previous string     `json:"previous,nullable" format:"uri"`
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
func (r WorkflowListResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowNewParams struct {
	WorkflowRequest WorkflowRequestParam
	paramObj
}

func (r WorkflowNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.WorkflowRequest)
}
func (r *WorkflowNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.WorkflowRequest)
}

type WorkflowUpdateParams struct {
	Enabled  param.Opt[bool]               `json:"enabled,omitzero"`
	Name     param.Opt[string]             `json:"name,omitzero"`
	Order    param.Opt[int64]              `json:"order,omitzero"`
	Actions  []WorkflowActionRequestParam  `json:"actions,omitzero"`
	Triggers []WorkflowTriggerRequestParam `json:"triggers,omitzero"`
	paramObj
}

func (r WorkflowUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowListParams struct {
	// A page number within the paginated result set.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Number of results to return per page.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkflowListParams]'s query parameters as `url.Values`.
func (r WorkflowListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
