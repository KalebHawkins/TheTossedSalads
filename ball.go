package main

import "github.com/hajimehoshi/ebiten/v2"

type Vec2D struct {
	x, y float64
}

type Ball struct {
	x, y   float64
	radius float64
	vel    Vec2D
	speed  float64
	sprite *ebiten.Image
}

func (b *Ball) Draw(dst *ebiten.Image) {
	// ebitenutil.DrawCircle(dst, b.x, b.y, b.radius, color.White)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(b.sprite.Bounds().Dx()/2), -float64(b.sprite.Bounds().Dy()/2))
	op.GeoM.Scale(0.08, 0.08)
	op.GeoM.Translate(b.x, b.y)
	dst.DrawImage(b.sprite, op)
}

func (b *Ball) Update() {

	b.x += b.vel.x
	b.y += b.vel.y

	if b.x+b.radius > scrWidth || b.x-b.radius < 0 {
		b.vel.x = -b.vel.x
	}
	if b.y+b.radius > scrHeight || b.y+b.radius < 0 {
		b.vel.y = -b.vel.y
	}
}
