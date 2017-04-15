## instalation steps

mkdir GoUI && cd GoUI

set GOPATH=./GoUI

set PATH=%PATH%;C:\Users\JOEL\Desktop\Go\GoUI\bin

go get github.com/JoelTinx/GoUI

go get github.com/lxn/walk

go get github.com/Knetic/govaluate

go get github.com/akavel/rsrc

cd src/github.com/JoelTinx/GoUI

rsrc -manifest main.manifest -o rsrc.syso

go build

go build -ldflags="-H windowsgui"