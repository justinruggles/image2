package image2

import (
    "image"
    "image/color"
)

/*****************************************************************************/

type NBGRA struct {
    Pix []uint8
    Stride int
    Rect image.Rectangle
}

func (p *NBGRA) ColorModel() color.Model {
    return color.NRGBAModel
}

func (p *NBGRA) Bounds() image.Rectangle {
    return p.Rect
}

func (p *NBGRA) At(x, y int) color.Color {
    if !(image.Point{x, y}.In(p.Rect)) {
        return color.NRGBA{}
    }

    i := p.PixOffset(x, y)
    return color.NRGBA{
        B: p.Pix[i + 0],
        G: p.Pix[i + 1],
        R: p.Pix[i + 2],
        A: p.Pix[i + 3],
    }
}

func (p *NBGRA) PixOffset(x, y int) int {
    return (y - p.Rect.Min.Y) * p.Stride + (x - p.Rect.Min.X) * 4
}

func (p *NBGRA) Set(x, y int, c color.Color) {
    if !(image.Point{x, y}.In(p.Rect)) {
        return
    }

    i  := p.PixOffset(x, y)
    c1 := color.NRGBAModel.Convert(c).(color.NRGBA)
    p.Pix[i + 0] = c1.B
    p.Pix[i + 1] = c1.G
    p.Pix[i + 2] = c1.R
    p.Pix[i + 3] = c1.A
}

func (p *NBGRA) SetNRGBA(x, y int, c color.NRGBA) {
    if !(image.Point{x, y}.In(p.Rect)) {
        return
    }
    i := p.PixOffset(x, y)
    p.Pix[i + 0] = c.B
    p.Pix[i + 1] = c.G
    p.Pix[i + 2] = c.R
    p.Pix[i + 3] = c.A
}

func (p *NBGRA) SubImage(r image.Rectangle) image.Image {
    r = r.Intersect(p.Rect)
    if r.Empty() {
        return &NBGRA{}
    }
    i := p.PixOffset(r.Min.X, r.Min.Y)
    return &NBGRA{
        Pix:    p.Pix[i:],
        Stride: p.Stride,
        Rect:   r,
    }
}

func NewNBGRA(r image.Rectangle) *NBGRA {
    w, h := r.Dx(), r.Dy()
    buf := make([]uint8, 4 * w * h)
    return &NBGRA{
        Pix:    buf,
        Stride: w * 4,
        Rect:   r,
    }
}

/*****************************************************************************/
