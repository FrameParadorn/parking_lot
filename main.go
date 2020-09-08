package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/FrameParadorn/parkinglot/src/controller"
)

func main() {

	fmt.Println(len(os.Args))
	if len(os.Args) >= 2 {
		readFileContent()
		return
	} else {
		waitStdin()
	}

}

func readFileContent() {
	file, err := os.Open(os.Args[1])
	defer file.Close()
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		controller.Run(&line, false)
	}

	err = scanner.Err()
	if err != nil {
		panic(err)
	}

}

func waitStdin() {
	var command string
	for {
		fmt.Print("$ ")
		input := bufio.NewReader(os.Stdin)
		command, _ = input.ReadString('\n')
		controller.Run(&command, true)
	}
}
