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
	"github.com/swagger-girly/go/internal/param"
	"github.com/swagger-girly/go/internal/requestconfig"
	"github.com/swagger-girly/go/option"
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
	Options []option.RequestOption
}

// NewFileService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFileService(opts ...option.RequestOption) (r *FileService) {
	r = &FileService{}
	r.Options = opts
	return
}

// Updates a stored file with either a multipart file part or a URL.
func (r *FileService) Update(ctx context.Context, path string, body FileUpdateParams, opts ...option.RequestOption) (res *File, err error) {
	opts = slices.Concat(r.Options, opts)
	if path == "" {
		err = errors.New("missing required path parameter")
		return nil, err
	}
	path := fmt.Sprintf("files/%s", path)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Returns all files in descending order.
func (r *FileService) List(ctx context.Context, opts ...option.RequestOption) (res *Fileslist, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "files"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Deletes a file while preserving slashes in the file path parameter.
func (r *FileService) Delete(ctx context.Context, path string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if path == "" {
		err = errors.New("missing required path parameter")
		return err
	}
	path := fmt.Sprintf("files/%s", path)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Uploads multiple files and nested manifest metadata as multipart/form-data.
func (r *FileService) NewArchive(ctx context.Context, body FileNewArchiveParams, opts ...option.RequestOption) (res *Archive, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "files/archive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Downloads a file while preserving slashes in the file path parameter.
func (r *FileService) Download(ctx context.Context, path string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	if path == "" {
		err = errors.New("missing required path parameter")
		return nil, err
	}
	path := fmt.Sprintf("files/%s", path)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Uploads a file via multipart/form-data with metadata and optional storage path.
func (r *FileService) Upload(ctx context.Context, body FileUploadParams, opts ...option.RequestOption) (res *File, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "files"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Uploads a whole request body as application/octet-stream.
func (r *FileService) UploadDirect(ctx context.Context, body io.Reader, params FileUploadDirectParams, opts ...option.RequestOption) (res *File, err error) {
	if params.XFileName.Present {
		opts = append(opts, option.WithHeader("X-File-Name", fmt.Sprintf("%v", params.XFileName)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithRequestBody("application/octet-stream", body)}, opts...)
	path := "files/direct"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type Archive struct {
	ID        string        `json:"id" api:"required"`
	Status    ArchiveStatus `json:"status" api:"required"`
	FileCount int64         `json:"fileCount"`
	JSON      archiveJSON   `json:"-"`
}

// archiveJSON contains the JSON metadata for the struct [Archive]
type archiveJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	FileCount   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Archive) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r archiveJSON) RawJSON() string {
	return r.raw
}

type ArchiveStatus string

const (
	ArchiveStatusQueued     ArchiveStatus = "queued"
	ArchiveStatusProcessing ArchiveStatus = "processing"
	ArchiveStatusComplete   ArchiveStatus = "complete"
)

func (r ArchiveStatus) IsKnown() bool {
	switch r {
	case ArchiveStatusQueued, ArchiveStatusProcessing, ArchiveStatusComplete:
		return true
	}
	return false
}

type File struct {
	// Timestamp when the file was last modified.
	LastModified time.Time `json:"lastModified" api:"required" format:"date-time"`
	// Path to the file in storage.
	Path string `json:"path" api:"required"`
	// Size of the file in bytes.
	Size int64 `json:"size" api:"required"`
	// Detected MIME type for the file.
	ContentType string   `json:"contentType"`
	JSON        fileJSON `json:"-"`
}

// fileJSON contains the JSON metadata for the struct [File]
type fileJSON struct {
	LastModified apijson.Field
	Path         apijson.Field
	Size         apijson.Field
	ContentType  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *File) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fileJSON) RawJSON() string {
	return r.raw
}

type Fileslist struct {
	Data []File        `json:"data" api:"required"`
	JSON fileslistJSON `json:"-"`
}

// fileslistJSON contains the JSON metadata for the struct [Fileslist]
type fileslistJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Fileslist) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fileslistJSON) RawJSON() string {
	return r.raw
}

type FileUpdateParams struct {
	// Replacement file content.
	File param.Field[io.Reader] `json:"file" format:"binary"`
	// Metadata to merge onto the stored file.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// URL to fetch replacement file content from.
	URL param.Field[string] `json:"url" format:"uri"`
}

func (r FileUpdateParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
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
	Files    param.Field[[]io.Reader]                  `json:"files" api:"required" format:"binary"`
	Manifest param.Field[FileNewArchiveParamsManifest] `json:"manifest"`
}

func (r FileNewArchiveParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
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

type FileNewArchiveParamsManifest struct {
	Name          param.Field[string]   `json:"name" api:"required"`
	Labels        param.Field[[]string] `json:"labels"`
	PreservePaths param.Field[bool]     `json:"preservePaths"`
}

func (r FileNewArchiveParamsManifest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FileUploadParams struct {
	// The file to upload.
	File param.Field[io.Reader] `json:"file" api:"required" format:"binary"`
	// Optional expiration in seconds.
	ExpiresIn param.Field[int64] `json:"expiresIn"`
	// Additional metadata to store with the file.
	Metadata param.Field[map[string]string] `json:"metadata"`
	// Optional storage path for the uploaded file.
	Path param.Field[string] `json:"path"`
	// Whether the file should be publicly readable.
	Public param.Field[bool] `json:"public"`
	// Tags to associate with the file.
	Tags param.Field[[]string] `json:"tags"`
}

func (r FileUploadParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
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
	Path      param.Field[string] `query:"path"`
	XFileName param.Field[string] `header:"X-File-Name"`
}

func (r FileUploadDirectParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
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
func (r FileUploadDirectParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
