package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printSomething(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	t.Log("stdOut:", stdOut)
	t.Log("r:", r)
	t.Log("w:", w)

	var wg sync.WaitGroup
	wg.Add(1)

	go printSomething("first", &wg)

	wg.Wait()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	t.Logf("output: %v\n", result)
	t.Logf("output: %v\n", output)

	os.Stdout = stdOut

	if !strings.Contains(output, "first") {
		t.Errorf("printSomething() = %v, want %v", output, "first")
	}
}
