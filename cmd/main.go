package main

import (
	_ "image/png"
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
	Score        int
	MouseX       int
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Update() error {
	//Keep track of mouse location
	mx, _ := ebiten.CursorPosition()
	g.MouseX = mx
	g.dropCurrFruit()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//Draw top fruit
	drawCurrFruit(screen, g.Queue[0].Img, g.MouseX)
	//Draw played fruits
	drawPlayedFruits(screen, g.PlayedFruits)
}

func (g *Game) dropCurrFruit() {
	//Keep track of mouse click status
	clicked := inpututil.IsMouseButtonJustReleased(0)
	if clicked {
		g.Queue[0].X = g.MouseX
		g.PlayedFruits = append(g.PlayedFruits, g.Queue[0])
		g.Queue = g.Queue[1:]
		g.Queue = append(g.Queue, models.CreateFruit(rand.IntN(4)))
	}

}

func drawPlayedFruits(screen *ebiten.Image, playedFruits []*models.Fruit) {
	for _, fruit := range playedFruits {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(-frameWidth/2), 0)
		op.GeoM.Translate(float64(fruit.X), screenHeight-frameHeight)
		screen.DrawImage(fruit.Img, op)
	}
}

// Draws the image sent at the top of the screen
func drawCurrFruit(screen *ebiten.Image, img *ebiten.Image, mx int) {
	//X-axis collision
	if mx > screenWidth-frameWidth/2 {
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
		Score:        0,
		PlayedFruits: []*models.Fruit{},
	}

	ebiten.SetVsyncEnabled(true)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
