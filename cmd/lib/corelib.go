package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"os"
	"strings"
	"unsafe"

	"github.com/cryptrunner49/tulipscript/internal/runtime"
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

// Tulip_InterpretWithResult interprets TulipScript source code and returns the last value as a string.
//
//export Tulip_InterpretWithResult
func Tulip_InterpretWithResult(csrc, cname *C.char, exitCode *C.int) *C.char {
	src := C.GoString(csrc)
	name := C.GoString(cname)
	code := vm.Interpret(src, name)
	*exitCode = C.int(code)
	return valueToCString(vm.GetLastValue())
}

// Tulip_RunFileWithResult runs a TulipScript script from a file and returns the last value as a string.
//
//export Tulip_RunFileWithResult
func Tulip_RunFileWithResult(cpath *C.char, exitCode *C.int) *C.char {
	path := C.GoString(cpath)
	source, err := os.ReadFile(path)
	if err != nil {
		*exitCode = C.int(74) // File I/O error
		return valueToCString(runtime.Value{Type: runtime.VAL_NULL})
	}
	code := vm.Interpret(string(source), path)
	*exitCode = C.int(code)
	return valueToCString(vm.GetLastValue())
}

// valueToString converts a runtime.Value into a human-readable string, similar to how runtime.PrintObject displays values.
func valueToString(val runtime.Value) string {
	switch val.Type {
	case runtime.VAL_NULL:
		return "null"
	case runtime.VAL_BOOL:
		if val.Bool {
			return "true"
		}
		return "false"
	case runtime.VAL_NUMBER:
		return fmt.Sprintf("%g", val.Number)
	case runtime.VAL_OBJ:
		switch obj := val.Obj.(type) {
		case *runtime.ObjString:
			return obj.Chars
		case *runtime.ObjArray:
			elements := make([]string, len(obj.Elements))
			for i, elem := range obj.Elements {
				elements[i] = valueToString(elem)
			}
			return "[" + strings.Join(elements, ", ") + "]"
		case *runtime.ObjMap:
			entries := make([]string, 0, len(obj.Entries))
			for key, value := range obj.Entries {
				entries = append(entries, fmt.Sprintf("%s: %s", key.Chars, valueToString(value)))
			}
			return "{" + strings.Join(entries, ", ") + "}"
		case *runtime.ObjStruct:
			return fmt.Sprintf("<struct %s>", obj.Name.Chars)
		case *runtime.ObjInstance:
			var sb strings.Builder
			sb.WriteString("<(struct ")
			sb.WriteString(obj.Structure.Name.Chars)
			sb.WriteString(")")
			first := true
			for fieldName, fieldValue := range obj.Fields {
				if !first {
					sb.WriteString(",")
				}
				fmt.Fprintf(&sb, " %s=%s", fieldName.Chars, valueToString(fieldValue))
				first = false
			}
			sb.WriteString(">")
			return sb.String()
		case *runtime.ObjFunction:
			if obj.Name == nil {
				return "<script>"
			}
			return fmt.Sprintf("<fn %s>", obj.Name.Chars)
		case *runtime.ObjClosure:
			if obj.Function.Name == nil {
				return "<script>"
			}
			return fmt.Sprintf("<fn %s>", obj.Function.Name.Chars)
		case *runtime.ObjNative:
			return "<native fn>"
		case *runtime.ObjModule:
			return fmt.Sprintf("<mod %s>", obj.Name.Chars)
		case *runtime.ObjDate:
			return fmt.Sprintf("<Date %s>", obj.Time.Format("2006-01-02"))
		case *runtime.ObjTime:
			return fmt.Sprintf("<Time %s>", obj.Time.Format("15:04:05"))
		case *runtime.ObjDateTime:
			return fmt.Sprintf("<DateTime %s>", obj.Time.Format("2006-01-02 15:04:05"))
		case *runtime.ObjArrayIterator:
			return fmt.Sprintf("<array iterator at %d>", obj.Index)
		default:
			return "<unknown object>"
		}
	default:
		return "<unknown>"
	}
}

// valueToCString converts a runtime.Value to a C string. The caller is responsible for freeing it.
func valueToCString(val runtime.Value) *C.char {
	return C.CString(valueToString(val))
}

// Tulip_Free frees the TulipScript VM resources.
//
//export Tulip_Free
func Tulip_Free() {
	vm.FreeVM()
}

func main() {}
