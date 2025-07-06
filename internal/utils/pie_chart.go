package utils

import (
	"fmt"
	"image/color"
	"math"

	"github.com/Rodrigoos/stock-bot-telegram/internal/models"
	"github.com/fogleman/gg"
)

func CreatePieChart(portfolio models.Portfolio, path string) error {
	const (
		W       = 700
		H       = 600
		radius  = 200.0
		centerX = 300.0
		centerY = H / 2.0
	)

	dc := gg.NewContext(W, H)
	dc.SetColor(color.White)
	dc.Clear()

	if err := dc.LoadFontFace("/usr/share/fonts/truetype/dejavu/DejaVuSans-Bold.ttf", 8); err != nil {
		return err
	}

	var total float64
	for _, asset := range portfolio.Assets {
		total += asset.Price * float64(asset.Quantity)
	}
	if total == 0 {
		total = 1
	}

	colors := []color.RGBA{
		{255, 99, 132, 255}, {54, 162, 235, 255}, {255, 206, 86, 255},
		{75, 192, 192, 255}, {153, 102, 255, 255}, {255, 159, 64, 255},
		{100, 255, 100, 255}, {200, 200, 200, 255},
	}

	startAngle := 0.0
	legendX := 540.0
	legendY := 100.0
	colorIndex := 0

	for _, asset := range portfolio.Assets {
		value := asset.Price * float64(asset.Quantity)
		portion := value / total
		angle := portion * 2 * math.Pi
		endAngle := startAngle + angle

		// Cor
		c := colors[colorIndex%len(colors)]
		dc.SetColor(c)

		// Fatia
		dc.MoveTo(centerX, centerY)
		dc.DrawArc(centerX, centerY, radius, startAngle, endAngle)
		dc.ClosePath()
		dc.Fill()

		// Rótulo no centro da fatia
		midAngle := (startAngle + endAngle) / 2
		labelX := centerX + math.Cos(midAngle)*radius*0.6
		labelY := centerY + math.Sin(midAngle)*radius*0.6

		percent := portion * 100
		label := fmt.Sprintf("%s-%.1f%%", asset.Ticker, percent)
		dc.SetRGB(0.1, 0.1, 0.1)
		dc.DrawStringAnchored(label, labelX, labelY, 0.5, 0.5)

		// Legenda lateral
		dc.SetColor(c)
		dc.DrawRectangle(legendX, legendY-10, 14, 14)
		dc.Fill()

		dc.SetRGB(0.1, 0.1, 0.1)
		legendText := fmt.Sprintf("%s - %s (%.1f%%)", asset.Ticker, FormatBRL(value), percent)
		dc.DrawStringAnchored(legendText, legendX+20, legendY, 0, 0.5)

		legendY += 24
		startAngle = endAngle
		colorIndex++
	}

	// Título
	dc.SetRGB(0.1, 0.1, 0.1)
	dc.LoadFontFace("/usr/share/fonts/truetype/dejavu/DejaVuSans-Bold.ttf", 10)
	dc.DrawStringAnchored(fmt.Sprintf("Carteira: %s", portfolio.Name), W/2, 30, 0.5, 0.5)

	return dc.SavePNG(path)
}
