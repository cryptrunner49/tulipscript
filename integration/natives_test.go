package integration

import (
	"os"
	"strconv"
	"testing"

	"github.com/cryptrunner49/tulipscript/internal/core"
	"github.com/cryptrunner49/tulipscript/internal/vm"
)

func TestPrint(t *testing.T) {
	vm.InitVM([]string{"tulipscript"})
	t.Cleanup(vm.FreeVM)

	script := `
		print("Hello, world!")
	`
	expectedOutput := "Hello, world!"

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

func TestClock(t *testing.T) {
	vm.InitVM([]string{"tulipscript"})
	t.Cleanup(vm.FreeVM)

	script := `
		let time = clock()
		print(time)
	`

	output := captureOutput(t, func() {
		result := core.Interpret(script, "<script>")
		if result != 0 {
			t.Fatalf("Interpretation failed: %d", result)
		}
	})

	if _, err := strconv.ParseFloat(output, 64); err != nil {
		t.Errorf("Expected a float, got %q", output)
	}
}

func TestRandomBetween(t *testing.T) {
	vm.InitVM([]string{"tulipscript"})
	t.Cleanup(vm.FreeVM)

	script := `
		let rand = random_between(1, 10)
		print(rand)
	`

	output := captureOutput(t, func() {
		result := core.Interpret(script, "<script>")
		if result != 0 {
			t.Fatalf("Interpretation failed: %d", result)
		}
	})

	num, err := strconv.Atoi(output)
	if err != nil {
		t.Errorf("Expected an integer, got %q", output)
	}
	if num < 1 || num > 10 {
		t.Errorf("Expected number between 1 and 10, got %d", num)
	}
}

func TestFileOperations(t *testing.T) {
	vm.InitVM([]string{"tulipscript"})
	t.Cleanup(vm.FreeVM)

	script := `
		let filename = "test.txt"
		let content = "Hello, World!"
		write_file(filename, content)
		let readContent = read_file(filename)
		print(readContent)
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

	// Clean up
	os.Remove("test.txt")
}

func TestDateCreation(t *testing.T) {
	vm.InitVM([]string{"tulipscript"})
	t.Cleanup(vm.FreeVM)

	script := `
		let d = Date(2023, 10, 15)
		print(date_format_datetime(d, "2006-01-02"))
	`
	expectedOutput := "2023-10-15"

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

func TestTimeCreation(t *testing.T) {
	vm.InitVM([]string{"tulipscript"})
	t.Cleanup(vm.FreeVM)

	script := `
		let t = Time(14, 30, 45)
		print(time_format(t, "15:04:05"))
	`
	expectedOutput := "14:30:45"

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

func TestDateTimeCreation(t *testing.T) {
	vm.InitVM([]string{"tulipscript"})
	t.Cleanup(vm.FreeVM)

	script := `
		let dt = DateTime(2023, 10, 15, 14, 30, 45)
		print(datetime_format(dt, "2006-01-02 15:04:05"))
	`
	expectedOutput := "2023-10-15 14:30:45"

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

func TestRandomString(t *testing.T) {
	vm.InitVM([]string{"tulipscript"})
	t.Cleanup(vm.FreeVM)

	script := `
		let s = random_string(8)
		print(s)
	`

	output := captureOutput(t, func() {
		result := core.Interpret(script, "<script>")
		if result != 0 {
			t.Fatalf("Interpretation failed: %d", result)
		}
	})

	if len(output) != 8 {
		t.Errorf("Expected string of length 8, got %q (length %d)", output, len(output))
	}
}

func TestSprintf(t *testing.T) {
	vm.InitVM([]string{"tulipscript"})
	t.Cleanup(vm.FreeVM)

	script := `
		let formatted = sprintf("Value: %v", 42)
		print(formatted)
	`
	expectedOutput := "Value: 42"

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
