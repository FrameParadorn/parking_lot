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
	"leave":              leaveSlot,
	"status":             showStatusSlot,
	"registration_numbers_for_cars_with_colour": showRegNoByColour,
	"slot_numbers_for_cars_with_colour":         showSlotNoByColour,
	"slot_number_for_registration_number":       showSlotByRegNo,
}

var parking model.Parking

const cmdInv = "Command invalid."

func Run(command *string) {

	cmd, arg := splitCommand(command)
	if commands[cmd] == nil {
		fmt.Println(cmdInv)
		return
	}

	err := commands[cmd](arg)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println()

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
		fmt.Printf("Create slot error : %s\n", err)
		return err
	}
	parking.CreateSlot(slotQty)
	fmt.Printf("Created a parking lot with %d slots\n", slotQty)
	return nil
}

func allocateSlot(args []string) error {
	if len(args) != 2 {
		fmt.Println(cmdInv)
		return nil
	}
	car := model.Car{
		RegNo:  args[0],
		Colour: args[1],
	}

	return parking.Allocate(&car)

}

func leaveSlot(args []string) error {
	slotNo, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("Leave slot error : %s\n", err)
		return err
	}

	return parking.Leave(slotNo)
}

func showStatusSlot(args []string) error {
	parking.Status()
	return nil
}

func showRegNoByColour(args []string) error {
	if len(args) != 1 {
		fmt.Println(cmdInv)
		return nil
	}

	slots := parking.Find("Colour", args[0])
	for i, slot := range slots {
		if i != 0 {
			fmt.Printf(", ")
		}
		fmt.Printf("%s", slot.Car.RegNo)
	}
	fmt.Println()
	return nil
}

func showSlotNoByColour(args []string) error {
	if len(args) != 1 {
		fmt.Println(cmdInv)
		return nil
	}

	slots := parking.Find("SlotNo", args[0])
	for i, slot := range slots {
		if i != 0 {
			fmt.Printf(", ")
		}
		fmt.Printf("%d", slot.No)
	}
	fmt.Println()
	return nil
}

func showSlotByRegNo(args []string) error {
	if len(args) != 1 {
		fmt.Println(cmdInv)
		return nil
	}

	slots := parking.Find("RegNo", args[0])

	if len(slots) == 0 {
		return fmt.Errorf("Not found")
	}

	fmt.Printf("%d", slots[0].No)
	return nil

}
