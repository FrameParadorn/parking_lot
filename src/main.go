package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/FrameParadorn/parkinglot/controller"
)

func main() {

	var command string
	for {
		fmt.Print("$ ")
		input := bufio.NewReader(os.Stdin)
		command, _ = input.ReadString('\n')
		controller.Run(&command)
	}

}
