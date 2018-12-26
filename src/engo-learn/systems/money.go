package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
)

type CityType int

const (
	CityTypeNew = iota
	CityTypeTown
	CityTypeCity
	CityTypeMetro
)

type CityUpdateMessage struct {
	Old, New CityType
}

const CityUpdateMessageType string = "CityUpdateMessage"

func (CityUpdateMessage) Type() string {
	return CityUpdateMessageType
}

type AddOfficerMessage struct {
}

const AddOfficerMessageType string = "AddOfficerMessage"

func (AddOfficerMessage) Type() string {
	return AddOfficerMessageType
}

type MoneySystem struct {
	amount                int
	towns, cities, metros int
	officers              int
	elapsed               float32
}

func (m *MoneySystem) New(w *ecs.World) {
	engo.Mailbox.Listen(CityUpdateMessageType, func(msg engo.Message) {
		upd, ok := msg.(CityUpdateMessage)
		if !ok {
			return
		}

		switch upd.New {
		case CityTypeNew:
			m.towns++
		case CityTypeTown:
			m.towns++
			if upd.Old == CityTypeTown {
				m.towns--
			} else if upd.Old == CityTypeCity {
				m.cities--
			} else if upd.Old == CityTypeMetro {
				m.metros--
			}
		case CityTypeCity:
			m.cities++
			if upd.Old == CityTypeTown {
				m.towns--
			} else if upd.Old == CityTypeCity {
				m.cities--
			} else if upd.Old == CityTypeMetro {
				m.metros--
			}
		case CityTypeMetro:
			m.metros++
			if upd.Old == CityTypeTown {
				m.towns--
			} else if upd.Old == CityTypeCity {
				m.cities--
			} else if upd.Old == CityTypeMetro {
				m.metros--
			}
		}

		engo.Mailbox.Listen(AddOfficerMessageType, func(engo.Message) {
			m.officers++
		})
	})
}

func (m *MoneySystem) Update(dt float32) {
	m.elapsed += dt
	if m.elapsed > 10 {
		m.amount += m.towns*100 + m.cities*500 + m.metros*1000
		m.amount -= m.officers * 20
		engo.Mailbox.Dispatch(HUDMoneyMessage{
			Amount: m.amount,
		})

		m.elapsed = 0
	}
}

func (*MoneySystem) Remove(ecs.BasicEntity) {

}
