package main

import (
	"./util"
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/gxfont"
	"github.com/google/gxui/samples/flags"
)

var(
	White = gxui.Color{1.0, 1.0, 1.0, 1.0}
)

func onAddNewRepositoryButtonClicked() {
	
}

func appMain(driver gxui.Driver) {
	theme := flags.CreateTheme(driver)

	font, err := driver.CreateFont(gxfont.Default, 75)
	if err != nil {
		panic(err)
	}

	window := theme.CreateWindow(1000, 700,	util.NAME + " " + util.VERSION)
	window.SetBackgroundBrush(gxui.CreateBrush(gxui.Gray50))
	
	AddNewRepositoryButton := theme.CreateButton()
	AddNewRepositoryButton.SetText("+")
	AddNewRepositoryButton.OnClick(func(gxui.MouseEvent) {onAddNewRepositoryButtonClicked()})

	label := theme.CreateLabel()
	label.SetFont(font)
	label.SetText(util.NAME)
	label.SetColor(White)
	
	//field := theme.CreateField()

	window.AddChild(label)
	//window.AddChild(field)
	window.AddChild(AddNewRepositoryButton)

	window.OnClose(driver.Terminate)
}

func main() {
	gl.StartDriver(appMain)
	util.MakeAppDir()
}