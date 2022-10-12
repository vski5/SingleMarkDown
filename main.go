package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type myAppConfig struct {
	entryWindow *widget.Entry
	showWindow  *widget.RichText
}

func main() {
	a := app.New()
	w := a.NewWindow("SingleMarkDown")
	//w.SetContent(widget.NewLabel("我是标签,属于widget(组件)"))
	var myApp myAppConfig
	entryWindow, showWindow := myApp.makeUI()
	//自左而右按顺序展示，在容器中
	w.SetContent(container.NewHBox(entryWindow, showWindow))
	//设置大小
	w.Resize(fyne.Size{Height: 800, Width: 800})

	//w.ShowAndRun() //程序会在此停止，直到关闭再运行其他的。

	//展示视窗
	w.Show()
	//展示视窗在屏幕中央
	w.CenterOnScreen()
	//运行app //一个程序只能有一个运行。
	a.Run()
}
func (myApp *myAppConfig) makeUI() (*widget.Entry, *widget.RichText) {
	//编辑窗口
	myAppEntryWindow := widget.NewMultiLineEntry()
	//解析MD文档的展示窗口
	myAppShowWindow := widget.NewRichTextFromMarkdown("")
	//将两个窗口的特质分别赋值给struct，使其实例化。
	myApp.entryWindow = myAppEntryWindow
	myApp.showWindow = myAppShowWindow

	//展示窗口跟踪编辑窗口，实时刷新
	myAppEntryWindow.OnChanged = myAppShowWindow.ParseMarkdown

	return myApp.entryWindow, myApp.showWindow
}
