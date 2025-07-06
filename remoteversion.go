// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package paperlessngx

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/paperless-nix-go/internal/requestconfig"
	"github.com/stainless-sdks/paperless-nix-go/option"
)

// RemoteVersionService contains methods and other services that help with
// interacting with the paperless-ngx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRemoteVersionService] method instead.
type RemoteVersionService struct {
	Options []option.RequestOption
}

// NewRemoteVersionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRemoteVersionService(opts ...option.RequestOption) (r RemoteVersionService) {
	r = RemoteVersionService{}
	r.Options = opts
	return
}

// Get the current version of the Paperless-NGX server
func (r *RemoteVersionService) Get(ctx context.Context, opts ...option.RequestOption) (res *RemoteVersionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "api/remote_version/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type RemoteVersionGetResponse map[string]any
