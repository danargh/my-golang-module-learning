package say_hello

import (
	"testing"
)

func TestSayHello(t *testing.T) {
	result := SayHello("World")
	if result != "Hello World" {
		t.Error("Expected Hello World, got ", result)
	}
}
