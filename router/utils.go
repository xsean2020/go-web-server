package router

import (
	"net/http"
	"strings"
)

var (
	validHttpMethod = map[string]string{
		strings.ToLower(http.MethodGet):     http.MethodGet,
		strings.ToLower(http.MethodHead):    http.MethodHead,
		strings.ToLower(http.MethodPost):    http.MethodPost,
		strings.ToLower(http.MethodPut):     http.MethodPut,
		strings.ToLower(http.MethodPatch):   http.MethodPatch, // RFC 5789
		strings.ToLower(http.MethodDelete):  http.MethodDelete,
		strings.ToLower(http.MethodConnect): http.MethodConnect,
		strings.ToLower(http.MethodOptions): http.MethodOptions,
		strings.ToLower(http.MethodTrace):   http.MethodTrace,
	}
)
