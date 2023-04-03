package tiktoken_go

import "testing"

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
