package tiktoken_go

import (
	"context"
	"os"
	"testing"

	"github.com/sashabaranov/go-openai"
)

// https://platform.openai.com/tokenizer

func TestCountTokens(t *testing.T) {
	count := CountTokens("gpt-3.5-turbo", "hello world")
	if count != 2 {
		t.Errorf("GetCompletionMaxTokens() = %v, want %v", count, 2)
	}
}

func TestGetContextSize(t *testing.T) {
	count := GetContextSize("gpt-3.5-turbo")
	if count != 4096 {
		t.Errorf("GetContextSize() = %v, want %v", count, 4096)
	}
}

func BenchmarkCountTokens(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountTokens("gpt-3.5-turbo", "hello world")
	}
}

func TestCountMessagesTokens(t *testing.T) {
	if os.Getenv("OPENAI_API_KEY") == "" {
		t.Skip("OPENAI_API_KEY is not set")
	}
	c := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	messages := []openai.ChatCompletionMessage{
		{
			Role:    "system",
			Content: "You are a helpful, pattern-following assistant that translates corporate jargon into plain English.",
		},
		{
			Role:    "system",
			Name:    "example_user",
			Content: "New synergies will help drive top-line growth.",
		},
		{
			Role:    "system",
			Name:    "example_assistant",
			Content: "Things working well together will increase revenue.",
		},
		{
			Role:    "system",
			Name:    "example_user",
			Content: "Let's circle back when we have more bandwidth to touch base on opportunities for increased leverage.",
		},
		{
			Role:    "system",
			Name:    "example_assistant",
			Content: "Let's talk later when we're less busy about how to do better.",
		},
		{
			Role:    "user",
			Content: "This late pivot means we don't have time to boil the ocean for the client deliverable.",
		},
	}

	var testcases = []struct {
		Model string
		Count int
	}{
		{"gpt-3.5-turbo-0301", 127},
		{"gpt-4-0314", 129},
	}

	for _, tc := range testcases {
		t.Run(
			tc.Model, func(t *testing.T) {
				req := openai.ChatCompletionRequest{
					Model:     tc.Model,
					Messages:  messages,
					MaxTokens: 1,
				}
				count := CountMessagesTokens(tc.Model, messages)
				resp, err := c.CreateChatCompletion(context.Background(), req)
				if err != nil {
					t.Error(err)
				}
				if want := resp.Usage.PromptTokens; count != want {
					t.Errorf("CountMessagesTokens() = %v, want %v", count, want)
				}
			},
		)
	}
}
