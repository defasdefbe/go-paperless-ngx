// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package paperlessnix

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/paperless-nix-go/internal/apiform"
	"github.com/stainless-sdks/paperless-nix-go/internal/apijson"
	"github.com/stainless-sdks/paperless-nix-go/internal/apiquery"
	"github.com/stainless-sdks/paperless-nix-go/internal/requestconfig"
	"github.com/stainless-sdks/paperless-nix-go/option"
	"github.com/stainless-sdks/paperless-nix-go/packages/param"
	"github.com/stainless-sdks/paperless-nix-go/packages/respjson"
)

// DocumentService contains methods and other services that help with interacting
// with the paperless-nix API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDocumentService] method instead.
type DocumentService struct {
	Options []option.RequestOption
	Notes   DocumentNoteService
}

// NewDocumentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDocumentService(opts ...option.RequestOption) (r DocumentService) {
	r = DocumentService{}
	r.Options = opts
	r.Notes = NewDocumentNoteService(opts...)
	return
}

// Retrieve a single document
func (r *DocumentService) Get(ctx context.Context, id int64, query DocumentGetParams, opts ...option.RequestOption) (res *Document, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/documents/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Pass a user object to serializer
func (r *DocumentService) Update(ctx context.Context, id int64, body DocumentUpdateParams, opts ...option.RequestOption) (res *Document, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/documents/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Pass a user object to serializer
func (r *DocumentService) List(ctx context.Context, query DocumentListParams, opts ...option.RequestOption) (res *DocumentListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/documents/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Pass a user object to serializer
func (r *DocumentService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("api/documents/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

func (r *DocumentService) BulkDownload(ctx context.Context, body DocumentBulkDownloadParams, opts ...option.RequestOption) (res *DocumentBulkDownloadResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/documents/bulk_download/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Perform a bulk edit operation on a list of documents
func (r *DocumentService) BulkEdit(ctx context.Context, body DocumentBulkEditParams, opts ...option.RequestOption) (res *DocumentBulkEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/documents/bulk_edit/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Download the document
func (r *DocumentService) Download(ctx context.Context, id int64, query DocumentDownloadParams, opts ...option.RequestOption) (res *io.Reader, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/documents/%v/download/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Email the document to one or more recipients as an attachment.
func (r *DocumentService) Email(ctx context.Context, id int64, body DocumentEmailParams, opts ...option.RequestOption) (res *DocumentEmailResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/documents/%v/email/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// View the document history
func (r *DocumentService) History(ctx context.Context, id int64, query DocumentHistoryParams, opts ...option.RequestOption) (res *DocumentHistoryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/documents/%v/history/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// View the document metadata
func (r *DocumentService) Metadata(ctx context.Context, id int64, opts ...option.RequestOption) (res *DocumentMetadataResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/documents/%v/metadata/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get the next available Archive Serial Number (ASN) for a new document
func (r *DocumentService) NextAsn(ctx context.Context, opts ...option.RequestOption) (res *int64, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/documents/next_asn/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// View the document preview
func (r *DocumentService) Preview(ctx context.Context, id int64, opts ...option.RequestOption) (res *io.Reader, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/documents/%v/preview/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get selection data for the selected documents
func (r *DocumentService) SelectionData(ctx context.Context, body DocumentSelectionDataParams, opts ...option.RequestOption) (res *DocumentSelectionDataResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/documents/selection_data/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// View share links for the document
func (r *DocumentService) ShareLinks(ctx context.Context, id string, opts ...option.RequestOption) (res *[]DocumentShareLinksResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/documents/%s/share_links/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// View suggestions for the document
func (r *DocumentService) Suggestions(ctx context.Context, id int64, opts ...option.RequestOption) (res *DocumentSuggestionsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/documents/%v/suggestions/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// View the document thumbnail
func (r *DocumentService) Thumbnail(ctx context.Context, id int64, opts ...option.RequestOption) (res *io.Reader, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("api/documents/%v/thumb/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Upload a document via the API
func (r *DocumentService) Upload(ctx context.Context, body DocumentUploadParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/documents/post_document/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// - `none` - none
// - `deflated` - deflated
// - `bzip2` - bzip2
// - `lzma` - lzma
type CompressionEnum string

const (
	CompressionEnumNone     CompressionEnum = "none"
	CompressionEnumDeflated CompressionEnum = "deflated"
	CompressionEnumBzip2    CompressionEnum = "bzip2"
	CompressionEnumLzma     CompressionEnum = "lzma"
)

// - `archive` - archive
// - `originals` - originals
// - `both` - both
type ContentEnum string

const (
	ContentEnumArchive   ContentEnum = "archive"
	ContentEnumOriginals ContentEnum = "originals"
	ContentEnumBoth      ContentEnum = "both"
)

// The properties Field, Value are required.
type CustomFieldInstanceRequestParam struct {
	// Given the _incoming_ primitive data, return the value for this field that should
	// be validated and transformed to a native value.
	Value CustomFieldInstanceRequestValueUnionParam `json:"value,omitzero,required"`
	Field int64                                     `json:"field,required"`
	paramObj
}

func (r CustomFieldInstanceRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CustomFieldInstanceRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomFieldInstanceRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CustomFieldInstanceRequestValueUnionParam struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfAnyMap map[string]any     `json:",omitzero,inline"`
	paramUnion
}

func (u CustomFieldInstanceRequestValueUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfAnyMap)
}
func (u *CustomFieldInstanceRequestValueUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CustomFieldInstanceRequestValueUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	}
	return nil
}

// Adds update nested feature
type Document struct {
	ID                  int64               `json:"id,required"`
	Added               time.Time           `json:"added,required" format:"date-time"`
	ArchivedFileName    string              `json:"archived_file_name,required"`
	Correspondent       int64               `json:"correspondent,required"`
	DocumentType        int64               `json:"document_type,required"`
	IsSharedByRequester bool                `json:"is_shared_by_requester,required"`
	MimeType            string              `json:"mime_type,required"`
	Modified            time.Time           `json:"modified,required" format:"date-time"`
	Notes               []Notes             `json:"notes,required"`
	OriginalFileName    string              `json:"original_file_name,required"`
	PageCount           int64               `json:"page_count,required"`
	Permissions         DocumentPermissions `json:"permissions,required"`
	StoragePath         int64               `json:"storage_path,required"`
	Tags                []int64             `json:"tags,required"`
	UserCanChange       bool                `json:"user_can_change,required"`
	// The position of this document in your physical document archive.
	ArchiveSerialNumber int64 `json:"archive_serial_number,nullable"`
	// The raw, text-only data of the document. This field is primarily used for
	// searching.
	Content string    `json:"content"`
	Created time.Time `json:"created" format:"date"`
	// Deprecated: deprecated
	CreatedDate  time.Time             `json:"created_date" format:"date"`
	CustomFields []DocumentCustomField `json:"custom_fields"`
	DeletedAt    time.Time             `json:"deleted_at,nullable" format:"date-time"`
	Owner        int64                 `json:"owner,nullable"`
	Title        string                `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		Added               respjson.Field
		ArchivedFileName    respjson.Field
		Correspondent       respjson.Field
		DocumentType        respjson.Field
		IsSharedByRequester respjson.Field
		MimeType            respjson.Field
		Modified            respjson.Field
		Notes               respjson.Field
		OriginalFileName    respjson.Field
		PageCount           respjson.Field
		Permissions         respjson.Field
		StoragePath         respjson.Field
		Tags                respjson.Field
		UserCanChange       respjson.Field
		ArchiveSerialNumber respjson.Field
		Content             respjson.Field
		Created             respjson.Field
		CreatedDate         respjson.Field
		CustomFields        respjson.Field
		DeletedAt           respjson.Field
		Owner               respjson.Field
		Title               respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Document) RawJSON() string { return r.JSON.raw }
func (r *Document) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentPermissions struct {
	Change DocumentPermissionsChange `json:"change"`
	View   DocumentPermissionsView   `json:"view"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Change      respjson.Field
		View        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentPermissions) RawJSON() string { return r.JSON.raw }
func (r *DocumentPermissions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentPermissionsChange struct {
	Groups []int64 `json:"groups"`
	Users  []int64 `json:"users"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Groups      respjson.Field
		Users       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentPermissionsChange) RawJSON() string { return r.JSON.raw }
func (r *DocumentPermissionsChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentPermissionsView struct {
	Groups []int64 `json:"groups"`
	Users  []int64 `json:"users"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Groups      respjson.Field
		Users       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentPermissionsView) RawJSON() string { return r.JSON.raw }
func (r *DocumentPermissionsView) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentCustomField struct {
	Field int64 `json:"field,required"`
	// Given the _incoming_ primitive data, return the value for this field that should
	// be validated and transformed to a native value.
	Value DocumentCustomFieldValueUnion `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentCustomField) RawJSON() string { return r.JSON.raw }
func (r *DocumentCustomField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DocumentCustomFieldValueUnion contains all possible properties and values from
// [string], [float64], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfDocumentCustomFieldValueMapItem]
type DocumentCustomFieldValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfDocumentCustomFieldValueMapItem any `json:",inline"`
	JSON                              struct {
		OfString                          respjson.Field
		OfFloat                           respjson.Field
		OfDocumentCustomFieldValueMapItem respjson.Field
		raw                               string
	} `json:"-"`
}

func (u DocumentCustomFieldValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DocumentCustomFieldValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DocumentCustomFieldValueUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DocumentCustomFieldValueUnion) RawJSON() string { return u.JSON.raw }

func (r *DocumentCustomFieldValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Notes struct {
	ID      int64     `json:"id,required"`
	User    NotesUser `json:"user,required"`
	Created time.Time `json:"created" format:"date-time"`
	// Note for the document
	Note string `json:"note"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		User        respjson.Field
		Created     respjson.Field
		Note        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Notes) RawJSON() string { return r.JSON.raw }
func (r *Notes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotesUser struct {
	ID int64 `json:"id,required"`
	// Required. 150 characters or fewer. Letters, digits and @/./+/-/\_ only.
	Username  string `json:"username,required"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Username    respjson.Field
		FirstName   respjson.Field
		LastName    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NotesUser) RawJSON() string { return r.JSON.raw }
func (r *NotesUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentListResponse struct {
	Count    int64      `json:"count,required"`
	Results  []Document `json:"results,required"`
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
func (r DocumentListResponse) RawJSON() string { return r.JSON.raw }
func (r *DocumentListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentBulkDownloadResponse struct {
	// - `none` - none
	// - `deflated` - deflated
	// - `bzip2` - bzip2
	// - `lzma` - lzma
	//
	// Any of "none", "deflated", "bzip2", "lzma".
	Compression CompressionEnum `json:"compression"`
	// - `archive` - archive
	// - `originals` - originals
	// - `both` - both
	//
	// Any of "archive", "originals", "both".
	Content          ContentEnum `json:"content"`
	FollowFormatting bool        `json:"follow_formatting"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Compression      respjson.Field
		Content          respjson.Field
		FollowFormatting respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentBulkDownloadResponse) RawJSON() string { return r.JSON.raw }
func (r *DocumentBulkDownloadResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentBulkEditResponse struct {
	Result string `json:"result,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Result      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentBulkEditResponse) RawJSON() string { return r.JSON.raw }
func (r *DocumentBulkEditResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentEmailResponse struct {
	Message string `json:"message,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentEmailResponse) RawJSON() string { return r.JSON.raw }
func (r *DocumentEmailResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentHistoryResponse struct {
	Count    int64                           `json:"count,required"`
	Results  []DocumentHistoryResponseResult `json:"results,required"`
	All      []any                           `json:"all"`
	Next     string                          `json:"next,nullable" format:"uri"`
	Previous string                          `json:"previous,nullable" format:"uri"`
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
func (r DocumentHistoryResponse) RawJSON() string { return r.JSON.raw }
func (r *DocumentHistoryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentHistoryResponseResult struct {
	ID        int64                              `json:"id,required"`
	Action    string                             `json:"action,required"`
	Actor     DocumentHistoryResponseResultActor `json:"actor,required"`
	Changes   map[string]any                     `json:"changes,required"`
	Timestamp time.Time                          `json:"timestamp,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Action      respjson.Field
		Actor       respjson.Field
		Changes     respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentHistoryResponseResult) RawJSON() string { return r.JSON.raw }
func (r *DocumentHistoryResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentHistoryResponseResultActor struct {
	ID       int64  `json:"id,required"`
	Username string `json:"username,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Username    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentHistoryResponseResultActor) RawJSON() string { return r.JSON.raw }
func (r *DocumentHistoryResponseResultActor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentMetadataResponse struct {
	ArchiveChecksum      string         `json:"archive_checksum,required"`
	ArchiveMediaFilename string         `json:"archive_media_filename,required"`
	ArchiveMetadata      map[string]any `json:"archive_metadata,required"`
	ArchiveSize          int64          `json:"archive_size,required"`
	HasArchiveVersion    bool           `json:"has_archive_version,required"`
	Lang                 string         `json:"lang,required"`
	MediaFilename        string         `json:"media_filename,required"`
	OriginalChecksum     string         `json:"original_checksum,required"`
	OriginalFilename     string         `json:"original_filename,required"`
	OriginalMetadata     map[string]any `json:"original_metadata,required"`
	OriginalMimeType     string         `json:"original_mime_type,required"`
	OriginalSize         int64          `json:"original_size,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ArchiveChecksum      respjson.Field
		ArchiveMediaFilename respjson.Field
		ArchiveMetadata      respjson.Field
		ArchiveSize          respjson.Field
		HasArchiveVersion    respjson.Field
		Lang                 respjson.Field
		MediaFilename        respjson.Field
		OriginalChecksum     respjson.Field
		OriginalFilename     respjson.Field
		OriginalMetadata     respjson.Field
		OriginalMimeType     respjson.Field
		OriginalSize         respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentMetadataResponse) RawJSON() string { return r.JSON.raw }
func (r *DocumentMetadataResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentSelectionDataResponse struct {
	SelectedCorrespondents []DocumentSelectionDataResponseSelectedCorrespondent `json:"selected_correspondents,required"`
	SelectedCustomFields   []DocumentSelectionDataResponseSelectedCustomField   `json:"selected_custom_fields,required"`
	SelectedDocumentTypes  []DocumentSelectionDataResponseSelectedDocumentType  `json:"selected_document_types,required"`
	SelectedStoragePaths   []DocumentSelectionDataResponseSelectedStoragePath   `json:"selected_storage_paths,required"`
	SelectedTags           []DocumentSelectionDataResponseSelectedTag           `json:"selected_tags,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SelectedCorrespondents respjson.Field
		SelectedCustomFields   respjson.Field
		SelectedDocumentTypes  respjson.Field
		SelectedStoragePaths   respjson.Field
		SelectedTags           respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentSelectionDataResponse) RawJSON() string { return r.JSON.raw }
func (r *DocumentSelectionDataResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentSelectionDataResponseSelectedCorrespondent struct {
	ID            int64 `json:"id,required"`
	DocumentCount int64 `json:"document_count,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		DocumentCount respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentSelectionDataResponseSelectedCorrespondent) RawJSON() string { return r.JSON.raw }
func (r *DocumentSelectionDataResponseSelectedCorrespondent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentSelectionDataResponseSelectedCustomField struct {
	ID            int64 `json:"id,required"`
	DocumentCount int64 `json:"document_count,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		DocumentCount respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentSelectionDataResponseSelectedCustomField) RawJSON() string { return r.JSON.raw }
func (r *DocumentSelectionDataResponseSelectedCustomField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentSelectionDataResponseSelectedDocumentType struct {
	ID            int64 `json:"id,required"`
	DocumentCount int64 `json:"document_count,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		DocumentCount respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentSelectionDataResponseSelectedDocumentType) RawJSON() string { return r.JSON.raw }
func (r *DocumentSelectionDataResponseSelectedDocumentType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentSelectionDataResponseSelectedStoragePath struct {
	ID            int64 `json:"id,required"`
	DocumentCount int64 `json:"document_count,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		DocumentCount respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentSelectionDataResponseSelectedStoragePath) RawJSON() string { return r.JSON.raw }
func (r *DocumentSelectionDataResponseSelectedStoragePath) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentSelectionDataResponseSelectedTag struct {
	ID            int64 `json:"id,required"`
	DocumentCount int64 `json:"document_count,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		DocumentCount respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentSelectionDataResponseSelectedTag) RawJSON() string { return r.JSON.raw }
func (r *DocumentSelectionDataResponseSelectedTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentShareLinksResponse struct {
	ID         int64     `json:"id"`
	Created    time.Time `json:"created" format:"date-time"`
	Expiration time.Time `json:"expiration" format:"date-time"`
	Slug       string    `json:"slug"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Created     respjson.Field
		Expiration  respjson.Field
		Slug        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentShareLinksResponse) RawJSON() string { return r.JSON.raw }
func (r *DocumentShareLinksResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentSuggestionsResponse struct {
	Correspondents []int64  `json:"correspondents,required"`
	Dates          []string `json:"dates,required"`
	DocumentTypes  []int64  `json:"document_types,required"`
	StoragePaths   []int64  `json:"storage_paths,required"`
	Tags           []int64  `json:"tags,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Correspondents respjson.Field
		Dates          respjson.Field
		DocumentTypes  respjson.Field
		StoragePaths   respjson.Field
		Tags           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentSuggestionsResponse) RawJSON() string { return r.JSON.raw }
func (r *DocumentSuggestionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentGetParams struct {
	FullPerms param.Opt[bool] `query:"full_perms,omitzero" json:"-"`
	Fields    []string        `query:"fields,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DocumentGetParams]'s query parameters as `url.Values`.
func (r DocumentGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DocumentUpdateParams struct {
	// The position of this document in your physical document archive.
	ArchiveSerialNumber param.Opt[int64]     `json:"archive_serial_number,omitzero"`
	Correspondent       param.Opt[int64]     `json:"correspondent,omitzero"`
	DeletedAt           param.Opt[time.Time] `json:"deleted_at,omitzero" format:"date-time"`
	DocumentType        param.Opt[int64]     `json:"document_type,omitzero"`
	Owner               param.Opt[int64]     `json:"owner,omitzero"`
	RemoveInboxTags     param.Opt[bool]      `json:"remove_inbox_tags,omitzero"`
	StoragePath         param.Opt[int64]     `json:"storage_path,omitzero"`
	// The raw, text-only data of the document. This field is primarily used for
	// searching.
	Content        param.Opt[string]                  `json:"content,omitzero"`
	Created        param.Opt[time.Time]               `json:"created,omitzero" format:"date"`
	CreatedDate    param.Opt[time.Time]               `json:"created_date,omitzero" format:"date"`
	Title          param.Opt[string]                  `json:"title,omitzero"`
	CustomFields   []CustomFieldInstanceRequestParam  `json:"custom_fields,omitzero"`
	SetPermissions DocumentUpdateParamsSetPermissions `json:"set_permissions,omitzero"`
	Tags           []int64                            `json:"tags,omitzero"`
	paramObj
}

func (r DocumentUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow DocumentUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DocumentUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentUpdateParamsSetPermissions struct {
	Change DocumentUpdateParamsSetPermissionsChange `json:"change,omitzero"`
	View   DocumentUpdateParamsSetPermissionsView   `json:"view,omitzero"`
	paramObj
}

func (r DocumentUpdateParamsSetPermissions) MarshalJSON() (data []byte, err error) {
	type shadow DocumentUpdateParamsSetPermissions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DocumentUpdateParamsSetPermissions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentUpdateParamsSetPermissionsChange struct {
	Groups []int64 `json:"groups,omitzero"`
	Users  []int64 `json:"users,omitzero"`
	paramObj
}

func (r DocumentUpdateParamsSetPermissionsChange) MarshalJSON() (data []byte, err error) {
	type shadow DocumentUpdateParamsSetPermissionsChange
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DocumentUpdateParamsSetPermissionsChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentUpdateParamsSetPermissionsView struct {
	Groups []int64 `json:"groups,omitzero"`
	Users  []int64 `json:"users,omitzero"`
	paramObj
}

func (r DocumentUpdateParamsSetPermissionsView) MarshalJSON() (data []byte, err error) {
	type shadow DocumentUpdateParamsSetPermissionsView
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DocumentUpdateParamsSetPermissionsView) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentListParams struct {
	ID                           param.Opt[int64]     `query:"id,omitzero" json:"-"`
	AddedDateGt                  param.Opt[time.Time] `query:"added__date__gt,omitzero" format:"date" json:"-"`
	AddedDateGte                 param.Opt[time.Time] `query:"added__date__gte,omitzero" format:"date" json:"-"`
	AddedDateLt                  param.Opt[time.Time] `query:"added__date__lt,omitzero" format:"date" json:"-"`
	AddedDateLte                 param.Opt[time.Time] `query:"added__date__lte,omitzero" format:"date" json:"-"`
	AddedDay                     param.Opt[float64]   `query:"added__day,omitzero" json:"-"`
	AddedGt                      param.Opt[time.Time] `query:"added__gt,omitzero" format:"date-time" json:"-"`
	AddedGte                     param.Opt[time.Time] `query:"added__gte,omitzero" format:"date-time" json:"-"`
	AddedLt                      param.Opt[time.Time] `query:"added__lt,omitzero" format:"date-time" json:"-"`
	AddedLte                     param.Opt[time.Time] `query:"added__lte,omitzero" format:"date-time" json:"-"`
	AddedMonth                   param.Opt[float64]   `query:"added__month,omitzero" json:"-"`
	AddedYear                    param.Opt[float64]   `query:"added__year,omitzero" json:"-"`
	ArchiveSerialNumber          param.Opt[int64]     `query:"archive_serial_number,omitzero" json:"-"`
	ArchiveSerialNumberGt        param.Opt[int64]     `query:"archive_serial_number__gt,omitzero" json:"-"`
	ArchiveSerialNumberGte       param.Opt[int64]     `query:"archive_serial_number__gte,omitzero" json:"-"`
	ArchiveSerialNumberIsnull    param.Opt[bool]      `query:"archive_serial_number__isnull,omitzero" json:"-"`
	ArchiveSerialNumberLt        param.Opt[int64]     `query:"archive_serial_number__lt,omitzero" json:"-"`
	ArchiveSerialNumberLte       param.Opt[int64]     `query:"archive_serial_number__lte,omitzero" json:"-"`
	ChecksumIcontains            param.Opt[string]    `query:"checksum__icontains,omitzero" json:"-"`
	ChecksumIendswith            param.Opt[string]    `query:"checksum__iendswith,omitzero" json:"-"`
	ChecksumIexact               param.Opt[string]    `query:"checksum__iexact,omitzero" json:"-"`
	ChecksumIstartswith          param.Opt[string]    `query:"checksum__istartswith,omitzero" json:"-"`
	ContentIcontains             param.Opt[string]    `query:"content__icontains,omitzero" json:"-"`
	ContentIendswith             param.Opt[string]    `query:"content__iendswith,omitzero" json:"-"`
	ContentIexact                param.Opt[string]    `query:"content__iexact,omitzero" json:"-"`
	ContentIstartswith           param.Opt[string]    `query:"content__istartswith,omitzero" json:"-"`
	CorrespondentID              param.Opt[int64]     `query:"correspondent__id,omitzero" json:"-"`
	CorrespondentIDNone          param.Opt[int64]     `query:"correspondent__id__none,omitzero" json:"-"`
	CorrespondentIsnull          param.Opt[bool]      `query:"correspondent__isnull,omitzero" json:"-"`
	CorrespondentNameIcontains   param.Opt[string]    `query:"correspondent__name__icontains,omitzero" json:"-"`
	CorrespondentNameIendswith   param.Opt[string]    `query:"correspondent__name__iendswith,omitzero" json:"-"`
	CorrespondentNameIexact      param.Opt[string]    `query:"correspondent__name__iexact,omitzero" json:"-"`
	CorrespondentNameIstartswith param.Opt[string]    `query:"correspondent__name__istartswith,omitzero" json:"-"`
	CreatedDateGt                param.Opt[time.Time] `query:"created__date__gt,omitzero" format:"date" json:"-"`
	CreatedDateGte               param.Opt[time.Time] `query:"created__date__gte,omitzero" format:"date" json:"-"`
	CreatedDateLt                param.Opt[time.Time] `query:"created__date__lt,omitzero" format:"date" json:"-"`
	CreatedDateLte               param.Opt[time.Time] `query:"created__date__lte,omitzero" format:"date" json:"-"`
	CreatedDay                   param.Opt[float64]   `query:"created__day,omitzero" json:"-"`
	CreatedGt                    param.Opt[time.Time] `query:"created__gt,omitzero" format:"date" json:"-"`
	CreatedGte                   param.Opt[time.Time] `query:"created__gte,omitzero" format:"date" json:"-"`
	CreatedLt                    param.Opt[time.Time] `query:"created__lt,omitzero" format:"date" json:"-"`
	CreatedLte                   param.Opt[time.Time] `query:"created__lte,omitzero" format:"date" json:"-"`
	CreatedMonth                 param.Opt[float64]   `query:"created__month,omitzero" json:"-"`
	CreatedYear                  param.Opt[float64]   `query:"created__year,omitzero" json:"-"`
	CustomFieldQuery             param.Opt[string]    `query:"custom_field_query,omitzero" json:"-"`
	CustomFieldsIcontains        param.Opt[string]    `query:"custom_fields__icontains,omitzero" json:"-"`
	CustomFieldsIDAll            param.Opt[int64]     `query:"custom_fields__id__all,omitzero" json:"-"`
	CustomFieldsIDIn             param.Opt[int64]     `query:"custom_fields__id__in,omitzero" json:"-"`
	CustomFieldsIDNone           param.Opt[int64]     `query:"custom_fields__id__none,omitzero" json:"-"`
	DocumentTypeID               param.Opt[int64]     `query:"document_type__id,omitzero" json:"-"`
	DocumentTypeIDNone           param.Opt[int64]     `query:"document_type__id__none,omitzero" json:"-"`
	DocumentTypeIsnull           param.Opt[bool]      `query:"document_type__isnull,omitzero" json:"-"`
	DocumentTypeNameIcontains    param.Opt[string]    `query:"document_type__name__icontains,omitzero" json:"-"`
	DocumentTypeNameIendswith    param.Opt[string]    `query:"document_type__name__iendswith,omitzero" json:"-"`
	DocumentTypeNameIexact       param.Opt[string]    `query:"document_type__name__iexact,omitzero" json:"-"`
	DocumentTypeNameIstartswith  param.Opt[string]    `query:"document_type__name__istartswith,omitzero" json:"-"`
	FullPerms                    param.Opt[bool]      `query:"full_perms,omitzero" json:"-"`
	// Has custom field
	HasCustomFields param.Opt[bool] `query:"has_custom_fields,omitzero" json:"-"`
	IsInInbox       param.Opt[bool] `query:"is_in_inbox,omitzero" json:"-"`
	// Is tagged
	IsTagged        param.Opt[bool]      `query:"is_tagged,omitzero" json:"-"`
	MimeType        param.Opt[string]    `query:"mime_type,omitzero" json:"-"`
	ModifiedDateGt  param.Opt[time.Time] `query:"modified__date__gt,omitzero" format:"date" json:"-"`
	ModifiedDateGte param.Opt[time.Time] `query:"modified__date__gte,omitzero" format:"date" json:"-"`
	ModifiedDateLt  param.Opt[time.Time] `query:"modified__date__lt,omitzero" format:"date" json:"-"`
	ModifiedDateLte param.Opt[time.Time] `query:"modified__date__lte,omitzero" format:"date" json:"-"`
	ModifiedDay     param.Opt[float64]   `query:"modified__day,omitzero" json:"-"`
	ModifiedGt      param.Opt[time.Time] `query:"modified__gt,omitzero" format:"date-time" json:"-"`
	ModifiedGte     param.Opt[time.Time] `query:"modified__gte,omitzero" format:"date-time" json:"-"`
	ModifiedLt      param.Opt[time.Time] `query:"modified__lt,omitzero" format:"date-time" json:"-"`
	ModifiedLte     param.Opt[time.Time] `query:"modified__lte,omitzero" format:"date-time" json:"-"`
	ModifiedMonth   param.Opt[float64]   `query:"modified__month,omitzero" json:"-"`
	ModifiedYear    param.Opt[float64]   `query:"modified__year,omitzero" json:"-"`
	// Which field to use when ordering the results.
	Ordering                    param.Opt[string] `query:"ordering,omitzero" json:"-"`
	OriginalFilenameIcontains   param.Opt[string] `query:"original_filename__icontains,omitzero" json:"-"`
	OriginalFilenameIendswith   param.Opt[string] `query:"original_filename__iendswith,omitzero" json:"-"`
	OriginalFilenameIexact      param.Opt[string] `query:"original_filename__iexact,omitzero" json:"-"`
	OriginalFilenameIstartswith param.Opt[string] `query:"original_filename__istartswith,omitzero" json:"-"`
	OwnerID                     param.Opt[int64]  `query:"owner__id,omitzero" json:"-"`
	OwnerIDNone                 param.Opt[int64]  `query:"owner__id__none,omitzero" json:"-"`
	OwnerIsnull                 param.Opt[bool]   `query:"owner__isnull,omitzero" json:"-"`
	// A page number within the paginated result set.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Number of results to return per page.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// A search term.
	Search                     param.Opt[string] `query:"search,omitzero" json:"-"`
	SharedByID                 param.Opt[bool]   `query:"shared_by__id,omitzero" json:"-"`
	StoragePathID              param.Opt[int64]  `query:"storage_path__id,omitzero" json:"-"`
	StoragePathIDNone          param.Opt[int64]  `query:"storage_path__id__none,omitzero" json:"-"`
	StoragePathIsnull          param.Opt[bool]   `query:"storage_path__isnull,omitzero" json:"-"`
	StoragePathNameIcontains   param.Opt[string] `query:"storage_path__name__icontains,omitzero" json:"-"`
	StoragePathNameIendswith   param.Opt[string] `query:"storage_path__name__iendswith,omitzero" json:"-"`
	StoragePathNameIexact      param.Opt[string] `query:"storage_path__name__iexact,omitzero" json:"-"`
	StoragePathNameIstartswith param.Opt[string] `query:"storage_path__name__istartswith,omitzero" json:"-"`
	TagsID                     param.Opt[int64]  `query:"tags__id,omitzero" json:"-"`
	TagsIDAll                  param.Opt[int64]  `query:"tags__id__all,omitzero" json:"-"`
	TagsIDIn                   param.Opt[int64]  `query:"tags__id__in,omitzero" json:"-"`
	TagsIDNone                 param.Opt[int64]  `query:"tags__id__none,omitzero" json:"-"`
	TagsNameIcontains          param.Opt[string] `query:"tags__name__icontains,omitzero" json:"-"`
	TagsNameIendswith          param.Opt[string] `query:"tags__name__iendswith,omitzero" json:"-"`
	TagsNameIexact             param.Opt[string] `query:"tags__name__iexact,omitzero" json:"-"`
	TagsNameIstartswith        param.Opt[string] `query:"tags__name__istartswith,omitzero" json:"-"`
	TitleIcontains             param.Opt[string] `query:"title__icontains,omitzero" json:"-"`
	TitleIendswith             param.Opt[string] `query:"title__iendswith,omitzero" json:"-"`
	TitleIexact                param.Opt[string] `query:"title__iexact,omitzero" json:"-"`
	TitleIstartswith           param.Opt[string] `query:"title__istartswith,omitzero" json:"-"`
	TitleContent               param.Opt[string] `query:"title_content,omitzero" json:"-"`
	// Multiple values may be separated by commas.
	CorrespondentIDIn []int64 `query:"correspondent__id__in,omitzero" json:"-"`
	// Multiple values may be separated by commas.
	DocumentTypeIDIn []int64  `query:"document_type__id__in,omitzero" json:"-"`
	Fields           []string `query:"fields,omitzero" json:"-"`
	// Multiple values may be separated by commas.
	IDIn []int64 `query:"id__in,omitzero" json:"-"`
	// Multiple values may be separated by commas.
	OwnerIDIn []int64 `query:"owner__id__in,omitzero" json:"-"`
	// Multiple values may be separated by commas.
	StoragePathIDIn []int64 `query:"storage_path__id__in,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DocumentListParams]'s query parameters as `url.Values`.
func (r DocumentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DocumentBulkDownloadParams struct {
	Documents        []int64         `json:"documents,omitzero,required"`
	FollowFormatting param.Opt[bool] `json:"follow_formatting,omitzero"`
	// - `none` - none
	// - `deflated` - deflated
	// - `bzip2` - bzip2
	// - `lzma` - lzma
	//
	// Any of "none", "deflated", "bzip2", "lzma".
	Compression CompressionEnum `json:"compression,omitzero"`
	// - `archive` - archive
	// - `originals` - originals
	// - `both` - both
	//
	// Any of "archive", "originals", "both".
	Content ContentEnum `json:"content,omitzero"`
	paramObj
}

func (r DocumentBulkDownloadParams) MarshalJSON() (data []byte, err error) {
	type shadow DocumentBulkDownloadParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DocumentBulkDownloadParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentBulkEditParams struct {
	Documents []int64 `json:"documents,omitzero,required"`
	// - `set_correspondent` - set_correspondent
	// - `set_document_type` - set_document_type
	// - `set_storage_path` - set_storage_path
	// - `add_tag` - add_tag
	// - `remove_tag` - remove_tag
	// - `modify_tags` - modify_tags
	// - `modify_custom_fields` - modify_custom_fields
	// - `delete` - delete
	// - `reprocess` - reprocess
	// - `set_permissions` - set_permissions
	// - `rotate` - rotate
	// - `merge` - merge
	// - `split` - split
	// - `delete_pages` - delete_pages
	//
	// Any of "set_correspondent", "set_document_type", "set_storage_path", "add_tag",
	// "remove_tag", "modify_tags", "modify_custom_fields", "delete", "reprocess",
	// "set_permissions", "rotate", "merge", "split", "delete_pages".
	Method     DocumentBulkEditParamsMethod `json:"method,omitzero,required"`
	Parameters map[string]any               `json:"parameters,omitzero"`
	paramObj
}

func (r DocumentBulkEditParams) MarshalJSON() (data []byte, err error) {
	type shadow DocumentBulkEditParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DocumentBulkEditParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// - `set_correspondent` - set_correspondent
// - `set_document_type` - set_document_type
// - `set_storage_path` - set_storage_path
// - `add_tag` - add_tag
// - `remove_tag` - remove_tag
// - `modify_tags` - modify_tags
// - `modify_custom_fields` - modify_custom_fields
// - `delete` - delete
// - `reprocess` - reprocess
// - `set_permissions` - set_permissions
// - `rotate` - rotate
// - `merge` - merge
// - `split` - split
// - `delete_pages` - delete_pages
type DocumentBulkEditParamsMethod string

const (
	DocumentBulkEditParamsMethodSetCorrespondent   DocumentBulkEditParamsMethod = "set_correspondent"
	DocumentBulkEditParamsMethodSetDocumentType    DocumentBulkEditParamsMethod = "set_document_type"
	DocumentBulkEditParamsMethodSetStoragePath     DocumentBulkEditParamsMethod = "set_storage_path"
	DocumentBulkEditParamsMethodAddTag             DocumentBulkEditParamsMethod = "add_tag"
	DocumentBulkEditParamsMethodRemoveTag          DocumentBulkEditParamsMethod = "remove_tag"
	DocumentBulkEditParamsMethodModifyTags         DocumentBulkEditParamsMethod = "modify_tags"
	DocumentBulkEditParamsMethodModifyCustomFields DocumentBulkEditParamsMethod = "modify_custom_fields"
	DocumentBulkEditParamsMethodDelete             DocumentBulkEditParamsMethod = "delete"
	DocumentBulkEditParamsMethodReprocess          DocumentBulkEditParamsMethod = "reprocess"
	DocumentBulkEditParamsMethodSetPermissions     DocumentBulkEditParamsMethod = "set_permissions"
	DocumentBulkEditParamsMethodRotate             DocumentBulkEditParamsMethod = "rotate"
	DocumentBulkEditParamsMethodMerge              DocumentBulkEditParamsMethod = "merge"
	DocumentBulkEditParamsMethodSplit              DocumentBulkEditParamsMethod = "split"
	DocumentBulkEditParamsMethodDeletePages        DocumentBulkEditParamsMethod = "delete_pages"
)

type DocumentDownloadParams struct {
	Original param.Opt[bool] `query:"original,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DocumentDownloadParams]'s query parameters as `url.Values`.
func (r DocumentDownloadParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DocumentEmailParams struct {
	Addresses         string          `json:"addresses,required"`
	Message           string          `json:"message,required"`
	Subject           string          `json:"subject,required"`
	UseArchiveVersion param.Opt[bool] `json:"use_archive_version,omitzero"`
	paramObj
}

func (r DocumentEmailParams) MarshalJSON() (data []byte, err error) {
	type shadow DocumentEmailParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DocumentEmailParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentHistoryParams struct {
	// A page number within the paginated result set.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Number of results to return per page.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DocumentHistoryParams]'s query parameters as `url.Values`.
func (r DocumentHistoryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DocumentSelectionDataParams struct {
	Documents []int64 `json:"documents,omitzero,required"`
	paramObj
}

func (r DocumentSelectionDataParams) MarshalJSON() (data []byte, err error) {
	type shadow DocumentSelectionDataParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DocumentSelectionDataParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentUploadParams struct {
	Document            io.Reader            `json:"document,omitzero,required" format:"binary"`
	Correspondent       param.Opt[int64]     `json:"correspondent,omitzero"`
	Created             param.Opt[time.Time] `json:"created,omitzero" format:"date-time"`
	DocumentType        param.Opt[int64]     `json:"document_type,omitzero"`
	StoragePath         param.Opt[int64]     `json:"storage_path,omitzero"`
	ArchiveSerialNumber param.Opt[int64]     `json:"archive_serial_number,omitzero"`
	FromWebui           param.Opt[bool]      `json:"from_webui,omitzero"`
	Title               param.Opt[string]    `json:"title,omitzero"`
	CustomFields        []int64              `json:"custom_fields,omitzero"`
	Tags                []int64              `json:"tags,omitzero"`
	paramObj
}

func (r DocumentUploadParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}
