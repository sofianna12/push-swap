package main_test

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"
)

var pushSwapBin string

func TestMain(m *testing.M) {
	bin, err := os.MkdirTemp("", "push-swap-test-*")
	if err != nil {
		panic("failed to create temp dir: " + err.Error())
	}
	defer os.RemoveAll(bin) //nolint:errcheck

	pushSwapBin = bin + "/push-swap"
	out, err := exec.Command("go", "build", "-o", pushSwapBin, ".").CombinedOutput()
	if err != nil {
		panic("failed to build push-swap: " + err.Error() + "\n" + string(out))
	}

	os.Exit(m.Run())
}

func run(args ...string) (stdout, stderr string, exitCode int) {
	cmd := exec.Command(pushSwapBin, args...)
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

func countLines(s string) int {
	n := 0
	for _, line := range strings.Split(strings.TrimSpace(s), "\n") {
		if line != "" {
			n++
		}
	}
	return n
}

func TestPushSwap_NoArgs(t *testing.T) {
	stdout, stderr, code := run()
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

func TestPushSwap_AlreadySorted(t *testing.T) {
	stdout, _, code := run("0 1 2 3 4 5")
	if code != 0 {
		t.Fatalf("exit code = %d, want 0", code)
	}
	if stdout != "" {
		t.Fatalf("stdout = %q, want empty", stdout)
	}
}

func TestPushSwap_AuditInput(t *testing.T) {
	stdout, _, code := run("2 1 3 6 5 8")
	if code != 0 {
		t.Fatalf("exit code = %d, want 0", code)
	}
	n := countLines(stdout)
	if n >= 9 {
		t.Fatalf("op count = %d, want < 9\nops:\n%s", n, stdout)
	}
}

func TestPushSwap_FiveNumbers(t *testing.T) {
	tests := []string{
		"1 5 2 4 3",
		"5 4 3 2 1",
		"3 1 5 2 4",
	}
	for _, arg := range tests {
		t.Run(arg, func(t *testing.T) {
			stdout, _, code := run(arg)
			if code != 0 {
				t.Fatalf("exit code = %d, want 0", code)
			}
			n := countLines(stdout)
			if n >= 12 {
				t.Fatalf("op count = %d, want < 12\nops:\n%s", n, stdout)
			}
		})
	}
}

func TestPushSwap_InvalidArg(t *testing.T) {
	stdout, stderr, code := run("0 one 2 3")
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

func TestPushSwap_Duplicates(t *testing.T) {
	stdout, stderr, code := run("1 2 2 3")
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

func TestPushSwap_100Numbers(t *testing.T) {
	nums := make([]string, 100)
	perm := deterministicPerm(100)
	for i, v := range perm {
		nums[i] = strconv.Itoa(v)
	}
	arg := strings.Join(nums, " ")

	stdout, _, code := run(arg)
	if code != 0 {
		t.Fatalf("exit code = %d, want 0", code)
	}
	n := countLines(stdout)
	if n >= 700 {
		t.Fatalf("op count = %d, want < 700", n)
	}
}

func deterministicPerm(n int) []int {
	perm := make([]int, n)
	for i := range perm {
		perm[i] = i + 1
	}
	seed := uint64(12345)
	next := func() uint64 {
		seed ^= seed << 13
		seed ^= seed >> 7
		seed ^= seed << 17
		return seed
	}
	for i := n - 1; i > 0; i-- {
		j := int(next() % uint64(i+1))
		perm[i], perm[j] = perm[j], perm[i]
	}
	return perm
}
