package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"os"
	"unsafe"

	"github.com/cryptrunner49/tulipscript/internal/vm"
)

// Tulip_Init initializes the TulipScript VM with command-line arguments.
//
//export Tulip_Init
func Tulip_Init(argc C.int, argv **C.char) {
	n := int(argc)
	var args []string
	if n > 0 && argv != nil {
		args = make([]string, n)
		slice := (*[1 << 28]*C.char)(unsafe.Pointer(argv))[:n:n]
		for i, s := range slice {
			args[i] = C.GoString(s)
		}
	} else {
		args = []string{}
	}
	vm.InitVM(args)
}

// Tulip_Interpret interprets TulipScript source code with a given name.
//
//export Tulip_Interpret
func Tulip_Interpret(csrc, cname *C.char) C.int {
	src := C.GoString(csrc)
	name := C.GoString(cname)
	return C.int(vm.Interpret(src, name))
}

// Tulip_RunFile runs a TulipScript script from a file path.
//
//export Tulip_RunFile
func Tulip_RunFile(cpath *C.char) C.int {
	path := C.GoString(cpath)
	source, err := os.ReadFile(path)
	if err != nil {
		return C.int(74) // File I/O error
	}
	return C.int(vm.Interpret(string(source), path))
}

// Tulip_Free frees the TulipScript VM resources.
//
//export Tulip_Free
func Tulip_Free() {
	vm.FreeVM()
}

func main() {}
