package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/kviatkovsky/quiz/internal/api"
	"github.com/spf13/cobra"
)

func CompareCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "compare",
		Short: "Compare your score with others",
		Run: func(cmd *cobra.Command, args []string) {
			comparison := getComparisonResult()
			fmt.Printf(comparison)
		},
	}
}

func getComparisonResult() string {
	result := map[string]string{}
	url := fmt.Sprintf("%s%s", api.BaseUrl, api.GetCompareEndpoint)
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return ""
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return ""
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Printf("Error unmarshaling JSON: %v\n", err)
		return ""
	}

	return result["comparison"]
}
