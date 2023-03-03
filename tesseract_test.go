package main

import "testing"

func TestImageData_GetText(t *testing.T) {
	testFiles := map[string]string{
		"testFiles/1695.gif":               "1695",
		"testFiles/alphabet.jpeg":          "1695",
		"testFiles/elon-musk-bitcoin.webp": "1695",
	}
	image := NewImageData("testFiles/1695.gif", "eng", 70, 13, 3)
	got, err := image.GetText()
	want := "1695"
	if err != nil {
		t.Fatalf("get error %q", err)
	}
	if got != want {
		t.Fatalf("want %q, get %q", want, got)
	}
}

func TestArabicImageData_GetText(t *testing.T) {
	testFiles := map[string]string{
		"testFiles/is-android-apps-profitable.jpg": "1695",
	}
	image := NewImageData("testFiles/1695.gif", "eng", 70, 13, 3)
	got, err := image.GetText()
	want := "1695"
	if err != nil {
		t.Fatalf("get error %q", err)
	}
	if got != want {
		t.Fatalf("want %q, get %q", want, got)
	}
}
