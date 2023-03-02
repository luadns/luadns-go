package luadns

import (
	"strconv"
)

type ErrBadStatusCode struct {
	StatusCode int
}

func (e *ErrBadStatusCode) Error() string {
	return "Server returned bad status code (" + strconv.Itoa(e.StatusCode) + ")"
}

type ErrBadContentType struct {
	ContentType string
}

func (e *ErrBadContentType) Error() string {
	return "Server returned bad content type (" + e.ContentType + ")"
}

type ErrTooManyRequests struct {
	Limit int64
	Reset int64
}

func (e *ErrTooManyRequests) Error() string {
	return "Too many requests, retry after " + strconv.FormatInt(e.Reset, 10) + " unix time"
}
