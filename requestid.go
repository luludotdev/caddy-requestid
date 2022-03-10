package requestid

import (
	"net/http"
	"strconv"

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
type RequestID struct {
	// Length of standalone ID
	Length int `json:"length"`

	// Map of additional IDs to set
	Additional map[string]int `json:"additional,omitempty"`
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
	if m.Length < 1 {
		m.Length = 21
	}

	if m.Additional == nil {
		m.Additional = make(map[string]int)
	}

	return nil
}

// ServeHTTP implements caddyhttp.MiddlewareHandler.
func (m RequestID) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	repl := r.Context().Value(caddy.ReplacerCtxKey).(*caddy.Replacer)

	id := nanoid.Must(m.Length)
	repl.Set("http.request_id", id)

	for key, value := range m.Additional {
		id := nanoid.Must(value)
		repl.Set("http.request_id."+key, id)
	}

	return next.ServeHTTP(w, r)
}

// UnmarshalCaddyfile - this is a no-op
func (m *RequestID) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	arg1 := d.NextArg()
	arg2 := d.NextArg()

	// Parse standalone length
	if arg1 && arg2 {
		val := d.Val()
		len, err := strconv.Atoi(val)

		if err != nil {
			return d.Err("failed to convert length to int")
		}

		if len < 1 {
			return d.Err("length cannot be less than 1")
		}

		m.Length = len
	}

	if m.Additional == nil {
		m.Additional = make(map[string]int)
	}

	// Parse additional IDs
	for d.NextBlock(0) {
		key := d.Val()
		if !d.NextArg() {
			return d.ArgErr()
		}

		val := d.Val()
		len, err := strconv.Atoi(val)

		if err != nil {
			return d.Err("failed to convert length to int")
		}

		if len < 1 {
			return d.Err("length cannot be less than 1")
		}

		if _, ok := m.Additional[key]; ok {
			return d.Errf("duplicate key: %v\n", key)
		}

		m.Additional[key] = len
	}

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
	_ caddy.Provisioner           = (*RequestID)(nil)
	_ caddyhttp.MiddlewareHandler = (*RequestID)(nil)
	_ caddyfile.Unmarshaler       = (*RequestID)(nil)
)
