package helpers

import (
	"bytes"
	"log"
	"os"
	"os/exec"
)

const (
	quiet = true
)

// ExecCmd executes a command, and returns an error if command fails
func ExecCmd(name string, args []string) error {
	cmd := exec.Command(name, args...)

	var output bytes.Buffer
	if quiet {
		cmd.Stdout = &output
		cmd.Stderr = &output
	} else {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	if !quiet {
		log.Printf("$ %s", cmd.Args)
	}

	err := cmd.Run()
	if err != nil {
		log.Printf("Cmd error: %v\n%s", err, output.String())
	}

	return err
}

// MustExecCmd executes a command, and panic if command fails
func MustExecCmd(name string, args []string) {
	if err := ExecCmd(name, args); err != nil {
		log.Fatalln("Fatal cmd error")
	}
}
