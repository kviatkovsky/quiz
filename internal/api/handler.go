package api

import (
	"encoding/json"
	"net/http"

	"github.com/kviatkovsky/quiz/internal/quiz"
)

type APIHandler struct {
	quizService *quiz.QuizService
}

func NewAPIHandler(quizService *quiz.QuizService) *APIHandler {
	return &APIHandler{quizService: quizService}
}

func (h *APIHandler) GetQuestions(w http.ResponseWriter, r *http.Request) {
	questions := h.quizService.GetQuestions()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(questions)
}

func (h *APIHandler) SubmitAnswers(w http.ResponseWriter, r *http.Request) {
	var answers quiz.UserAnswers
	if err := json.NewDecoder(r.Body).Decode(&answers); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result := h.quizService.SubmitAnswers(&answers)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (h *APIHandler) GetComparisonResult(w http.ResponseWriter, r *http.Request) {
	comparison := h.quizService.GetComparisonResult()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"comparison": comparison})
}
