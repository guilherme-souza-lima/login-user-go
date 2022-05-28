package main

import (
	"context"
	"github.com/spf13/cobra"
	"loginUserGo/cmd"
	"loginUserGo/infrastructure"
	"os/signal"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	defer stop()

	config := infrastructure.NewConfig()
	containerDI := infrastructure.NewContainerDI(config)
	defer containerDI.ShutDown()

	cmdHttpServer := &cobra.Command{
		Use:   "httpserver",
		Short: "Run httpserver",
		Run: func(cli *cobra.Command, args []string) {
			go containerDI.Migration.MakeMigrations()
			cmd.StartHttp(ctx, containerDI)
		},
	}

	var rootCmd = &cobra.Command{Use: "APP"}
	rootCmd.AddCommand(cmdHttpServer)
	rootCmd.Execute()
}
