package vm

import (
	"github.com/cryptrunner49/tulipscript/internal/runtime"
)

// defineArgs creates a global variable "args" as an array of strings.
// Each command-line argument is converted into an ObjString and stored in the array.
func defineArgs(args []string) {
	elements := make([]runtime.Value, len(args))
	for i, arg := range args {
		elements[i] = runtime.ObjVal(runtime.NewObjString(arg))
	}
	argsName := runtime.NewObjString("args")
	vm.globals[argsName] = GlobalVar{Value: runtime.ObjVal(runtime.NewArray(elements)), IsConst: false}
}
