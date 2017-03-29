package main

import (
	"fmt"
	//"log"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	var inTE, outTE *walk.TextEdit
	//out, _ := exec.Command("wmic","logicaldisk", "get", "name").Output()

	regex, err := regexp.Compile("\n")
	if err != nil {
		return
	}

	width, err := exec.Command("wmic", "desktopmonitor", "get", "screenwidth").Output()
	if err != nil {
		panic(err)
	}

	height, err := exec.Command("wmic", "desktopmonitor", "get", "screenheight").Output()
	if err != nil {
		panic(err)
	}

	h, _ := strconv.Atoi(strings.Replace(strings.Replace(regex.ReplaceAllString(string(height), ""), "ScreenHeight", "", 1), " ", "", -1))
	w, _ := strconv.Atoi(strings.Replace(strings.Replace(regex.ReplaceAllString(string(width), ""), "ScreenWidth", "", 1), " ", "", -1))

	fmt.Println(strings.Replace(strings.Replace(regex.ReplaceAllString(string(height), ""), "ScreenHeight", "", 1), " ", "", -1))
	fmt.Println(strings.Replace(strings.Replace(regex.ReplaceAllString(string(width), ""), "ScreenWidth", "", 1), " ", "", -1))

	MainWindow{
		Title:   "SCREAMO",
		MinSize: Size{600, 400},
		Layout:  VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &inTE, Text: strconv.Itoa(h) + ":" + strconv.Itoa(w)},
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
