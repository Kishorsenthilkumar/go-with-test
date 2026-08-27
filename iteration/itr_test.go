package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if expected != repeated {
		t.Errorf("repeated %q is not expected %q", repeated, expected)
	}
}
