package protocolstate

import (
	"github.com/onaio/nuclei/v2/pkg/types"
	"github.com/pkg/errors"
	"github.com/projectdiscovery/fastdialer/fastdialer"
)

// Dialer is a shared fastdialer instance for host DNS resolution
var Dialer *fastdialer.Dialer

// Init creates the Dialer instance based on user configuration
func Init(options *types.Options) error {
	opts := fastdialer.DefaultOptions
	if options.SystemResolvers {
		opts.EnableFallback = true
	}
	if options.ResolversFile != "" {
		opts.BaseResolvers = options.InternalResolversList
	}
	dialer, err := fastdialer.NewDialer(opts)
	if err != nil {
		return errors.Wrap(err, "could not create dialer")
	}
	Dialer = dialer
	return nil
}

// Close closes the global shared fastdialer
func Close() {
	if Dialer != nil {
		Dialer.Close()
	}
}
