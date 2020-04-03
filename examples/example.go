package main

import (
	"fmt"

	"github.com/aellwein/config/pkg/config"
)

func main() {
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

	fmt.Println(cfg.GetAll())
}
