package drawing

import (
	"github.com/fogleman/gg"
	"go-wp/pkg/image"
	"log"
)

func DrawTextIntoImage(imagePath string, fontPath string, fontSize float64, x float64, y float64, text string) {
	im, err := gg.LoadImage(imagePath)
	if err != nil {
		log.Fatal(err)
	}
	with, height := image.GetImageDimensions(imagePath)
	dc := gg.NewContext(with, height)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)

	if err := dc.LoadFontFace(fontPath, fontSize); err != nil {
		panic(err)
	}
	dc.DrawImage(im, 0, 0)

	w, h := dc.MeasureString(text)
	margin := 20.0
	dc.DrawRectangle(x-margin-w/2.0, y-margin-h/2.0, w+2*margin, h+2*margin)
	dc.SetRGBA(1, 1, 1, 128)
	dc.Fill()

	dc.SetRGB(0, 0, 0)
	//dc.DrawStringWrapped(text, x, y, 0, 0, 1000, 2, gg.AlignLeft)

	dc.DrawStringAnchored(text, x, y, 0.5, 0.5)

	dc.Clip()
	dc.SavePNG(imagePath)
}
