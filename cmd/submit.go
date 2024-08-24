package cmd

import (
	"fmt"

	"github.com/kviatkovsky/quiz/internal/quiz"

	"github.com/spf13/cobra"
)

func SubmitCmd(quizService *quiz.QuizService) *cobra.Command {
	return &cobra.Command{
		Use:   "submit",
		Short: "Submit your quiz answers",
		Run: func(cmd *cobra.Command, args []string) {
			// For simplicity, hardcoding some answers here
			userAnswers := quiz.UserAnswers{
				Answers: []int{1, 2, 1},
			}
			result := quizService.SubmitAnswers(&userAnswers)
			fmt.Printf("You got %d out of %d correct!\n", result.CorrectAnswers, result.TotalQuestions)
		},
	}
}
