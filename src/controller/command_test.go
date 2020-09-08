package controller

import (
	"strings"
	"testing"
)

func TestSplitCommand(t *testing.T) {
	command := "create_parking_lot 6"
	cmd, args := splitCommand(&command)

	if cmd != "create_parking_lot" {
		t.Errorf("Create parking split command fail. command is : %s expect : create_parking_lot", cmd)
		return
	}

	if len(args) != 1 {
		t.Errorf("Create parking split argrument fail. len(args) is : %d expect : 1", len(args))
		return
	}

	if !strings.Contains(args[0], "6") {
		t.Errorf("Create parking split argrument fail. args[0] is : %s expect : 6", args[0])
		return
	}

}
