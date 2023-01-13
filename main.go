package main

import (
	"bytes"
	_ "embed"
	"image"
	"image/color"
	_ "image/png"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
)

//go:embed assets/textures/background.png
var backgroundSprite []byte

//go:embed assets/textures/ball.png
var ballSprite []byte

//go:embed assets/textures/salads.png
var saladsSprite []byte

//go:embed assets/audio/airhorn.mp3
var airHornSound []byte

const (
	scrWidth  = 1139
	scrHeight = 572
)

var (
	red = color.RGBA{255, 0, 0, 255}
)

type Game struct {
	*Salad
	*Ball
	audioContext *audio.Context
	audioPlayer  *audio.Player
	background   *ebiten.Image
}

func (g *Game) DetectCollision() {
	backOfBall := g.Ball.x - g.Ball.radius
	frontOfBall := g.Ball.x + g.Ball.radius
	topOfBall := g.Ball.y - g.Ball.radius
	bottomOfBall := g.Ball.y + g.Ball.radius

	if g.Salad.x+float64(g.Salad.sprite.Bounds().Dx())/2 > backOfBall &&
		g.Salad.x-float64(g.Salad.sprite.Bounds().Dx())/2 < frontOfBall &&
		g.Salad.y+float64(g.Salad.sprite.Bounds().Dy())/2 > topOfBall &&
		g.Salad.y-float64(g.Salad.sprite.Bounds().Dy())/2 < bottomOfBall &&
		g.Salad.isFlipping {

		g.Ball.vel.x += rand.Float64()
		g.Ball.vel.y += rand.Float64()
	}
}

func (g *Game) Update() error {
	g.Salad.Update()
	g.Ball.Update()
	g.DetectCollision()

	if g.Ball.x-g.Ball.radius <= 0 {
		g.audioPlayer.Play()
		g.audioPlayer.Rewind()
	}

	return nil
}

func (g *Game) Draw(scr *ebiten.Image) {
	scr.DrawImage(g.background, nil)
	g.Salad.Draw(scr)
	g.Ball.Draw(scr)
}

func (g *Game) Layout(outWidth, outHeight int) (int, int) {
	return scrWidth, scrHeight
}

func NewGame() *Game {
	backgroundImg, _, err := image.Decode(bytes.NewReader(backgroundSprite))
	if err != nil {
		panic(err)
	}

	ballImg, _, err := image.Decode(bytes.NewReader(ballSprite))
	if err != nil {
		panic(err)
	}

	saladImg, _, err := image.Decode(bytes.NewReader(saladsSprite))
	if err != nil {
		panic(err)
	}

	sound, err := mp3.DecodeWithoutResampling(bytes.NewReader(airHornSound))
	if err != nil {
		panic(err)
	}

	ctx := audio.NewContext(48000)
	p, err := ctx.NewPlayer(sound)
	if err != nil {
		panic(err)
	}

	return &Game{
		background: ebiten.NewImageFromImage(backgroundImg),
		Salad:      &Salad{rotationRate: 0, x: scrWidth / 3, y: scrHeight / 3, sprite: ebiten.NewImageFromImage(saladImg), saladSpeed: 10, isFlipping: false},
		Ball: &Ball{
			x:      scrWidth / 2,
			y:      scrHeight / 2,
			radius: 10,
			vel:    Vec2D{0, 0},
			speed:  5,
			sprite: ebiten.NewImageFromImage(ballImg),
		},
		audioContext: ctx,
		audioPlayer:  p,
	}
}

func main() {
	rand.Seed(time.Now().Unix())

	ebiten.SetWindowSize(scrWidth, scrHeight)
	ebiten.SetWindowTitle("The Tossed Salads")

	if err := ebiten.RunGame(NewGame()); err != nil {
		panic(err)
	}
}
