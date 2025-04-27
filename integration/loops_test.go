package integration

import (
	"testing"

	"github.com/cryptrunner49/tulipscript/internal/core"
	"github.com/cryptrunner49/tulipscript/internal/vm"
)

func TestForLoop(t *testing.T) {
	vm.InitVM([]string{"tulipscript"})
	t.Cleanup(vm.FreeVM)

	script := `
		for (let i = 0; i < 3; i = i + 1) {
			print(i)
		}
	`
	expectedOutput := "012"

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

func TestWhileLoop(t *testing.T) {
	vm.InitVM([]string{"tulipscript"})
	t.Cleanup(vm.FreeVM)

	script := `
		let count = 0;
		while (count < 3) {
			print(count)
			count = count + 1
		}
	`
	expectedOutput := "012"

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

func TestIteratorLoop(t *testing.T) {
	vm.InitVM([]string{"tulipscript"})
	t.Cleanup(vm.FreeVM)

	script := `
		iter (let item in [10, 20, 30]) {
			print(item)
		}
	`
	expectedOutput := "102030"

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
