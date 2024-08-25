package cmd

import (
	"github.com/spf13/cobra"
)

func Execute() error {

	rootCmd := &cobra.Command{Use: "quiz-app"}
	rootCmd.AddCommand(StartCmd())
	rootCmd.AddCommand(QuestionsCmd())
	rootCmd.AddCommand(SubmitCmd())
	rootCmd.AddCommand(CompareCmd())

	return rootCmd.Execute()
}
