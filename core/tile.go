package core

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
)

// TileSet tile set
type TileSet struct {
	img           *ebiten.Image
	width, height int
}

func NewTileSet() *TileSet {
	t := &TileSet{
		width:  16, // tile unit size
		height: 16,
	}

	// load tile set image
	f, err := ebitenutil.OpenFile("assets/imgs/tile.png")
	if err != nil {
		return nil
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil
	}
	t.img = ebiten.NewImageFromImage(img)
	return t
}

func (t *TileSet) Image(name string) *ebiten.Image {
	switch name {
	case "player":
		return t.subImage(12, 38)
	case "monster":
		return t.subImage(12, 44)
	case "wall":
		return t.subImage(0, 29)
	case "floor":
		return t.subImage(1, 16)
	default:
		return t.subImage(28, 0)
	}
}

func (t *TileSet) subImage(xIndex, yIndex int) *ebiten.Image {
	x := xIndex * t.width
	y := yIndex * t.width
	return t.img.SubImage(image.Rect(x, y, x+t.width, y+t.height)).(*ebiten.Image)
}
