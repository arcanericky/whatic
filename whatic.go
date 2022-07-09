package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
)

const envWhatIC = "WHATIC"

var versionText = "<unknown>"

func bashExamples() {
	fmt.Println(`Bash Examples

$ whatic arg1 "arg 2" arg\"3
Command Line
------------
whatic arg1 arg 2 arg"3

Arguments
---------
0 whatic
1 arg1
2 arg 2
3 arg"3`)
}

func windowsExamples() {
	fmt.Println(`Windows Examples

Command Prompt (cmd.exe)
C:\>whatic arg1 "arg 2" arg\"3
Command Line
------------
whatic arg1 arg 2 arg"3

Arguments
---------
0 whatic
1 arg1
2 arg 2
3 arg"3

PowerShell
PS C:\WhatIC> whatic 'arg1\"arg2'
Command Line
------------
C:\WhatIC\whatic.exe arg1"arg2

Arguments
---------
0 C:\WhatIC\whatic.exe
1 arg1"arg2`)
}

func examples(system string) {
	switch system {
	case "linux", "darwin":
		bashExamples()
	case "windows":
		windowsExamples()
	}
}

func help() {
	fmt.Printf("WhatIC version %s %s/%s\n", versionText, runtime.GOOS, runtime.GOARCH)
	fmt.Println(`  Because the command-line interface is difficult.
  https://github.com/arcanericky/whatic

  Execute this application with parameters to reveal how the
  command-line interface parses and transforms them. See the
  README for more information and examples.`)
	fmt.Println()

	examples(runtime.GOOS)
}

func arguments(opener, closer string, args []string) {
	// Show the command line as a single line
	fmt.Println("Command Line")
	fmt.Println("------------")
	fmt.Printf("%s%s%s", opener, args[0], closer)
	for _, argument := range args[1:] {
		fmt.Printf(" %s%s%s", opener, argument, closer)
	}
	fmt.Printf("\n\n")

	// Show each argument individually
	digits := int(math.Floor(math.Log10(float64(len(args)))) + 1)
	fmt.Println("Arguments")
	fmt.Println("---------")
	for i, argument := range args {
		fmt.Printf("%*d %s%s%s\n", digits, i, opener, argument, closer)
	}
}

func delimiters() (string, string) {
	delimiter := os.Getenv(envWhatIC)
	delimiterLen := len(delimiter)

	if delimiterLen > 0 {
		fmt.Println("Environment")
		fmt.Println("-----------")
		fmt.Printf(envWhatIC+"=%s\n\n", delimiter)
	}

	switch delimiterLen {
	case 0:
		return "", ""
	case 1:
		return delimiter, delimiter
	}

	return string(delimiter[0]), string(delimiter[1])
}

func main() {
	if len(os.Args) == 1 {
		help()
		return
	}

	opener, closer := delimiters()
	arguments(opener, closer, os.Args)
}
