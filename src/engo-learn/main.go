package main

import (
	"engo-learn/systems"
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"image"
	"image/color"
)

const (
	KeyBoardScrollSpeed = 400
	EdgeScrollSpeed     = 400
	EdgeScrollMargin    = 20
	MouseZoomSpeed      = -0.125
)

type HUD struct {
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

	// 添加渲染系统
	world.AddSystem(&common.RenderSystem{})

	// 添加监控鼠标系统
	world.AddSystem(&common.MouseSystem{})

	kbs := common.NewKeyboardScroller(
		KeyBoardScrollSpeed,
		engo.DefaultHorizontalAxis,
		engo.DefaultVerticalAxis)
	world.AddSystem(kbs)

	world.AddSystem(&common.EdgeScroller{EdgeScrollSpeed, EdgeScrollMargin})
	world.AddSystem(&common.MouseZoomer{MouseZoomSpeed})

	// 添加自定义的城市系统
	world.AddSystem(&systems.CityBuildingSystem{})

	// 实例一个基础信息框
	hud := HUD{BasicEntity: ecs.NewBasic()}
	hud.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{
			0,
			engo.WindowHeight() - 200,
		},
		Width:  200,
		Height: 200,
	}
	// 处理基础信息框纹理
	hudImage := image.NewUniform(color.RGBA{205, 205, 205, 255})
	hudNRGBA := common.ImageToNRGBA(hudImage, 200, 200)
	hudImageObj := common.NewImageObject(hudNRGBA)
	hudTexture := common.NewTextureSingle(hudImageObj)

	hud.RenderComponent = common.RenderComponent{
		Drawable: hudTexture,
		Scale:    engo.Point{1, 1},
		Repeat:   common.Repeat,
	}
	// 添加着色器
	hud.RenderComponent.SetShader(common.HUDShader)
	hud.RenderComponent.SetZIndex(1)

	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(
				&hud.BasicEntity,
				&hud.RenderComponent,
				&hud.SpaceComponent)
		}
	}

	// 设置场景背景颜色为白色
	common.SetBackground(color.White)
}

func main() {
	opts := engo.RunOptions{
		Title:          "hello engo-learn", // 设置窗口标题
		Width:          800,                // 设置窗口宽度
		Height:         800,                // 设置窗口高度
		Fullscreen:     false,              // 开启全屏
		AssetsRoot:     "resources",        // 设置资源目录
		StandardInputs: true,               // 是否开启标准输入
	}

	// 运行引擎
	engo.Run(opts, &myScene{})
}
