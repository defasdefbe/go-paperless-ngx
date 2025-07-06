// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package paperlessngx

import (
	"context"
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

// DocumentNoteService contains methods and other services that help with
// interacting with the paperless-ngx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDocumentNoteService] method instead.
type DocumentNoteService struct {
	Options []option.RequestOption
}

// NewDocumentNoteService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDocumentNoteService(opts ...option.RequestOption) (r DocumentNoteService) {
	r = DocumentNoteService{}
	r.Options = opts
	return
}

// View, add, or delete notes for the document
func (r *DocumentNoteService) New(ctx context.Context, id int64, params DocumentNoteNewParams, opts ...option.RequestOption) (res *PaginatedNotesList, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/documents/%v/notes/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// View, add, or delete notes for the document
func (r *DocumentNoteService) List(ctx context.Context, id int64, query DocumentNoteListParams, opts ...option.RequestOption) (res *PaginatedNotesList, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/documents/%v/notes/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// View, add, or delete notes for the document
func (r *DocumentNoteService) Delete(ctx context.Context, id int64, body DocumentNoteDeleteParams, opts ...option.RequestOption) (res *PaginatedNotesList, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/documents/%v/notes/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type PaginatedNotesList struct {
	Count    int64   `json:"count,required"`
	Results  []Notes `json:"results,required"`
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
func (r PaginatedNotesList) RawJSON() string { return r.JSON.raw }
func (r *PaginatedNotesList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentNoteNewParams struct {
	Note string `json:"note,required"`
	// Note ID to delete (used only for DELETE requests)
	ID param.Opt[int64] `query:"id,omitzero" json:"-"`
	// A page number within the paginated result set.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Number of results to return per page.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	paramObj
}

func (r DocumentNoteNewParams) MarshalJSON() (data []byte, err error) {
	type shadow DocumentNoteNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DocumentNoteNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [DocumentNoteNewParams]'s query parameters as `url.Values`.
func (r DocumentNoteNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DocumentNoteListParams struct {
	// Note ID to delete (used only for DELETE requests)
	ID param.Opt[int64] `query:"id,omitzero" json:"-"`
	// A page number within the paginated result set.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Number of results to return per page.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DocumentNoteListParams]'s query parameters as `url.Values`.
func (r DocumentNoteListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DocumentNoteDeleteParams struct {
	// Note ID to delete (used only for DELETE requests)
	ID param.Opt[int64] `query:"id,omitzero" json:"-"`
	// A page number within the paginated result set.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Number of results to return per page.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DocumentNoteDeleteParams]'s query parameters as
// `url.Values`.
func (r DocumentNoteDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
