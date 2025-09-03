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

// MailRuleService contains methods and other services that help with interacting
// with the paperless-ngx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMailRuleService] method instead.
type MailRuleService struct {
	Options []option.RequestOption
}

// NewMailRuleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMailRuleService(opts ...option.RequestOption) (r MailRuleService) {
	r = MailRuleService{}
	r.Options = opts
	return
}

func (r *MailRuleService) New(ctx context.Context, body MailRuleNewParams, opts ...option.RequestOption) (res *MailRule, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/mail_rules/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

func (r *MailRuleService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *MailRule, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/mail_rules/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

func (r *MailRuleService) Update(ctx context.Context, id int64, body MailRuleUpdateParams, opts ...option.RequestOption) (res *MailRule, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/mail_rules/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

func (r *MailRuleService) List(ctx context.Context, query MailRuleListParams, opts ...option.RequestOption) (res *MailRuleListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/mail_rules/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

func (r *MailRuleService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("api/mail_rules/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// - `1` - Do not assign a correspondent
// - `2` - Use mail address
// - `3` - Use name (or mail address if not available)
// - `4` - Use correspondent selected below
type AssignCorrespondent int64

const (
	AssignCorrespondent1 AssignCorrespondent = 1
	AssignCorrespondent2 AssignCorrespondent = 2
	AssignCorrespondent3 AssignCorrespondent = 3
	AssignCorrespondent4 AssignCorrespondent = 4
)

// - `1` - Use subject as title
// - `2` - Use attachment filename as title
// - `3` - Do not assign title from rule
type AssignTitle int64

const (
	AssignTitle1 AssignTitle = 1
	AssignTitle2 AssignTitle = 2
	AssignTitle3 AssignTitle = 3
)

// - `1` - Only process attachments.
// - `2` - Process all files, including 'inline' attachments.
type AttachmentType int64

const (
	AttachmentType1 AttachmentType = 1
	AttachmentType2 AttachmentType = 2
)

//   - `1` - Only process attachments.
//   - `2` - Process full Mail (with embedded attachments in file) as .eml
//   - `3` - Process full Mail (with embedded attachments in file) as .eml + process
//     attachments as separate documents
type ConsumptionScope int64

const (
	ConsumptionScope1 ConsumptionScope = 1
	ConsumptionScope2 ConsumptionScope = 2
	ConsumptionScope3 ConsumptionScope = 3
)

type MailRule struct {
	ID            int64  `json:"id,required"`
	Account       int64  `json:"account,required"`
	Name          string `json:"name,required"`
	UserCanChange bool   `json:"user_can_change,required"`
	// - `1` - Delete
	// - `2` - Move to specified folder
	// - `3` - Mark as read, don't process read mails
	// - `4` - Flag the mail, don't process flagged mails
	// - `5` - Tag the mail with specified tag, don't process tagged mails
	//
	// Any of 1, 2, 3, 4, 5.
	Action              int64  `json:"action"`
	ActionParameter     string `json:"action_parameter,nullable"`
	AssignCorrespondent int64  `json:"assign_correspondent,nullable"`
	// - `1` - Do not assign a correspondent
	// - `2` - Use mail address
	// - `3` - Use name (or mail address if not available)
	// - `4` - Use correspondent selected below
	//
	// Any of 1, 2, 3, 4.
	AssignCorrespondentFrom int64   `json:"assign_correspondent_from"`
	AssignDocumentType      int64   `json:"assign_document_type,nullable"`
	AssignOwnerFromRule     bool    `json:"assign_owner_from_rule"`
	AssignTags              []int64 `json:"assign_tags"`
	// - `1` - Use subject as title
	// - `2` - Use attachment filename as title
	// - `3` - Do not assign title from rule
	//
	// Any of 1, 2, 3.
	AssignTitleFrom int64 `json:"assign_title_from"`
	// - `1` - Only process attachments.
	// - `2` - Process all files, including 'inline' attachments.
	//
	// Any of 1, 2.
	AttachmentType int64 `json:"attachment_type"`
	//   - `1` - Only process attachments.
	//   - `2` - Process full Mail (with embedded attachments in file) as .eml
	//   - `3` - Process full Mail (with embedded attachments in file) as .eml + process
	//     attachments as separate documents
	//
	// Any of 1, 2, 3.
	ConsumptionScope int64 `json:"consumption_scope"`
	Enabled          bool  `json:"enabled"`
	// Do not consume documents which entirely match this filename if specified.
	// Wildcards such as *.pdf or *invoice\* are allowed. Case insensitive.
	FilterAttachmentFilenameExclude string `json:"filter_attachment_filename_exclude,nullable"`
	// Only consume documents which entirely match this filename if specified.
	// Wildcards such as *.pdf or *invoice\* are allowed. Case insensitive.
	FilterAttachmentFilenameInclude string `json:"filter_attachment_filename_include,nullable"`
	FilterBody                      string `json:"filter_body,nullable"`
	FilterFrom                      string `json:"filter_from,nullable"`
	FilterSubject                   string `json:"filter_subject,nullable"`
	FilterTo                        string `json:"filter_to,nullable"`
	// Subfolders must be separated by a delimiter, often a dot ('.') or slash ('/'),
	// but it varies by mail server.
	Folder string `json:"folder"`
	// Specified in days.
	MaximumAge int64 `json:"maximum_age"`
	Order      int64 `json:"order"`
	Owner      int64 `json:"owner,nullable"`
	// - `0` - System default
	// - `1` - Text, then HTML
	// - `2` - HTML, then text
	// - `3` - HTML only
	// - `4` - Text only
	//
	// Any of 0, 1, 2, 3, 4.
	PdfLayout int64 `json:"pdf_layout"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                              respjson.Field
		Account                         respjson.Field
		Name                            respjson.Field
		UserCanChange                   respjson.Field
		Action                          respjson.Field
		ActionParameter                 respjson.Field
		AssignCorrespondent             respjson.Field
		AssignCorrespondentFrom         respjson.Field
		AssignDocumentType              respjson.Field
		AssignOwnerFromRule             respjson.Field
		AssignTags                      respjson.Field
		AssignTitleFrom                 respjson.Field
		AttachmentType                  respjson.Field
		ConsumptionScope                respjson.Field
		Enabled                         respjson.Field
		FilterAttachmentFilenameExclude respjson.Field
		FilterAttachmentFilenameInclude respjson.Field
		FilterBody                      respjson.Field
		FilterFrom                      respjson.Field
		FilterSubject                   respjson.Field
		FilterTo                        respjson.Field
		Folder                          respjson.Field
		MaximumAge                      respjson.Field
		Order                           respjson.Field
		Owner                           respjson.Field
		PdfLayout                       respjson.Field
		ExtraFields                     map[string]respjson.Field
		raw                             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MailRule) RawJSON() string { return r.JSON.raw }
func (r *MailRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// - `1` - Delete
// - `2` - Move to specified folder
// - `3` - Mark as read, don't process read mails
// - `4` - Flag the mail, don't process flagged mails
// - `5` - Tag the mail with specified tag, don't process tagged mails
type MailRuleAction int64

const (
	MailRuleAction1 MailRuleAction = 1
	MailRuleAction2 MailRuleAction = 2
	MailRuleAction3 MailRuleAction = 3
	MailRuleAction4 MailRuleAction = 4
	MailRuleAction5 MailRuleAction = 5
)

// The properties Account, Name are required.
type MailRuleRequestParam struct {
	Account             int64             `json:"account,required"`
	Name                string            `json:"name,required"`
	ActionParameter     param.Opt[string] `json:"action_parameter,omitzero"`
	AssignCorrespondent param.Opt[int64]  `json:"assign_correspondent,omitzero"`
	AssignDocumentType  param.Opt[int64]  `json:"assign_document_type,omitzero"`
	// Do not consume documents which entirely match this filename if specified.
	// Wildcards such as *.pdf or *invoice\* are allowed. Case insensitive.
	FilterAttachmentFilenameExclude param.Opt[string] `json:"filter_attachment_filename_exclude,omitzero"`
	// Only consume documents which entirely match this filename if specified.
	// Wildcards such as *.pdf or *invoice\* are allowed. Case insensitive.
	FilterAttachmentFilenameInclude param.Opt[string] `json:"filter_attachment_filename_include,omitzero"`
	FilterBody                      param.Opt[string] `json:"filter_body,omitzero"`
	FilterFrom                      param.Opt[string] `json:"filter_from,omitzero"`
	FilterSubject                   param.Opt[string] `json:"filter_subject,omitzero"`
	FilterTo                        param.Opt[string] `json:"filter_to,omitzero"`
	Owner                           param.Opt[int64]  `json:"owner,omitzero"`
	AssignOwnerFromRule             param.Opt[bool]   `json:"assign_owner_from_rule,omitzero"`
	Enabled                         param.Opt[bool]   `json:"enabled,omitzero"`
	// Subfolders must be separated by a delimiter, often a dot ('.') or slash ('/'),
	// but it varies by mail server.
	Folder param.Opt[string] `json:"folder,omitzero"`
	// Specified in days.
	MaximumAge param.Opt[int64] `json:"maximum_age,omitzero"`
	Order      param.Opt[int64] `json:"order,omitzero"`
	// - `1` - Delete
	// - `2` - Move to specified folder
	// - `3` - Mark as read, don't process read mails
	// - `4` - Flag the mail, don't process flagged mails
	// - `5` - Tag the mail with specified tag, don't process tagged mails
	//
	// Any of 1, 2, 3, 4, 5.
	Action int64 `json:"action,omitzero"`
	// - `1` - Do not assign a correspondent
	// - `2` - Use mail address
	// - `3` - Use name (or mail address if not available)
	// - `4` - Use correspondent selected below
	//
	// Any of 1, 2, 3, 4.
	AssignCorrespondentFrom int64   `json:"assign_correspondent_from,omitzero"`
	AssignTags              []int64 `json:"assign_tags,omitzero"`
	// - `1` - Use subject as title
	// - `2` - Use attachment filename as title
	// - `3` - Do not assign title from rule
	//
	// Any of 1, 2, 3.
	AssignTitleFrom int64 `json:"assign_title_from,omitzero"`
	// - `1` - Only process attachments.
	// - `2` - Process all files, including 'inline' attachments.
	//
	// Any of 1, 2.
	AttachmentType int64 `json:"attachment_type,omitzero"`
	//   - `1` - Only process attachments.
	//   - `2` - Process full Mail (with embedded attachments in file) as .eml
	//   - `3` - Process full Mail (with embedded attachments in file) as .eml + process
	//     attachments as separate documents
	//
	// Any of 1, 2, 3.
	ConsumptionScope int64 `json:"consumption_scope,omitzero"`
	// - `0` - System default
	// - `1` - Text, then HTML
	// - `2` - HTML, then text
	// - `3` - HTML only
	// - `4` - Text only
	//
	// Any of 0, 1, 2, 3, 4.
	PdfLayout      int64                              `json:"pdf_layout,omitzero"`
	SetPermissions MailRuleRequestSetPermissionsParam `json:"set_permissions,omitzero"`
	paramObj
}

func (r MailRuleRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow MailRuleRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MailRuleRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MailRuleRequestParam](
		"action", 1, 2, 3, 4, 5,
	)
	apijson.RegisterFieldValidator[MailRuleRequestParam](
		"assign_correspondent_from", 1, 2, 3, 4,
	)
	apijson.RegisterFieldValidator[MailRuleRequestParam](
		"assign_title_from", 1, 2, 3,
	)
	apijson.RegisterFieldValidator[MailRuleRequestParam](
		"attachment_type", 1, 2,
	)
	apijson.RegisterFieldValidator[MailRuleRequestParam](
		"consumption_scope", 1, 2, 3,
	)
	apijson.RegisterFieldValidator[MailRuleRequestParam](
		"pdf_layout", 0, 1, 2, 3, 4,
	)
}

type MailRuleRequestSetPermissionsParam struct {
	Change MailRuleRequestSetPermissionsChangeParam `json:"change,omitzero"`
	View   MailRuleRequestSetPermissionsViewParam   `json:"view,omitzero"`
	paramObj
}

func (r MailRuleRequestSetPermissionsParam) MarshalJSON() (data []byte, err error) {
	type shadow MailRuleRequestSetPermissionsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MailRuleRequestSetPermissionsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MailRuleRequestSetPermissionsChangeParam struct {
	Groups []int64 `json:"groups,omitzero"`
	Users  []int64 `json:"users,omitzero"`
	paramObj
}

func (r MailRuleRequestSetPermissionsChangeParam) MarshalJSON() (data []byte, err error) {
	type shadow MailRuleRequestSetPermissionsChangeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MailRuleRequestSetPermissionsChangeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MailRuleRequestSetPermissionsViewParam struct {
	Groups []int64 `json:"groups,omitzero"`
	Users  []int64 `json:"users,omitzero"`
	paramObj
}

func (r MailRuleRequestSetPermissionsViewParam) MarshalJSON() (data []byte, err error) {
	type shadow MailRuleRequestSetPermissionsViewParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MailRuleRequestSetPermissionsViewParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// - `0` - System default
// - `1` - Text, then HTML
// - `2` - HTML, then text
// - `3` - HTML only
// - `4` - Text only
type PdfLayout int64

const (
	PdfLayout0 PdfLayout = 0
	PdfLayout1 PdfLayout = 1
	PdfLayout2 PdfLayout = 2
	PdfLayout3 PdfLayout = 3
	PdfLayout4 PdfLayout = 4
)

type MailRuleListResponse struct {
	Count    int64      `json:"count,required"`
	Results  []MailRule `json:"results,required"`
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
func (r MailRuleListResponse) RawJSON() string { return r.JSON.raw }
func (r *MailRuleListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MailRuleNewParams struct {
	MailRuleRequest MailRuleRequestParam
	paramObj
}

func (r MailRuleNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.MailRuleRequest)
}
func (r *MailRuleNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.MailRuleRequest)
}

type MailRuleUpdateParams struct {
	ActionParameter     param.Opt[string] `json:"action_parameter,omitzero"`
	AssignCorrespondent param.Opt[int64]  `json:"assign_correspondent,omitzero"`
	AssignDocumentType  param.Opt[int64]  `json:"assign_document_type,omitzero"`
	// Do not consume documents which entirely match this filename if specified.
	// Wildcards such as *.pdf or *invoice\* are allowed. Case insensitive.
	FilterAttachmentFilenameExclude param.Opt[string] `json:"filter_attachment_filename_exclude,omitzero"`
	// Only consume documents which entirely match this filename if specified.
	// Wildcards such as *.pdf or *invoice\* are allowed. Case insensitive.
	FilterAttachmentFilenameInclude param.Opt[string] `json:"filter_attachment_filename_include,omitzero"`
	FilterBody                      param.Opt[string] `json:"filter_body,omitzero"`
	FilterFrom                      param.Opt[string] `json:"filter_from,omitzero"`
	FilterSubject                   param.Opt[string] `json:"filter_subject,omitzero"`
	FilterTo                        param.Opt[string] `json:"filter_to,omitzero"`
	Owner                           param.Opt[int64]  `json:"owner,omitzero"`
	Account                         param.Opt[int64]  `json:"account,omitzero"`
	AssignOwnerFromRule             param.Opt[bool]   `json:"assign_owner_from_rule,omitzero"`
	Enabled                         param.Opt[bool]   `json:"enabled,omitzero"`
	// Subfolders must be separated by a delimiter, often a dot ('.') or slash ('/'),
	// but it varies by mail server.
	Folder param.Opt[string] `json:"folder,omitzero"`
	// Specified in days.
	MaximumAge param.Opt[int64]  `json:"maximum_age,omitzero"`
	Name       param.Opt[string] `json:"name,omitzero"`
	Order      param.Opt[int64]  `json:"order,omitzero"`
	// - `1` - Delete
	// - `2` - Move to specified folder
	// - `3` - Mark as read, don't process read mails
	// - `4` - Flag the mail, don't process flagged mails
	// - `5` - Tag the mail with specified tag, don't process tagged mails
	//
	// Any of 1, 2, 3, 4, 5.
	Action int64 `json:"action,omitzero"`
	// - `1` - Do not assign a correspondent
	// - `2` - Use mail address
	// - `3` - Use name (or mail address if not available)
	// - `4` - Use correspondent selected below
	//
	// Any of 1, 2, 3, 4.
	AssignCorrespondentFrom int64   `json:"assign_correspondent_from,omitzero"`
	AssignTags              []int64 `json:"assign_tags,omitzero"`
	// - `1` - Use subject as title
	// - `2` - Use attachment filename as title
	// - `3` - Do not assign title from rule
	//
	// Any of 1, 2, 3.
	AssignTitleFrom int64 `json:"assign_title_from,omitzero"`
	// - `1` - Only process attachments.
	// - `2` - Process all files, including 'inline' attachments.
	//
	// Any of 1, 2.
	AttachmentType int64 `json:"attachment_type,omitzero"`
	//   - `1` - Only process attachments.
	//   - `2` - Process full Mail (with embedded attachments in file) as .eml
	//   - `3` - Process full Mail (with embedded attachments in file) as .eml + process
	//     attachments as separate documents
	//
	// Any of 1, 2, 3.
	ConsumptionScope int64 `json:"consumption_scope,omitzero"`
	// - `0` - System default
	// - `1` - Text, then HTML
	// - `2` - HTML, then text
	// - `3` - HTML only
	// - `4` - Text only
	//
	// Any of 0, 1, 2, 3, 4.
	PdfLayout      int64                              `json:"pdf_layout,omitzero"`
	SetPermissions MailRuleUpdateParamsSetPermissions `json:"set_permissions,omitzero"`
	paramObj
}

func (r MailRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow MailRuleUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MailRuleUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MailRuleUpdateParamsSetPermissions struct {
	Change MailRuleUpdateParamsSetPermissionsChange `json:"change,omitzero"`
	View   MailRuleUpdateParamsSetPermissionsView   `json:"view,omitzero"`
	paramObj
}

func (r MailRuleUpdateParamsSetPermissions) MarshalJSON() (data []byte, err error) {
	type shadow MailRuleUpdateParamsSetPermissions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MailRuleUpdateParamsSetPermissions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MailRuleUpdateParamsSetPermissionsChange struct {
	Groups []int64 `json:"groups,omitzero"`
	Users  []int64 `json:"users,omitzero"`
	paramObj
}

func (r MailRuleUpdateParamsSetPermissionsChange) MarshalJSON() (data []byte, err error) {
	type shadow MailRuleUpdateParamsSetPermissionsChange
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MailRuleUpdateParamsSetPermissionsChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MailRuleUpdateParamsSetPermissionsView struct {
	Groups []int64 `json:"groups,omitzero"`
	Users  []int64 `json:"users,omitzero"`
	paramObj
}

func (r MailRuleUpdateParamsSetPermissionsView) MarshalJSON() (data []byte, err error) {
	type shadow MailRuleUpdateParamsSetPermissionsView
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MailRuleUpdateParamsSetPermissionsView) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MailRuleListParams struct {
	// A page number within the paginated result set.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Number of results to return per page.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MailRuleListParams]'s query parameters as `url.Values`.
func (r MailRuleListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
