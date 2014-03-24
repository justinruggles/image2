package image2

import (
    "image"
    "image/color"

    "github.com/justinruggles/image2/color2"
)

/*****************************************************************************/

type RGB struct {
    Pix []uint8
    Stride int
    Rect image.Rectangle
}

func (p *RGB) ColorModel() color.Model { return color2.RGBModel }

func (p *RGB) Bounds() image.Rectangle { return p.Rect }

func (p *RGB) At(x, y int) color.Color {
    if !(image.Point{x, y}.In(p.Rect)) {
        return color2.RGB{}
    }
    i := p.PixOffset(x, y)
    return color2.RGB{p.Pix[i+0], p.Pix[i+1], p.Pix[i+2]}
}

func (p *RGB) PixOffset(x, y int) int {
    return (y-p.Rect.Min.Y)*p.Stride + (x-p.Rect.Min.X)*3
}

func (p *RGB) Set(x, y int, c color.Color) {
    if !(image.Point{x, y}.In(p.Rect)) {
        return
    }
    i := p.PixOffset(x, y)
    c1 := color2.RGBModel.Convert(c).(color2.RGB)
    p.Pix[i+0] = c1[0]
    p.Pix[i+1] = c1[1]
    p.Pix[i+2] = c1[2]
}

func (p *RGB) SetRGB(x, y int, c color2.RGB) {
    if !(image.Point{x, y}.In(p.Rect)) {
        return
    }
    i := p.PixOffset(x, y)
    p.Pix[i+0] = c[0]
    p.Pix[i+1] = c[1]
    p.Pix[i+2] = c[2]
}

func (p *RGB) SubImage(r image.Rectangle) image.Image {
    r = r.Intersect(p.Rect)
    if r.Empty() {
        return &RGB{}
    }
    i := p.PixOffset(r.Min.X, r.Min.Y)
    return &RGB{
        Pix:    p.Pix[i:],
        Stride: p.Stride,
        Rect:   r,
    }
}

func NewRGB(r image.Rectangle) *RGB {
    w, h := r.Dx(), r.Dy()
    buf := make([]uint8, 3*w*h)
    return &RGB{buf, 3 * w, r}
}

/*****************************************************************************/

type RGB48 struct {
    Pix []uint8
    Stride int
    Rect image.Rectangle
}

func (p *RGB48) ColorModel() color.Model { return color2.RGB48Model }

func (p *RGB48) Bounds() image.Rectangle { return p.Rect }

func (p *RGB48) At(x, y int) color.Color {
    if !(image.Point{x, y}.In(p.Rect)) {
        return color2.RGB48{}
    }
    i := p.PixOffset(x, y)
    return color2.RGB48 {
        (uint16(p.Pix[i+0]) << 8) | uint16(p.Pix[i+1]),
        (uint16(p.Pix[i+2]) << 8) | uint16(p.Pix[i+3]),
        (uint16(p.Pix[i+4]) << 8) | uint16(p.Pix[i+5]),
    }
}

func (p *RGB48) PixOffset(x, y int) int {
    return (y-p.Rect.Min.Y)*p.Stride + (x-p.Rect.Min.X)*6
}

func (p *RGB48) Set(x, y int, c color.Color) {
    if !(image.Point{x, y}.In(p.Rect)) {
        return
    }
    i := p.PixOffset(x, y)
    c1 := color2.RGB48Model.Convert(c).(color2.RGB48)
    p.Pix[i+0] = uint8(c1[0] >> 8)
    p.Pix[i+1] = uint8(c1[0] & 0xFF)
    p.Pix[i+2] = uint8(c1[1] >> 8)
    p.Pix[i+3] = uint8(c1[1] & 0xFF)
    p.Pix[i+4] = uint8(c1[2] >> 8)
    p.Pix[i+5] = uint8(c1[2] & 0xFF)
}

func (p *RGB48) SetRGB48(x, y int, c color2.RGB48) {
    if !(image.Point{x, y}.In(p.Rect)) {
        return
    }
    i := p.PixOffset(x, y)
    p.Pix[i+0] = uint8(c[0] >> 8)
    p.Pix[i+1] = uint8(c[0] & 0xFF)
    p.Pix[i+2] = uint8(c[1] >> 8)
    p.Pix[i+3] = uint8(c[1] & 0xFF)
    p.Pix[i+4] = uint8(c[2] >> 8)
    p.Pix[i+5] = uint8(c[2] & 0xFF)
}

func (p *RGB48) SubImage(r image.Rectangle) image.Image {
    r = r.Intersect(p.Rect)
    if r.Empty() {
        return &RGB48{}
    }
    i := p.PixOffset(r.Min.X, r.Min.Y)
    return &RGB48{
        Pix:    p.Pix[i:],
        Stride: p.Stride,
        Rect:   r,
    }
}

func NewRGB48(r image.Rectangle) *RGB48 {
    w, h := r.Dx(), r.Dy()
    buf := make([]uint8, 6*w*h)
    return &RGB48{buf, 6 * w, r}
}

/*****************************************************************************/
