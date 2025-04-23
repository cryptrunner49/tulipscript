//go:build cgo
// +build cgo

package main

// #cgo CFLAGS: -I${SRCDIR}/../../../bin
// #cgo LDFLAGS: -L${SRCDIR}/../../../bin -ltulip
// #include <stdlib.h>
// #include "libtulip.h"
import "C"
import (
	"fmt"
	"os"
	"unsafe"
)

func main() {
	// Convert Go command-line arguments to C-style argc/argv
	argc := C.int(len(os.Args))
	argv := make([]*C.char, argc)
	for i, arg := range os.Args {
		argv[i] = C.CString(arg)
		defer C.free(unsafe.Pointer(argv[i]))
	}

	// Initialize TulipScript VM
	C.Tulip_Init(argc, &argv[0])

	// Run TulipScript
	var result C.int
	if len(os.Args) > 1 {
		// Run script from file
		path := C.CString(os.Args[1])
		defer C.free(unsafe.Pointer(path))
		result = C.Tulip_RunFile(path)
	} else {
		// Run inline script
		source := C.CString("println(\"Hello from TulipScript!\");")
		name := C.CString("<test>")
		defer C.free(unsafe.Pointer(source))
		defer C.free(unsafe.Pointer(name))
		result = C.Tulip_Interpret(source, name)
	}

	// Print result
	fmt.Printf("Result: %d\n", result)

	// Free TulipScript VM
	C.Tulip_Free()
}
