// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package helloworldtestingggg

import (
	"context"
	"net/http"
	"os"
	"slices"
	"strings"

	"github.com/swagger-girly/go/internal/requestconfig"
	"github.com/swagger-girly/go/option"
)

// Client creates a struct with services and top level methods that help with
// interacting with the hello-world-testingggg API. You should not instantiate this
// client directly, and instead use the [NewClient] method instead.
type Client struct {
	options []option.RequestOption
	// Everything about your Pets
	Pet PetService
	// File storage operations
	Files FileService
	// Pet owner profile and compliance operations
	Profiles ProfileService
	// Adoption policies and applications
	Adoptions AdoptionService
	// Post-adoption placement tracking
	Placements    PlacementService
	Veterinary    VeterinaryService
	Webhooks      WebhookService
	Notifications NotificationService
	// Access to Petstore orders
	Store StoreService
	// Operations about user
	User  UserService
	AI    AIService
	Media MediaService
}

// DefaultClientOptions read from the environment (API_KEY, BASIC_AUTH_USERNAME,
// BASIC_AUTH_PASSWORD, PETSTORE_WEBHOOK_SECRET, HELLO_WORLD_TESTINGGGG_BASE_URL).
// This should be used to initialize new clients.
func DefaultClientOptions() []option.RequestOption {
	defaults := []option.RequestOption{option.WithHTTPClient(defaultHTTPClient()), option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("HELLO_WORLD_TESTINGGGG_BASE_URL"); ok {
		defaults = append(defaults, option.WithBaseURL(o))
	}
	if o, ok := os.LookupEnv("API_KEY"); ok {
		defaults = append(defaults, option.WithAPIKey(o))
	}
	if o, ok := os.LookupEnv("BASIC_AUTH_USERNAME"); ok {
		defaults = append(defaults, option.WithBasicAuthUsername(o))
	}
	if o, ok := os.LookupEnv("BASIC_AUTH_PASSWORD"); ok {
		defaults = append(defaults, option.WithBasicAuthPassword(o))
	}
	if o, ok := os.LookupEnv("PETSTORE_WEBHOOK_SECRET"); ok {
		defaults = append(defaults, option.WithWebhookSecret(o))
	}
	if o, ok := os.LookupEnv("HELLO_WORLD_TESTINGGGG_CUSTOM_HEADERS"); ok {
		for _, line := range strings.Split(o, "\n") {
			colon := strings.Index(line, ":")
			if colon >= 0 {
				defaults = append(defaults, option.WithHeader(strings.TrimSpace(line[:colon]), strings.TrimSpace(line[colon+1:])))
			}
		}
	}
	return defaults
}

// NewClient generates a new client with the default option read from the
// environment (API_KEY, BASIC_AUTH_USERNAME, BASIC_AUTH_PASSWORD,
// PETSTORE_WEBHOOK_SECRET, HELLO_WORLD_TESTINGGGG_BASE_URL). The option passed in
// as arguments are applied after these default arguments, and all option will be
// passed down to the services and requests that this client makes.
func NewClient(opts ...option.RequestOption) (r Client) {
	opts = append(DefaultClientOptions(), opts...)

	r = Client{options: opts}

	r.Pet = NewPetService(opts...)
	r.Files = NewFileService(opts...)
	r.Profiles = NewProfileService(opts...)
	r.Adoptions = NewAdoptionService(opts...)
	r.Placements = NewPlacementService(opts...)
	r.Veterinary = NewVeterinaryService(opts...)
	r.Webhooks = NewWebhookService(opts...)
	r.Notifications = NewNotificationService(opts...)
	r.Store = NewStoreService(opts...)
	r.User = NewUserService(opts...)
	r.AI = NewAIService(opts...)
	r.Media = NewMediaService(opts...)

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
	opts = slices.Concat(r.options, opts)
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

// Returns the current API health, including per-service statuses.
func (r *Client) Health(ctx context.Context, opts ...option.RequestOption) (res *SystemHealth, err error) {
	opts = slices.Concat(r.options, opts)
	path := "health"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns the caller's current rate-limit budget.
func (r *Client) GetRateLimits(ctx context.Context, opts ...option.RequestOption) (res *GetRateLimitsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "rate_limits"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}
