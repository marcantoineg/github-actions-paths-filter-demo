package app

import "testing"

func TestApp(t *testing.T) {
	if app() != "Man, I really love cats, in golang." {
		t.Error("app() did not return the expected string")
	}
}
