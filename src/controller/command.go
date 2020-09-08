package controller

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/FrameParadorn/parkinglot/model"
)

var commands = map[string]func(args []string) error{
	"create_parking_lot": createSlot,
}

var parking model.Parking

func Run(command *string) {

	cmd, arg := splitCommand(command)
	if commands[cmd] != nil {
		err := commands[cmd](arg)
		if err == nil {
			return
		}
	}

	fmt.Println("Command invalid.")
}

func splitCommand(command *string) (string, []string) {
	s := strings.TrimSuffix(*command, "\n")
	result := strings.Split(s, " ")
	if len(result) == 1 {
		return result[0], nil
	}
	args := result[1:]
	return result[0], args
}

func createSlot(args []string) error {
	slotQty, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("Create slot error : %s", err)
		return err
	}
	parking.CreateSlot(slotQty)
	fmt.Printf("Created a parking lot with %d slots\n", slotQty)
	return nil
}
