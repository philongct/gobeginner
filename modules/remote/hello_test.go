// Must be the same with hello.go
package remote

import "testing"

func TestHello(t *testing.T) {
    want := "Hi, world."
    if got := Hello(); got != want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}