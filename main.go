package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func makebar(number int, iteration int) {

	progress := " "

	if iteration%2 == 0 {
		for col := 1; col < number; col++ {
			r := strconv.Itoa(col)
			g := strconv.Itoa(0)
			b := strconv.Itoa(col)
			fmt.Print("\033[48;2;" + r + ";" + g + ";" + b + "m " + progress)
			time.Sleep(5 * time.Millisecond)
		}
	}

	if iteration%2 == 1 {
		for col := 1; col < number; col++ {
			r := strconv.Itoa(0)
			g := strconv.Itoa(col)
			b := strconv.Itoa(col)
			fmt.Print("\033[48;2;" + r + ";" + g + ";" + b + "m " + progress)
			time.Sleep(5 * time.Millisecond)
		}
	}
	fmt.Printf("\033[0m")
}

func main() {
	fmt.Print("\033[0;0H")
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args) > 1 {
		iter, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Iteration count must be a number")
			os.Exit(1)
		}
		for iterations := 0; iterations < iter; iterations++ {
			makebar(rand.Intn(255), iterations)
		}
	} else {
		fmt.Print("You must provide a number of wave iterations")
		os.Exit(1)
	}

}
