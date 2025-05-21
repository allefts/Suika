package models

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var GrapeImg, StrawberryImg, OrangeImg, WatermelonImg *ebiten.Image

func init() {
	var err error
	GrapeImg, _, err = ebitenutil.NewImageFromFile("../assets/Grape.png")
	StrawberryImg, _, err = ebitenutil.NewImageFromFile("../assets/Strawberry.png")
	OrangeImg, _, err = ebitenutil.NewImageFromFile("../assets/Orange.png")
	WatermelonImg, _, err = ebitenutil.NewImageFromFile("../assets/Watermelon.png")
	if err != nil {
		log.Fatal(err)
	}
}

type Fruit struct {
	Lvl  int
	Name string
	Val  int
	Img  *ebiten.Image
	X    int
	Y    int
}

func NewFruit(lvl int, name string, val int, img *ebiten.Image) *Fruit {
	return &Fruit{Lvl: lvl, Name: name, Val: val, Img: img, X: 0, Y: 5}
}

func CreateFruit(num int) *Fruit {
	switch num {
	case 0:
		return NewFruit(0, "Grape", 1, GrapeImg)
	case 1:
		return NewFruit(1, "Strawberry", 2, StrawberryImg)
	case 2:
		return NewFruit(2, "Orange", 4, OrangeImg)
	case 3:
		return NewFruit(3, "Watermelon", 8, WatermelonImg)
	}

	return nil
}
