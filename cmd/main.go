package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/offerni/guruchatgpt/repository"
	"github.com/offerni/guruchatgpt/search"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	e := echo.New()

	guruRepo, err := repository.NewGuruRepository(repository.NewGuruRepoOpts{
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
	})
	if err != nil {
		e.Logger.Fatal("repository.NewGuruRepository", err)
	}

	// chatGptRepo:= repository.NewChatGptRepository()

	searchService, err := search.NewService(search.NewServiceOpts{
		GuruRepository: guruRepo,
	})
	if err != nil {
		e.Logger.Fatal("search.NewService", err)
	}

	e.GET("/search", searchService.Handler)
	e.GET("/", func(c echo.Context) error {
		fmt.Sprintln("hello world")
		return c.NoContent(http.StatusOK)
	})

	e.Logger.Fatal((e.Start(":9091")))
}
