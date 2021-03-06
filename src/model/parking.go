package model

import "fmt"

type Parking struct {
	Slots      []slot
	CountLeave int
}

type slot struct {
	No  int
	Car Car
}

func (p *Parking) CreateSlot(quantity int) {
	for i := 0; i < quantity; i++ {
		p.Slots = append(p.Slots, slot{
			No:  i + 1,
			Car: Car{},
		})
	}
}

func (p *Parking) isAlready(car *Car) bool {
	for _, slot := range p.Slots {
		if slot.Car.RegNo == car.RegNo {
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
		if (slot.Car == Car{}) {
			p.Slots[i].Car = *car
			slotFull = false
			fmt.Printf("Allocated slot number: %d\n", slot.No)
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
		if slot.No == slotNo {
			p.Slots[i].Car = Car{}
			found = true
		}
	}

	if found {
		fmt.Printf("Slot number %d is free\n", slotNo)
		p.CountLeave++
		return nil
	}

	return fmt.Errorf("Slot number %d not found", slotNo)

}

func (p *Parking) Status() []slot {

	result := []slot{}
	for _, slot := range p.Slots {
		if (slot.Car != Car{}) {
			result = append(result, slot)
		}
	}
	return result

}

func (p *Parking) FindByColour(colour string) []slot {

	result := []slot{}
	for _, slot := range p.Slots {
		if slot.Car.Colour == colour {
			result = append(result, slot)
		}
	}
	return result

}

func (p *Parking) FindByRegNo(regNo string) []slot {
	result := []slot{}
	for _, slot := range p.Slots {
		if slot.Car.RegNo == regNo {
			result = append(result, slot)
		}
	}
	return result

}

func (p *Parking) Find(column string, keyword string) []slot {

	switch column {
	case "Colour":
		fallthrough
	case "SlotNo":
		return p.FindByColour(keyword)
	default:
		return p.FindByRegNo(keyword)

	}

	return []slot{}

}
