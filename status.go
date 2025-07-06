// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package paperlessngx

import (
	"context"
	"net/http"
	"time"

	"github.com/stainless-sdks/paperless-nix-go/internal/apijson"
	"github.com/stainless-sdks/paperless-nix-go/internal/requestconfig"
	"github.com/stainless-sdks/paperless-nix-go/option"
	"github.com/stainless-sdks/paperless-nix-go/packages/respjson"
)

// StatusService contains methods and other services that help with interacting
// with the paperless-ngx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStatusService] method instead.
type StatusService struct {
	Options []option.RequestOption
}

// NewStatusService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewStatusService(opts ...option.RequestOption) (r StatusService) {
	r = StatusService{}
	r.Options = opts
	return
}

// Get the current system status of the Paperless-NGX server
func (r *StatusService) Get(ctx context.Context, opts ...option.RequestOption) (res *StatusGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/status/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type StatusGetResponse struct {
	Classifier  StatusGetResponseClassifier  `json:"classifier,required"`
	Database    StatusGetResponseDatabase    `json:"database,required"`
	Index       StatusGetResponseIndex       `json:"index,required"`
	InstallType string                       `json:"install_type,required"`
	PngxVersion string                       `json:"pngx_version,required"`
	SanityCheck StatusGetResponseSanityCheck `json:"sanity_check,required"`
	ServerOs    string                       `json:"server_os,required"`
	Storage     StatusGetResponseStorage     `json:"storage,required"`
	Tasks       StatusGetResponseTasks       `json:"tasks,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Classifier  respjson.Field
		Database    respjson.Field
		Index       respjson.Field
		InstallType respjson.Field
		PngxVersion respjson.Field
		SanityCheck respjson.Field
		ServerOs    respjson.Field
		Storage     respjson.Field
		Tasks       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StatusGetResponse) RawJSON() string { return r.JSON.raw }
func (r *StatusGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StatusGetResponseClassifier struct {
	Error       string    `json:"error,required"`
	LastTrained time.Time `json:"last_trained,required" format:"date-time"`
	Status      string    `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		LastTrained respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StatusGetResponseClassifier) RawJSON() string { return r.JSON.raw }
func (r *StatusGetResponseClassifier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StatusGetResponseDatabase struct {
	Error           string                                   `json:"error,required"`
	MigrationStatus StatusGetResponseDatabaseMigrationStatus `json:"migration_status,required"`
	Status          string                                   `json:"status,required"`
	Type            string                                   `json:"type,required"`
	URL             string                                   `json:"url,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error           respjson.Field
		MigrationStatus respjson.Field
		Status          respjson.Field
		Type            respjson.Field
		URL             respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StatusGetResponseDatabase) RawJSON() string { return r.JSON.raw }
func (r *StatusGetResponseDatabase) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StatusGetResponseDatabaseMigrationStatus struct {
	LatestMigration     string   `json:"latest_migration,required"`
	UnappliedMigrations []string `json:"unapplied_migrations,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LatestMigration     respjson.Field
		UnappliedMigrations respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StatusGetResponseDatabaseMigrationStatus) RawJSON() string { return r.JSON.raw }
func (r *StatusGetResponseDatabaseMigrationStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StatusGetResponseIndex struct {
	Error        string    `json:"error,required"`
	LastModified time.Time `json:"last_modified,required" format:"date-time"`
	Status       string    `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error        respjson.Field
		LastModified respjson.Field
		Status       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StatusGetResponseIndex) RawJSON() string { return r.JSON.raw }
func (r *StatusGetResponseIndex) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StatusGetResponseSanityCheck struct {
	Error   string    `json:"error,required"`
	LastRun time.Time `json:"last_run,required" format:"date-time"`
	Status  string    `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		LastRun     respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StatusGetResponseSanityCheck) RawJSON() string { return r.JSON.raw }
func (r *StatusGetResponseSanityCheck) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StatusGetResponseStorage struct {
	Available int64 `json:"available,required"`
	Total     int64 `json:"total,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Available   respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StatusGetResponseStorage) RawJSON() string { return r.JSON.raw }
func (r *StatusGetResponseStorage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StatusGetResponseTasks struct {
	CeleryStatus string `json:"celery_status,required"`
	RedisError   string `json:"redis_error,required"`
	RedisStatus  string `json:"redis_status,required"`
	RedisURL     string `json:"redis_url,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CeleryStatus respjson.Field
		RedisError   respjson.Field
		RedisStatus  respjson.Field
		RedisURL     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StatusGetResponseTasks) RawJSON() string { return r.JSON.raw }
func (r *StatusGetResponseTasks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
