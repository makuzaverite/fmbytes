package fmbytes

import "testing"

func TestFormat(t *testing.T) {
	expected := "1.00 KB"
	actual := FormatBytes(1001, 2)

	if actual != expected {
		t.Errorf("Fomat function Expected %s, and got %s", expected, actual)
	}
}
