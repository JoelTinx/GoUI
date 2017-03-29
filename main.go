package main

import (
	"os/exec"
	"strconv"
	"strings"
	"unicode"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	//var inTE, outTE *walk.TextEdit
	var wv *walk.WebView

	width, _ := exec.Command("wmic", "desktopmonitor", "get", "screenwidth").Output()
	height, _ := exec.Command("wmic", "desktopmonitor", "get", "screenheight").Output()

	h, _ := strconv.Atoi(strings.Replace(strings.TrimSpace(stringMinifier(string(height))), "ScreenHeight", "", 1))
	w, _ := strconv.Atoi(strings.Replace(strings.TrimSpace(stringMinifier(string(width))), "ScreenWidth", "", 1))

	MainWindow{
		Title:   "MangaRead",
		MinSize: Size{Width: w - 10, Height: h - 50},
		Layout:  VBox{},
		Children: []Widget{
			WebView{
				AssignTo: &wv,
				Name:     "wv",
				URL:      "http://127.0.0.1:8080", // Ojo esta seccion todavía pertenece a otro proyecto
			},
		},
	}.Run()
}

// @write by: Kim Ilyong (http://intogooglego.blogspot.pe)
func stringMinifier(in string) (out string) {
	white := false
	for _, c := range in {
		if unicode.IsSpace(c) {
			if !white {
				out = out + ""
			}
			white = true
		} else {
			out = out + string(c)
			white = false
		}
	}
	return
}
