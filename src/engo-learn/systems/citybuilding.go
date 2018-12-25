package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"fmt"
	"log"
)

type City struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}
type CityBuildingSystem struct {
	world        *ecs.World
	mouseTracker MouseTracker
}

type MouseTracker struct {
	ecs.BasicEntity
	common.MouseComponent
}

// 从场景内删除实体时触发
func (*CityBuildingSystem) Remove(ecs.BasicEntity) {

}

// 每帧刷新
func (cb *CityBuildingSystem) Update(dt float32) {
	if engo.Input.Button("AddCity").JustPressed() {
		fmt.Println("一位路过的二五仔按下了F1")

		// 实例一个基础实体
		city := City{BasicEntity: ecs.NewBasic()}
		// 空间组件设置位置与大小, 位置使用鼠标所在坐标
		city.SpaceComponent = common.SpaceComponent{
			Position: engo.Point{
				cb.mouseTracker.MouseX,
				cb.mouseTracker.MouseY},
			Width:  30,
			Height: 64,
		}

		// 加载一个雪碧
		texture, err := common.LoadedSprite("textures/city.png")
		if err != nil {
			log.Println("Unable to load texture: " + err.Error())
		}

		// 渲染组件绘制内容与大小
		city.RenderComponent = common.RenderComponent{
			Drawable: texture,
			Scale:    engo.Point{0.1, 0.1},
		}

		// 向世界中添加实体
		for _, system := range cb.world.Systems() {
			switch sys := system.(type) {
			case *common.RenderSystem:
				sys.Add(
					&city.BasicEntity,
					&city.RenderComponent,
					&city.SpaceComponent)
			}
		}
	}
}

// 初始化一个系统
func (cb *CityBuildingSystem) New(w *ecs.World) {
	cb.world = w
	fmt.Println("实例CityBuildingSystem加入场景中")

	cb.mouseTracker.BasicEntity = ecs.NewBasic()
	cb.mouseTracker.MouseComponent = common.MouseComponent{
		Track: true,
	}

	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.MouseSystem:
			sys.Add(
				&cb.mouseTracker.BasicEntity,
				&cb.mouseTracker.MouseComponent,
				nil,
				nil)
		}
	}
}
