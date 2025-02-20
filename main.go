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
	piece     *ebiten.Image
)

type Game struct {
	pieces []*pieces.Piece
}

func NewGame() *Game {
	return &Game{pieces.CreatePieces()}
}

func (g *Game) Update() error {
	mouseX, mouseY := ebiten.CursorPosition()

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		for _, piece := range g.pieces {
			if piece.Contains(mouseX, mouseY) {
				piece.IsDragging = true
				piece.DragOffset.X = float64(mouseX) - piece.Position.X
				piece.DragOffset.Y = float64(mouseY) - piece.Position.Y

			}
		}
	} else {
		// Release drag when mouse button is released
		for _, piece := range g.pieces {
			piece.IsDragging = false
		}
	}

	// Move the dragging piece
	for _, piece := range g.pieces {
		if piece.IsDragging {
			piece.Position.X = float64(mouseX) - piece.DragOffset.X
			piece.Position.Y = float64(mouseY) - piece.DragOffset.Y
		}
	}
	if len(g.pieces) > 0 {
		g.pieces[0].Move(100, 100)
	}
	if len(g.pieces) > 1 {
		g.pieces[1].Capture()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	gameBoard.Draw(screen)
	for _, piece := range g.pieces {
		piece.Draw(screen)
	}

}

func (g *Game) Layout(Width, Height int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	gameBoard.Initialize(screenWidth, screenHeight)

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("chess")

	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
