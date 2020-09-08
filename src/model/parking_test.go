package model

import (
	"testing"
)

func TestCreateSlot(t *testing.T) {
	p := Parking{}
	p.CreateSlot(6)

	if len(p.Slots) != 6 {
		t.Errorf("Create slot fails len(p.Slots) is %d expected : %d", len(p.Slots), 6)
	}

	p.CreateSlot(3)

	if len(p.Slots) != 9 {
		t.Errorf("Create slot fails len(p.Slots) is %d expected : %d", len(p.Slots), 9)
	}

}

func createParking(qty int) *Parking {
	p := Parking{}
	p.CreateSlot(qty)
	return &p
}

func TestAllocateSlot(t *testing.T) {
	p := createParking(6)
	car := Car{
		RegNo:  "KA-01-HH-1234",
		Colour: "White",
	}

	p.Allocate(&car)

	if (p.Slots[0].Car == Car{}) {
		t.Errorf("Allotcate parking lot fail p.Slots[0] is %v expected : %v", p.Slots[0], Car{})
	}
}

func TestAllocateSlotAlready(t *testing.T) {
	p := createParking(6)
	car := Car{
		RegNo:  "KA-01-HH-1234",
		Colour: "White",
	}
	p.Allocate(&car)
	p.Allocate(&car)
	if p.Slots[1].Car == car {
		t.Errorf("Allotcate parking lot already fail p.Slots[1] is %v expected : %v", p.Slots[1].Car, car)
	}

}

func TestAllocateSlotFull(t *testing.T) {
	p := createParking(1)
	car1 := Car{
		RegNo:  "KA-01-HH-1231",
		Colour: "White",
	}
	car2 := Car{
		RegNo:  "KA-01-HH-1232",
		Colour: "White",
	}
	p.Allocate(&car1)
	p.Allocate(&car2)

	if len(p.Slots) > 1 {
		t.Errorf("Allotcate parking lot full fail len(p.Slots) is %d expected : %d", len(p.Slots), 1)
	}
}

func TestLeaveSlot(t *testing.T) {
	p := createParking(1)
	car := Car{
		RegNo:  "KA-01-HH-1231",
		Colour: "White",
	}
	p.Allocate(&car)
	p.Leave(1)

	if (p.Slots[0].Car != Car{}) {
		t.Errorf("Leave parking lot fail p.Slots[0] is %v expected : %v", p.Slots[0].Car, Car{})
	}
}

func TestShowStatusParkingSlot(t *testing.T) {
	p := createParking(1)
	car := Car{
		RegNo:  "KA-01-HH-1231",
		Colour: "White",
	}
	p.Allocate(&car)
	res := p.Status()

	if len(res) == 0 {
		t.Errorf("Show status parking lot fail p.Status() return is %v expected : %v", res, 1)
	}

}

func TestFindSlotByColourCar(t *testing.T) {
	p := createParking(2)
	car1 := Car{
		RegNo:  "KA-01-HH-1231",
		Colour: "White",
	}
	car2 := Car{
		RegNo:  "KA-01-HH-1232",
		Colour: "Black",
	}
	p.Allocate(&car1)
	p.Allocate(&car2)
	res := p.FindByColour("Black")
	t.Log(len(res))

	if res[0].Car.RegNo != car2.RegNo {
		t.Errorf("Find slot by car colour car fail p.FindByColour return is %v expected : %v", res[0].Car.RegNo, car2.RegNo)
	}

}

func TestFindSlotByRegNo(t *testing.T) {
	p := createParking(2)
	car1 := Car{
		RegNo:  "KA-01-HH-1231",
		Colour: "White",
	}
	car2 := Car{
		RegNo:  "KA-01-HH-1232",
		Colour: "Black",
	}
	p.Allocate(&car1)
	p.Allocate(&car2)
	res := p.FindByRegNo("KA-01-HH-1232")
	t.Log(len(res))

	if res[0].Car.RegNo != car2.RegNo {
		t.Errorf("Find slot by car colour car fail p.FindByRegNo return is %v expected : %v", res[0].Car.RegNo, car2.RegNo)
	}

}
