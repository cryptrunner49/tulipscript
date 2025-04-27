//go:build cgo
// +build cgo

package main

/*
#cgo pkg-config: readline
#include <stdio.h>
#include <stdlib.h>
#include <readline/readline.h>
#include <readline/history.h>
static int self_insert_wrapper(int count, int key) { return rl_tab_insert(count, key); }
static void bind_tab_key() { rl_bind_key('\t', self_insert_wrapper); }
*/
import "C"

import (
	"fmt"
	"os"
	"strings"
	"unsafe"

	"github.com/cryptrunner49/tulipscript/internal/common"
	"github.com/cryptrunner49/tulipscript/internal/vm"
)

func main() {
	C.bind_tab_key()

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "-h", "--help":
			showUsage()
			os.Exit(0)
		case "-v", "--version":
			showVersion()
			os.Exit(0)
		}
	}

	vm.InitVM(os.Args)
	defer vm.FreeVM()

	if len(os.Args) == 1 {
		fmt.Println("tulip REPL - TulipScript Virtual Machine (type Ctrl+D to exit)")
		repl()
	} else {
		runFile(os.Args[1])
	}
}

// showUsage prints detailed help and usage instructions.
func showUsage() {
	usage := `tulip - A TulipScript Virtual Machine Interpreter

Usage: tulip [options] [script]

Options:
  -h, --help       Display this help message and exit
  -v, --version    Show version information and exit

Modes:
  - If no script is provided, tulip starts an interactive REPL (Read-Eval-Print Loop)
    where you can type commands and execute them immediately.
  - If a script file is provided, tulip executes the script and exits.

REPL Usage:
  - Type TulipScript commands at the ">>> " prompt
  - For multi-line constructs (like if-statements or loops), continue typing on the 
    next line — you'll see the "... " prompt until all opening braces '{' are matched
  - Use arrow keys to navigate command history
  - Press Ctrl+D or Ctrl+C to exit

Script Execution:
  - Provide a path to a TulipScript file to execute it
  - Example: tulip myscript.tlp

Exit Codes:
  0   Successful execution
  65  Compilation error
  70  Runtime error
  74  File I/O error
`
	fmt.Print(usage)
}

// showVersion prints version information.
func showVersion() {
	fmt.Printf("tulip version %s\n", common.Version)
}

// countBlocks returns the difference between the number of '{' and '}' characters to track block nesting depth.
func countBlocks(input string) int {
	count := 0
	for _, char := range input {
		if char == '{' {
			count++
		} else if char == '}' {
			count--
		}
	}
	return count
}

// repl runs the interactive Read-Eval-Print Loop (REPL) for TulipScript
func repl() {
	var buffer strings.Builder
	blockDepth := 0

	for {
		// Set prompt to ">>> " for new input or "... " when continuing a multi-line block
		prompt := ">>> "
		if blockDepth > 0 {
			prompt = "... "
		}
		cPrompt := C.CString(prompt)
		line := C.readline(cPrompt)
		C.free(unsafe.Pointer(cPrompt))

		if line == nil {
			fmt.Println("\nExiting REPL")
			break
		}

		input := strings.TrimSpace(C.GoString(line))
		C.free(unsafe.Pointer(line))

		if len(input) == 0 && blockDepth == 0 {
			continue
		}

		if buffer.Len() > 0 {
			buffer.WriteString("\n")
		}
		buffer.WriteString(input)

		blockDepth += countBlocks(input)
		if blockDepth < 0 {
			fmt.Fprintf(os.Stderr, "REPL error: Unmatched closing brace '}'\n")
			buffer.Reset()
			blockDepth = 0
			continue
		}

		if blockDepth == 0 {
			source := buffer.String()
			historyEntry := C.CString(source)
			C.add_history(historyEntry)
			C.free(unsafe.Pointer(historyEntry))

			result := vm.Interpret(source, "<repl>")
			switch result {
			case vm.INTERPRET_OK:
				// Successful execution
			case vm.INTERPRET_COMPILE_ERROR:
				fmt.Fprintf(os.Stderr, "Compilation error in REPL\n")
			case vm.INTERPRET_RUNTIME_ERROR:
				fmt.Fprintf(os.Stderr, "Runtime error in REPL\n")
			default:
				fmt.Fprintf(os.Stderr, "Unknown error in REPL: %d\n", result)
			}
			buffer.Reset()
		}
	}
}

// runFile reads a TulipScript source file from disk and executes it using the interpreter.
// It exits the process with an appropriate error code if an error occurs.
func runFile(path string) {
	source, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file '%s': %v\n", path, err)
		os.Exit(74)
	}
	result := vm.Interpret(string(source), path)
	switch result {
	case vm.INTERPRET_OK:
		// Successful execution
	case vm.INTERPRET_COMPILE_ERROR:
		fmt.Fprintf(os.Stderr, "Compilation error in '%s'\n", path)
		os.Exit(65)
	case vm.INTERPRET_RUNTIME_ERROR:
		fmt.Fprintf(os.Stderr, "Runtime error in '%s'\n", path)
		os.Exit(70)
	default:
		fmt.Fprintf(os.Stderr, "Unknown error: %d\n", result)
		os.Exit(1)
	}
}
