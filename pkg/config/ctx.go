package config

import (
	"time"
)

const (
	// ContextConfigValue can be used as key to query the config object
	ContextConfigValue = "github.com/aellwein/config/pkg/config.Cfg"
)

// Context is used to pass context with underlying config object to commands.
type Context struct {
	cfg Cfg
}

// NewContext creates a context containing the given config object
func NewContext(cfg Cfg) *Context {
	return &Context{
		cfg: cfg,
	}
}

// Deadline has a no-op implementation of context.Context.Deadline()
func (c *Context) Deadline() (deadline time.Time, ok bool) {
	return time.Time{}, false
}

// Done has a no-op implementation of context.Context.Done()
func (c *Context) Done() <-chan struct{} {
	return nil
}

// Err has a no-op implementation of context.Context.Err()
func (c *Context) Err() error {
	return nil
}

// Value provides the actual method to get the config object from context
func (c *Context) Value(key interface{}) interface{} {
	if k, ok := key.(string); ok {
		switch k {
		case ContextConfigValue:
			return c.cfg
		default:
			// unknown key
			return nil
		}
	} else {
		panic("key must be of type string")
	}
}
