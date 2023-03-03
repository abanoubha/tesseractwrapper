package main

import "testing"

func TestImageData_GetText(t *testing.T) {
	image := NewImageData("testFile/1695.gif", "eng", 70, 13, 3)
	got, err := image.GetText()
	want := "1695"
	if err != nil {
		t.Fatalf("get error %q", err)
	}
	if got != want {
		t.Fatalf("want %q, get %q", want, got)
	}
}
