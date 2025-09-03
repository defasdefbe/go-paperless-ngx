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

// WorkflowTriggerService contains methods and other services that help with
// interacting with the paperless-ngx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkflowTriggerService] method instead.
type WorkflowTriggerService struct {
	Options []option.RequestOption
}

// NewWorkflowTriggerService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkflowTriggerService(opts ...option.RequestOption) (r WorkflowTriggerService) {
	r = WorkflowTriggerService{}
	r.Options = opts
	return
}

func (r *WorkflowTriggerService) New(ctx context.Context, body WorkflowTriggerNewParams, opts ...option.RequestOption) (res *WorkflowTrigger, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/workflow_triggers/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *WorkflowTriggerService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *WorkflowTrigger, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/workflow_triggers/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *WorkflowTriggerService) Update(ctx context.Context, id int64, body WorkflowTriggerUpdateParams, opts ...option.RequestOption) (res *WorkflowTrigger, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/workflow_triggers/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

func (r *WorkflowTriggerService) List(ctx context.Context, query WorkflowTriggerListParams, opts ...option.RequestOption) (res *WorkflowTriggerListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/workflow_triggers/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

func (r *WorkflowTriggerService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("api/workflow_triggers/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// - `added` - Added
// - `created` - Created
// - `modified` - Modified
// - `custom_field` - Custom Field
type ScheduleDateFieldEnum string

const (
	ScheduleDateFieldEnumAdded       ScheduleDateFieldEnum = "added"
	ScheduleDateFieldEnumCreated     ScheduleDateFieldEnum = "created"
	ScheduleDateFieldEnumModified    ScheduleDateFieldEnum = "modified"
	ScheduleDateFieldEnumCustomField ScheduleDateFieldEnum = "custom_field"
)

// - `1` - Consume Folder
// - `2` - Api Upload
// - `3` - Mail Fetch
// - `4` - Web UI
type SourcesEnum int64

const (
	SourcesEnum1 SourcesEnum = 1
	SourcesEnum2 SourcesEnum = 2
	SourcesEnum3 SourcesEnum = 3
	SourcesEnum4 SourcesEnum = 4
)

type WorkflowTrigger struct {
	// - `1` - Consumption Started
	// - `2` - Document Added
	// - `3` - Document Updated
	// - `4` - Scheduled
	//
	// Any of 1, 2, 3, 4.
	Type int64 `json:"type,required"`
	ID   int64 `json:"id,nullable"`
	// Only consume documents which entirely match this filename if specified.
	// Wildcards such as *.pdf or *invoice\* are allowed. Case insensitive.
	FilterFilename         string  `json:"filter_filename,nullable"`
	FilterHasCorrespondent int64   `json:"filter_has_correspondent,nullable"`
	FilterHasDocumentType  int64   `json:"filter_has_document_type,nullable"`
	FilterHasTags          []int64 `json:"filter_has_tags"`
	FilterMailrule         int64   `json:"filter_mailrule,nullable"`
	// Only consume documents with a path that matches this if specified. Wildcards
	// specified as \* are allowed. Case insensitive.
	FilterPath    string `json:"filter_path,nullable"`
	IsInsensitive bool   `json:"is_insensitive"`
	Match         string `json:"match"`
	// - `0` - None
	// - `1` - Any word
	// - `2` - All words
	// - `3` - Exact match
	// - `4` - Regular expression
	// - `5` - Fuzzy word
	//
	// Any of 0, 1, 2, 3, 4, 5.
	MatchingAlgorithm       int64 `json:"matching_algorithm"`
	ScheduleDateCustomField int64 `json:"schedule_date_custom_field,nullable"`
	// - `added` - Added
	// - `created` - Created
	// - `modified` - Modified
	// - `custom_field` - Custom Field
	//
	// Any of "added", "created", "modified", "custom_field".
	ScheduleDateField ScheduleDateFieldEnum `json:"schedule_date_field"`
	// If the schedule should be recurring.
	ScheduleIsRecurring bool `json:"schedule_is_recurring"`
	// The number of days to offset the schedule trigger by.
	ScheduleOffsetDays int64 `json:"schedule_offset_days"`
	// The number of days between recurring schedule triggers.
	ScheduleRecurringIntervalDays int64 `json:"schedule_recurring_interval_days"`
	// Any of 1, 2, 3, 4.
	Sources []int64 `json:"sources"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type                          respjson.Field
		ID                            respjson.Field
		FilterFilename                respjson.Field
		FilterHasCorrespondent        respjson.Field
		FilterHasDocumentType         respjson.Field
		FilterHasTags                 respjson.Field
		FilterMailrule                respjson.Field
		FilterPath                    respjson.Field
		IsInsensitive                 respjson.Field
		Match                         respjson.Field
		MatchingAlgorithm             respjson.Field
		ScheduleDateCustomField       respjson.Field
		ScheduleDateField             respjson.Field
		ScheduleIsRecurring           respjson.Field
		ScheduleOffsetDays            respjson.Field
		ScheduleRecurringIntervalDays respjson.Field
		Sources                       respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowTrigger) RawJSON() string { return r.JSON.raw }
func (r *WorkflowTrigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// - `0` - None
// - `1` - Any word
// - `2` - All words
// - `3` - Exact match
// - `4` - Regular expression
// - `5` - Fuzzy word
type WorkflowTriggerMatchingAlgorithmEnum int64

const (
	WorkflowTriggerMatchingAlgorithmEnum0 WorkflowTriggerMatchingAlgorithmEnum = 0
	WorkflowTriggerMatchingAlgorithmEnum1 WorkflowTriggerMatchingAlgorithmEnum = 1
	WorkflowTriggerMatchingAlgorithmEnum2 WorkflowTriggerMatchingAlgorithmEnum = 2
	WorkflowTriggerMatchingAlgorithmEnum3 WorkflowTriggerMatchingAlgorithmEnum = 3
	WorkflowTriggerMatchingAlgorithmEnum4 WorkflowTriggerMatchingAlgorithmEnum = 4
	WorkflowTriggerMatchingAlgorithmEnum5 WorkflowTriggerMatchingAlgorithmEnum = 5
)

// The property Type is required.
type WorkflowTriggerRequestParam struct {
	// - `1` - Consumption Started
	// - `2` - Document Added
	// - `3` - Document Updated
	// - `4` - Scheduled
	//
	// Any of 1, 2, 3, 4.
	Type int64            `json:"type,omitzero,required"`
	ID   param.Opt[int64] `json:"id,omitzero"`
	// Only consume documents which entirely match this filename if specified.
	// Wildcards such as *.pdf or *invoice\* are allowed. Case insensitive.
	FilterFilename         param.Opt[string] `json:"filter_filename,omitzero"`
	FilterHasCorrespondent param.Opt[int64]  `json:"filter_has_correspondent,omitzero"`
	FilterHasDocumentType  param.Opt[int64]  `json:"filter_has_document_type,omitzero"`
	FilterMailrule         param.Opt[int64]  `json:"filter_mailrule,omitzero"`
	// Only consume documents with a path that matches this if specified. Wildcards
	// specified as \* are allowed. Case insensitive.
	FilterPath              param.Opt[string] `json:"filter_path,omitzero"`
	ScheduleDateCustomField param.Opt[int64]  `json:"schedule_date_custom_field,omitzero"`
	IsInsensitive           param.Opt[bool]   `json:"is_insensitive,omitzero"`
	Match                   param.Opt[string] `json:"match,omitzero"`
	// If the schedule should be recurring.
	ScheduleIsRecurring param.Opt[bool] `json:"schedule_is_recurring,omitzero"`
	// The number of days to offset the schedule trigger by.
	ScheduleOffsetDays param.Opt[int64] `json:"schedule_offset_days,omitzero"`
	// The number of days between recurring schedule triggers.
	ScheduleRecurringIntervalDays param.Opt[int64] `json:"schedule_recurring_interval_days,omitzero"`
	FilterHasTags                 []int64          `json:"filter_has_tags,omitzero"`
	// - `0` - None
	// - `1` - Any word
	// - `2` - All words
	// - `3` - Exact match
	// - `4` - Regular expression
	// - `5` - Fuzzy word
	//
	// Any of 0, 1, 2, 3, 4, 5.
	MatchingAlgorithm int64 `json:"matching_algorithm,omitzero"`
	// - `added` - Added
	// - `created` - Created
	// - `modified` - Modified
	// - `custom_field` - Custom Field
	//
	// Any of "added", "created", "modified", "custom_field".
	ScheduleDateField ScheduleDateFieldEnum `json:"schedule_date_field,omitzero"`
	// Any of 1, 2, 3, 4.
	Sources []int64 `json:"sources,omitzero"`
	paramObj
}

func (r WorkflowTriggerRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowTriggerRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowTriggerRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowTriggerRequestParam](
		"type", 1, 2, 3, 4,
	)
	apijson.RegisterFieldValidator[WorkflowTriggerRequestParam](
		"matching_algorithm", 0, 1, 2, 3, 4, 5,
	)
}

// - `1` - Consumption Started
// - `2` - Document Added
// - `3` - Document Updated
// - `4` - Scheduled
type WorkflowTriggerTypeEnum int64

const (
	WorkflowTriggerTypeEnum1 WorkflowTriggerTypeEnum = 1
	WorkflowTriggerTypeEnum2 WorkflowTriggerTypeEnum = 2
	WorkflowTriggerTypeEnum3 WorkflowTriggerTypeEnum = 3
	WorkflowTriggerTypeEnum4 WorkflowTriggerTypeEnum = 4
)

type WorkflowTriggerListResponse struct {
	Count    int64             `json:"count,required"`
	Results  []WorkflowTrigger `json:"results,required"`
	All      []any             `json:"all"`
	Next     string            `json:"next,nullable" format:"uri"`
	Previous string            `json:"previous,nullable" format:"uri"`
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
func (r WorkflowTriggerListResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowTriggerListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowTriggerNewParams struct {
	WorkflowTriggerRequest WorkflowTriggerRequestParam
	paramObj
}

func (r WorkflowTriggerNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.WorkflowTriggerRequest)
}
func (r *WorkflowTriggerNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.WorkflowTriggerRequest)
}

type WorkflowTriggerUpdateParams struct {
	ID param.Opt[int64] `json:"id,omitzero"`
	// Only consume documents which entirely match this filename if specified.
	// Wildcards such as *.pdf or *invoice\* are allowed. Case insensitive.
	FilterFilename         param.Opt[string] `json:"filter_filename,omitzero"`
	FilterHasCorrespondent param.Opt[int64]  `json:"filter_has_correspondent,omitzero"`
	FilterHasDocumentType  param.Opt[int64]  `json:"filter_has_document_type,omitzero"`
	FilterMailrule         param.Opt[int64]  `json:"filter_mailrule,omitzero"`
	// Only consume documents with a path that matches this if specified. Wildcards
	// specified as \* are allowed. Case insensitive.
	FilterPath              param.Opt[string] `json:"filter_path,omitzero"`
	ScheduleDateCustomField param.Opt[int64]  `json:"schedule_date_custom_field,omitzero"`
	IsInsensitive           param.Opt[bool]   `json:"is_insensitive,omitzero"`
	Match                   param.Opt[string] `json:"match,omitzero"`
	// If the schedule should be recurring.
	ScheduleIsRecurring param.Opt[bool] `json:"schedule_is_recurring,omitzero"`
	// The number of days to offset the schedule trigger by.
	ScheduleOffsetDays param.Opt[int64] `json:"schedule_offset_days,omitzero"`
	// The number of days between recurring schedule triggers.
	ScheduleRecurringIntervalDays param.Opt[int64] `json:"schedule_recurring_interval_days,omitzero"`
	FilterHasTags                 []int64          `json:"filter_has_tags,omitzero"`
	// - `0` - None
	// - `1` - Any word
	// - `2` - All words
	// - `3` - Exact match
	// - `4` - Regular expression
	// - `5` - Fuzzy word
	//
	// Any of 0, 1, 2, 3, 4, 5.
	MatchingAlgorithm int64 `json:"matching_algorithm,omitzero"`
	// - `added` - Added
	// - `created` - Created
	// - `modified` - Modified
	// - `custom_field` - Custom Field
	//
	// Any of "added", "created", "modified", "custom_field".
	ScheduleDateField ScheduleDateFieldEnum `json:"schedule_date_field,omitzero"`
	// Any of 1, 2, 3, 4.
	Sources []int64 `json:"sources,omitzero"`
	// - `1` - Consumption Started
	// - `2` - Document Added
	// - `3` - Document Updated
	// - `4` - Scheduled
	//
	// Any of 1, 2, 3, 4.
	Type int64 `json:"type,omitzero"`
	paramObj
}

func (r WorkflowTriggerUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowTriggerUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowTriggerUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowTriggerListParams struct {
	// A page number within the paginated result set.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Number of results to return per page.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkflowTriggerListParams]'s query parameters as
// `url.Values`.
func (r WorkflowTriggerListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
