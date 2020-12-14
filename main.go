package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args) > 1 {

		v := ""
		for yarg := 1; yarg < len(os.Args); yarg++ {
			r := strconv.Itoa(rand.Intn(255))
			g := strconv.Itoa(rand.Intn(255))
			b := strconv.Itoa(rand.Intn(255))
			v += fmt.Sprint("\033[38;2;" + r + ";" + g + ";" + b + "m " + os.Args[yarg] + "\033[0m")
		}

		for i := 0; i < 45; i++ {
			num := rand.Intn(250)
			space := ""
			for c := 0; c < num; c++ {
				space += " "
			}

			fmt.Print(space, v)
		}

	} else {
		fmt.Printf("Usage is \033[38;2;200;0;0mpatternmaker \033[38;2;0;200;0mmessage\033[0m")
	}

}
