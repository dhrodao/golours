package golours_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/dhrodao/golours"
)

// Test Golours Printf service is working properly
func TestGoloursPrintf(t *testing.T) {
	expectOutput(t, func() (int, error) {
		return golours.Printf("%CHello with %C%Ceffects%C!%C\n",
			golours.FgRGBColor{RGBColor: golours.RGBColor{R: 255, G: 56, B: 0}},
			golours.FgRGBColor{RGBColor: golours.RGBColor{R: 255, G: 255, B: 0}}, golours.Bold,
			golours.FgRGBColor{RGBColor: golours.RGBColor{R: 255, G: 56, B: 0}}, golours.Reset)
	}, "\033[38;2;255;56;0mHello with \033[38;2;255;255;0m\033[1meffects\033[38;2;255;56;0m!\033[0m\n")
}

// Test Golours Printf service is working properly
func TestGoloursSprintf(t *testing.T) {
	expectString(t, func() (string, error) {
		return golours.Sprintf("%CHello with %C%Ceffects%C!%C\n",
			golours.FgRGBColor{RGBColor: golours.RGBColor{R: 255, G: 56, B: 0}},
			golours.FgRGBColor{RGBColor: golours.RGBColor{R: 255, G: 255, B: 0}}, golours.Bold,
			golours.FgRGBColor{RGBColor: golours.RGBColor{R: 255, G: 56, B: 0}}, golours.Reset)
	}, "\033[38;2;255;56;0mHello with \033[38;2;255;255;0m\033[1meffects\033[38;2;255;56;0m!\033[0m\n")
}

func expectOutput(t *testing.T, printingFunc func() (int, error), expectedOutput string) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	if _, err := printingFunc(); err != nil {
		t.Error(err)
	}

	outChannel := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outChannel <- buf.String()
	}()

	w.Close()
	os.Stdout = old
	out := <-outChannel

	if out != expectedOutput {
		t.Errorf("Expected %q, got %q", expectedOutput, out)
	} else {
		t.Logf("Output: %v", string(out))
	}
}

func expectString(t *testing.T, printingFunc func() (string, error), expectedString string) {
	s, err := printingFunc()
	if err != nil {
		t.Error(err)
	}

	if s != expectedString {
		t.Errorf("Expected %q, got %q", expectedString, s)
	} else {
		t.Logf("Output: %v", s)
	}
}
