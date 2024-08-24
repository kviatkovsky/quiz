package cmd

import (
	"github.com/kviatkovsky/quiz/internal/api"
	"github.com/spf13/cobra"
)

func StartCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Start API server",
		Run: func(cmd *cobra.Command, args []string) {
			api.Init()
		},
	}
}
