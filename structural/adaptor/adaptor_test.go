package adaptor

import "testing"

func TestNewLogger_Info(t *testing.T) {

	newLogger := &NewLogger{
		oldLogger: &OldLogger{name: "oldLogger"},
	}

	got := newLogger.Info("Hello", " ", "World", "!")
	want := "Hello World!"
	if got != want {
		t.Fatalf("Expected %q, got %q", want, got)
	}
}
