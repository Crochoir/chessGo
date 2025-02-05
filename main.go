package main

import (
	_ "image/png"
	"log"

	"github.com/Crochoir/chessGo/board"
	"github.com/Crochoir/chessGo/pieces"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 800
	screenHeight = 800
)

var (
	gameBoard board.Board
	piece     pieces.Piece
)

type Game struct {
	turn int
}

func (g *Game) Update() error {
	if err := gameBoard.Update(); err != nil {
		return err
	}
	if g.turn == 1 {
		g.turn = 2
	} else if g.turn == 2 {
		g.turn = 1
	} else {
		g.turn = 1
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	gameBoard.Draw(screen)
	piece.Draw(screen)

}

func (g *Game) Layout(Width, Height int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	if err := gameBoard.Initialize("board/board.png"); err != nil {
		log.Fatal(err)
	}

	if err := piece.Initialize("pieces/black-rook.png", 10, 0); err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("chess")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
