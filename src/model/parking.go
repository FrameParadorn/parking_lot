package model

type Parking struct {
	slots []slot
}

type slot struct {
	no  int
	car Car
}

func (p *Parking) CreateSlot(quantity int) {
	for i := 0; i < quantity; i++ {
		p.slots = append(p.slots, slot{
			no:  i + 1,
			car: Car{},
		})
	}

}
