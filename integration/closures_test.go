package integration

import (
	"testing"

	"github.com/cryptrunner49/tulipscript/internal/core"
	"github.com/cryptrunner49/tulipscript/internal/vm"
)

func TestClosureBasic(t *testing.T) {
	vm.InitVM([]string{"tulipscript"})
	t.Cleanup(vm.FreeVM)

	script := `
		function outer() {
			let a = 1;
			let b = 2;
			function middle() {
				let c = 3;
				let d = 4;
				function inner() {
					print(a + c + b + d);
				}
				inner();
			}
			middle();
		}
		outer();
	`
	expectedOutput := "10"

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

func TestClosureCounter(t *testing.T) {
	vm.InitVM([]string{"tulipscript"})
	t.Cleanup(vm.FreeVM)

	script := `
		function makeCounter() {
			let value = 0
			function increment() {
				value = value + 1
				return value
			}
			return increment
		}
		let counter = makeCounter()
		print(counter())
		print(counter())
	`
	expectedOutput := "12"

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
