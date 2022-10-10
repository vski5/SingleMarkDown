package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type myApp struct {
	lab *widget.Label
}

func main() {
	a := app.New()
	w := a.NewWindow("创造一个window,此处是window的标题")
	//w.SetContent(widget.NewLabel("我是标签,属于widget(组件)"))
	var myapp myApp
	lab, btm, entryLine := myapp.changeTitle()
	//自上而下按顺序展示，在容器中
	w.SetContent(container.NewVBox(lab, entryLine, btm))
	//设置大小
	w.Resize(fyne.Size{Height: 500, Width: 500})
	//w.ShowAndRun() //程序在此停止，直到关闭再运行其他的。
	//展示视窗
	w.Show()
	//运行app //一个程序只能有一个运行。
	a.Run()
}
func (myApp *myApp) changeTitle() (*widget.Label, *widget.Button, *widget.Entry) {

	lab := widget.NewLabel("new label")
	myApp.lab = lab

	entryLine := widget.NewEntry()

	btm := widget.NewButton("i'm button", func() {
		myApp.lab.SetText(entryLine.Text)
	})

	return lab, btm, entryLine
}
