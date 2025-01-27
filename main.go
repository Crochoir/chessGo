package main

import (
	_ "image/png"
	"log"
	"os"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth = 800
	screenHeight = 800
)

var (
	board *ebiten.Image
	piece *ebiten.Image
)

type Game struct {
	turn int
}

func (g *Game) Update() error {
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
	opBoard := &ebiten.DrawImageOptions{}
	screen.DrawImage(board, opBoard)

	opPiece := &ebiten.DrawImageOptions{}
	opPiece.GeoM.Translate(200, 200)
	opPiece.GeoM.Scale(0.75, 0.75)
	screen.DrawImage(piece, opPiece)

}

func (g *Game) Layout(Width, Height int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	boardfile, err := os.Open("board/board.png")
	if err != nil {
		log.Fatal(err)
	}
	defer boardfile.Close()


	img, _, err := ebitenutil.NewImageFromReader(boardfile)
	if err != nil {
		log.Fatal(err)
	}
	board = img

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
