#!/bin/sh

CGO_ENABLED=1 \
GOOS=windows \
GOARCH=amd64 \
CC="zig cc -target x86_64-windows-gnu" \
CXX="zig c++ -target x86_64-windows-gnu" \
go build -trimpath -ldflags='-H=windowsgui' -o dist/windows-amd64/caster-gui.exe src/main.go
