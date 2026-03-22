package main_test

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

// checkerBin is the path to the compiled checker binary used by all tests.
var checkerBin string

func TestMain(m *testing.M) {
	bin, err := os.MkdirTemp("", "checker-test-*")
	if err != nil {
		panic("failed to create temp dir: " + err.Error())
	}
	defer os.RemoveAll(bin)

	checkerBin = bin + "/checker"
	out, err := exec.Command("go", "build", "-o", checkerBin, ".").CombinedOutput()
	if err != nil {
		panic("failed to build checker: " + err.Error() + "\n" + string(out))
	}

	os.Exit(m.Run())
}

// runChecker executes checker with the given arguments and stdin, returns stdout, stderr, exit code.
func runChecker(stdin string, args ...string) (stdout, stderr string, exitCode int) {
	cmd := exec.Command(checkerBin, args...)
	cmd.Stdin = strings.NewReader(stdin)
	var outBuf, errBuf strings.Builder
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf
	err := cmd.Run()
	stdout = outBuf.String()
	stderr = errBuf.String()
	exitCode = 0
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitCode = exitErr.ExitCode()
		}
	}
	return
}

// TestChecker_NoArgs — no arguments: exit 0, no output.
func TestChecker_NoArgs(t *testing.T) {
	stdout, stderr, code := runChecker("")
	if code != 0 {
		t.Fatalf("exit code = %d, want 0", code)
	}
	if stdout != "" {
		t.Fatalf("stdout = %q, want empty", stdout)
	}
	if stderr != "" {
		t.Fatalf("stderr = %q, want empty", stderr)
	}
}

// TestChecker_InvalidArg — non-integer argument: exit 1, stderr = "Error\n".
func TestChecker_InvalidArg(t *testing.T) {
	stdout, stderr, code := runChecker("", "0 one 2 3")
	if code != 1 {
		t.Fatalf("exit code = %d, want 1", code)
	}
	if stderr != "Error\n" {
		t.Fatalf("stderr = %q, want \"Error\\n\"", stderr)
	}
	if stdout != "" {
		t.Fatalf("stdout = %q, want empty", stdout)
	}
}

// TestChecker_OK — valid instructions that sort the stack: stdout = "OK\n".
func TestChecker_OK(t *testing.T) {
	// echo -e "pb\nra\npb\nra\nsa\nra\npa\npa\n" | ./checker "0 9 1 8 2"
	stdin := "pb\nra\npb\nra\nsa\nra\npa\npa\n"
	stdout, _, code := runChecker(stdin, "0 9 1 8 2")
	if code != 0 {
		t.Fatalf("exit code = %d, want 0", code)
	}
	if stdout != "OK\n" {
		t.Fatalf("stdout = %q, want \"OK\\n\"", stdout)
	}
}

// TestChecker_KO — valid instructions that do not sort the stack: stdout = "KO\n".
func TestChecker_KO(t *testing.T) {
	// echo -e "sa\npb\nrrr\n" | ./checker "0 9 1 8 2 7 3 6 4 5"
	stdin := "sa\npb\nrrr\n"
	stdout, _, code := runChecker(stdin, "0 9 1 8 2 7 3 6 4 5")
	if code != 0 {
		t.Fatalf("exit code = %d, want 0", code)
	}
	if stdout != "KO\n" {
		t.Fatalf("stdout = %q, want \"KO\\n\"", stdout)
	}
}

// TestChecker_UnknownInstruction — invalid instruction: exit 1, stderr = "Error\n".
func TestChecker_UnknownInstruction(t *testing.T) {
	stdout, stderr, code := runChecker("sa\nXX\n", "1 2 3")
	if code != 1 {
		t.Fatalf("exit code = %d, want 1", code)
	}
	if stderr != "Error\n" {
		t.Fatalf("stderr = %q, want \"Error\\n\"", stderr)
	}
	if stdout != "" {
		t.Fatalf("stdout = %q, want empty", stdout)
	}
}

// TestChecker_AlreadySortedNoOps — sorted stack, no instructions: stdout = "OK\n".
func TestChecker_AlreadySortedNoOps(t *testing.T) {
	stdout, _, code := runChecker("", "1 2 3 4 5")
	if code != 0 {
		t.Fatalf("exit code = %d, want 0", code)
	}
	if stdout != "OK\n" {
		t.Fatalf("stdout = %q, want \"OK\\n\"", stdout)
	}
}

// TestChecker_UnsortedNoOps — unsorted stack, no instructions: stdout = "KO\n".
func TestChecker_UnsortedNoOps(t *testing.T) {
	stdout, _, code := runChecker("", "3 2 1")
	if code != 0 {
		t.Fatalf("exit code = %d, want 0", code)
	}
	if stdout != "KO\n" {
		t.Fatalf("stdout = %q, want \"KO\\n\"", stdout)
	}
}

// TestChecker_EmptyLines — empty lines in stdin are silently skipped.
func TestChecker_EmptyLines(t *testing.T) {
	stdin := "\nsa\n\n"
	stdout, _, code := runChecker(stdin, "2 1")
	if code != 0 {
		t.Fatalf("exit code = %d, want 0", code)
	}
	if stdout != "OK\n" {
		t.Fatalf("stdout = %q, want \"OK\\n\"", stdout)
	}
}

// TestChecker_Duplicates — duplicate args: exit 1, stderr = "Error\n".
func TestChecker_Duplicates(t *testing.T) {
	stdout, stderr, code := runChecker("", "1 2 2 3")
	if code != 1 {
		t.Fatalf("exit code = %d, want 1", code)
	}
	if stderr != "Error\n" {
		t.Fatalf("stderr = %q, want \"Error\\n\"", stderr)
	}
	if stdout != "" {
		t.Fatalf("stdout = %q, want empty", stdout)
	}
}
