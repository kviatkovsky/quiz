package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/kviatkovsky/quiz/internal/api"
	"github.com/kviatkovsky/quiz/internal/quiz"
	"github.com/spf13/cobra"
)

func QuestionsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "questions",
		Short: "Get quiz questions",
		Run: func(cmd *cobra.Command, args []string) {
			questions := getQuestions()
			for i, q := range questions {
				fmt.Printf("Question %d: %s\n", i+1, q.Question)
				for j, a := range q.Answers {
					fmt.Printf("  %d: %s\n", j+1, a)
				}
			}
		},
	}
}

func getQuestions() []quiz.Question {
	var questions []quiz.Question
	url := fmt.Sprintf("%s%s", api.BaseUrl, api.GetQuestionsEndpoint)
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return questions
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return questions
	}

	err = json.Unmarshal(body, &questions)
	if err != nil {
		fmt.Printf("Error unmarshaling JSON: %v\n", err)
		return questions
	}

	return questions
}
