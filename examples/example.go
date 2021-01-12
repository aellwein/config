package main

import (
	"context"
	"fmt"

	"github.com/aellwein/config/v1/config"
)

func getConfig() config.Cfg {
	cfg, err := config.NewBuilder().
		ConfigName("example").
		ConfigType("yaml").
		ConfigPaths(
			"./examples/configs/",
			"$HOME/",
		).
		Defaults(map[string]interface{}{
			"foo":   2,
			"bar":   42,
			"baz":   []string{"BAZ", "na"},
			"error": false,
		}).
		Options(config.NoErrorOnMissingCfgFileOption).
		Build()

	if err != nil {
		panic(err)
	}
	return cfg
}

func getConfigContext() context.Context {
	return config.NewContext(getConfig())
}

func main() {
	// Use case #1: use config interface directly
	cfg := getConfig()
	fmt.Println(cfg.GetAll())

	// Use case #2: use context.Context object which can be used e.g. with CLI parsers
	ctx := getConfigContext()

	// get config and values from it
	cfg = ctx.Value(config.ContextConfigValue).(config.Cfg)
	fmt.Println(cfg.GetAll())
}
