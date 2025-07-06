// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package paperlessngx

import (
	"context"
	"net/http"

	"github.com/defasdefbe/go-paperless-ngx/internal/apijson"
	"github.com/defasdefbe/go-paperless-ngx/internal/requestconfig"
	"github.com/defasdefbe/go-paperless-ngx/option"
	"github.com/defasdefbe/go-paperless-ngx/packages/param"
	"github.com/defasdefbe/go-paperless-ngx/packages/respjson"
)

// BulkEditObjectService contains methods and other services that help with
// interacting with the paperless-ngx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBulkEditObjectService] method instead.
type BulkEditObjectService struct {
	Options []option.RequestOption
}

// NewBulkEditObjectService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBulkEditObjectService(opts ...option.RequestOption) (r BulkEditObjectService) {
	r = BulkEditObjectService{}
	r.Options = opts
	return
}

// Perform a bulk edit operation on a list of objects
func (r *BulkEditObjectService) New(ctx context.Context, body BulkEditObjectNewParams, opts ...option.RequestOption) (res *BulkEditObjectNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/bulk_edit_objects/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type BulkEditObjectNewResponse struct {
	Result string `json:"result,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Result      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BulkEditObjectNewResponse) RawJSON() string { return r.JSON.raw }
func (r *BulkEditObjectNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BulkEditObjectNewParams struct {
	// - `tags` - tags
	// - `correspondents` - correspondents
	// - `document_types` - document_types
	// - `storage_paths` - storage_paths
	//
	// Any of "tags", "correspondents", "document_types", "storage_paths".
	ObjectType BulkEditObjectNewParamsObjectType `json:"object_type,omitzero,required"`
	Objects    []int64                           `json:"objects,omitzero,required"`
	// - `set_permissions` - set_permissions
	// - `delete` - delete
	//
	// Any of "set_permissions", "delete".
	Operation   BulkEditObjectNewParamsOperation `json:"operation,omitzero,required"`
	Owner       param.Opt[int64]                 `json:"owner,omitzero"`
	Merge       param.Opt[bool]                  `json:"merge,omitzero"`
	Permissions map[string]any                   `json:"permissions,omitzero"`
	paramObj
}

func (r BulkEditObjectNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BulkEditObjectNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BulkEditObjectNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// - `tags` - tags
// - `correspondents` - correspondents
// - `document_types` - document_types
// - `storage_paths` - storage_paths
type BulkEditObjectNewParamsObjectType string

const (
	BulkEditObjectNewParamsObjectTypeTags           BulkEditObjectNewParamsObjectType = "tags"
	BulkEditObjectNewParamsObjectTypeCorrespondents BulkEditObjectNewParamsObjectType = "correspondents"
	BulkEditObjectNewParamsObjectTypeDocumentTypes  BulkEditObjectNewParamsObjectType = "document_types"
	BulkEditObjectNewParamsObjectTypeStoragePaths   BulkEditObjectNewParamsObjectType = "storage_paths"
)

// - `set_permissions` - set_permissions
// - `delete` - delete
type BulkEditObjectNewParamsOperation string

const (
	BulkEditObjectNewParamsOperationSetPermissions BulkEditObjectNewParamsOperation = "set_permissions"
	BulkEditObjectNewParamsOperationDelete         BulkEditObjectNewParamsOperation = "delete"
)
