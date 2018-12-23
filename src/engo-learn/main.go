package main

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"image/color"
	"log"
)

type City struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

// 一个场景
type myScene struct {
}

func (*myScene) Type() string {
	return "myGame"
}

// preload函数内处理资源
func (*myScene) Preload() {
	// 设置资源目录
	// engo.Files.SetRoot("resources")

	// 加载一个资源文件
	engo.Files.Load("textures/city.png")
}

// setup函数内添加实体与系统设置
func (*myScene) Setup(u engo.Updater) {
	world, _ := u.(*ecs.World)
	world.AddSystem(&common.RenderSystem{})

	// 实例一个基础实体
	city := City{BasicEntity: ecs.NewBasic()}
	// 空间组件设置位置与大小
	city.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{10, 10},
		Width:    303,
		Height:   641,
	}

	// 加载一个雪碧
	texture, err := common.LoadedSprite("textures/city.png")
	if err != nil {
		log.Println("Unable to load texture: " + err.Error())
	}

	// 渲染组件绘制内容与大小
	city.RenderComponent = common.RenderComponent{
		Drawable: texture,
		Scale:    engo.Point{1, 1},
	}

	// 向世界中添加实体
	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&city.BasicEntity, &city.RenderComponent, &city.SpaceComponent)
		}
	}

	// 设置场景背景颜色为白色
	common.SetBackground(color.White)
}

func main() {
	opts := engo.RunOptions{
		Title:      "hello engo-learn", // 设置窗口标题
		Width:      400,                // 设置窗口宽度
		Height:     400,                // 设置窗口高度
		Fullscreen: false,              // 开启全屏
		AssetsRoot: "resources",        // 设置资源目录
	}

	// 运行引擎
	engo.Run(opts, &myScene{})
}
