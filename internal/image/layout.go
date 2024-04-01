package image

import (
	"github.com/fogleman/gg"
	"github.com/ostapneko/muso-matrix/internal/muso"
)

const (
	MATRIX_SIZE = 64
	MARGIN      = 4
	FONT_SIZE   = 12
)

// offset is the number of pixels to offset the image to the lef
func Draw(track *muso.Track, offset int) {
	// only used to measure stuff
	dc := gg.NewContext(0, 0)
	var imageW float64
	// dc.LoadFontFace("/usr/share/fonts/truetype/dejavu/DejaVuSansMono.ttf", FONT_SIZE)

	for _, str := range []string{track.Title, track.Artist, track.Album} {
		w, _ := dc.MeasureString(str)
		if w > imageW {
			imageW = w
		}
	}

	// dc = gg.NewContext(int(imageW)+2*MARGIN, MATRIX_SIZE)
	dc = gg.NewContext(MATRIX_SIZE, MATRIX_SIZE)

	dc.SetRGB(0, 0, 0)
	dc.Clear()
	dc.SetRGB(1, 1, 1)

	dc.LoadFontFace("/usr/share/fonts/truetype/dejavu/DejaVuSansMono.ttf", FONT_SIZE)

	x := MARGIN - float64(offset)

	dc.DrawString(track.Title, x, FONT_SIZE)

	dc.SetRGB(1, 0, 0)
	dc.DrawString(track.Album, x, (MATRIX_SIZE+FONT_SIZE-MARGIN)/2)

	dc.SetRGB(0, 1, 0)
	dc.DrawString(track.Artist, x, MATRIX_SIZE-MARGIN)

	dc.SavePNG("out.png")
}
