package integration

import (
	"testing"

	"github.com/cryptrunner49/tulipscript/internal/core"
	"github.com/cryptrunner49/tulipscript/internal/vm"
)

func TestArrayCreationAndAccess(t *testing.T) {
	vm.InitVM([]string{"tulipscript"})
	t.Cleanup(vm.FreeVM)

	script := `
		let arr = [1, 2, 3]
		print(arr[0])
		print(arr[1])
		print(arr[2])
	`
	expectedOutput := "123"

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

func TestArrayPushAndPop(t *testing.T) {
	vm.InitVM([]string{"tulipscript"})
	t.Cleanup(vm.FreeVM)

	script := `
		let arr = [1, 2, 3]
		push(arr, 4)
		print(arr[3])
		print(pop(arr))
		print(len(arr))
	`
	expectedOutput := "443"

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

func TestArrayLength(t *testing.T) {
	vm.InitVM([]string{"tulipscript"})
	t.Cleanup(vm.FreeVM)

	script := `
		let arr = [1, 2, 3, 4, 5]
		print(len(arr))
	`
	expectedOutput := "5"

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

func TestArraySorting(t *testing.T) {
	vm.InitVM([]string{"tulipscript"})
	t.Cleanup(vm.FreeVM)

	script := `
		let arr = [5, 3, 8, 1, 42, 10]
		array_sort(arr)
		print(arr)
	`
	expectedOutput := "[1, 3, 5, 8, 10, 42]"

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

func TestArraySplit(t *testing.T) {
	vm.InitVM([]string{"tulipscript"})
	t.Cleanup(vm.FreeVM)

	script := `
		let arr = [1, 2, "sep", 3, 4, "sep", 5, 6]
		let split = array_split(arr, "sep")
		print(array_to_string(split[0]))
		print(array_to_string(split[1]))
		print(array_to_string(split[2]))
	`
	expectedOutput := "[1, 2][3, 4][5, 6]"

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

func TestArrayJoin(t *testing.T) {
	vm.InitVM([]string{"tulipscript"})
	t.Cleanup(vm.FreeVM)

	script := `
		let a1 = [1, 2]
		let a2 = [3, 4]
		let a3 = [5, 6]
		let joined = array_join(a1, a2, a3)
		print(array_to_string(joined))
	`
	expectedOutput := "[1, 2, 3, 4, 5, 6]"

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

func TestArraySearch(t *testing.T) {
	vm.InitVM([]string{"tulipscript"})
	t.Cleanup(vm.FreeVM)

	script := `
		let arr = ["cat", "dog", "bird", "dog"]
		let linear = array_linear_search(arr, "dog")
		let sorted = ["apple", "banana", "cherry", "date"]
		let binary = array_binary_search(sorted, "cherry")
		print(linear)
		print(binary)
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

func TestArraySlices(t *testing.T) {
	vm.InitVM([]string{"tulipscript"})
	t.Cleanup(vm.FreeVM)

	script := `
		let arr = [1, 2, 3, 4, 5]
		let slice1 = arr[1:3]
		let slice2 = arr[2:]
		print(array_to_string(slice1))
		print(array_to_string(slice2))
	`
	expectedOutput := "[2, 3][3, 4, 5]"

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
