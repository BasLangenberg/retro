package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

var question_source = "questions.txt"

func main() {
	run()
}

func run() {
	f, err := ioutil.ReadFile(question_source)
	if err != nil {
		log.Fatalf("Unable to open questions file: %s", err)
		os.Exit(1)
	}

	q := strings.Split(string(f), "\n")

	if len(q) == 0 {
		log.Fatal("Did not load any questions from file, please check questions.txt")
		os.Exit(1)
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(q), func(i, j int) { q[i], q[j] = q[j], q[i] })

	for _, qs := range q {
		fmt.Println(qs)
	}

	var curr int
	total := len(q)
	input := bufio.NewScanner(os.Stdin)

	for {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		if curr == total {
			fmt.Printf("\n\n\nThanks for playing!\n\n")
			fmt.Printf("\n\n\nSource: https://github.com/BasLangenberg/retro!\n\n")
			os.Exit(0)
		}
		fmt.Printf("\033[33mQuestion %d of %d\n\n\033[0m", curr+1, total)
		fmt.Printf("\t\t\033[32m%s\n\033[0m\n\n\n", q[curr])
		fmt.Println("\033[36mPress enter to continue\033[0m")
		input.Scan()
		curr++
	}
}
