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
import "unsafe"

func CountTokens(model, prompt string) int {
	m := C.CString(model)
	p := C.CString(prompt)
	count := C.count_tokens(m, p)
	C.free(unsafe.Pointer(m))
	C.free(unsafe.Pointer(p))
	return int(count)
}

func GetContextSize(model string) int {
	m := C.CString(model)
	count := C.get_context_size(m)
	C.free(unsafe.Pointer(m))
	return int(count)
}
