package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/offerni/guruchatgpt"
	"github.com/offerni/guruchatgpt/repository"
	"github.com/offerni/guruchatgpt/search"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	e := echo.New()

	guruRepo, err := repository.NewGuruRepository(repository.NewGuruRepoOpts{
		Username: os.Getenv("GURU_USERNAME"),
		Password: os.Getenv("GURU_PASSWORD"),
	})
	if err != nil {
		e.Logger.Fatal("repository.NewGuruRepository", err)
	}

	openAIRepo, err := repository.NewOpenAIRepository(repository.NewOpenAIRepoOpts{
		Credentials: repository.Credentials{
			BearerToken: os.Getenv("OPEN_AI_USER_TOKEN"),
		},
		Chat: repository.Chat{CompletionRequestOpts: &repository.OpenAIChatCompletionRequest{
			Model: os.Getenv("CHAT_GPT_MODEL"),
		}},
	})

	searchService, err := search.NewService(search.NewServiceOpts{
		GuruRepository:   guruRepo,
		OpenAiRepository: openAIRepo,
		Credentials:      guruchatgpt.Credentials(openAIRepo.Credentials),
	})
	if err != nil {
		e.Logger.Fatal("search.NewService", err)
	}

	e.GET("/search", searchService.Handler)
	e.GET("/", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	port := os.Getenv("PORT")
	defaultPort := "8080"
	if port == "" {
		port = defaultPort
	}

	e.Logger.Fatal((e.Start(fmt.Sprintf(":%s", port))))
}
