package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func makebar(number int, iteration int) {

	progress := " "
	if iteration == 0 {
		for col := 1; col < number; col++ {
			r := strconv.Itoa(col)
			g := strconv.Itoa(0)
			b := strconv.Itoa(col)
			fmt.Print("\033[48;2;" + r + ";" + g + ";" + b + "m " + progress)
			time.Sleep(5 * time.Millisecond)
		}
	}
	if iteration == 1 {
		for col := 1; col < number; col++ {
			r := strconv.Itoa(col)
			g := strconv.Itoa(50)
			b := strconv.Itoa(10)
			fmt.Print("\033[48;2;" + r + ";" + g + ";" + b + "m " + progress)
			time.Sleep(5 * time.Millisecond)
		}
	}
	if iteration == 2 {
		for col := 1; col < number; col++ {
			r := strconv.Itoa(col)
			g := strconv.Itoa(0)
			b := strconv.Itoa(75)
			fmt.Print("\033[48;2;" + r + ";" + g + ";" + b + "m " + progress)
			time.Sleep(5 * time.Millisecond)
		}
	}
	if iteration == 3 {
		for col := 1; col < number; col++ {
			r := strconv.Itoa(col)
			g := strconv.Itoa(col)
			b := strconv.Itoa(0)
			fmt.Print("\033[48;2;" + r + ";" + g + ";" + b + "m " + progress)
			time.Sleep(5 * time.Millisecond)
		}
	}
	if iteration == 4 {
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

	if len(os.Args) > 1 {
		iter, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Iteration count must be a number")
			os.Exit(1)
		}
		for iterations := 0; iterations < iter; iterations++ {
			switch iterations {
			case 0:
				makebar(155, 0)
			case 1:
				makebar(200, 1)
			case 2:
				makebar(110, 2)
			case 3:
				makebar(155, 3)
			case 4:
				makebar(175, 4)
			default:
				fmt.Println("De-fault!")
			}
		}
	} else {
		fmt.Print("You must provide a number of wave iterations")
		os.Exit(1)
	}

}
