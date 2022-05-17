package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var (
	// defaults
	min = 0
	max = 10
)

func printHelp() {
	fmt.Println("Usage:")
	fmt.Println("sleepr\t\tsleeps random seconds between 0 and 10")
	fmt.Println("sleepr foo\t\tsleeps random seconds between 0 and foo")
	fmt.Println("sleepr bar foo\t\tsleeps random seconds between bar and foo")
}

func main() {
	// show help
	help := flag.Bool("help", false, "display help")
	h := flag.Bool("h", false, "display help")
	flag.Parse()

	if *help || *h {
		printHelp()
	}

	// starting rand
	rand.Seed(time.Now().UnixNano())

	switch lonx := len(os.Args[1:]); lonx {
	case 1:
		max, _ = strconv.Atoi(os.Args[1])
	case 2:
		min, _ = strconv.Atoi(os.Args[1])
		max, _ = strconv.Atoi(os.Args[2])
	}

	n := rand.Intn(max-min+1) + min

	// fmt.Printf("Durmindo %d segundos...\n", n)
	time.Sleep(time.Duration(n) * time.Second)
	// fmt.Println("Feito")
}
