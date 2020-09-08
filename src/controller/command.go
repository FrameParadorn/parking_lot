package controller

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/FrameParadorn/parkinglot/model"
)

var commands = map[string]func(args []string) error{
	"create_parking_lot": createSlot,
	"park":               allocateSlot,
}

var parking model.Parking

func Run(command *string) {

	cmd, arg := splitCommand(command)
	if commands[cmd] == nil {
		fmt.Println("Command invalid.")
		return
	}

	err := commands[cmd](arg)
	if err != nil {
		fmt.Println(err)
	}

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

func allocateSlot(args []string) error {
	car := model.Car{
		RegNo:  args[0],
		Colour: args[1],
	}

	err := parking.Allocate(&car)
	if err != nil {
		return err
	}

	fmt.Println(parking.Slots)
	return nil

}
