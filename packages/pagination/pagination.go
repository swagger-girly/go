// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"net/http"

	"github.com/swagger-girly/go/internal/apijson"
	"github.com/swagger-girly/go/internal/requestconfig"
	"github.com/swagger-girly/go/option"
)

type CustomCursorPage[T any] struct {
	Items      []T                  `json:"items"`
	NextCursor string               `json:"next_cursor" api:"nullable"`
	JSON       customCursorPageJSON `json:"-"`
	cfg        *requestconfig.RequestConfig
	res        *http.Response
}

// customCursorPageJSON contains the JSON metadata for the struct
// [CustomCursorPage[T]]
type customCursorPageJSON struct {
	Items       apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomCursorPage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customCursorPageJSON) RawJSON() string {
	return r.raw
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
	Data    []T                 `json:"data"`
	HasMore bool                `json:"has_more" api:"required"`
	JSON    xFakeSinglePageJSON `json:"-"`
	cfg     *requestconfig.RequestConfig
	res     *http.Response
}

// xFakeSinglePageJSON contains the JSON metadata for the struct
// [XFakeSinglePage[T]]
type xFakeSinglePageJSON struct {
	Data        apijson.Field
	HasMore     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *XFakeSinglePage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r xFakeSinglePageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *XFakeSinglePage[T]) GetNextPage() (res *XFakeSinglePage[T], err error) {
	if len(r.Data) == 0 {
		return nil, nil
	}

	if !r.JSON.HasMore.IsMissing() && r.HasMore == false {
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
