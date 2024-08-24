package quiz

type Question struct {
	Question string   `json:"question"`
	Answers  []string `json:"answers"`
	Correct  int      `json:"-"`
}

// Result represents the result of the quiz
type Result struct {
	TotalQuestions int `json:"total_questions"`
	CorrectAnswers int `json:"correct_answers"`
}

type UserAnswers struct {
	Answers []int `json:"answers"`
}
