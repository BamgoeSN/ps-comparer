package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/schollz/progressbar/v3"
)

var fp, _ = os.Create("output.txt")
var outputWriter = bufio.NewWriter(fp)

var (
	timeOut       = 4000 * time.Millisecond
	testcases int = 100
)

func main() {
	inputDebugNum := flag.Int("id", 0, "true if it's for checking inputgen")
	testCaseFlag := flag.Int("tc", testcases, fmt.Sprintf("Number of testcases; default is %d", testcases))
	isJava := flag.Bool("java", false, "Checks if the wrong case is in .jar file")
	isPython := flag.Bool("pypy", false, "Checks if the wrong case is in .py file")
	isJs := flag.Bool("js", false, "Checks if the wrong case is in .js file")
	flag.Parse()

	if *inputDebugNum > 0 {
		fmt.Print("Input generator debug mode\n\n")
		inputDebug(*inputDebugNum)
		return
	}

	testcases = *testCaseFlag

	// runtimeProcs := runtime.NumCPU()
	runtimeProcs := 3
	fmt.Printf("Using %d processors\n", runtimeProcs)
	runtime.GOMAXPROCS(runtimeProcs)

	rand.Seed(time.Now().UnixNano())
	// rand.Seed(1)
	defer outputWriter.Flush()

	wrong := make(map[string]int)

	cnt := 0
	jumps := runtimeProcs * 3

	bar := progressbar.Default(int64(testcases))
	for testcase := 0; testcase < testcases; testcase += jumps {
		// fmt.Printf("%.2f%% done out of %d\r", float64(testcase)/float64(testcases)*100, testcases)

		ins := make([]string, jumps)
		crStr := make([]string, jumps)
		wrStr := make([]string, jumps)
		wrongs := make([]bool, jumps)

		var wg sync.WaitGroup
		for i := range crStr {
			ins[i] = CreateInput()
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				crOut, crErr := Run(ins[i], timeOut, "./cr.exe")
				if crErr == "" {
					crStr[i] = crOut
				} else {
					crStr[i] = crErr
				}

				if *isJava {
					wrOut, wrErr := Run(ins[i], timeOut*2+time.Second, "java", "-jar", "./wr.jar")
					if wrErr == "" {
						wrStr[i] = wrOut
					} else {
						wrStr[i] = wrErr
					}
				} else if *isPython {
					wrOut, wrErr := Run(ins[i], timeOut*3+time.Second*2, "pypy", "./wr.py")
					if wrErr == "" {
						wrStr[i] = wrOut
					} else {
						wrStr[i] = wrErr
					}
				} else if *isJs {
					wrOut, wrErr := Run(ins[i], timeOut*3+time.Second*2, "node", "./wr.js")
					if wrErr == "" {
						wrStr[i] = wrOut
					} else {
						wrStr[i] = wrErr
					}
				} else {
					wrOut, wrErr := Run(ins[i], timeOut, "./wr.exe")
					if wrErr == "" {
						wrStr[i] = wrOut
					} else {
						wrStr[i] = wrErr
					}
				}
			}(i)
		}
		wg.Wait()

		for i := range wrongs {
			if TrimSpaces(crStr[i]) != TrimSpaces(wrStr[i]) {
				wrongs[i] = true
			}
		}

		for i, b := range wrongs {
			if b {
				in, crString, wrString := ins[i], crStr[i], wrStr[i]
				_, e := wrong[in]
				if !e {
					wrong[in] = 1
					cnt++
					fmt.Fprintln(outputWriter, in)
					fmt.Fprintln(outputWriter)
					fmt.Fprintf(outputWriter, "정답: %s\n", crString)
					fmt.Fprintf(outputWriter, "출력: %s\n", wrString)
					fmt.Fprintln(outputWriter)
					// fmt.Println("Found!", cnt, "                          ")
				}
			}
		}
		outputWriter.Flush()

		err := bar.Add(jumps)
		if err != nil {
			continue
		}
	}
	fmt.Printf("\nDone... %d out of %d                               \n", cnt, testcases)
}

func inputDebug(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(CreateInput())
		fmt.Println()
	}
}
