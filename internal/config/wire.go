package config

import (
	"fmt"
	"github.com/google/wire"
	"github.com/spf13/cobra"
)

var ConfigSet = wire.NewSet(
	ProvideConfig,
)

func ProvideConfig(cmd *cobra.Command) (*Config, error) {
	InitConfig(cmd)
	cfg := GetConfig()
	if cfg == nil {
		return nil, fmt.Errorf("failed to initialize config")
	}
	return cfg, nil
}
