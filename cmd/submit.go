package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/kviatkovsky/quiz/internal/api"
	"github.com/kviatkovsky/quiz/internal/quiz"
	"github.com/spf13/cobra"
)

func SubmitCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "submit",
		Short: "Submit your quiz answers interactively",
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

			result := submitAnswers(&answers)
			fmt.Printf("You got %d out of %d correct!\n", result.CorrectAnswers, result.TotalQuestions)
		},
	}
}

func submitAnswers(answers *quiz.UserAnswers) quiz.Result {
	var submitResult quiz.Result

	url := fmt.Sprintf("%s%s", api.BaseUrl, api.PostSubmitEndpoint)

	jsonData, err := json.Marshal(answers)
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		return submitResult
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return submitResult
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		return submitResult
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return submitResult
	}

	err = json.Unmarshal(body, &submitResult)
	if err != nil {
		fmt.Printf("Error unmarshaling JSON: %v\n", err)
		return submitResult
	}

	return submitResult
}
