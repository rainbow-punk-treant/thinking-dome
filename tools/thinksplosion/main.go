package main

import (
	"os"
)

main() {

doop := `terminator -T 'LOADING THOUGHTS' --geometry=250x500+750+500 --profile='thinking'`



command := os.Command(doop)

for i := 0;i < 9;i++ {
	command.Run()
}
	
}
