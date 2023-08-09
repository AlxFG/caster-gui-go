package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	currdir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	caster := currdir + "\\cccaster.v3.1.exe"
	const (
		screenWidth  = 1280
		screenHeight = 720
	)

	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.SetConfigFlags(rl.FlagWindowUndecorated)
	rl.InitWindow(screenWidth, screenHeight, "caster-gui")

	var (
		mousePosition  = rl.Vector2{0, 0}
		windowPosition = rl.Vector2{500, 200}
		panOffset      = mousePosition
		dragWindow     = false
	)

	rl.SetWindowPosition(int(windowPosition.X), int(windowPosition.Y))

	exitWindow := false

	rl.SetTargetFPS(60)
	var host string = ""
	var spectate string = ""
	host_state := true
	spectate_state := false

	var host_bounds = rl.Rectangle{200, 200, 300, 30}
	var spectate_bounds = rl.Rectangle{200, 300, 300, 30}


	for !exitWindow && !rl.WindowShouldClose() {

		mousePosition = rl.GetMousePosition()

		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			if (rl.CheckCollisionPointRec(mousePosition, rl.Rectangle{0, 0, screenWidth, 20})) {
				dragWindow = true
				panOffset = mousePosition
			}
		}

		if dragWindow {
			windowPosition.X += (mousePosition.X - panOffset.X)
			windowPosition.Y += (mousePosition.Y - panOffset.Y)

			if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
				dragWindow = false
			}

			rl.SetWindowPosition(int(windowPosition.X), int(windowPosition.Y))
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		exitWindow = rg.WindowBox(rl.Rectangle{0, 0, screenWidth, screenHeight}, "#198# CASTER-GUI")

		rl.DrawText(fmt.Sprintf("IP:Port"), 10, 200, 30, rl.DarkGray)
		rg.TextBox(host_bounds, &host, 1024, host_state)
		if rg.Button(rl.Rectangle{500, 200, 80, 30}, "Connect/Host") {
			cmd := exec.Command(caster, "-n", host)
			err := cmd.Run()

			if err != nil {
				log.Fatal(err)
			}
			host = ""
		}
		if rg.Button(rl.Rectangle{580, 200, 60, 30}, "Paste") {
			cmd := exec.Command(caster, "-n", rl.GetClipboardText())
			err := cmd.Run()
			if err != nil {
				log.Fatal(err)
			}
		}
		if rl.CheckCollisionPointRec(mousePosition, host_bounds) && rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			host_state = true
			spectate_state = false
		}

		rl.DrawText(fmt.Sprintf("Spectate"), 10, 300, 30, rl.DarkGray)
		rg.TextBox(spectate_bounds, &spectate, 1024,
			spectate_state)
		if rg.Button(rl.Rectangle{500, 300, 80, 30}, "Connect") {
			cmd := exec.Command(caster, "-ns", spectate)
			err := cmd.Run()
			if err != nil {
				log.Fatal(err)
			}
			spectate = ""
		}
		if rg.Button(rl.Rectangle{580, 300, 60, 30}, "Paste") {
			cmd := exec.Command(caster, "-ns", rl.GetClipboardText())
			err := cmd.Run()
			if err != nil {
				log.Fatal(err)
			}
		}
		if rl.CheckCollisionPointRec(mousePosition, spectate_bounds) && rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			spectate_state = true
			host_state = false
		}

		rl.DrawText(fmt.Sprintf("Offline"), 10, 400, 30, rl.DarkGray)
		if rg.Button(rl.Rectangle{200, 400, 80, 30}, "Training") {
			cmd := exec.Command(caster, "-not")
			err := cmd.Run()
			if err != nil {
				log.Fatal(err)
			}
		}

		if !rl.CheckCollisionPointRec(mousePosition, host_bounds) &&
			!rl.CheckCollisionPointRec(mousePosition, spectate_bounds) &&
			rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			host_state = false
			spectate_state = false
		}
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
