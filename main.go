package main

import (
	_ "image/png"
	"log"
<<<<<<< HEAD
	"os"

	"github.com/Crochoir/chessGo/board"
=======

	"github.com/Crochoir/chessGo/board"
	"github.com/Crochoir/chessGo/pieces"
>>>>>>> c7fa942b5252e2d25417aa39e682288a11c76269
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 800
	screenHeight = 800
)

var (
	gameBoard board.Board
<<<<<<< HEAD
	piece     *ebiten.Image
=======
>>>>>>> c7fa942b5252e2d25417aa39e682288a11c76269
)

type Game struct {
	pieces []*pieces.Piece
}

func NewGame() *Game {
	return &Game{pieces.CreatePieces()}
}

func (g *Game) Update() error {
<<<<<<< HEAD
	if err := gameBoard.Update(); err != nil {
		return err
	}
	if g.turn == 1 {
		g.turn = 2
	} else if g.turn == 2 {
		g.turn = 1
=======
	mouseX, mouseY := ebiten.CursorPosition()

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		for _, piece := range g.pieces {
			if piece.Contains(mouseX, mouseY) {
				piece.IsDragging = true
				piece.DragOffset.X = float64(mouseX) - piece.Position.X
				piece.DragOffset.Y = float64(mouseY) - piece.Position.Y

			}
		}
>>>>>>> c7fa942b5252e2d25417aa39e682288a11c76269
	} else {
		// Release drag when mouse button is released
		for _, piece := range g.pieces {
			piece.IsDragging = false
		}
	}
<<<<<<< HEAD
=======

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

>>>>>>> c7fa942b5252e2d25417aa39e682288a11c76269
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	gameBoard.Draw(screen)
<<<<<<< HEAD

	opPiece := &ebiten.DrawImageOptions{}
	opPiece.GeoM.Translate(15, 900)
	opPiece.GeoM.Scale(0.75, 0.75)
	screen.DrawImage(piece, opPiece)
=======
	for _, piece := range g.pieces {
		piece.Draw(screen)
	}
>>>>>>> c7fa942b5252e2d25417aa39e682288a11c76269

}

func (g *Game) Layout(Width, Height int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
<<<<<<< HEAD
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
=======
	gameBoard.Initialize(screenWidth, screenHeight)
>>>>>>> c7fa942b5252e2d25417aa39e682288a11c76269

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("chess")

	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
