package tiktoken_go

/*
#cgo LDFLAGS: ./tiktoken-cffi/libtitoken.a -ldl -framework Security -framework CoreFoundation
#include <stdlib.h>

extern unsigned int count_tokens(const char*, const char*);
extern unsigned int get_context_size(const char*);
*/
import "C"
