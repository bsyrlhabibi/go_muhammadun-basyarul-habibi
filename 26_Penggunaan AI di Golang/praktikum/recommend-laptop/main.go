package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

var (
	debug = false
)

type RecommendationRequest struct {
	Budget  int    `json:"budget"`
	Purpose string `json:"purpose"`
}

type RecommendationResponse struct {
	Recommendations []openai.ChatCompletionMessage `json:"recommendations"`
}

func getCompletionFromMessages(
	ctx context.Context,
	client *openai.Client,
	messages []openai.ChatCompletionMessage,
	model string,
) (openai.ChatCompletionResponse, error) {
	if model == "" {
		model = openai.GPT3Dot5Turbo
	}

	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:    model,
			Messages: messages,
		},
	)
	return resp, err
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Cannot load env file. Err: %s", err)
	}

	ctx := context.Background()
	client := openai.NewClient(os.Getenv("KEY"))

	http.HandleFunc("/recommends", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		decoder := json.NewDecoder(r.Body)
		var request RecommendationRequest
		if err := decoder.Decode(&request); err != nil {
			http.Error(w, "Invalid request data", http.StatusBadRequest)
			return
		}

		budget := fmt.Sprintf("%d", request.Budget)

		messages := []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: fmt.Sprintf("Hi, my name is Habibi, and I want a %s laptop under Rp. %s", request.Purpose, budget),
			},
		}
		model := openai.GPT3Dot5Turbo

		resp, err := getCompletionFromMessages(ctx, client, messages, model)
		if err != nil {
			http.Error(w, "Error when interacting with OpenAI", http.StatusInternalServerError)
			return
		}

		answer := openai.ChatCompletionMessage{
			Role:    resp.Choices[0].Message.Role,
			Content: resp.Choices[0].Message.Content,
		}
		messages = append(messages, answer)

		recommendations := RecommendationResponse{
			Recommendations: messages,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(recommendations)
	})

	fmt.Println("Starting recommendation laptop...")
	http.ListenAndServe(":8080", nil)
}
