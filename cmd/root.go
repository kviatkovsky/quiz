package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// RootCmd is the base command for the CLI
var RootCmd = &cobra.Command{
	Use:   "quiz-app",
	Short: "Quiz application CLI",
	Long:  `This CLI allows you to interact with the Quiz application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to the Quiz CLI!")
	},
}

func Execute() error {
	return RootCmd.Execute()
}
