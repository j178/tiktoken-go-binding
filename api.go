package tiktoken_go

//go:generate cargo -C tiktoken-cffi build --release

/*
#cgo LDFLAGS: ${SRCDIR}/tiktoken-cffi/target/release/libtiktoken.a
#cgo darwin LDFLAGS: -framework Security -framework CoreFoundation
#cgo windows LDFLAGS: -lws2_32
#cgo linux LDFLAGS: -ldl

#include <stdlib.h>

extern unsigned int count_tokens(const char*, const char*);
extern unsigned int get_context_size(const char*);
*/
import "C"
import (
	"strings"
	"unsafe"

	"github.com/sashabaranov/go-openai"
)

func CountTokens(model, prompt string) int {
	m := C.CString(model)
	p := C.CString(prompt)
	count := C.count_tokens(m, p)
	C.free(unsafe.Pointer(m))
	C.free(unsafe.Pointer(p))
	return int(count)
}

// GetContextSize Returns the context size of a specified model.
// The context size represents the maximum number of tokens a model can process in a single input.
// This function checks the model name and returns the corresponding context size.
// See <https://platform.openai.com/docs/models> for up-to-date information.
// It returns a default value of 4096 if the model is not recognized.
func GetContextSize(model string) int {
	switch {
	case strings.HasPrefix(model, "gpt-4-32k"):
		return 32768
	case strings.HasPrefix(model, "gpt-4"):
		return 8192
	case strings.HasPrefix(model, "gpt-3.5-turbo"):
		return 4096
	case strings.HasPrefix(model, "text-davinci-002"), strings.HasPrefix(model, "text-davinci-003"):
		return 4097
	case strings.HasPrefix(model, "ada"), strings.HasPrefix(model, "babbage"), strings.HasPrefix(model, "curie"):
		return 2049
	case strings.HasPrefix(model, "code-cushman-001"):
		return 2048
	case strings.HasPrefix(model, "code-davinci-002"):
		return 8001
	case strings.HasPrefix(model, "davinci"):
		return 2049
	case strings.HasPrefix(model, "text-ada-001"), strings.HasPrefix(
		model,
		"text-babbage-001",
	), strings.HasPrefix(model, "text-curie-001"):
		return 2049
	default:
		return 4096
	}
}

// CountMessagesTokens based on https://github.com/openai/openai-cookbook/blob/main/examples/How_to_count_tokens_with_tiktoken.ipynb
func CountMessagesTokens(model string, messages []openai.ChatCompletionMessage) int {
	var tokens int
	var tokensPerMessage int
	var tokensPerName int

	switch model {
	case openai.GPT3Dot5Turbo, openai.GPT3Dot5Turbo0301:
		tokensPerMessage = 4 // every message follows <|start|>{role/name}\n{content}<|end|>\n
		tokensPerName = -1   // if there's a name, the role is omitted
	case openai.GPT4, openai.GPT40314, openai.GPT432K, openai.GPT432K0314:
		tokensPerMessage = 3
		tokensPerName = 1
	}

	for k := range messages {
		tokens += tokensPerMessage

		tokens += CountTokens(model, messages[k].Role)
		tokens += CountTokens(model, messages[k].Content)
		tokens += CountTokens(model, messages[k].Name)
		if messages[k].Name != "" {
			tokens += tokensPerName
		}
	}

	tokens += 3 // every reply is primed with <|start|>assistant<|message|>

	return tokens
}
