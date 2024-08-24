package cmd

import (
	"fmt"

	"github.com/kviatkovsky/quiz/internal/quiz"

	"github.com/spf13/cobra"
)

func CompareCmd(quizService *quiz.QuizService) *cobra.Command {
	return &cobra.Command{
		Use:   "compare",
		Short: "Compare your score with others",
		Run: func(cmd *cobra.Command, args []string) {
			comparison := quizService.GetComparisonResult()
			fmt.Printf("Comparison: %s\n", comparison)
		},
	}
}
