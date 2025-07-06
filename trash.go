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
)

// TrashService contains methods and other services that help with interacting with
// the paperless-ngx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTrashService] method instead.
type TrashService struct {
	Options []option.RequestOption
}

// NewTrashService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTrashService(opts ...option.RequestOption) (r TrashService) {
	r = TrashService{}
	r.Options = opts
	return
}

func (r *TrashService) New(ctx context.Context, body TrashNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "api/trash/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

func (r *TrashService) List(ctx context.Context, query TrashListParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "api/trash/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

type TrashNewParams struct {
	// - `restore` - restore
	// - `empty` - empty
	//
	// Any of "restore", "empty".
	Action    TrashNewParamsAction `json:"action,omitzero,required"`
	Documents []int64              `json:"documents,omitzero"`
	paramObj
}

func (r TrashNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TrashNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TrashNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// - `restore` - restore
// - `empty` - empty
type TrashNewParamsAction string

const (
	TrashNewParamsActionRestore TrashNewParamsAction = "restore"
	TrashNewParamsActionEmpty   TrashNewParamsAction = "empty"
)

type TrashListParams struct {
	// A page number within the paginated result set.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Number of results to return per page.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TrashListParams]'s query parameters as `url.Values`.
func (r TrashListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
