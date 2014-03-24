package image2

import (
    "image"
    "image/color"

    "github.com/justinruggles/image2/color2"
)

/*****************************************************************************/

type GrayNA struct {
    Pix []uint8
    Stride int
    Rect image.Rectangle
}

func (p *GrayNA) ColorModel() color.Model { return color2.GrayNAModel }

func (p *GrayNA) Bounds() image.Rectangle { return p.Rect }

func (p *GrayNA) At(x, y int) color.Color {
    if !(image.Point{x, y}.In(p.Rect)) {
        return color2.GrayNA{}
    }
    i := p.PixOffset(x, y)
    return color2.GrayNA{p.Pix[i+0], p.Pix[i+1]}
}

func (p *GrayNA) PixOffset(x, y int) int {
    return (y-p.Rect.Min.Y)*p.Stride + (x-p.Rect.Min.X)*2
}

func (p *GrayNA) Set(x, y int, c color.Color) {
    if !(image.Point{x, y}.In(p.Rect)) {
        return
    }
    i := p.PixOffset(x, y)
    c1 := color2.GrayNAModel.Convert(c).(color2.GrayNA)
    p.Pix[i+0] = c1.Y
    p.Pix[i+1] = c1.A
}

func (p *GrayNA) SetGrayNA(x, y int, c color2.GrayNA) {
    if !(image.Point{x, y}.In(p.Rect)) {
        return
    }
    i := p.PixOffset(x, y)
    p.Pix[i+0] = c.Y
    p.Pix[i+1] = c.A
}

func (p *GrayNA) SubImage(r image.Rectangle) image.Image {
    r = r.Intersect(p.Rect)
    if r.Empty() {
        return &GrayNA{}
    }
    i := p.PixOffset(r.Min.X, r.Min.Y)
    return &GrayNA{
        Pix:    p.Pix[i:],
        Stride: p.Stride,
        Rect:   r,
    }
}

func NewGrayNA(r image.Rectangle) *GrayNA {
    w, h := r.Dx(), r.Dy()
    buf := make([]uint8, 2*w*h)
    return &GrayNA{buf, 2 * w, r}
}

/*****************************************************************************/

type GrayA struct {
    Pix []uint8
    Stride int
    Rect image.Rectangle
}

func (p *GrayA) ColorModel() color.Model { return color2.GrayAModel }

func (p *GrayA) Bounds() image.Rectangle { return p.Rect }

func (p *GrayA) At(x, y int) color.Color {
    if !(image.Point{x, y}.In(p.Rect)) {
        return color2.GrayA{}
    }
    i := p.PixOffset(x, y)
    return color2.GrayA{p.Pix[i+0], p.Pix[i+1]}
}

func (p *GrayA) PixOffset(x, y int) int {
    return (y-p.Rect.Min.Y)*p.Stride + (x-p.Rect.Min.X)*2
}

func (p *GrayA) Set(x, y int, c color.Color) {
    if !(image.Point{x, y}.In(p.Rect)) {
        return
    }
    i := p.PixOffset(x, y)
    c1 := color2.GrayAModel.Convert(c).(color2.GrayA)
    p.Pix[i+0] = c1.Y
    p.Pix[i+1] = c1.A
}

func (p *GrayA) SetGrayA(x, y int, c color2.GrayA) {
    if !(image.Point{x, y}.In(p.Rect)) {
        return
    }
    i := p.PixOffset(x, y)
    p.Pix[i+0] = c.Y
    p.Pix[i+1] = c.A
}

func (p *GrayA) SubImage(r image.Rectangle) image.Image {
    r = r.Intersect(p.Rect)
    if r.Empty() {
        return &GrayA{}
    }
    i := p.PixOffset(r.Min.X, r.Min.Y)
    return &GrayA{
        Pix:    p.Pix[i:],
        Stride: p.Stride,
        Rect:   r,
    }
}

func NewGrayA(r image.Rectangle) *GrayA {
    w, h := r.Dx(), r.Dy()
    buf := make([]uint8, 2*w*h)
    return &GrayA{buf, 2 * w, r}
}

/*****************************************************************************/

type GrayNA32 struct {
    Pix []uint8
    Stride int
    Rect image.Rectangle
}

func (p *GrayNA32) ColorModel() color.Model { return color2.GrayNA32Model }

func (p *GrayNA32) Bounds() image.Rectangle { return p.Rect }

func (p *GrayNA32) At(x, y int) color.Color {
    if !(image.Point{x, y}.In(p.Rect)) {
        return color2.GrayNA32{}
    }
    i  := p.PixOffset(x, y)
    y0 := (uint16(p.Pix[i+0]) << 8) | uint16(p.Pix[i+1])
    a  := (uint16(p.Pix[i+2]) << 8) | uint16(p.Pix[i+3])
    return color2.GrayNA32{y0, a}
}

func (p *GrayNA32) PixOffset(x, y int) int {
    return (y-p.Rect.Min.Y)*p.Stride + (x-p.Rect.Min.X)*4
}

func (p *GrayNA32) Set(x, y int, c color.Color) {
    if !(image.Point{x, y}.In(p.Rect)) {
        return
    }
    i := p.PixOffset(x, y)
    c1 := color2.GrayNA32Model.Convert(c).(color2.GrayNA32)
    p.Pix[i+0] = uint8(c1.Y >> 8)
    p.Pix[i+1] = uint8(c1.Y & 0xFF)
    p.Pix[i+3] = uint8(c1.A >> 8)
    p.Pix[i+4] = uint8(c1.A & 0xFF)
}

func (p *GrayNA32) SetGrayNA32(x, y int, c color2.GrayNA32) {
    if !(image.Point{x, y}.In(p.Rect)) {
        return
    }
    i := p.PixOffset(x, y)
    p.Pix[i+0] = uint8(c.Y >> 8)
    p.Pix[i+1] = uint8(c.Y & 0xFF)
    p.Pix[i+3] = uint8(c.A >> 8)
    p.Pix[i+4] = uint8(c.A & 0xFF)
}

func (p *GrayNA32) SubImage(r image.Rectangle) image.Image {
    r = r.Intersect(p.Rect)
    if r.Empty() {
        return &GrayNA32{}
    }
    i := p.PixOffset(r.Min.X, r.Min.Y)
    return &GrayNA32{
        Pix:    p.Pix[i:],
        Stride: p.Stride,
        Rect:   r,
    }
}

func NewGrayNA32(r image.Rectangle) *GrayNA32 {
    w, h := r.Dx(), r.Dy()
    buf := make([]uint8, 4*w*h)
    return &GrayNA32{buf, 4 * w, r}
}


/*****************************************************************************/

type GrayA32 struct {
    Pix []uint8
    Stride int
    Rect image.Rectangle
}

func (p *GrayA32) ColorModel() color.Model { return color2.GrayA32Model }

func (p *GrayA32) Bounds() image.Rectangle { return p.Rect }

func (p *GrayA32) At(x, y int) color.Color {
    if !(image.Point{x, y}.In(p.Rect)) {
        return color2.GrayA32{}
    }
    i := p.PixOffset(x, y)
    y0 := (uint16(p.Pix[i+0]) << 8) | uint16(p.Pix[i+1])
    a  := (uint16(p.Pix[i+2]) << 8) | uint16(p.Pix[i+3])
    return color2.GrayA32{y0, a}
}

func (p *GrayA32) PixOffset(x, y int) int {
    return (y-p.Rect.Min.Y)*p.Stride + (x-p.Rect.Min.X)*4
}

func (p *GrayA32) Set(x, y int, c color.Color) {
    if !(image.Point{x, y}.In(p.Rect)) {
        return
    }
    i := p.PixOffset(x, y)
    c1 := color2.GrayA32Model.Convert(c).(color2.GrayA32)
    p.Pix[i+0] = uint8(c1.Y >> 8)
    p.Pix[i+1] = uint8(c1.Y & 0xFF)
    p.Pix[i+3] = uint8(c1.A >> 8)
    p.Pix[i+4] = uint8(c1.A & 0xFF)
}

func (p *GrayA32) SetGrayA32(x, y int, c color2.GrayA32) {
    if !(image.Point{x, y}.In(p.Rect)) {
        return
    }
    i := p.PixOffset(x, y)
    p.Pix[i+0] = uint8(c.Y >> 8)
    p.Pix[i+1] = uint8(c.Y & 0xFF)
    p.Pix[i+3] = uint8(c.A >> 8)
    p.Pix[i+4] = uint8(c.A & 0xFF)
}

func (p *GrayA32) SubImage(r image.Rectangle) image.Image {
    r = r.Intersect(p.Rect)
    if r.Empty() {
        return &GrayA32{}
    }
    i := p.PixOffset(r.Min.X, r.Min.Y)
    return &GrayA32{
        Pix:    p.Pix[i:],
        Stride: p.Stride,
        Rect:   r,
    }
}

func NewGrayA32(r image.Rectangle) *GrayA32 {
    w, h := r.Dx(), r.Dy()
    buf := make([]uint8, 4*w*h)
    return &GrayA32{buf, 4 * w, r}
}

/*****************************************************************************/
