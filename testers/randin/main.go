package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"sync"
	"time"

	"github.com/schollz/progressbar/v3"
)

var fp, _ = os.Create("output.txt")
var writer = bufio.NewWriter(fp)

var (
	timeOut   int = 4000 // In ms
	testcases int = 10
)

func main() {
	inputDebugNum := flag.Int("id", 0, "true if it's for checking inputgen")
	testCaseFlag := flag.Int("tc", testcases, fmt.Sprintf("Number of testcases; default is %d", testcases))
	isJava := flag.Bool("java", false, "Checks if the wrong case is in .jar file")
	isPython := flag.Bool("pypy", false, "Checks if the wrong case is in .py file")
	flag.Parse()

	if *inputDebugNum > 0 {
		fmt.Print("Input generator debug mode\n\n")
		inputDebug(*inputDebugNum)
		return
	}

	testcases = *testCaseFlag

	runtimeProcs := runtime.NumCPU()
	fmt.Printf("Using %d processors\n", runtimeProcs)
	runtime.GOMAXPROCS(runtimeProcs)

	rand.Seed(time.Now().UnixNano())
	// rand.Seed(1)
	defer writer.Flush()
	start := time.Now()

	cnt := 0
	jumps := runtimeProcs * 3

	bar := progressbar.Default(int64(testcases))
	for testcase := 0; testcase < testcases; testcase += jumps {
		// fmt.Printf("%.2f%% done out of %d\r", float64(testcase)/float64(testcases)*100, testcases)

		ins := make([]string, jumps)
		outs := make([]string, jumps)

		var wg sync.WaitGroup
		for i := range outs {
			ins[i] = CreateInput()
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				var wr *exec.Cmd
				if *isJava {
					wr = exec.Command(`java`, `-jar`, `.\wr.jar`)
				} else if *isPython {
					wr = exec.Command(`pypy`, `.\wr.py`)
				} else {
					wr = exec.Command(".\\wr.exe")
				}
				outs[i], _ = Run(wr, ins[i])
			}(i)
		}
		wg.Wait()

		for i := range ins {
			fmt.Fprintln(writer, ins[i])
			fmt.Fprintln(writer)
			fmt.Fprintln(writer, outs[i])
			fmt.Fprint(writer, "\n\n")
		}

		err := bar.Add(jumps)
		if err != nil {
			continue
		}
	}
	fmt.Printf("\nDone... %d out of %d                               \n", cnt, testcases)
	fmt.Printf("Done in %v\n", time.Since(start))
}

func inputDebug(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(CreateInput())
		fmt.Println()
	}
}
