package cmd

import (
	"fmt"
	"strconv"

	"github.com/kviatkovsky/quiz/internal/quiz"

	"github.com/spf13/cobra"
)

func SubmitAllCmd(quizService *quiz.QuizService) *cobra.Command {
	return &cobra.Command{
		Use:   "submitAll",
		Short: "Submit all your quiz answers",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("Please provide your answers as command-line arguments.")
				return
			}

			var userAnswers []int
			for _, arg := range args {
				answer, err := strconv.Atoi(arg)
				if err != nil {
					fmt.Printf("Invalid answer: %s. Please provide numeric answers.\n", arg)
					return
				}
				userAnswers = append(userAnswers, answer)
			}

			answers := quiz.UserAnswers{
				Answers: userAnswers,
			}

			result := quizService.SubmitAnswers(&answers)
			fmt.Printf("You got %d out of %d correct!\n", result.CorrectAnswers, result.TotalQuestions)
		},
	}
}
