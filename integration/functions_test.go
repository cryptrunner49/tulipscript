package integration

import (
	"testing"

	"github.com/cryptrunner49/tulipscript/internal/core"
	"github.com/cryptrunner49/tulipscript/internal/vm"
)

func TestFunctionCall(t *testing.T) {
	vm.InitVM([]string{"tulipscript"})
	t.Cleanup(vm.FreeVM)

	script := `
		function add(a, b) { return a + b }
		print(add(5, 3))
	`
	expectedOutput := "8"

	output := captureOutput(t, func() {
		result := core.Interpret(script, "<script>")
		if result != 0 {
			t.Fatalf("Interpretation failed: %d", result)
		}
	})

	if output != expectedOutput {
		t.Errorf("Expected %q, got %q", expectedOutput, output)
	}
}

func TestFunctionWithString(t *testing.T) {
	vm.InitVM([]string{"tulipscript"})
	t.Cleanup(vm.FreeVM)

	script := `
		function greet(name) {
			return "Hello, " + name + "!"
		}
		print(greet("World"))
	`
	expectedOutput := "Hello, World!"

	output := captureOutput(t, func() {
		result := core.Interpret(script, "<script>")
		if result != 0 {
			t.Fatalf("Interpretation failed: %d", result)
		}
	})

	if output != expectedOutput {
		t.Errorf("Expected %q, got %q", expectedOutput, output)
	}
}
