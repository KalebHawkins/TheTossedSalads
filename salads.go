package main

import "github.com/hajimehoshi/ebiten/v2"

type Salad struct {
	rotationRate float64
	x, y         float64
	sprite       *ebiten.Image
	saladSpeed   float64
	isFlipping   bool
}

func (s *Salad) Draw(dst *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(s.sprite.Bounds().Dx()/2), -float64(s.sprite.Bounds().Dy()/2))
	op.GeoM.Rotate(s.rotationRate)
	op.GeoM.Translate(s.x, s.y)

	dst.DrawImage(s.sprite, op)
}

func (s *Salad) Update() {
	s.isFlipping = false
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		s.rotationRate += 50
		s.isFlipping = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		s.y -= s.saladSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		s.x -= s.saladSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		s.y += s.saladSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		s.x += s.saladSpeed
	}
}
