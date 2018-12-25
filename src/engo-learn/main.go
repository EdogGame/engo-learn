package main

import (
	. "engo-learn/systems"
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"image/color"
)

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
	// 注册按键
	engo.Input.RegisterButton("AddCity", engo.KeyF1)

	world, _ := u.(*ecs.World)

	// 添加渲染系统
	world.AddSystem(&common.RenderSystem{})

	// 添加监控鼠标系统
	world.AddSystem(&common.MouseSystem{})

	// 添加自定义的城市系统
	world.AddSystem(&CityBuildingSystem{})

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
