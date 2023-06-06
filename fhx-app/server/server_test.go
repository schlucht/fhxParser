package server

import "testing"

func TestRun(t *testing.T) {
	_, err := run()
	if err != nil {
		t.Error("failed run")
	}
}
