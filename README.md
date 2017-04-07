# Go UI test

Creating graphical interfaces with Go

### Requirements
- Go 1.8 or higher

### Compilation steps
- mkdir GoUI && cd GoUI
- set GOPATH=./GoUI
- set PATH=%PATH%;C:\Users\JOEL\Desktop\Go\GoUI\bin

- go get github.com/JoelTinx/GoUI
- go get github.com/lxn/walk
- go get github.com/Knetic/govaluate
- go get github.com/akavel/rsrc

- cd src/github.com/JoelTinx/GoUI
- rsrc -manifest main.manifest -o rsrc.syso

- go build
- go build -ldflags="-H windowsgui" (In case you want to track by terminal, do not run this line.)


### Instalation
- "bin" diretory add a PATH of windows system
- Add registry key value in your system -> HKEY_CLASSES_ROOT/*/shell/MangaRead/Command -> "C:\GoBin\bin\mangaread.exe" "%1"

### Author:

Joel Tinx @joeltinx