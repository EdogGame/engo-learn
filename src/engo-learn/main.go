package main

import (
	"engo.io/engo"
)

// 一个场景
type myScene struct {
}

func (*myScene) Type() string {
	return "myGame"
}
func (*myScene) Preload() {

}
func (*myScene) Setup(engo.Updater) {

}

func main() {
	opts := engo.RunOptions{
		Title:      "hello engo-learn", //设置窗口标题
		Width:      400,                // 设置窗口宽度
		Height:     400,                // 设置窗口高度
		Fullscreen: false,              // 开启全屏
	}

	// 运行引擎
	engo.Run(opts, &myScene{})
}
