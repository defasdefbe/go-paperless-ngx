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

// WorkflowActionService contains methods and other services that help with
// interacting with the paperless-ngx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkflowActionService] method instead.
type WorkflowActionService struct {
	Options []option.RequestOption
}

// NewWorkflowActionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkflowActionService(opts ...option.RequestOption) (r WorkflowActionService) {
	r = WorkflowActionService{}
	r.Options = opts
	return
}

func (r *WorkflowActionService) New(ctx context.Context, body WorkflowActionNewParams, opts ...option.RequestOption) (res *WorkflowAction, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/workflow_actions/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *WorkflowActionService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *WorkflowAction, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/workflow_actions/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *WorkflowActionService) Update(ctx context.Context, id int64, body WorkflowActionUpdateParams, opts ...option.RequestOption) (res *WorkflowAction, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/workflow_actions/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

func (r *WorkflowActionService) List(ctx context.Context, query WorkflowActionListParams, opts ...option.RequestOption) (res *WorkflowActionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/workflow_actions/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

func (r *WorkflowActionService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("api/workflow_actions/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type WorkflowAction struct {
	ID                  int64   `json:"id,nullable"`
	AssignChangeGroups  []int64 `json:"assign_change_groups"`
	AssignChangeUsers   []int64 `json:"assign_change_users"`
	AssignCorrespondent int64   `json:"assign_correspondent,nullable"`
	AssignCustomFields  []int64 `json:"assign_custom_fields"`
	// Optional values to assign to the custom fields.
	AssignCustomFieldsValues any     `json:"assign_custom_fields_values"`
	AssignDocumentType       int64   `json:"assign_document_type,nullable"`
	AssignOwner              int64   `json:"assign_owner,nullable"`
	AssignStoragePath        int64   `json:"assign_storage_path,nullable"`
	AssignTags               []int64 `json:"assign_tags"`
	// Assign a document title, can include some placeholders, see documentation.
	AssignTitle             string              `json:"assign_title,nullable"`
	AssignViewGroups        []int64             `json:"assign_view_groups"`
	AssignViewUsers         []int64             `json:"assign_view_users"`
	Email                   WorkflowActionEmail `json:"email,nullable"`
	RemoveAllCorrespondents bool                `json:"remove_all_correspondents"`
	RemoveAllCustomFields   bool                `json:"remove_all_custom_fields"`
	RemoveAllDocumentTypes  bool                `json:"remove_all_document_types"`
	RemoveAllOwners         bool                `json:"remove_all_owners"`
	RemoveAllPermissions    bool                `json:"remove_all_permissions"`
	RemoveAllStoragePaths   bool                `json:"remove_all_storage_paths"`
	RemoveAllTags           bool                `json:"remove_all_tags"`
	RemoveChangeGroups      []int64             `json:"remove_change_groups"`
	RemoveChangeUsers       []int64             `json:"remove_change_users"`
	RemoveCorrespondents    []int64             `json:"remove_correspondents"`
	RemoveCustomFields      []int64             `json:"remove_custom_fields"`
	RemoveDocumentTypes     []int64             `json:"remove_document_types"`
	RemoveOwners            []int64             `json:"remove_owners"`
	RemoveStoragePaths      []int64             `json:"remove_storage_paths"`
	RemoveTags              []int64             `json:"remove_tags"`
	RemoveViewGroups        []int64             `json:"remove_view_groups"`
	RemoveViewUsers         []int64             `json:"remove_view_users"`
	// - `1` - Assignment
	// - `2` - Removal
	// - `3` - Email
	// - `4` - Webhook
	//
	// Any of 1, 2, 3, 4.
	Type    int64                 `json:"type"`
	Webhook WorkflowActionWebhook `json:"webhook,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                       respjson.Field
		AssignChangeGroups       respjson.Field
		AssignChangeUsers        respjson.Field
		AssignCorrespondent      respjson.Field
		AssignCustomFields       respjson.Field
		AssignCustomFieldsValues respjson.Field
		AssignDocumentType       respjson.Field
		AssignOwner              respjson.Field
		AssignStoragePath        respjson.Field
		AssignTags               respjson.Field
		AssignTitle              respjson.Field
		AssignViewGroups         respjson.Field
		AssignViewUsers          respjson.Field
		Email                    respjson.Field
		RemoveAllCorrespondents  respjson.Field
		RemoveAllCustomFields    respjson.Field
		RemoveAllDocumentTypes   respjson.Field
		RemoveAllOwners          respjson.Field
		RemoveAllPermissions     respjson.Field
		RemoveAllStoragePaths    respjson.Field
		RemoveAllTags            respjson.Field
		RemoveChangeGroups       respjson.Field
		RemoveChangeUsers        respjson.Field
		RemoveCorrespondents     respjson.Field
		RemoveCustomFields       respjson.Field
		RemoveDocumentTypes      respjson.Field
		RemoveOwners             respjson.Field
		RemoveStoragePaths       respjson.Field
		RemoveTags               respjson.Field
		RemoveViewGroups         respjson.Field
		RemoveViewUsers          respjson.Field
		Type                     respjson.Field
		Webhook                  respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowAction) RawJSON() string { return r.JSON.raw }
func (r *WorkflowAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowActionEmail struct {
	// The body (message) of the email, can include some placeholders, see
	// documentation.
	Body string `json:"body,required"`
	// The subject of the email, can include some placeholders, see documentation.
	Subject string `json:"subject,required"`
	// The destination email addresses, comma separated.
	To              string `json:"to,required"`
	ID              int64  `json:"id,nullable"`
	IncludeDocument bool   `json:"include_document"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body            respjson.Field
		Subject         respjson.Field
		To              respjson.Field
		ID              respjson.Field
		IncludeDocument respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowActionEmail) RawJSON() string { return r.JSON.raw }
func (r *WorkflowActionEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowActionWebhook struct {
	// The destination URL for the notification.
	URL    string `json:"url,required"`
	ID     int64  `json:"id,nullable"`
	AsJson bool   `json:"as_json"`
	// The body to send with the webhook URL if parameters not used.
	Body string `json:"body,nullable"`
	// The headers to send with the webhook URL.
	Headers         any  `json:"headers"`
	IncludeDocument bool `json:"include_document"`
	// The parameters to send with the webhook URL if body not used.
	Params    any  `json:"params"`
	UseParams bool `json:"use_params"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL             respjson.Field
		ID              respjson.Field
		AsJson          respjson.Field
		Body            respjson.Field
		Headers         respjson.Field
		IncludeDocument respjson.Field
		Params          respjson.Field
		UseParams       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowActionWebhook) RawJSON() string { return r.JSON.raw }
func (r *WorkflowActionWebhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Body, Subject, To are required.
type WorkflowActionEmailParam struct {
	// The body (message) of the email, can include some placeholders, see
	// documentation.
	Body string `json:"body,required"`
	// The subject of the email, can include some placeholders, see documentation.
	Subject string `json:"subject,required"`
	// The destination email addresses, comma separated.
	To              string           `json:"to,required"`
	ID              param.Opt[int64] `json:"id,omitzero"`
	IncludeDocument param.Opt[bool]  `json:"include_document,omitzero"`
	paramObj
}

func (r WorkflowActionEmailParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowActionEmailParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowActionEmailParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowActionRequestParam struct {
	ID                  param.Opt[int64] `json:"id,omitzero"`
	AssignCorrespondent param.Opt[int64] `json:"assign_correspondent,omitzero"`
	AssignDocumentType  param.Opt[int64] `json:"assign_document_type,omitzero"`
	AssignOwner         param.Opt[int64] `json:"assign_owner,omitzero"`
	AssignStoragePath   param.Opt[int64] `json:"assign_storage_path,omitzero"`
	// Assign a document title, can include some placeholders, see documentation.
	AssignTitle             param.Opt[string] `json:"assign_title,omitzero"`
	RemoveAllCorrespondents param.Opt[bool]   `json:"remove_all_correspondents,omitzero"`
	RemoveAllCustomFields   param.Opt[bool]   `json:"remove_all_custom_fields,omitzero"`
	RemoveAllDocumentTypes  param.Opt[bool]   `json:"remove_all_document_types,omitzero"`
	RemoveAllOwners         param.Opt[bool]   `json:"remove_all_owners,omitzero"`
	RemoveAllPermissions    param.Opt[bool]   `json:"remove_all_permissions,omitzero"`
	RemoveAllStoragePaths   param.Opt[bool]   `json:"remove_all_storage_paths,omitzero"`
	RemoveAllTags           param.Opt[bool]   `json:"remove_all_tags,omitzero"`
	AssignChangeGroups      []int64           `json:"assign_change_groups,omitzero"`
	AssignChangeUsers       []int64           `json:"assign_change_users,omitzero"`
	AssignCustomFields      []int64           `json:"assign_custom_fields,omitzero"`
	// Optional values to assign to the custom fields.
	AssignCustomFieldsValues any                      `json:"assign_custom_fields_values,omitzero"`
	AssignTags               []int64                  `json:"assign_tags,omitzero"`
	AssignViewGroups         []int64                  `json:"assign_view_groups,omitzero"`
	AssignViewUsers          []int64                  `json:"assign_view_users,omitzero"`
	Email                    WorkflowActionEmailParam `json:"email,omitzero"`
	RemoveChangeGroups       []int64                  `json:"remove_change_groups,omitzero"`
	RemoveChangeUsers        []int64                  `json:"remove_change_users,omitzero"`
	RemoveCorrespondents     []int64                  `json:"remove_correspondents,omitzero"`
	RemoveCustomFields       []int64                  `json:"remove_custom_fields,omitzero"`
	RemoveDocumentTypes      []int64                  `json:"remove_document_types,omitzero"`
	RemoveOwners             []int64                  `json:"remove_owners,omitzero"`
	RemoveStoragePaths       []int64                  `json:"remove_storage_paths,omitzero"`
	RemoveTags               []int64                  `json:"remove_tags,omitzero"`
	RemoveViewGroups         []int64                  `json:"remove_view_groups,omitzero"`
	RemoveViewUsers          []int64                  `json:"remove_view_users,omitzero"`
	// - `1` - Assignment
	// - `2` - Removal
	// - `3` - Email
	// - `4` - Webhook
	//
	// Any of 1, 2, 3, 4.
	Type    int64                      `json:"type,omitzero"`
	Webhook WorkflowActionWebhookParam `json:"webhook,omitzero"`
	paramObj
}

func (r WorkflowActionRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowActionRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowActionRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowActionRequestParam](
		"type", 1, 2, 3, 4,
	)
}

// - `1` - Assignment
// - `2` - Removal
// - `3` - Email
// - `4` - Webhook
type WorkflowActionType int64

const (
	WorkflowActionType1 WorkflowActionType = 1
	WorkflowActionType2 WorkflowActionType = 2
	WorkflowActionType3 WorkflowActionType = 3
	WorkflowActionType4 WorkflowActionType = 4
)

// The property URL is required.
type WorkflowActionWebhookParam struct {
	// The destination URL for the notification.
	URL string           `json:"url,required"`
	ID  param.Opt[int64] `json:"id,omitzero"`
	// The body to send with the webhook URL if parameters not used.
	Body            param.Opt[string] `json:"body,omitzero"`
	AsJson          param.Opt[bool]   `json:"as_json,omitzero"`
	IncludeDocument param.Opt[bool]   `json:"include_document,omitzero"`
	UseParams       param.Opt[bool]   `json:"use_params,omitzero"`
	// The headers to send with the webhook URL.
	Headers any `json:"headers,omitzero"`
	// The parameters to send with the webhook URL if body not used.
	Params any `json:"params,omitzero"`
	paramObj
}

func (r WorkflowActionWebhookParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowActionWebhookParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowActionWebhookParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowActionListResponse struct {
	Count    int64            `json:"count,required"`
	Results  []WorkflowAction `json:"results,required"`
	All      []any            `json:"all"`
	Next     string           `json:"next,nullable" format:"uri"`
	Previous string           `json:"previous,nullable" format:"uri"`
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
func (r WorkflowActionListResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowActionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowActionNewParams struct {
	WorkflowActionRequest WorkflowActionRequestParam
	paramObj
}

func (r WorkflowActionNewParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.WorkflowActionRequest)
}
func (r *WorkflowActionNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.WorkflowActionRequest)
}

type WorkflowActionUpdateParams struct {
	ID                  param.Opt[int64] `json:"id,omitzero"`
	AssignCorrespondent param.Opt[int64] `json:"assign_correspondent,omitzero"`
	AssignDocumentType  param.Opt[int64] `json:"assign_document_type,omitzero"`
	AssignOwner         param.Opt[int64] `json:"assign_owner,omitzero"`
	AssignStoragePath   param.Opt[int64] `json:"assign_storage_path,omitzero"`
	// Assign a document title, can include some placeholders, see documentation.
	AssignTitle             param.Opt[string] `json:"assign_title,omitzero"`
	RemoveAllCorrespondents param.Opt[bool]   `json:"remove_all_correspondents,omitzero"`
	RemoveAllCustomFields   param.Opt[bool]   `json:"remove_all_custom_fields,omitzero"`
	RemoveAllDocumentTypes  param.Opt[bool]   `json:"remove_all_document_types,omitzero"`
	RemoveAllOwners         param.Opt[bool]   `json:"remove_all_owners,omitzero"`
	RemoveAllPermissions    param.Opt[bool]   `json:"remove_all_permissions,omitzero"`
	RemoveAllStoragePaths   param.Opt[bool]   `json:"remove_all_storage_paths,omitzero"`
	RemoveAllTags           param.Opt[bool]   `json:"remove_all_tags,omitzero"`
	AssignChangeGroups      []int64           `json:"assign_change_groups,omitzero"`
	AssignChangeUsers       []int64           `json:"assign_change_users,omitzero"`
	AssignCustomFields      []int64           `json:"assign_custom_fields,omitzero"`
	// Optional values to assign to the custom fields.
	AssignCustomFieldsValues any                      `json:"assign_custom_fields_values,omitzero"`
	AssignTags               []int64                  `json:"assign_tags,omitzero"`
	AssignViewGroups         []int64                  `json:"assign_view_groups,omitzero"`
	AssignViewUsers          []int64                  `json:"assign_view_users,omitzero"`
	Email                    WorkflowActionEmailParam `json:"email,omitzero"`
	RemoveChangeGroups       []int64                  `json:"remove_change_groups,omitzero"`
	RemoveChangeUsers        []int64                  `json:"remove_change_users,omitzero"`
	RemoveCorrespondents     []int64                  `json:"remove_correspondents,omitzero"`
	RemoveCustomFields       []int64                  `json:"remove_custom_fields,omitzero"`
	RemoveDocumentTypes      []int64                  `json:"remove_document_types,omitzero"`
	RemoveOwners             []int64                  `json:"remove_owners,omitzero"`
	RemoveStoragePaths       []int64                  `json:"remove_storage_paths,omitzero"`
	RemoveTags               []int64                  `json:"remove_tags,omitzero"`
	RemoveViewGroups         []int64                  `json:"remove_view_groups,omitzero"`
	RemoveViewUsers          []int64                  `json:"remove_view_users,omitzero"`
	// - `1` - Assignment
	// - `2` - Removal
	// - `3` - Email
	// - `4` - Webhook
	//
	// Any of 1, 2, 3, 4.
	Type    int64                      `json:"type,omitzero"`
	Webhook WorkflowActionWebhookParam `json:"webhook,omitzero"`
	paramObj
}

func (r WorkflowActionUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowActionUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowActionUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowActionListParams struct {
	// A page number within the paginated result set.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Number of results to return per page.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkflowActionListParams]'s query parameters as
// `url.Values`.
func (r WorkflowActionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
