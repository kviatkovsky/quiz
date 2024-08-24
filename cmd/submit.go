package cmd

import (
	"fmt"

	"github.com/kviatkovsky/quiz/internal/quiz"

	"github.com/spf13/cobra"
)

func SubmitCmd(quizService *quiz.QuizService) *cobra.Command {
	return &cobra.Command{
		Use:   "submit",
		Short: "Submit your quiz answers interactively",
		Run: func(cmd *cobra.Command, args []string) {
			var userAnswers []int
			questions := quizService.GetQuestions()
			for i, q := range questions {
				var answer int
				fmt.Printf("question-%d %s: \n", i+1, q.Question)
				for key, value := range q.Answers {
					fmt.Printf("%v: %s\n", key+1, value)
				}
				for {
					if isValidAnswer(&answer) {
						break
					}
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

func isValidAnswer(answer *int) bool {
	_, err := fmt.Scan(answer)
	if err != nil {
		fmt.Printf("Invalid input, please enter a number.\n")
		return false
	}

	if *answer < 1 || *answer > 4 {
		fmt.Printf("Invalid answer, please enter a number between 1 and 4\n")
		return false
	}

	return true
}
