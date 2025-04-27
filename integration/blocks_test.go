package integration

import (
	"testing"

	"github.com/cryptrunner49/tulipscript/internal/core"
	"github.com/cryptrunner49/tulipscript/internal/vm"
)

func TestBlockScope(t *testing.T) {
	vm.InitVM([]string{"tulipscript"})
	t.Cleanup(vm.FreeVM)

	script := `
		let x = 10
		{
			let x = 20
			print(x)
		}
		print(x)
	`
	expectedOutput := "2010"

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
