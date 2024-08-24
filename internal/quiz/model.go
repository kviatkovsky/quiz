package quiz

// Question represents a quiz question
type Question struct {
	Question string   `json:"question"`
	Answers  []string `json:"answers"`
	Correct  int      `json:"-"` // This should not be exposed via API
}

// Result represents the result of the quiz
type Result struct {
	TotalQuestions int `json:"total_questions"`
	CorrectAnswers int `json:"correct_answers"`
}

type UserAnswers struct {
	Answers []int `json:"answers"`
}
