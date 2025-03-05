package golours

import (
	"fmt"
)

type TextEffector interface {
	Sequence() string
}

type BasicEffect string

func (e BasicEffect) Sequence() string {
	return string(e)
}

const (
	Reset         BasicEffect = "\033[0m"
	Bold          BasicEffect = "\033[1m"
	Italic        BasicEffect = "\033[3m"
	Underline     BasicEffect = "\033[4m"
	Strikethrough BasicEffect = "\033[9m"
	FontBlack     BasicEffect = "\033[30m"
	FontRed       BasicEffect = "\033[31m"
	FontGreen     BasicEffect = "\033[32m"
	FontYellow    BasicEffect = "\033[33m"
	FontBlue      BasicEffect = "\033[34m"
	FontMagenta   BasicEffect = "\033[35m"
	FontCyan      BasicEffect = "\033[36m"
	FontWhite     BasicEffect = "\033[37m"
	BgBlack       BasicEffect = "\033[40m"
	BgRed         BasicEffect = "\033[41m"
	BgGreen       BasicEffect = "\033[42m"
	BgYellow      BasicEffect = "\033[43m"
	BgBlue        BasicEffect = "\033[44m"
	BgMagenta     BasicEffect = "\033[45m"
	BgCyan        BasicEffect = "\033[46m"
	BgWhite       BasicEffect = "\033[47m"
)

type RGBColor struct {
	R, G, B int
}

func (e RGBColor) String() string {
	return fmt.Sprintf("%d;%d;%d", e.R, e.G, e.B)
}

type FgRGBColor struct {
	RGBColor
}

func (e FgRGBColor) Sequence() string {
	return fmt.Sprintf("\033[38;2;%sm", e.RGBColor)
}

type BgRGBColor struct {
	RGBColor
}

func (e BgRGBColor) Sequence() string {
	return fmt.Sprintf("\033[48;2;%sm", e.RGBColor)
}

// Printf formats according to the format specifier '%C' and prints the resulting string.
// It returns the number of bytes written or an error if occurs.
func Printf(format string, args ...TextEffector) (int, error) {
	buf, err := doPrintf(format, args...)
	if err != nil {
		return 0, err
	}
	return fmt.Print(string(buf))
}

// Sprintf formats according to the format specifier '%C' and returns the resulting string or an error if occurs.
func Sprintf(format string, args ...TextEffector) (string, error) {
	buf, err := doPrintf(format, args...)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

// Internal implementation of the printf service
func doPrintf(format string, args ...TextEffector) (buf []byte, err error) {
	countFmtSequences := 0
	currArg := 0
	ilast := 0
	for i := 0; i < len(format); {
		if format[i] == '%' {
			buf = append(buf, format[ilast:i]...)

			if i+1 >= len(format) {
				continue
			}

			if format[i+1] == 'C' {
				countFmtSequences++
				if countFmtSequences > len(args) {
					return nil, fmt.Errorf("not enough arguments")
				}

				buf = append(buf, args[currArg].Sequence()...)

				currArg++
				i += 2
				ilast = i
			} else {
				return nil, fmt.Errorf("invalid formatting sequence: %v ", format[i:i+2])
			}
		} else {
			if i == len(format)-1 {
				buf = append(buf, format[ilast:i+1]...)
			} else {
				buf = append(buf, format[ilast:i]...)
			}
			ilast = i
			i++
		}
	}

	if countFmtSequences < len(args) {
		return nil, fmt.Errorf("not enough arguments")
	}

	return buf, nil
}
