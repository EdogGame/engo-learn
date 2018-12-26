package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"fmt"
	"image/color"
)

type mouseTracker struct {
	ecs.BasicEntity
	common.MouseComponent
}

type Text struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

type HUDTextMessage struct {
	ecs.BasicEntity
	common.SpaceComponent
	common.MouseComponent
	Line1, Line2, Line3, Line4 string
}

const HUDTextMessageType string = "HUDTextMessage"

func (HUDTextMessage) Type() string {
	return HUDTextMessageType
}

type HUDMoneyMessage struct {
	Amount int
}

const HUDMoneyMessageType string = "HUDMoneyMessage"

func (HUDMoneyMessage) Type() string {
	return HUDMoneyMessageType
}

type HUDTextEntity struct {
	*ecs.BasicEntity
	*common.SpaceComponent
	*common.MouseComponent
	Line1, Line2, Line3, Line4 string
}

type HUDTextSystem struct {
	text1, text2, text3, text4, money Text
	entities                          []HUDTextEntity
	updateMoney                       bool
	amount                            int
}

func (h *HUDTextSystem) New(w *ecs.World) {
	fnt := &common.Font{
		URL:  "go.ttf",
		FG:   color.Black,
		Size: 24,
	}

	fnt.CreatePreloaded()

	// text1
	h.text1 = Text{BasicEntity: ecs.NewBasic()}
	h.text1.RenderComponent.Drawable = common.Text{
		Font: fnt,
		Text: "Nothing Selected!",
	}
	h.text1.SetShader(common.TextHUDShader)
	h.text1.RenderComponent.SetZIndex(1001)
	h.text1.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{
			X: 0,
			Y: engo.WindowWidth() - 200,
		},
	}

	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(
				&h.text1.BasicEntity,
				&h.text1.RenderComponent,
				&h.text1.SpaceComponent)
		}
	}

	// text2
	h.text2 = Text{BasicEntity: ecs.NewBasic()}
	h.text2.RenderComponent.Drawable = common.Text{
		Font: fnt,
		Text: "click on an element",
	}
	h.text2.SetShader(common.TextHUDShader)
	h.text2.RenderComponent.SetZIndex(1001)
	h.text2.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{
			X: 0,
			Y: engo.WindowWidth() - 180,
		},
	}

	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(
				&h.text2.BasicEntity,
				&h.text2.RenderComponent,
				&h.text2.SpaceComponent)
		}
	}

	// text3
	h.text3 = Text{BasicEntity: ecs.NewBasic()}
	h.text3.RenderComponent.Drawable = common.Text{
		Font: fnt,
		Text: "to get info",
	}
	h.text3.SetShader(common.TextHUDShader)
	h.text3.RenderComponent.SetZIndex(1001)
	h.text3.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{
			X: 0,
			Y: engo.WindowHeight() - 160,
		},
	}

	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(
				&h.text3.BasicEntity,
				&h.text3.RenderComponent,
				&h.text3.SpaceComponent)
		}
	}

	// text4
	h.text4 = Text{BasicEntity: ecs.NewBasic()}
	h.text4.RenderComponent.Drawable = common.Text{
		Font: fnt,
		Text: "about it.",
	}
	h.text4.SetShader(common.TextHUDShader)
	h.text4.RenderComponent.SetZIndex(1001)
	h.text4.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{
			X: 0,
			Y: engo.WindowHeight() - 140,
		},
	}
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(
				&h.text4.BasicEntity,
				&h.text4.RenderComponent,
				&h.text4.SpaceComponent)
		}
	}

	h.money = Text{BasicEntity: ecs.NewBasic()}
	h.money.RenderComponent.Drawable = common.Text{
		Font: fnt,
		Text: "$0",
	}
	h.money.SetShader(common.TextHUDShader)
	h.money.RenderComponent.SetZIndex(1001)
	h.money.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{
			X: 0,
			Y: engo.WindowHeight() - 40,
		},
	}
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(
				&h.money.BasicEntity,
				&h.money.RenderComponent,
				&h.money.SpaceComponent)
		}
	}

	engo.Mailbox.Listen(HUDTextMessageType, func(m engo.Message) {
		msg, ok := m.(HUDTextMessage)
		if !ok {
			return
		}

		for _, system := range w.Systems() {
			switch sys := system.(type) {
			case *common.MouseSystem:
				sys.Add(
					&msg.BasicEntity,
					&msg.MouseComponent,
					&msg.SpaceComponent,
					nil)
			case *HUDTextSystem:
				sys.Add(
					&msg.BasicEntity,
					&msg.SpaceComponent,
					&msg.MouseComponent,
					msg.Line1,
					msg.Line2,
					msg.Line3,
					msg.Line4,
				)
			}
		}
	})

	engo.Mailbox.Listen(HUDMoneyMessageType, func(m engo.Message) {
		msg, ok := m.(HUDMoneyMessage)
		if !ok {
			return
		}

		h.amount = msg.Amount
		h.updateMoney = true
	})
}

func (h *HUDTextSystem) Add(b *ecs.BasicEntity, s *common.SpaceComponent, m *common.MouseComponent, l1, l2, l3, l4 string) {
	h.entities = append(
		h.entities, HUDTextEntity{b, s, m, l1, l2, l3, l4})
}

func (h *HUDTextSystem) Update(dt float32) {
	for _, e := range h.entities {
		if e.MouseComponent.Clicked {
			txt := h.text1.RenderComponent.Drawable.(common.Text)
			txt.Text = e.Line1
			h.text1.RenderComponent.Drawable = txt

			txt = h.text2.RenderComponent.Drawable.(common.Text)
			txt.Text = e.Line2
			h.text2.RenderComponent.Drawable = txt

			txt = h.text3.RenderComponent.Drawable.(common.Text)
			txt.Text = e.Line3
			h.text3.RenderComponent.Drawable = txt

			txt = h.text4.RenderComponent.Drawable.(common.Text)
			txt.Text = e.Line4
			h.text4.RenderComponent.Drawable = txt
		}
	}

	if h.updateMoney {
		txt := h.money.RenderComponent.Drawable.(common.Text)
		txt.Text = fmt.Sprintf("$%v", h.amount)
		h.money.RenderComponent.Drawable = txt
	}
}

func (h *HUDTextSystem) Remove(basic ecs.BasicEntity) {
	delete := -1
	for index, e := range h.entities {
		if e.BasicEntity.ID() == basic.ID() {
			delete = index
			break
		}
	}

	if delete >= 0 {
		h.entities = append(h.entities[:delete], h.entities[delete+1:]...)
	}
}
