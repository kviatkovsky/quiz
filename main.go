package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kviatkovsky/quiz/cmd"
	"github.com/kviatkovsky/quiz/internal/api"
	"github.com/kviatkovsky/quiz/internal/quiz"
	"github.com/spf13/cobra"
)

func main() {
	quizService := quiz.NewQuizService()
	handler := api.NewAPIHandler(quizService)

	// Set up HTTP Router using chi
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

	// Set up CLI using Cobra
	rootCmd := &cobra.Command{Use: "quiz-app"}
	rootCmd.AddCommand(cmd.QuestionsCmd(quizService))
	rootCmd.AddCommand(cmd.SubmitCmd(quizService))
	rootCmd.AddCommand(cmd.CompareCmd(quizService))

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("could not start CLI: %v\n", err)
	}
	time.Sleep(8 * time.Second)
}
