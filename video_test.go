package video

import "testing"

func TestExtension(t *testing.T) {
	curLen := 19
	if len(Extensions) != curLen {
		t.Fatalf("Length doesn't match. Expected %d, Got %d", curLen, len(Extensions))
	}
}

func TestIs(t *testing.T) {
	if Is("a/b/c/bar.css") {
		t.Fatal("Wrong detection. Must not be video")
	}

	if !Is("a/b/c/bar.wmv") {
		t.Fatal("Wrong detection. Must be video")
	}

	if !Is("a/b/c/bar.WMV") {
		t.Fatal("Wrong detection. Must be video")
	}

	if Is("a/b/c/barwmv") {
		t.Fatal("Wrong detection. Must not be video")
	}
}
