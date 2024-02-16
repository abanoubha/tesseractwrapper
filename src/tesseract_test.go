package main

import "testing"

func TestImageData_GetText(t *testing.T) {
	testFiles := map[string]string{
		"../testFiles/1695.gif":               "1695",
		"../testFiles/alphabet.jpeg":          "a",
		"../testFiles/elon-musk-bitcoin.webp": "=",
	}
	for imagePath, expected := range testFiles {
		image := NewImageData(imagePath, "eng", 70, 13, 3)
		got, err := image.GetText()
		if err != nil {
			t.Fatalf("get error %q", err)
		}
		if got != expected {
			t.Fatalf("want %q, get %q", expected, got)
		}
	}
}

func TestArabicImageData_GetText(t *testing.T) {
	testFiles := map[string]string{
		"../testFiles/is-android-apps-profitable.jpg": "1695",
	}
	for imagePath, expected := range testFiles {
		image := NewImageData(imagePath, "ara", 70, 13, 3)
		got, err := image.GetText()
		if err != nil {
			t.Fatalf("get error %q", err)
		}
		if got != expected {
			t.Fatalf("want %q, get %q", expected, got)
		}
	}
}
