package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

var whitePixel *ebiten.Image

type Game struct {
	angle float64
	t     float64
}

func (g *Game) Update() error {
	g.angle += 0.005
	g.t += 0.02
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{50, 50, 50, 255})

	width, height := ebiten.WindowSize()
	screenWidth, screenHeight := float64(width), float64(height)
	cx, cy := screenWidth/2, screenHeight/2
	baseRadius := math.Min(screenWidth, screenHeight) * 0.3
	drawFractal(screen, cx, cy, baseRadius, 0, 3, g.angle, g.t)
}

func drawFractal(screen *ebiten.Image, x, y, radius float64, depth, maxDepth int, baseAngle, time float64) {
	if depth > maxDepth {
		return
	}

	petalcount := 6 + depth*2
	angleOffset := baseAngle + math.Sin(time*0.3+float64(depth)*0.5)*0.8

	for i := 0; i < petalcount; i++ {
		theta := float64(i)*2*math.Pi/float64(petalcount) + angleOffset
		nx := x + radius*math.Sin(theta)
		ny := y + radius*math.Cos(theta)

		r := float32(1.0)                                   //red
		g := float32(0.5 + 0.5*math.Sin(time+float64(i)+2)) //green
		b := float32(0.0)                                   //no blue
		alpha := float32(1.0)                               //opaque

		vertices := []ebiten.Vertex{
			{
				DstX: float32(x),
				DstY: float32(y),
				SrcX: 0, SrcY: 0,
				ColorR: r, ColorG: g, ColorB: b, ColorA: alpha,
			},
			{
				DstX: float32(nx),
				DstY: float32(ny),
				SrcX: 0, SrcY: 0,
				ColorR: r, ColorG: g, ColorB: b, ColorA: alpha,
			},
			{
				DstX: float32(x + radius*0.5*math.Cos(theta+math.Pi/2)),
				DstY: float32(y + radius*0.5*math.Sin(theta+math.Pi/2)),
				SrcX: 0, SrcY: 0,
				ColorR: r, ColorG: g, ColorB: b, ColorA: alpha,
			},
		}

		indices := []uint16{0, 1, 2}
		screen.DrawTriangles(vertices, indices, whitePixel, nil)

		drawFractal(screen, nx, ny, radius*0.5, depth+1, maxDepth, -baseAngle*1.5, time)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {
	whitePixel = ebiten.NewImage(1, 1)
	whitePixel.Fill(color.White)

	ebiten.SetWindowSize(1920, 1080)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Flower Art")

	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
