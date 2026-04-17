package main_test

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

var checkerBin string

func TestMain(m *testing.M) {
	bin, err := os.MkdirTemp("", "checker-test-*")
	if err != nil {
		panic("failed to create temp dir: " + err.Error())
	}
	defer os.RemoveAll(bin) //nolint:errcheck

	checkerBin = bin + "/checker"
	out, err := exec.Command("go", "build", "-o", checkerBin, ".").CombinedOutput()
	if err != nil {
		panic("failed to build checker: " + err.Error() + "\n" + string(out))
	}

	os.Exit(m.Run())
}

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

func TestChecker_OK(t *testing.T) {
	stdin := "pb\nra\npb\nra\nsa\nra\npa\npa\n"
	stdout, _, code := runChecker(stdin, "0 9 1 8 2")
	if code != 0 {
		t.Fatalf("exit code = %d, want 0", code)
	}
	if stdout != "OK\n" {
		t.Fatalf("stdout = %q, want \"OK\\n\"", stdout)
	}
}

func TestChecker_KO(t *testing.T) {
	stdin := "sa\npb\nrrr\n"
	stdout, _, code := runChecker(stdin, "0 9 1 8 2 7 3 6 4 5")
	if code != 0 {
		t.Fatalf("exit code = %d, want 0", code)
	}
	if stdout != "KO\n" {
		t.Fatalf("stdout = %q, want \"KO\\n\"", stdout)
	}
}

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

func TestChecker_AlreadySortedNoOps(t *testing.T) {
	stdout, _, code := runChecker("", "1 2 3 4 5")
	if code != 0 {
		t.Fatalf("exit code = %d, want 0", code)
	}
	if stdout != "OK\n" {
		t.Fatalf("stdout = %q, want \"OK\\n\"", stdout)
	}
}

func TestChecker_UnsortedNoOps(t *testing.T) {
	stdout, _, code := runChecker("", "3 2 1")
	if code != 0 {
		t.Fatalf("exit code = %d, want 0", code)
	}
	if stdout != "KO\n" {
		t.Fatalf("stdout = %q, want \"KO\\n\"", stdout)
	}
}

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
