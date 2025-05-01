package config

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var (
	// cfgFile config file
	cfgFile string
	// globalConfig global configuration instance
	globalConfig Config
	// isInitialized whether it has been initialized
	isInitialized bool
)

func InitConfig(rootCmd *cobra.Command) {
	if isInitialized {
		return
	}
	isInitialized = true
	cobra.OnInitialize(loadConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file")
	bindFlags(&globalConfig, rootCmd, "")
}

// loadConfig loads configuration file
func loadConfig() {
	// Set configuration file path
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("configs")
	}

	// Set Viper
	viper.AutomaticEnv()                                   // Automatically read environment variables
	viper.SetEnvPrefix("CINEMAGO")                         // Set environment variable prefix
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // Replace '.' with '_' in environment variable keys

	// Read configuration file
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("[LoadConfig] Failed to read configuration file: %v\n", err)
	}

	// Unmarshal configuration to struct
	if err := viper.Unmarshal(&globalConfig); err != nil {
		fmt.Printf("[LoadConfig] Failed to unmarshal configuration: %v\n", err)
	}

	// Set environment variable
	env := os.Getenv("APP_ENV")
	if env != "" {
		globalConfig.Env = env
	}
	if globalConfig.Env == "" {
		globalConfig.Env = "dev"
	}
}

// GetConfig returns global configuration
func GetConfig() *Config {
	return &globalConfig
}
