package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kviatkovsky/quiz/internal/quiz"
)

const (
	BaseUrl              = "http://localhost:8080/"
	GetCompareEndpoint   = "compare"
	PostSubmitEndpoint   = "submit"
	GetQuestionsEndpoint = "questions"
)

func Init() {
	quizService := quiz.NewQuizService()

	startServer(quizService)
}

func startServer(quizService *quiz.QuizService) {
	handler := NewAPIHandler(quizService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/questions", handler.GetQuestions)
	r.Post("/submit", handler.SubmitAnswers)
	r.Get("/compare", handler.GetComparisonResult)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("could not start server: %v\n", err)
	}
}
