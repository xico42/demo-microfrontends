package canary

import (
	"strconv"

	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

func init() {
	httpcaddyfile.RegisterHandlerDirective("canary", parseCaddyfile)
	httpcaddyfile.RegisterDirectiveOrder("canary", httpcaddyfile.Before, "reverse_proxy")
}

func parseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	c := new(Canary)
	err := c.UnmarshalCaddyfile(h.Dispenser)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Canary) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	d.Next() // consume directive name

	for d.NextBlock(d.Nesting()) {
		switch d.Val() {
		case "weight":
			args := d.RemainingArgs()
			if len(args) != 1 {
				return d.ArgErr()
			}
			w, err := strconv.ParseFloat(args[0], 64)
			if err != nil {
				return d.Errf("invalid weight value: %v", err)
			}
			c.Weight = w
		default:
			return d.Errf("unrecognized subdirective '%s'", d.Val())
		}
	}
	return nil
}
