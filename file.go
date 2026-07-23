// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package helloworldtestingggg

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/swagger-girly/go/internal/apiform"
	"github.com/swagger-girly/go/internal/apijson"
	"github.com/swagger-girly/go/internal/apiquery"
	"github.com/swagger-girly/go/internal/requestconfig"
	"github.com/swagger-girly/go/option"
	"github.com/swagger-girly/go/packages/param"
	"github.com/swagger-girly/go/packages/respjson"
)

// File storage operations
//
// FileService contains methods and other services that help with interacting with
// the hello-world-testingggg API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFileService] method instead.
type FileService struct {
	options []option.RequestOption
}

// NewFileService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFileService(opts ...option.RequestOption) (r FileService) {
	r = FileService{}
	r.options = opts
	return
}

// Updates a stored file with either a multipart file part or a URL.
func (r *FileService) Update(ctx context.Context, path string, body FileUpdateParams, opts ...option.RequestOption) (res *File, err error) {
	opts = slices.Concat(r.options, opts)
	if path == "" {
		err = errors.New("missing required path parameter")
		return nil, err
	}
	path := fmt.Sprintf("files/%s", url.PathEscape(path))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Returns all files in descending order.
func (r *FileService) List(ctx context.Context, opts ...option.RequestOption) (res *Fileslist, err error) {
	opts = slices.Concat(r.options, opts)
	path := "files"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Deletes a file while preserving slashes in the file path parameter.
func (r *FileService) Delete(ctx context.Context, path string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if path == "" {
		err = errors.New("missing required path parameter")
		return err
	}
	path := fmt.Sprintf("files/%s", url.PathEscape(path))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Uploads multiple files and nested manifest metadata as multipart/form-data.
func (r *FileService) NewArchive(ctx context.Context, body FileNewArchiveParams, opts ...option.RequestOption) (res *Archive, err error) {
	opts = slices.Concat(r.options, opts)
	path := "files/archive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Downloads a file while preserving slashes in the file path parameter.
func (r *FileService) Download(ctx context.Context, path string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	if path == "" {
		err = errors.New("missing required path parameter")
		return nil, err
	}
	path := fmt.Sprintf("files/%s", url.PathEscape(path))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Uploads a file via multipart/form-data with metadata and optional storage path.
func (r *FileService) Upload(ctx context.Context, body FileUploadParams, opts ...option.RequestOption) (res *File, err error) {
	opts = slices.Concat(r.options, opts)
	path := "files"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Uploads a whole request body as application/octet-stream.
func (r *FileService) UploadDirect(ctx context.Context, body io.Reader, params FileUploadDirectParams, opts ...option.RequestOption) (res *File, err error) {
	if !param.IsOmitted(params.XFileName) {
		opts = append(opts, option.WithHeader("X-File-Name", fmt.Sprintf("%v", params.XFileName.Value)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithRequestBody("application/octet-stream", body)}, opts...)
	path := "files/direct"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type Archive struct {
	ID string `json:"id" api:"required"`
	// Any of "queued", "processing", "complete".
	Status    ArchiveStatus `json:"status" api:"required"`
	FileCount int64         `json:"fileCount"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Status      respjson.Field
		FileCount   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Archive) RawJSON() string { return r.JSON.raw }
func (r *Archive) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ArchiveStatus string

const (
	ArchiveStatusQueued     ArchiveStatus = "queued"
	ArchiveStatusProcessing ArchiveStatus = "processing"
	ArchiveStatusComplete   ArchiveStatus = "complete"
)

type File struct {
	// Timestamp when the file was last modified.
	LastModified time.Time `json:"lastModified" api:"required" format:"date-time"`
	// Path to the file in storage.
	Path string `json:"path" api:"required"`
	// Size of the file in bytes.
	Size int64 `json:"size" api:"required"`
	// Detected MIME type for the file.
	ContentType string `json:"contentType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LastModified respjson.Field
		Path         respjson.Field
		Size         respjson.Field
		ContentType  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r File) RawJSON() string { return r.JSON.raw }
func (r *File) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Fileslist struct {
	Data []File `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Fileslist) RawJSON() string { return r.JSON.raw }
func (r *Fileslist) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileUpdateParams struct {
	// URL to fetch replacement file content from.
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	// Replacement file content.
	File io.Reader `json:"file,omitzero" format:"binary"`
	// Metadata to merge onto the stored file.
	Metadata map[string]string `json:"metadata,omitzero"`
	paramObj
}

func (r FileUpdateParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

type FileNewArchiveParams struct {
	// Files to include in the archive.
	Files    []io.Reader                  `json:"files,omitzero" api:"required" format:"binary"`
	Manifest FileNewArchiveParamsManifest `json:"manifest,omitzero"`
	paramObj
}

func (r FileNewArchiveParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

// The property Name is required.
type FileNewArchiveParamsManifest struct {
	Name          string          `json:"name" api:"required"`
	PreservePaths param.Opt[bool] `json:"preservePaths,omitzero"`
	Labels        []string        `json:"labels,omitzero"`
	paramObj
}

func (r FileNewArchiveParamsManifest) MarshalJSON() (data []byte, err error) {
	type shadow FileNewArchiveParamsManifest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileNewArchiveParamsManifest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileUploadParams struct {
	// The file to upload.
	File io.Reader `json:"file,omitzero" api:"required" format:"binary"`
	// Optional expiration in seconds.
	ExpiresIn param.Opt[int64] `json:"expiresIn,omitzero"`
	// Optional storage path for the uploaded file.
	Path param.Opt[string] `json:"path,omitzero"`
	// Whether the file should be publicly readable.
	Public param.Opt[bool] `json:"public,omitzero"`
	// Additional metadata to store with the file.
	Metadata map[string]string `json:"metadata,omitzero"`
	// Tags to associate with the file.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r FileUploadParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

type FileUploadDirectParams struct {
	// Optional storage path for the raw file.
	Path      param.Opt[string] `query:"path,omitzero" json:"-"`
	XFileName param.Opt[string] `header:"X-File-Name,omitzero" json:"-"`
	paramObj
}

func (r FileUploadDirectParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

// URLQuery serializes [FileUploadDirectParams]'s query parameters as `url.Values`.
func (r FileUploadDirectParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
