package main

import (
	"os/exec"
	"strings"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	var inTE, outTE *walk.TextEdit
	//out, _ := exec.Command("wmic","logicaldisk", "get", "name").Output()

	out, err := exec.Command("wmic", "desktopmonitor", "get", "screenheight").Output()
	if err != nil {
		panic(err)
	}

	//screenwidth

	MainWindow{
		Title:   "SCREAMO",
		MinSize: Size{600, 400},
		Layout:  VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &inTE, Text: string(out)},
					TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},
			PushButton{
				Text: "SCREAM",
				OnClicked: func() {
					outTE.SetText(strings.ToUpper(inTE.Text()))
				},
			},
		},
	}.Run()
}
