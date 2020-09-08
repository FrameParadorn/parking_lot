package model

import "fmt"

type Parking struct {
	Slots []slot
}

type slot struct {
	no  int
	car Car
}

func (p *Parking) CreateSlot(quantity int) {
	for i := 0; i < quantity; i++ {
		p.Slots = append(p.Slots, slot{
			no:  i + 1,
			car: Car{},
		})
	}
}

func (p *Parking) isAlready(car *Car) bool {
	for _, slot := range p.Slots {
		if slot.car.RegNo == car.RegNo {
			return true
		}
	}
	return false

}

func (p *Parking) Allocate(car *Car) error {

	if p.isAlready(car) {
		return fmt.Errorf("Car is already allocated")
	}

	slotFull := true
	for i, slot := range p.Slots {
		if (slot.car == Car{}) {
			p.Slots[i].car = *car
			slotFull = false
			fmt.Printf("Allocated slot number: %d\n", slot.no)
			break
		}
	}

	if slotFull {
		return fmt.Errorf("Sorry, parking lot is full")
	}

	return nil

}

func (p *Parking) Leave(slotNo int) error {

	found := false
	for i, slot := range p.Slots {
		if slot.no == slotNo {
			p.Slots[i].car = Car{}
			found = true
		}
	}

	if found {
		fmt.Printf("Slot number %d is free\n", slotNo)
		return nil
	}

	return fmt.Errorf("Slot number %d not found", slotNo)

}
