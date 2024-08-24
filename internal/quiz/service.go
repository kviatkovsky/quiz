package quiz

import (
	"strconv"
	"sync"
)

type QuizService struct {
	questions   []Question
	results     []int
	resultsLock sync.Mutex
}

func NewQuizService() *QuizService {
	return &QuizService{
		questions: []Question{
			{
				Question: "What is the capital of France?",
				Answers:  []string{"Berlin", "Madrid", "Paris", "Lisbon"},
				Correct:  2,
			},
			{
				Question: "What is 2 + 2?",
				Answers:  []string{"3", "4", "5", "6"},
				Correct:  1,
			},
			{
				Question: "What is the largest planet in our solar system?",
				Answers:  []string{"Earth", "Mars", "Jupiter", "Saturn"},
				Correct:  2,
			},
		},
		results: []int{},
	}
}

func (s *QuizService) GetQuestions() []Question {
	return s.questions
}

func (s *QuizService) SubmitAnswers(answers *UserAnswers) *Result {
	correctCount := 0
	for i, answer := range answers.Answers {
		if s.questions[i].Correct == answer-1 {
			correctCount++
		}
	}

	s.resultsLock.Lock()
	s.results = append(s.results, correctCount)
	s.resultsLock.Unlock()

	return &Result{
		TotalQuestions: len(s.questions),
		CorrectAnswers: correctCount,
	}
}

func (s *QuizService) GetComparisonResult() string {
	s.resultsLock.Lock()
	defer s.resultsLock.Unlock()

	if len(s.results) == 0 {
		return "No results found"
	}

	userScore := s.results[len(s.results)-1]
	higherCount := 0

	for _, result := range s.results {
		if result < userScore {
			higherCount++
		}
	}

	percentage := (float64(higherCount) / float64(len(s.results))) * 100
	return "You were better than " + strconv.FormatFloat(percentage, 'f', 2, 64) + "% of all quizzers"
}
