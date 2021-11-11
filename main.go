package main

import (
	_ "image/png"
	"log"
	"saga/core"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	g := core.NewGame()
	ebiten.SetWindowResizable(true)

	ebiten.SetWindowTitle("Tower")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
