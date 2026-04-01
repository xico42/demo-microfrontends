package canary

import (
	"math/rand"
	"net/http"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"go.uber.org/zap"
)

// Canary is a middleware that randomly sets a placeholder variable
// indicating which upstream should handle the request.
// The Caddyfile uses expression matchers on {canary.upstream}
// to route via standard reverse_proxy.
type Canary struct {
	// Weight is the probability (0.0–1.0) of selecting the canary upstream.
	Weight float64 `json:"weight"`
	logger *zap.Logger
}

func init() {
	caddy.RegisterModule(new(Canary))
}

func (c *Canary) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID: "http.handlers.canary",
		New: func() caddy.Module {
			return new(Canary)
		},
	}
}

func (c *Canary) Provision(ctx caddy.Context) error {
	c.logger = ctx.Logger()
	if c.Weight == 0 {
		c.Weight = 0.5
	}
	return nil
}

func (c *Canary) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	repl := r.Context().Value(caddy.ReplacerCtxKey).(*caddy.Replacer)

	upstream := "primary"
	if rand.Float64() < c.Weight {
		upstream = "canary"
	}

	repl.Set("canary.upstream", upstream)

	c.logger.Info("canary decision",
		zap.String("upstream", upstream),
		zap.String("path", r.URL.Path),
	)

	return next.ServeHTTP(w, r)
}

var (
	_ caddy.Provisioner           = (*Canary)(nil)
	_ caddy.Module                = (*Canary)(nil)
	_ caddyhttp.MiddlewareHandler = (*Canary)(nil)
	_ caddyfile.Unmarshaler       = (*Canary)(nil)
)
