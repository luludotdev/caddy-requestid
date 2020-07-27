package requestid

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"github.com/google/uuid"
)

func init() {
	caddy.RegisterModule(RequestID{})
	httpcaddyfile.RegisterHandlerDirective("request_id", parseCaddyfile)
}

// RequestID implements an HTTP handler that writes a
// unique request ID to response headers.
type RequestID struct {
	// The name of the header to write to.
	// Defaults to "x-request-id"
	Header string `json:"header,omitempty"`

	h string
}

// CaddyModule returns the Caddy module information.
func (RequestID) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.request_id",
		New: func() caddy.Module { return new(RequestID) },
	}
}

// Provision implements caddy.Provisioner.
func (m *RequestID) Provision(ctx caddy.Context) error {
	m.h = m.Header

	if m.h == "" {
		m.h = "x-request-id"
	}

	return nil
}

// Validate implements caddy.Validator.
func (m *RequestID) Validate() error {
	if m.h == "" {
		return fmt.Errorf("no header")
	}

	return nil
}

// ServeHTTP implements caddyhttp.MiddlewareHandler.
func (m RequestID) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	uid := uuid.New().String()
	w.Header().Set(m.h, strings.ReplaceAll(uid, "-", ""))

	return next.ServeHTTP(w, r)
}

// UnmarshalCaddyfile implements caddyfile.Unmarshaler.
func (m *RequestID) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			// optional arg
			m.Header = d.Val()
		}

		if d.NextArg() {
			// too many args
			return d.ArgErr()
		}
	}

	return nil
}

// parseCaddyfile unmarshals tokens from h into a new Middleware.
func parseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	var m RequestID
	err := m.UnmarshalCaddyfile(h.Dispenser)

	return m, err
}

// Interface guards
var (
	_ caddy.Provisioner           = (*RequestID)(nil)
	_ caddy.Validator             = (*RequestID)(nil)
	_ caddyhttp.MiddlewareHandler = (*RequestID)(nil)
	_ caddyfile.Unmarshaler       = (*RequestID)(nil)
)
