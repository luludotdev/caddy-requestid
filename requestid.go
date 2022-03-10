package requestid

import (
	"net/http"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	nanoid "github.com/matoous/go-nanoid/v2"
)

func init() {
	caddy.RegisterModule(RequestID{})
	httpcaddyfile.RegisterHandlerDirective("request_id", parseCaddyfile)
}

// RequestID implements an HTTP handler that writes a
// unique request ID to response headers.
type RequestID struct{}

// CaddyModule returns the Caddy module information.
func (RequestID) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.request_id",
		New: func() caddy.Module { return new(RequestID) },
	}
}

// ServeHTTP implements caddyhttp.MiddlewareHandler.
func (m RequestID) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	repl := r.Context().Value(caddy.ReplacerCtxKey).(*caddy.Replacer)

	id := nanoid.Must()
	repl.Set("http.request_id", id)

	return next.ServeHTTP(w, r)
}

// UnmarshalCaddyfile - this is a no-op
func (m *RequestID) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	return nil
}

func parseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	m := new(RequestID)
	err := m.UnmarshalCaddyfile(h.Dispenser)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// Interface guards
var (
	_ caddyhttp.MiddlewareHandler = (*RequestID)(nil)
	_ caddyfile.Unmarshaler       = (*RequestID)(nil)
)
