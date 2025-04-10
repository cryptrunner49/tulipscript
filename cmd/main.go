//go:build cgo
// +build cgo

package main

/*
#cgo pkg-config: readline
#include <stdio.h>
#include <stdlib.h>
#include <readline/readline.h>
#include <readline/history.h>
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
  - Type TulipScript commands at the "> " prompt
  - For multi-line constructs (like if-statements or loops), continue typing on the 
  next line — you'll see the "... " prompt until all opening braces '{' are matched
  - Use arrow keys to navigate command history
  - Press Ctrl+D or Ctrl+C to exit

Script Execution:
  - Provide a path to a TulipScript script file to execute it
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

func main() {
	// Handle command-line arguments
	if len(os.Args) > 1 {
		switch arg := os.Args[1]; arg {
		case "-h", "--help":
			showUsage()
			os.Exit(0)
		case "-v", "--version":
			showVersion()
			os.Exit(0)
		}
	}

	// Initialize VM with arguments
	vm.InitVM(os.Args)
	defer vm.FreeVM()

	// Determine mode: REPL or script execution
	if len(os.Args) == 1 {
		fmt.Println("tulip REPL - TulipScript Virtual Machine (type Ctrl+D to exit)")
		repl()
	} else {
		runFile(os.Args[1])
	}
}

// countBlocks counts the net number of open blocks (unmatched '{' minus '}')
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

func repl() {
	var buffer strings.Builder // Accumulate multi-line input
	blockDepth := 0            // Track open blocks

	for {
		// Adjust prompt based on whether we're in a block
		prompt := "> "
		if blockDepth > 0 {
			prompt = strings.Repeat("  ", blockDepth) + "... "
		}
		cPrompt := C.CString(prompt)
		line := C.readline(cPrompt)
		C.free(unsafe.Pointer(cPrompt))

		if line == nil { // EOF (Ctrl+D)
			fmt.Println("\nExiting REPL")
			break
		}

		input := strings.TrimSpace(C.GoString(line))
		C.free(unsafe.Pointer(line))

		if len(input) == 0 && blockDepth == 0 {
			continue // Skip empty lines unless in a block
		}

		// Add input to buffer with a newline
		if buffer.Len() > 0 {
			buffer.WriteString("\n")
		}
		buffer.WriteString(input)

		// Update block depth
		blockDepth += countBlocks(input)

		if blockDepth < 0 {
			fmt.Fprintf(os.Stderr, "REPL error: Unmatched closing brace '}'\n")
			buffer.Reset()
			blockDepth = 0
			continue
		}

		// If all blocks are closed, interpret the accumulated input
		if blockDepth == 0 {
			source := buffer.String()

			// Add to history only when block is complete
			historyEntry := C.CString(source)
			C.add_history(historyEntry)
			C.free(unsafe.Pointer(historyEntry))

			result := vm.Interpret(source, "<repl>")
			switch result {
			case vm.INTERPRET_OK:
				// Successful execution, no output needed
			case vm.INTERPRET_COMPILE_ERROR:
				fmt.Fprintf(os.Stderr, "Compilation error in REPL\n")
			case vm.INTERPRET_RUNTIME_ERROR:
				fmt.Fprintf(os.Stderr, "Runtime error in REPL\n")
			default:
				fmt.Fprintf(os.Stderr, "Unknown error in REPL: %v\n", result)
			}

			// Reset buffer after interpretation
			buffer.Reset()
		}
	}
}

func runFile(path string) {
	source, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file '%s': %v\n", path, err)
		os.Exit(74)
	}

	result := vm.Interpret(string(source), path)
	switch result {
	case vm.INTERPRET_OK:
		// Successful execution, exit silently
	case vm.INTERPRET_COMPILE_ERROR:
		fmt.Fprintf(os.Stderr, "Compilation error in '%s'\n", path)
		os.Exit(65)
	case vm.INTERPRET_RUNTIME_ERROR:
		fmt.Fprintf(os.Stderr, "Runtime error in '%s'\n", path)
		os.Exit(70)
	default:
		fmt.Fprintf(os.Stderr, "Unknown error: %v\n", result)
		os.Exit(1)
	}
}
