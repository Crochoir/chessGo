package pieces

import (
	_ "image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Coords represents a Position on the board.
type Coords struct {
	X float64
	Y float64
}

// Rank represents the type of chess piece.
type Rank struct {
	rank   string
	motion *Coords
	image  *ebiten.Image
}

// Piece represents a chess piece with a rank and Position.
type Piece struct {
	rank       *Rank
	Position   *Coords
	IsDragging bool
	DragOffset *Coords
	captured   bool
	hasMoved   bool
}

func (p *Piece) Initialize(imagePath string, rankName string, motionX, motionY, X, Y float64) error {
	piecefile, err := os.Open(imagePath)
	if err != nil {
		log.Printf("Error opening file %s: %v", imagePath, err)
		return err
	}
	defer piecefile.Close()

	img, _, err := ebitenutil.NewImageFromReader(piecefile)
	if err != nil {
		log.Printf("Error loading image from %s: %v", imagePath, err)
		return err
	}

	p.rank = &Rank{
		rank:   rankName,
		motion: &Coords{X: motionX, Y: motionY},
		image:  img,
	}
	p.rank.image = img
	p.Position = &Coords{X: X, Y: Y}

	p.captured = false
	p.hasMoved = false

	//log.Printf("Loaded image %s successfully", imagePath)
	return nil
}

// Draw renders the piece on the screen.
func (p *Piece) Draw(screen *ebiten.Image) {
	if p.rank == nil || p.rank.image == nil {
		log.Printf("Skipping draw for piece at (%f, %f) because image is nil", p.Position.X, p.Position.Y)
		return // Prevents drawing if the image isn't loaded properlY
	}
	//log.Printf("Drawing piece at (%f, %f)", p.Position.X, p.Position.Y)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.75, 0.75)
	op.GeoM.Translate(p.Position.X, p.Position.Y)
	screen.DrawImage(p.rank.image, op)
}

func (p *Piece) Move(newX, newY float64) {
	p.Position.X = newX
	p.Position.Y = newY
	p.hasMoved = true
}

func (p *Piece) Capture() {
	p.captured = true
}

// Update handles any piece updates.
func (p *Piece) Update() error {
	return nil
}

func CreatePieces() []*Piece {
	pieceData := []struct {
		imagePath string
		rank      string
		motionX   float64
		motionY   float64
		X         float64
		Y         float64
	}{
		// Black pieces
		{"pieces/black-rook.png", "BRook", 3, 2, 0, 0},
		{"pieces/black-knight.png", "BKnight", 2.5, 2.5, 100, 0},
		{"pieces/black-bishop.png", "BBishop", 3, 3, 200, 0},
		{"pieces/black-queen.png", "BQueen", 3, 3, 300, 0},
		{"pieces/black-king.png", "BKing", 2, 2, 400, 0},
		{"pieces/black-bishop.png", "BBishop", 3, 3, 500, 0},
		{"pieces/black-knight.png", "BKnight", 2.5, 2.5, 600, 0},
		{"pieces/black-rook.png", "BRook", 3, 2, 700, 0},
		{"pieces/black-pawn.png", "BPawn", 1, 2, 10, 100},
		{"pieces/black-pawn.png", "BPawn", 1, 2, 100, 100},
		{"pieces/black-pawn.png", "BPawn", 1, 2, 200, 100},
		{"pieces/black-pawn.png", "BPawn", 1, 2, 300, 100},
		{"pieces/black-pawn.png", "BPawn", 1, 2, 400, 100},
		{"pieces/black-pawn.png", "BPawn", 1, 2, 500, 100},
		{"pieces/black-pawn.png", "BPawn", 1, 2, 600, 100},
		{"pieces/black-pawn.png", "BPawn", 1, 2, 700, 100},

		// White pieces
		{"pieces/white-rook.png", "WRook", 3, 2, 0, 700},
		{"pieces/white-knight.png", "WKnight", 2.5, 2.5, 100, 700},
		{"pieces/white-bishop.png", "WBishop", 3, 3, 200, 700},
		{"pieces/white-queen.png", "WQueen", 3, 3, 300, 700},
		{"pieces/white-king.png", "WKing", 2, 2, 400, 700},
		{"pieces/white-bishop.png", "WBishop", 3, 3, 500, 700},
		{"pieces/white-knight.png", "WKnight", 2.5, 2.5, 600, 700},
		{"pieces/white-rook.png", "WRook", 3, 2, 700, 700},
		{"pieces/white-pawn.png", "WPawn", 1, 2, 10, 600},
		{"pieces/white-pawn.png", "WPawn", 1, 2, 100, 600},
		{"pieces/white-pawn.png", "WPawn", 1, 2, 200, 600},
		{"pieces/white-pawn.png", "WPawn", 1, 2, 300, 600},
		{"pieces/white-pawn.png", "WPawn", 1, 2, 400, 600},
		{"pieces/white-pawn.png", "WPawn", 1, 2, 500, 600},
		{"pieces/white-pawn.png", "WPawn", 1, 2, 600, 600},
		{"pieces/white-pawn.png", "WPawn", 1, 2, 700, 600},
	}

	var pieces []*Piece
	for _, data := range pieceData {
		p := &Piece{}
		err := p.Initialize(data.imagePath, data.rank, data.motionX, data.motionY, data.X, data.Y)
		if err != nil {
			log.Fatalf("Failed to initialize piece %s: %v", data.rank, err)
		}
		pieces = append(pieces, p)
	}
	return pieces
}

func (p *Piece) Contains(x, y int) bool {
	width, height := p.rank.image.Bounds().Dx(), p.rank.image.Bounds().Dy()
	posX, posY := int(p.Position.X), int(p.Position.Y)

	inside := x >= posX && x <= posX+width &&
		y >= posY && y <= posY+height

	if inside {
		//log.Printf("Piece %s clicked at (%d, %d)", p.rank.rank, x, y)
	}
	return inside

}
