package main

import "testing"

func TestAdd(t *testing.T) {
	if ans := add(1, 2); ans != 3 {
		t.Error("add(1, 2) should be equal to 3")
	}
}

// === RUN   TestAdd
// --- PASS: TestAdd (0.00s)
// PASS
// ok      _/Users/itzhouq/dev/golang/hello/go_test        0.571s
