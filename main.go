package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type ImageData struct {
	Path string
	Lang string
	Dpi  int
	Psm  int
	Oem  int
}

type ImageService interface {
	GetText() (string, error)
}

func NewImageData(path string, lang string, dpi int, psm int, oem int) ImageData {
	return ImageData{
		Path: path,
		Lang: lang,
		Dpi:  dpi,
		Psm:  psm,
		Oem:  oem,
	}
}

func (i ImageData) GetText() (string, error) {
	command := fmt.Sprintf("tesseract -l %s --dpi %d --psm %d --oem %d %s stdout", i.Lang, i.Dpi, i.Psm, i.Oem, i.Path)
	cmd := exec.Command("tesseract", "-l", i.Lang, "--dpi", strconv.Itoa(i.Dpi), "--psm", strconv.Itoa(i.Psm), "--oem", strconv.Itoa(i.Oem), i.Path, "stdout")
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		return command, err
	}
	strStdOut := string(stdout)
	return strings.TrimSpace(strStdOut), nil
}

func main() {
	text, err := NewImageData("testFiles/1695.gif", "eng", 70, 13, 3).GetText()
	if err != nil {
		println("error : ", err.Error())
		os.Exit(1)
	}
	println("Extracted: ", text)
	os.Exit(0)
}
