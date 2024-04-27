package main

import (
	"sync"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	var wq sync.WaitGroup

	wq.Add(1)
	go updateMessage("Hello, 1!", &wq)
	wq.Wait()

	if msg != "Hello, 1!" {
		t.Errorf("msg = %s; want Hello, 1!", msg)
	}
}
