package utils

import (
	"image/color"

	"github.com/fogleman/gg"
)

func CreateStartImage(path string) error {
	const W, H = 600, 600
	dc := gg.NewContext(W, H)

	// Fundo branco
	dc.SetColor(color.White)
	dc.Clear()

	// Título
	if err := dc.LoadFontFace("/usr/share/fonts/truetype/dejavu/DejaVuSans-Bold.ttf", 30); err != nil {

		return err
	}
	dc.SetRGB(0.1, 0.1, 0.1)
	dc.DrawStringAnchored("📈 Stock Bot", W/2, 60, 0.5, 0.5)

	// Subtítulo
	if err := dc.LoadFontFace("/usr/share/fonts/truetype/dejavu/DejaVuSans-Bold.ttf", 28); err != nil {

		return err
	}
	dc.SetRGB(0.3, 0.3, 0.3)
	dc.DrawStringAnchored("Seu assistente de investimentos", W/2, 100, 0.5, 0.5)

	// Blocos de comandos
	commands := []string{
		"💹 /stock ou /acao",
		"🏢 /fundo",
		"🪙 /cripto",
		"📊 /portfolio",
		"🗂️ /all-portfolios",
	}

	boxY := 160.0
	boxHeight := 50.0

	dc.SetRGB(0.3, 0.3, 0.3)

	for _, cmd := range commands {

		dc.SetRGB(0.9, 0.9, 0.9)
		// Caixa
		dc.DrawRoundedRectangle(60, boxY, W-120, boxHeight, 10)
		dc.Fill()

		// Texto
		dc.SetRGB(0, 0, 0)
		dc.DrawStringAnchored(cmd, W/2, boxY+boxHeight/2, 0.5, 0.5)

		boxY += boxHeight + 15
	}

	// Rodapé
	dc.SetRGB(0.6, 0.6, 0.6)
	dc.LoadFontFace("/usr/share/fonts/truetype/dejavu/DejaVuSans-Bold.ttf", 10)

	dc.DrawStringAnchored("by Rodrigo • github.com/Rodrigoos", W/2, H-30, 0.5, 0.5)

	// Salvar imagem
	return dc.SavePNG(path)
}
