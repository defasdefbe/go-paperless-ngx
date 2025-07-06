// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package paperlessngx

import (
	"context"
	"net/http"
	"os"

	"github.com/stainless-sdks/paperless-nix-go/internal/requestconfig"
	"github.com/stainless-sdks/paperless-nix-go/option"
)

// Client creates a struct with services and top level methods that help with
// interacting with the paperless-ngx API. You should not instantiate this client
// directly, and instead use the [NewClient] method instead.
type Client struct {
	Options          []option.RequestOption
	BulkEditObjects  BulkEditObjectService
	Config           ConfigService
	Correspondents   CorrespondentService
	CustomFields     CustomFieldService
	DocumentTypes    DocumentTypeService
	Documents        DocumentService
	Groups           GroupService
	Logs             LogService
	MailAccounts     MailAccountService
	MailRules        MailRuleService
	OAuth            OAuthService
	Profile          ProfileService
	RemoteVersion    RemoteVersionService
	SavedViews       SavedViewService
	Search           SearchService
	ShareLinks       ShareLinkService
	Statistics       StatisticService
	Status           StatusService
	StoragePaths     StoragePathService
	Tags             TagService
	Tasks            TaskService
	Token            TokenService
	Trash            TrashService
	UiSettings       UiSettingService
	Users            UserService
	WorkflowActions  WorkflowActionService
	WorkflowTriggers WorkflowTriggerService
	Workflows        WorkflowService
}

// DefaultClientOptions read from the environment (PAPERLESS_NIX_API_KEY,
// PAPERLESS_NIX_USERNAME, PAPERLESS_NIX_PASSWORD, PAPERLESS_NGX_BASE_URL). This
// should be used to initialize new clients.
func DefaultClientOptions() []option.RequestOption {
	defaults := []option.RequestOption{option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("PAPERLESS_NGX_BASE_URL"); ok {
		defaults = append(defaults, option.WithBaseURL(o))
	}
	if o, ok := os.LookupEnv("PAPERLESS_NIX_USERNAME"); ok {
		defaults = append(defaults, option.WithUsername(o))
	}
	if o, ok := os.LookupEnv("PAPERLESS_NIX_PASSWORD"); ok {
		defaults = append(defaults, option.WithPassword(o))
	}
	if o, ok := os.LookupEnv("PAPERLESS_NIX_API_KEY"); ok {
		defaults = append(defaults, option.WithAPIKey(o))
	}
	return defaults
}

// NewClient generates a new client with the default option read from the
// environment (PAPERLESS_NIX_API_KEY, PAPERLESS_NIX_USERNAME,
// PAPERLESS_NIX_PASSWORD, PAPERLESS_NGX_BASE_URL). The option passed in as
// arguments are applied after these default arguments, and all option will be
// passed down to the services and requests that this client makes.
func NewClient(opts ...option.RequestOption) (r Client) {
	opts = append(DefaultClientOptions(), opts...)

	r = Client{Options: opts}

	r.BulkEditObjects = NewBulkEditObjectService(opts...)
	r.Config = NewConfigService(opts...)
	r.Correspondents = NewCorrespondentService(opts...)
	r.CustomFields = NewCustomFieldService(opts...)
	r.DocumentTypes = NewDocumentTypeService(opts...)
	r.Documents = NewDocumentService(opts...)
	r.Groups = NewGroupService(opts...)
	r.Logs = NewLogService(opts...)
	r.MailAccounts = NewMailAccountService(opts...)
	r.MailRules = NewMailRuleService(opts...)
	r.OAuth = NewOAuthService(opts...)
	r.Profile = NewProfileService(opts...)
	r.RemoteVersion = NewRemoteVersionService(opts...)
	r.SavedViews = NewSavedViewService(opts...)
	r.Search = NewSearchService(opts...)
	r.ShareLinks = NewShareLinkService(opts...)
	r.Statistics = NewStatisticService(opts...)
	r.Status = NewStatusService(opts...)
	r.StoragePaths = NewStoragePathService(opts...)
	r.Tags = NewTagService(opts...)
	r.Tasks = NewTaskService(opts...)
	r.Token = NewTokenService(opts...)
	r.Trash = NewTrashService(opts...)
	r.UiSettings = NewUiSettingService(opts...)
	r.Users = NewUserService(opts...)
	r.WorkflowActions = NewWorkflowActionService(opts...)
	r.WorkflowTriggers = NewWorkflowTriggerService(opts...)
	r.Workflows = NewWorkflowService(opts...)

	return
}

// Execute makes a request with the given context, method, URL, request params,
// response, and request options. This is useful for hitting undocumented endpoints
// while retaining the base URL, auth, retries, and other options from the client.
//
// If a byte slice or an [io.Reader] is supplied to params, it will be used as-is
// for the request body.
//
// The params is by default serialized into the body using [encoding/json]. If your
// type implements a MarshalJSON function, it will be used instead to serialize the
// request. If a URLQuery method is implemented, the returned [url.Values] will be
// used as query strings to the url.
//
// If your params struct uses [param.Field], you must provide either [MarshalJSON],
// [URLQuery], and/or [MarshalForm] functions. It is undefined behavior to use a
// struct uses [param.Field] without specifying how it is serialized.
//
// Any "…Params" object defined in this library can be used as the request
// argument. Note that 'path' arguments will not be forwarded into the url.
//
// The response body will be deserialized into the res variable, depending on its
// type:
//
//   - A pointer to a [*http.Response] is populated by the raw response.
//   - A pointer to a byte array will be populated with the contents of the request
//     body.
//   - A pointer to any other type uses this library's default JSON decoding, which
//     respects UnmarshalJSON if it is defined on the type.
//   - A nil value will not read the response body.
//
// For even greater flexibility, see [option.WithResponseInto] and
// [option.WithResponseBodyInto].
func (r *Client) Execute(ctx context.Context, method string, path string, params any, res any, opts ...option.RequestOption) error {
	opts = append(r.Options, opts...)
	return requestconfig.ExecuteNewRequest(ctx, method, path, params, res, opts...)
}

// Get makes a GET request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Get(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodGet, path, params, res, opts...)
}

// Post makes a POST request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Post(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPost, path, params, res, opts...)
}

// Put makes a PUT request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Put(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPut, path, params, res, opts...)
}

// Patch makes a PATCH request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Patch(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPatch, path, params, res, opts...)
}

// Delete makes a DELETE request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Delete(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodDelete, path, params, res, opts...)
}
