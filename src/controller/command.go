package controller

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/FrameParadorn/parkinglot/model"
)

var commands = map[string]func(arg interface{}) error{
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

	fmt.Println("Command invalid.\n")
}

func createSlot(arg interface{}) error {
	slotQty, err := strconv.Atoi(arg.(string))
	if err != nil {
		fmt.Println("Create slot error : %s", err)
		return err
	}
	parking.CreateSlot(slotQty)
	fmt.Printf("Created a parking lot with %d slots\n", slotQty)
	return nil
}

func splitCommand(command *string) (string, string) {
	s := strings.TrimSuffix(*command, "\n")
	result := strings.Split(s, " ")
	if len(result) == 1 {
		return result[0], ""
	}
	return result[0], result[1]
}
