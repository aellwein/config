package config

import (
	"time"

	"github.com/spf13/viper"
)

type myCfg struct {
	v       *viper.Viper
	options []CfgOption
}

type newConfigT struct {
	cfg *myCfg
}

type configNameT struct {
	cfg *myCfg
}

type configTypeT struct {
	cfg *myCfg
}

type defaultsT struct {
	cfg *myCfg
}

type configPathsT struct {
	cfg *myCfg
}

type optionsT struct {
	cfg *myCfg
}

type CfgOption uint

const (
	NoErrorOnMissingCfgFileOption CfgOption = iota
)

// NewWithOptions creates a new config
func NewConfig() *newConfigT {
	return &newConfigT{
		cfg: &myCfg{
			v:       viper.New(),
			options: make([]CfgOption, 0),
		},
	}
}

func (t *newConfigT) ConfigName(name string) *configNameT {
	t.cfg.v.SetConfigName(name)
	return &configNameT{
		cfg: t.cfg,
	}
}

func (t *configNameT) ConfigType(typ string) *configTypeT {
	t.cfg.v.SetConfigType(typ)
	return &configTypeT{
		cfg: t.cfg,
	}
}

func configPaths(cfg *myCfg, paths ...string) *configPathsT {
	for _, p := range paths {
		cfg.v.AddConfigPath(p)
	}
	return &configPathsT{
		cfg: cfg,
	}
}

func (t *configNameT) ConfigPaths(path ...string) *configPathsT {
	return configPaths(t.cfg, path...)
}

func (t *configTypeT) ConfigPaths(path ...string) *configPathsT {
	return configPaths(t.cfg, path...)
}

func defaults(c *myCfg, defs map[string]interface{}) *defaultsT {
	for k, v := range defs {
		c.v.SetDefault(string(k), v)
	}
	return &defaultsT{
		cfg: c,
	}
}

func (t *configPathsT) Defaults(defs map[string]interface{}) *defaultsT {
	return defaults(t.cfg, defs)
}

func (t *configPathsT) NoDefaults() *defaultsT {
	return &defaultsT{
		cfg: t.cfg,
	}
}

func (t *defaultsT) Options(options ...CfgOption) *optionsT {
	for _, i := range options {
		t.cfg.options = append(t.cfg.options, i)
	}
	return &optionsT{
		cfg: t.cfg,
	}
}

func (c *myCfg) HasOption(option CfgOption) bool {
	for _, o := range c.options {
		if option == o {
			return true
		}
	}
	return false
}

func (t *optionsT) Build() (Cfg, error) {
	err := t.cfg.v.ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			if !t.cfg.HasOption(NoErrorOnMissingCfgFileOption) {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	return t.cfg, nil
}

func (c *myCfg) GetInt(key string) int {
	return c.v.GetInt(key)
}

func (c *myCfg) GetInt32(key string) int32 {
	return c.v.GetInt32(key)
}

func (c *myCfg) GetInt64(key string) int64 {
	return c.v.GetInt64(key)
}

func (c *myCfg) GetUint(key string) uint {
	return c.v.GetUint(key)
}

func (c *myCfg) GetUint32(key string) uint32 {
	return c.v.GetUint32(key)
}

func (c *myCfg) GetUint64(key string) uint64 {
	return c.v.GetUint64(key)
}

func (c *myCfg) GetString(key string) string {
	return c.v.GetString(key)
}

func (c *myCfg) GetBool(key string) bool {
	return c.v.GetBool(key)
}

func (c *myCfg) Get(key string) interface{} {
	return c.v.Get(key)
}

func (c *myCfg) GetFloat64(key string) float64 {
	return c.v.GetFloat64(key)
}

func (c *myCfg) GetAll() map[string]interface{} {
	return c.v.AllSettings()
}

func (c *myCfg) GetDuration(key string) time.Duration {
	return c.v.GetDuration(key)
}

func (c *myCfg) GetIntSlice(key string) []int {
	return c.v.GetIntSlice(key)
}

func (c *myCfg) GetSizeInBytes(key string) uint {
	return c.v.GetSizeInBytes(key)
}

func (c *myCfg) GetTime(key string) time.Time {
	return c.v.GetTime(key)
}

func (c *myCfg) GetStringSlice(key string) []string {
	return c.v.GetStringSlice(key)
}
