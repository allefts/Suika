package main

import (
	"fmt"
	"image"
	"math/rand/v2"

	"github.com/allefts/suika/models"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 300
	screenHeight = 300
	frameWidth   = 32
	frameHeight  = 32
)

type Game struct {
	Queue        []*models.Fruit
	PlayedFruits []*models.Fruit
	Colliders    []*image.Rectangle
	Score        int
	MouseX       int
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Update() error {
	mx, _ := ebiten.CursorPosition() //Mouse Location
	g.MouseX = mx                    //Sets game mouseX position
	g.dropCurrFruit()                //Handles drop of fruit

	if len(g.PlayedFruits) != 0 {
		lastDroppedFruit := g.PlayedFruits[len(g.PlayedFruits)-1]
		for _, collider := range g.Colliders {
			if collider.Overlaps(image.Rect(lastDroppedFruit.X, lastDroppedFruit.Y, lastDroppedFruit.X+16, lastDroppedFruit.Y+32)) {
				fmt.Println("Fruit Collision")
			}
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	drawCurrFruit(screen, g.Queue[0].Img, g.MouseX)       //Draw Top Fruit
	drawPlayedFruits(screen, g.PlayedFruits, g.Colliders) //Draw Played Fruits
}

func (g *Game) dropCurrFruit() {
	clicked := inpututil.IsMouseButtonJustReleased(0) //Mouse was clicked
	if clicked {
		g.Queue[0].X = g.MouseX
		g.Queue[0].Y = 5

		g.PlayedFruits = append(g.PlayedFruits, g.Queue[0])                                     //Add to played fruits
		newCollider := image.Rect(g.Queue[0].X, g.Queue[0].Y, g.Queue[0].X+32, g.Queue[0].Y+32) //New collider
		g.Colliders = append(g.Colliders, &newCollider)                                         //Append to colliders

		g.Queue = g.Queue[1:]                                       //Remove from Queue
		g.Queue = append(g.Queue, models.CreateFruit(rand.IntN(4))) //Append new fruit to queue
	}

}

func drawPlayedFruits(screen *ebiten.Image, playedFruits []*models.Fruit, colliders []*image.Rectangle) {
	for _, fruit := range playedFruits {
		if fruit.Y < (screenHeight - frameHeight) {
			fruit.Y += 1
		}

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(-frameWidth/2), 0)
		op.GeoM.Translate(float64(fruit.X), float64(fruit.Y))
		screen.DrawImage(fruit.Img, op)
	}
}

// Draws the image sent at the top of the screen
func drawCurrFruit(screen *ebiten.Image, img *ebiten.Image, mx int) {
	if mx > screenWidth-frameWidth/2 { //X-axis collision
		mx = screenWidth - frameWidth/2
	} else if mx < frameWidth/2 {
		mx = 0 + frameWidth/2
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(-frameWidth)/2, 0) // Move origin to center of image
	op.GeoM.Translate(float64(mx), 0)
	screen.DrawImage(img, op)
}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Suika Game")

	game := &Game{
		Queue:        []*models.Fruit{models.CreateFruit(rand.IntN(4)), models.CreateFruit(rand.IntN(4)), models.CreateFruit(rand.IntN(4)), models.CreateFruit(rand.IntN(4)), models.CreateFruit(rand.IntN(4))},
		PlayedFruits: []*models.Fruit{},
		Colliders:    []*image.Rectangle{},
		Score:        0,
		MouseX:       0,
	}

	ebiten.SetVsyncEnabled(true)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
