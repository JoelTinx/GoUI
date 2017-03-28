package main

import (
	"log"
	"os/exec"
	"strings"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	var inTE, outTE *walk.TextEdit
	//out, _ := exec.Command("wmic","logicaldisk", "get", "name").Output()

	out, err := exec.Command("wmic", "desktopmonitor", "get", "screenheight", ",", "screenwidth").Output()
	if err != nil {
		panic(err)
	}

	height, err := exec.Command("wmic", "desktopmonitor", "get", "screenwidth").Output()
	if err != nil {
		panic(err)
	}

	width, err := exec.Command("wmic", "desktopmonitor", "get", "screenheight").Output()
	if err != nil {
		panic(err)
	}

	// log.Println(strings.Replace(string(out), "ScreenHeight  ScreenWidth \n", "", 1))

	log.Println(strings.Split(string(height), ` `)[2])
	log.Println(strings.Split(string(width), ` `)[2])
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
