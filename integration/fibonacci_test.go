package integration

import (
	"testing"

	"github.com/cryptrunner49/tulipscript/internal/core"
	"github.com/cryptrunner49/tulipscript/internal/vm"
)

func TestFibonacciRecursive(t *testing.T) {
	vm.InitVM([]string{"tulipscript"})
	t.Cleanup(vm.FreeVM)

	script := `
		function fib(n) {
			if (n < 2) return n
			return fib(n - 2) + fib(n - 1)
		}
		print(fib(16))
	`
	expectedOutput := "987"

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

func TestFibonacciIterative(t *testing.T) {
	vm.InitVM([]string{"tulipscript"})
	t.Cleanup(vm.FreeVM)

	script := `
		function fib(n) {
			if (n < 2) return n;
			let a = 0
			let b = 1
			for (let i = 2; i <= n; i = i + 1) {
				let temp = a + b
				a = b
				b = temp
			}
			return b
		}
		print(fib(16))
	`
	expectedOutput := "987"

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
