package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
	"time"
)

func main() {
	n := flag.Int("rate", 1, "rate limit per second")
	p := flag.Int("inflight", 1, "max parallel inflight commands")

	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		log.Fatalf("failed to get a command argument\n")
	}
	command := args[0]

	var cliArgs []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		cliArgs = append(cliArgs, scanner.Text())
	}

	if len(cliArgs) == 0 {
		fmt.Printf("no cli arguments\n")
		return
	}

	burstyLimit := make(chan time.Time, *p)
	for i := 0; i < *p; i++ {
		burstyLimit <- time.Now()
	}

	go func() {
		duration := time.Duration(1000 / *n) * time.Millisecond
		for t := range time.Tick(duration) {
			burstyLimit <- t
		}
	}()

	wg := sync.WaitGroup{}

	for _, arg := range cliArgs {
		wg.Add(1)
		<-burstyLimit
		go func(arg string) {
			out, err := exec.Command(command, arg).Output()
			if err != nil {
				log.Fatalf("failed to run command. Error: %v", err)
			}
			fmt.Printf("%s", out)
			wg.Done()
		}(arg)

	}

	wg.Wait()
}
