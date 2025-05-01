package main

import (
	"cinemago/internal/config"
	"cinemago/internal/server"
	"cinemago/internal/wire"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:   "cinemago",
		Short: "Cinemago CLI",
		Run: func(cmd *cobra.Command, args []string) {
			app, cleanup, err := wire.InitializeApp(cmd)
			if err != nil {
				fmt.Printf("Error initializing app: %v\n", err)
			}
			defer cleanup()
			server.StartServerWithGracefulShutdown(app.Config, app.App)
		},
	}
)

func init() {
	config.InitConfig(rootCmd)
}

//	@title			Cinemago
//	@version		1.0
//	@description	This is an API for Cinemago

// @host		127.0.0.1:8000
// @BasePath	/api
func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
