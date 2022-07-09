package main

import (
	"bytes"
	"context"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

// New
func Run(input string, timeLimit time.Duration, process string, args ...string) (outstr string, errstr string) {
	ctx, cancel := context.WithTimeout(context.Background(), timeLimit)
	defer cancel()
	cmd := exec.CommandContext(ctx, process, args...)

	infname := genFile(input)
	f, err := os.Open(infname)
	if err != nil {
		log.Fatal(err, "Error while opening input file")
	}
	defer func() {
		f.Close()
		os.Remove(infname)
	}()
	cmd.Stdin = f

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		if err.Error() == "signal: killed" {
			return "", "Timeout!"
		}
		log.Fatal(err, "Error while running a process", process, args)
	}

	return stdout.String(), stderr.String()
}

func genFile(input string) string {
	l := rand.Intn(10) + 10
	h := sha256.New()
	for i := 0; i < l; i++ {
		v := rand.Int63()
		h.Write([]byte(fmt.Sprintf("%d", v)))
	}
	fname := fmt.Sprintf("input_%x", h.Sum(nil))
	stdin, err := os.Create(fname)
	defer stdin.Close()
	if err != nil {
		log.Fatal(err, "Error while creating a file for an input")
	}
	stdin.WriteString(input)
	return fname
}

// Old2

func getOutput(path string, input string) (stdout string, stderr string) {
	l := rand.Intn(10) + 10
	h := sha256.New()
	for i := 0; i < l; i++ {
		v := rand.Int63()
		h.Write([]byte(fmt.Sprintf("%d", v)))
	}
	fname := fmt.Sprintf("input_%x", h.Sum(nil))

	stdin, _ := os.Create(fname)
	defer func() {
		stdin.Close()
		os.Remove(fname)
	}()
	stdin.WriteString(input)

	timer := time.NewTimer(time.Duration(timeOut) * time.Millisecond)
	done := make(chan bool)

	go func() {
		stdout, stderr = runProc(path, stdin)
		stdout = TrimSpaces(stdout)
		done <- true
	}()

	select {
	case <-timer.C:
		return "", "Timeout!"

	case <-done:
		timer.Stop()
		return
	}
}

func runProc(path string, input *os.File) (stdout string, stderr string) {
	r, w, _ := os.Pipe()
	er, ew, _ := os.Pipe()

	attr := &os.ProcAttr{Env: os.Environ(), Files: []*os.File{input, w, ew}, Sys: nil}
	os.StartProcess(path, []string{path}, attr)

	w.Close()
	ew.Close()

	out, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	stdout = string(out)
	eout, err := ioutil.ReadAll(er)
	if err != nil {
		log.Fatal(err)
	}
	stderr = string(eout)

	return
}

// Old

func RunOld(p *exec.Cmd, in string) string {
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
			return "Timeout, " + err.Error()
		}
		return "Timeout!"
	case err = <-done:
		timer.Stop()
		if err != nil {
			close(done)
			return "Process returned an error, " + err.Error() + "\n" + TrimSpaces(buf.String())
		}
		return TrimSpaces(buf.String())
	}
}
