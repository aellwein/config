package main

import (
	"fmt"

	"github.com/aellwein/config/pkg/config"
)

func main() {
	cfg, err := config.NewConfig().
		ConfigName("example").
		ConfigType("yaml").
		ConfigPaths(".").
		Defaults(map[string]interface{}{
			"foo":   1,
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
