// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"net/http"
	"reflect"

	"github.com/swagger-girly/go/internal/apijson"
	"github.com/swagger-girly/go/internal/requestconfig"
	"github.com/swagger-girly/go/option"
	"github.com/swagger-girly/go/packages/param"
	"github.com/swagger-girly/go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type CustomCursorPage[T any] struct {
	Items      []T    `json:"items"`
	NextCursor string `json:"next_cursor" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r CustomCursorPage[T]) RawJSON() string { return r.JSON.raw }
func (r *CustomCursorPage[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *CustomCursorPage[T]) GetNextPage() (res *CustomCursorPage[T], err error) {
	if len(r.Items) == 0 {
		return nil, nil
	}
	next := r.NextCursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *CustomCursorPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &CustomCursorPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type CustomCursorPageAutoPager[T any] struct {
	page *CustomCursorPage[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewCustomCursorPageAutoPager[T any](page *CustomCursorPage[T], err error) *CustomCursorPageAutoPager[T] {
	return &CustomCursorPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *CustomCursorPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Items) == 0 {
		return false
	}
	if r.idx >= len(r.page.Items) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Items) == 0 {
			return false
		}
	}
	r.cur = r.page.Items[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *CustomCursorPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *CustomCursorPageAutoPager[T]) Err() error {
	return r.err
}

func (r *CustomCursorPageAutoPager[T]) Index() int {
	return r.run
}

type XFakeSinglePage[T any] struct {
	Data    []T  `json:"data"`
	HasMore bool `json:"has_more" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		HasMore     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r XFakeSinglePage[T]) RawJSON() string { return r.JSON.raw }
func (r *XFakeSinglePage[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *XFakeSinglePage[T]) GetNextPage() (res *XFakeSinglePage[T], err error) {
	if len(r.Data) == 0 {
		return nil, nil
	}

	if r.JSON.HasMore.Valid() && r.HasMore == false {
		return nil, nil
	}
	// This page represents a response that isn't actually paginated at the API level
	// so there will never be a next page.
	cfg := (*requestconfig.RequestConfig)(nil)
	if cfg == nil {
		return nil, nil
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *XFakeSinglePage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &XFakeSinglePage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type XFakeSinglePageAutoPager[T any] struct {
	page *XFakeSinglePage[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewXFakeSinglePageAutoPager[T any](page *XFakeSinglePage[T], err error) *XFakeSinglePageAutoPager[T] {
	return &XFakeSinglePageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *XFakeSinglePageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Data) == 0 {
		return false
	}
	if r.idx >= len(r.page.Data) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Data) == 0 {
			return false
		}
	}
	r.cur = r.page.Data[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *XFakeSinglePageAutoPager[T]) Current() T {
	return r.cur
}

func (r *XFakeSinglePageAutoPager[T]) Err() error {
	return r.err
}

func (r *XFakeSinglePageAutoPager[T]) Index() int {
	return r.run
}

type ReportCursorPage[T any] struct {
	Data    []T  `json:"data"`
	HasMore bool `json:"has_more" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		HasMore     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r ReportCursorPage[T]) RawJSON() string { return r.JSON.raw }
func (r *ReportCursorPage[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *ReportCursorPage[T]) GetNextPage() (res *ReportCursorPage[T], err error) {
	if len(r.Data) == 0 {
		return nil, nil
	}

	if r.JSON.HasMore.Valid() && r.HasMore == false {
		return nil, nil
	}
	items := r.Data
	if items == nil || len(items) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	value := reflect.ValueOf(items[len(items)-1])
	field := value.FieldByName("ID")
	err = cfg.Apply(option.WithQuery("starting_after", field.Interface().(string)))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *ReportCursorPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &ReportCursorPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type ReportCursorPageAutoPager[T any] struct {
	page *ReportCursorPage[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewReportCursorPageAutoPager[T any](page *ReportCursorPage[T], err error) *ReportCursorPageAutoPager[T] {
	return &ReportCursorPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *ReportCursorPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Data) == 0 {
		return false
	}
	if r.idx >= len(r.page.Data) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Data) == 0 {
			return false
		}
	}
	r.cur = r.page.Data[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *ReportCursorPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *ReportCursorPageAutoPager[T]) Err() error {
	return r.err
}

func (r *ReportCursorPageAutoPager[T]) Index() int {
	return r.run
}
