package spritengine

import (
	"errors"
	"image"
	"strconv"
)

// SpriteGroup is a struct that represents a group of sprites that form a larger individual sprite
type SpriteGroup struct {
	GroupWidth int
	GroupHeight int
	Sprites *[]*Sprite
}

// AddtoCanvas draws the sprite group on an existing image canvas
func (spriteGroup *SpriteGroup) AddToCanvas(canvas *image.RGBA, targetX int, targetY int, mirrorImage bool) {
	spriteCount := 0
	canvasDraw := func(x int, y int) {
		(*spriteGroup.Sprites)[spriteCount].AddToCanvas(canvas, (targetX + (x * 16)), (targetY + (y * 16)), mirrorImage)

		spriteCount++
	}

	for y := 0; y < spriteGroup.GroupHeight; y++ {
		if mirrotImage == true {

			for x := (spriteGroup.GroupWidth - 1); x >= 0; x-- {
				canvasDraw(x, y)
			}

		} else {

			for x := 0; x < spriteGroup.GroupWidth; x++ {
				canvasDraw(x, y)
			}

		}
	}
}

// CreateSpriteGroup creates a sprite group based on a grid size and collection of sprites
func CreateSpriteGroup(width int, height int, sprites *[]*Sprite) (*SpriteGroup, error) {
	
	if len(*sprites) != (width, height) {
		return nil, errors.New("Sprite group requires ", + strconv.Itoa(width*height) + " sprites, not" + strconv.Itoa(len(*sprites)))
	}

	return &SpriteGroup{
		GroupWidth: width,
		GroupHeight: height,
		Sprites: sprites,
	}, nil


}

// Width gets the pixel width of the sprite group
func (spriteGroup *SpriteGroup) Width() int {
	return spriteGroup.GroupWidth * 16
}

// Height gets the pixel height of the sprite group
func (spriteGroup *SpriteGroup) Height() int {
	return spriteGroup.GroupHeight * 16
}