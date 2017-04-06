// Write by: Joel Tinx
package main

import (
	"archive/zip"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"unicode"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

const templ = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Manga Read</title>
  </head>
  <body style="background-color: #333; text-align: center;">
    <div style="margin: 0 auto;">
      {{ range $key, $value := . }}
        <img src="/public/{{ $value }}" alt="" style="margin: 0 auto;" />
      {{ end }}
    </div>
  </body>
</html>
`

var destino, origen string

func init() {
	// Get variable "temp" windows: storage temporal files
	destino = os.Getenv("temp") + "\\mangaread"
}

func main() {
	// Verify parameter with filename
	if len(os.Args) >= 2 {
		origen = os.Args[1]
	} else {
		log.Println("Ha ocurrido un error interno")
		os.Exit(1)
	}

	// Check and delete files on destination path
	if val, _ := exists(destino); val {
		if err := os.RemoveAll(destino); err != nil {
			log.Println("Ha ocurrido un error interno")
			os.Exit(1)
		}
	}

	// Create directory to save the uncompressed images
	os.MkdirAll(destino, 0777)
	err := Unzip(origen, destino)
	if err != nil {
		log.Println("No se ha podido descomprimir")
		os.Exit(1)
	}
	defer os.RemoveAll(destino)

	// Array storage files name (images uncompressed)
	Images := make([]string, 0)

	//
	filepath.Walk(destino, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && IsImage(path) {
			strx, _ := filepath.Rel(destino, path)
			strx = filepath.ToSlash(strx)
			Images = append(Images, strx)
		}
		return err
	})

	t, err := template.New("Mangarad template").Parse(templ)
	if err != nil {
		log.Println("Error al Renderizar Manga")
		os.Exit(1)
	}

	go RunServer(destino, t, Images)

	//var inTE, outTE *walk.TextEdit
	var frm *walk.MainWindow
	var wv *walk.WebView

	width, err := exec.Command("wmic", "desktopmonitor", "get", "screenwidth").Output()
	if err != nil {
		walk.MsgBox(nil, "Error", "Ola que ase")
	}
	height, _ := exec.Command("wmic", "desktopmonitor", "get", "screenheight").Output()

	h, _ := strconv.Atoi(strings.Replace(strings.TrimSpace(stringMinifier(string(height))), "ScreenHeight", "", 1))
	w, _ := strconv.Atoi(strings.Replace(strings.TrimSpace(stringMinifier(string(width))), "ScreenWidth", "", 1))

	MainWindow{
		Title:    "MangaRead",
		MinSize:  Size{Width: w - 10, Height: h - 50},
		Layout:   VBox{},
		AssignTo: &frm,
		Children: []Widget{
			WebView{
				AssignTo: &wv,
				Name:     "wv",
				URL:      "http://127.0.0.1:3000", // Ojo esta seccion todavía pertenece a otro proyecto
			},
		},
	}.Run()
}

// Async function Go Golang
func RunServer(destino string, tmpl *template.Template, Images []string) {
	fs := http.FileServer(http.Dir(destino))
	http.Handle("/public/", http.StripPrefix("/public/", fs))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, Images)
	})

	// Ejecutar este proceso en paralelo
	http.ListenAndServe(":3000", nil)
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

// Verifica que un archivo sea imagen
func IsImage(path string) bool {
	output := false
	switch strings.ToUpper(filepath.Ext(path)) {
	case ".JPG", ".JPEG", ".PNG":
		output = true
	}
	return output
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func Unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()

	os.MkdirAll(dest, 0755)

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}
