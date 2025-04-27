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

	// Initialize the Tulip scripting environment
	C.Tulip_Init(argc, &argv[0])

	if len(os.Args) > 1 {
		path := C.CString(os.Args[1])
		defer C.free(unsafe.Pointer(path))

		// Run Tulip script from a file
		C.Tulip_RunFile(path)
	} else {
		source := C.CString("1 + 2;")
		name := C.CString("<test>")
		defer C.free(unsafe.Pointer(source))
		defer C.free(unsafe.Pointer(name))
		var exitCode C.int

		// Interpret a Tulip script and capture the result
		result := C.Tulip_InterpretWithResult(source, name, &exitCode)
		if exitCode == 0 {
			fmt.Printf("Last value: %s\n", C.GoString(result))
		} else {
			fmt.Printf("Execution failed with code %d\n", exitCode)
		}

		// Free the result string to prevent memory leaks
		C.free(unsafe.Pointer(result))
	}

	// Clean up Tulip scripting environment resources
	C.Tulip_Free()
}
