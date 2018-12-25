package systems

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"fmt"
	"log"
	"math/rand"
	"time"
)

var Spritesheet *common.Spritesheet

var cities = [...][12]int{
	{99, 100, 101,
		454, 269, 455,
		415, 195, 416,
		452, 306, 453,
	},
	{99, 100, 101,
		268, 269, 270,
		268, 269, 270,
		305, 306, 307,
	},
	{75, 76, 77,
		446, 261, 447,
		446, 261, 447,
		444, 298, 445,
	},
	{75, 76, 77,
		407, 187, 408,
		407, 187, 408,
		444, 298, 445,
	},
	{75, 76, 77,
		186, 150, 188,
		186, 150, 188,
		297, 191, 299,
	},
	{83, 84, 85,
		413, 228, 414,
		411, 191, 412,
		448, 302, 449,
	},
	{83, 84, 85,
		227, 228, 229,
		190, 191, 192,
		301, 302, 303,
	},
	{91, 92, 93,
		241, 242, 243,
		278, 279, 280,
		945, 946, 947,
	},
	{91, 92, 93,
		241, 242, 243,
		278, 279, 280,
		945, 803, 947,
	},
	{91, 92, 93,
		238, 239, 240,
		238, 239, 240,
		312, 313, 314,
	},
}

type City struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}
type CityBuildingSystem struct {
	world        *ecs.World
	mouseTracker MouseTracker
	usedTiles    []int
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
			panic(err)
		}

		// 渲染组件绘制内容与大小
		city.RenderComponent = common.RenderComponent{
			Drawable: texture,
			Scale:    engo.Point{0.5, 0.5},
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

	// 注册按键
	//engo.Input.RegisterButton("AddCity", engo.KeyF1)

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

	Spritesheet = common.NewSpritesheetWithBorderFromFile("textures/citySheet.png", 16, 16, 1, 1)
	rand.Seed(time.Now().UnixNano())
}

func (cb *CityBuildingSystem) generateCity() {
	x := rand.Intn(18)
	y := rand.Intn(18)

	t := x + y*18

	for cb.isTileUsed(t) {
		if len(cb.usedTiles) > 300 {
			break
		}

		x = rand.Intn(18)
		y = rand.Intn(18)
		t = x + y*18
	}
	cb.usedTiles = append(cb.usedTiles, t)

	city := rand.Intn(len(cities))
	cityTiles := make([]*City, 0)
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			tile := &City{BasicEntity: ecs.NewBasic()}
			tile.SpaceComponent.Position = engo.Point{
				X: float32(((x+1)*64)+8) + float32(i*16),
				Y: float32((y+1)*64) + float32(j*16),
			}
			tile.RenderComponent.Drawable = Spritesheet.Cell(
				cities[city][i+3*j])
			tile.RenderComponent.SetZIndex(1)
			cityTiles = append(cityTiles, tile)
		}
	}
}
