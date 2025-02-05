package main

import (
	_ "image/png"
	"log"
	"os"

	"github.com/Crochoir/chessGo/board"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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

	opPiece := &ebiten.DrawImageOptions{}
	opPiece.GeoM.Translate(15, 900)
	opPiece.GeoM.Scale(0.75, 0.75)
	screen.DrawImage(piece, opPiece)

}

func (g *Game) Layout(Width, Height int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	if err := gameBoard.Initialize("board/board.png"); err != nil {
		log.Fatal(err)
	}

	pieceFile, err := os.Open("pieces/black-rook.png")
	if err != nil {
		log.Fatal(err)
	}
	defer pieceFile.Close()

	imgPiece, _, err := ebitenutil.NewImageFromReader(pieceFile)
	if err != nil {
		log.Fatal(err)
	}
	piece = imgPiece

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("chess")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
