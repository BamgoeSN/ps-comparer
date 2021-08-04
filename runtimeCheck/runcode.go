package main

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
	"time"
)

func Run(p *exec.Cmd, in string) (time.Duration, string) {
	p.Stdin = strings.NewReader(in)
	stdout, err := p.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	// Start process
	start := time.Now()
	err = p.Start()
	if err != nil {
		log.Fatal(err)
	}

	// Setup a buffer to capture stdout
	var buf bytes.Buffer

	// Capture any errors
	done := make(chan error)
	go func() {
		if _, err := buf.ReadFrom(stdout); err != nil {
			log.Fatal(err)
		}
		done <- p.Wait()
	}()

	// Switch based on actions received
	select {
	case err = <-done:
		if err != nil {
			close(done)
			return time.Since(start), "Process returned an error, " + err.Error() + "\n" + TrimSpaces(buf.String())
		}
		return time.Since(start), TrimSpaces(buf.String())
	}
}
