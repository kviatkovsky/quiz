package cmd

import (
	"fmt"

	"github.com/kviatkovsky/quiz/internal/quiz"

	"github.com/spf13/cobra"
)

func QuestionsCmd(quizService *quiz.QuizService) *cobra.Command {
	return &cobra.Command{
		Use:   "questions",
		Short: "Get quiz questions",
		Run: func(cmd *cobra.Command, args []string) {
			questions := quizService.GetQuestions()
			for i, q := range questions {
				fmt.Printf("Question %d: %s\n", i+1, q.Question)
				for j, a := range q.Answers {
					fmt.Printf("  %d: %s\n", j+1, a)
				}
			}
		},
	}
}
