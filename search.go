// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package paperlessngx

import (
	"context"
	"net/http"
	"net/url"

	"github.com/defasdefbe/go-paperless-ngx/internal/apijson"
	"github.com/defasdefbe/go-paperless-ngx/internal/apiquery"
	"github.com/defasdefbe/go-paperless-ngx/internal/requestconfig"
	"github.com/defasdefbe/go-paperless-ngx/option"
	"github.com/defasdefbe/go-paperless-ngx/packages/param"
	"github.com/defasdefbe/go-paperless-ngx/packages/respjson"
)

// SearchService contains methods and other services that help with interacting
// with the paperless-ngx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSearchService] method instead.
type SearchService struct {
	Options []option.RequestOption
}

// NewSearchService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSearchService(opts ...option.RequestOption) (r SearchService) {
	r = SearchService{}
	r.Options = opts
	return
}

// Global search
func (r *SearchService) Get(ctx context.Context, query SearchGetParams, opts ...option.RequestOption) (res *SearchGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/search/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get a list of all available tags
func (r *SearchService) ListTags(ctx context.Context, query SearchListTagsParams, opts ...option.RequestOption) (res *[]string, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/search/autocomplete/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type SearchGetResponse struct {
	Correspondents []Correspondent `json:"correspondents,required"`
	CustomFields   []CustomField   `json:"custom_fields,required"`
	DocumentTypes  []DocumentType  `json:"document_types,required"`
	Documents      []Document      `json:"documents,required"`
	Groups         []Group         `json:"groups,required"`
	MailAccounts   []MailAccount   `json:"mail_accounts,required"`
	MailRules      []MailRule      `json:"mail_rules,required"`
	SavedViews     []SavedView     `json:"saved_views,required"`
	StoragePaths   []StoragePath   `json:"storage_paths,required"`
	Tags           []Tag           `json:"tags,required"`
	Total          int64           `json:"total,required"`
	Users          []User          `json:"users,required"`
	Workflows      []Workflow      `json:"workflows,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Correspondents respjson.Field
		CustomFields   respjson.Field
		DocumentTypes  respjson.Field
		Documents      respjson.Field
		Groups         respjson.Field
		MailAccounts   respjson.Field
		MailRules      respjson.Field
		SavedViews     respjson.Field
		StoragePaths   respjson.Field
		Tags           respjson.Field
		Total          respjson.Field
		Users          respjson.Field
		Workflows      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SearchGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SearchGetParams struct {
	// Query to search for
	Query string `query:"query,required" json:"-"`
	// Search only the database
	DBOnly param.Opt[bool] `query:"db_only,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SearchGetParams]'s query parameters as `url.Values`.
func (r SearchGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SearchListTagsParams struct {
	// Number of completions to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Term to search for
	Term param.Opt[string] `query:"term,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SearchListTagsParams]'s query parameters as `url.Values`.
func (r SearchListTagsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
