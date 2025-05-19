package main

import (
	"image"
	"image/color"
	"math/rand/v2"

	"github.com/allefts/suika/models"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
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
		for i := 0; i < len(g.PlayedFruits)-1; i++ {
			if g.Colliders[i].Overlaps(image.Rect(lastDroppedFruit.X-16, lastDroppedFruit.Y, lastDroppedFruit.X+16, lastDroppedFruit.Y+32)) {
				// fmt.Println(lastDroppedFruit.Name, g.PlayedFruits[i].Name)
			}
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	drawCurrFruit(screen, g.Queue[0].Img, g.MouseX)       //Draw Top Fruit
	drawPlayedFruits(screen, g.PlayedFruits, g.Colliders) //Draw Played Fruits

	for _, collider := range g.Colliders {
		vector.StrokeRect(screen, float32(collider.Min.X), float32(collider.Min.Y), float32(collider.Dx()), float32(collider.Dy()), 1.0, color.RGBA{255, 0, 0, 255}, true)
	}
}

func (g *Game) dropCurrFruit() {
	clicked := inpututil.IsMouseButtonJustReleased(0) //Mouse was clicked
	if clicked {
		g.Queue[0].X = g.MouseX
		g.Queue[0].Y = 5

		g.PlayedFruits = append(g.PlayedFruits, g.Queue[0])                                        //Add to played fruits
		newCollider := image.Rect(g.Queue[0].X-16, g.Queue[0].Y, g.Queue[0].X+16, g.Queue[0].Y+32) //New collider
		g.Colliders = append(g.Colliders, &newCollider)                                            //Append to colliders

		g.Queue = g.Queue[1:]                                       //Remove from Queue
		g.Queue = append(g.Queue, models.CreateFruit(rand.IntN(4))) //Append new fruit to queue
	}

}

func drawPlayedFruits(screen *ebiten.Image, playedFruits []*models.Fruit, colliders []*image.Rectangle) {
	for idx, fruit := range playedFruits {
		if fruit.Y < (screenHeight - frameHeight) {
			fruit.Y += 1
		}

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(-frameWidth/2), 0)
		op.GeoM.Translate(float64(fruit.X), float64(fruit.Y))
		screen.DrawImage(fruit.Img, op)

		colliders[idx].Min = image.Point{X: fruit.X - 16, Y: fruit.Y}
		colliders[idx].Max = image.Point{X: fruit.X + 16, Y: fruit.Y + 32}
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
