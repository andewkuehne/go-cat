package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	var (
		showAll   bool
		showEnds  bool
		showTabs  bool
		numBlanks bool
	)
	flag.BoolVar(&showAll, "A", false, "equivalent to -vET")
	flag.BoolVar(&showEnds, "E", false, "display $ at end of each line")
	flag.BoolVar(&showTabs, "T", false, "display TAB characters as ^I")
	flag.BoolVar(&numBlanks, "s", false, "squeeze multiple blank lines into one")
	flag.Parse()

	if flag.NArg() == 0 {
		copyStdin()
		return
	}

	for _, filename := range flag.Args() {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: could not open file %s: %v\n", filename, err)
			os.Exit(1)
		}
		defer f.Close()

		buf := make([]byte, 4096)
		lastCharWasNewline := true
		lastLineWasBlank := true

		for {
			n, err := f.Read(buf)
			if err != nil {
				if err == io.EOF {
					break
				} else {
					fmt.Fprintf(os.Stderr, "error: could not read file %s: %v\n", filename, err)
					os.Exit(1)
				}
			}
			for i := 0; i < n; i++ {
				if showAll {
					if buf[i] == '\n' {
						if showEnds && lastCharWasNewline {
							fmt.Print("$")
						}
						fmt.Print("\n")
						lastCharWasNewline = true
						lastLineWasBlank = true
					} else if buf[i] == '\t' && showTabs {
						fmt.Print("^I")
						lastCharWasNewline = false
						lastLineWasBlank = false
					} else if buf[i] == '\n' {
						if numBlanks {
							if !lastLineWasBlank {
								fmt.Print("\n")
							}
							lastLineWasBlank = true
						} else {
							fmt.Print("\n")
						}
						lastCharWasNewline = true
					} else if buf[i] == ' ' {
						if numBlanks {
							if !lastLineWasBlank {
								fmt.Print(" ")
							}
							lastLineWasBlank = true
						} else {
							fmt.Print(" ")
						}
						lastCharWasNewline = false
					} else {
						fmt.Print(string(buf[i]))
						lastCharWasNewline = false
						lastLineWasBlank = false
					}
				} else {
					fmt.Print(string(buf[i]))
				}
			}
		}
	}
}

func copyStdin() {
	buf := make([]byte, 4096)
	for {
		n, err := os.Stdin.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Fprintf(os.Stderr, "error: could not read from stdin: %v\n", err)
				os.Exit(1)
			}
		}
		fmt.Print(string(buf[:n]))
	}
}
