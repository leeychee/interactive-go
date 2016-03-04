package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/bobappleyard/readline"
)

const usage = "icli cmd"

var _debug bool

func init() {
	flag.BoolVar(&_debug, "d", false, "Debug Mode")
}

func main() {
	flag.Parse()
	cmd := flag.Arg(0)
	if _debug {
		fmt.Printf("CMD: \"%s\"\n", cmd)
	}
	if cmd == "" {
		fmt.Printf("Usage: \n%s\n", usage)
		os.Exit(1)
	}
	for {
		subCmd, err := readline.String("> ")
		if _debug {
			fmt.Printf("SubCMD: \"%s\"\n", subCmd)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s", err)
			break
		}
		readline.AddHistory(subCmd)
		var cmdRun *exec.Cmd
		s := strings.TrimSpace(subCmd)
		if strings.HasPrefix(s, "!") {
			command := subCmd[1:]
			if command == "quit" {
				os.Exit(0)
			} else {
				cmdRun = exec.Command("bash", "-c", subCmd[1:])
			}
		} else {
			// Run (cmd subCmd)
			command := strings.Join(flag.Args(), " ") + " " + subCmd
			cmdRun = exec.Command("bash", "-c", command)
		}
		output, err := cmdRun.CombinedOutput()
		if err != nil {
			fmt.Println("error: ", err)
		}
		fmt.Print(string(output))
	}
}
