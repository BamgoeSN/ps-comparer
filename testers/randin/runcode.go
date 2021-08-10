package main

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
	"time"
)

func Run(p *exec.Cmd, in string) (string, bool) {
	p.Stdin = strings.NewReader(in)

	// Get pipe to stdout
	stdout, err := p.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	// Start process
	timer := time.NewTimer(time.Duration(timeOut) * time.Millisecond)
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
	case <-timer.C:
		if err := p.Process.Kill(); err != nil {
			return "Timeout, " + err.Error(), false
		}
		return "Timeout!", false
	case err = <-done:
		timer.Stop()
		if err != nil {
			close(done)
			return "Process returned an error", false
		}
		return TrimSpaces(buf.String()), true
	}
}
