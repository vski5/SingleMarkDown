package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("创造一个window,此处是window的标题")
	w.SetContent(widget.NewLabel("我是标签,属于widget(组件)"))
	w.ShowAndRun()

}
