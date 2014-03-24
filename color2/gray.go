package color2

import (
    "image/color"
)

/*****************************************************************************/

type GrayNA struct {
    Y uint8
    A uint8
}

func (c GrayNA) RGBA() (r, g, b, a uint32) {
    y := uint32(c.Y)
    y |= y << 8
    a  = uint32(c.A)
    a |= a << 8
    y  = y * a / 0xffff
    r = y
    g = y
    b = y
    return
}

func graynaModel(c color.Color) color.Color {
    if _, ok := c.(GrayNA); ok {
        return c
    }
    r, g, b, a := c.RGBA()
    r = (r * 0xffff) / a
    g = (g * 0xffff) / a
    b = (b * 0xffff) / a
    y := (299 * r + 587 * g + 114 * b + 500) / 1000
    return GrayNA { uint8(y >> 8), uint8(a >> 8) }
}

var GrayNAModel color.Model = color.ModelFunc(graynaModel)

/*****************************************************************************/

type GrayA struct {
    Y uint8
    A uint8
}

func (c GrayA) RGBA() (r, g, b, a uint32) {
    y := uint32(c.Y)
    y |= y << 8
    a  = uint32(c.A)
    a |= a << 8
    r = y
    g = y
    b = y
    return
}

func grayaModel(c color.Color) color.Color {
    if _, ok := c.(GrayA); ok {
        return c
    }
    r, g, b, a := c.RGBA()
    y := (299 * r + 587 * g + 114 * b + 500) / 1000
    return GrayA { uint8(y >> 8), uint8(a >> 8) }
}

var GrayAModel color.Model = color.ModelFunc(grayaModel)

/*****************************************************************************/

type GrayNA32 struct {
    Y uint16
    A uint16
}

func (c GrayNA32) RGBA() (r, g, b, a uint32) {
    y := uint32(c.Y)
    a  = uint32(c.A)
    y  = y * a / 0xffff
    r  = y
    g  = y
    b  = y
    return
}

func grayna32Model(c color.Color) color.Color {
    if _, ok := c.(GrayNA32); ok {
        return c
    }
    r, g, b, a := c.RGBA()
    r = (r * 0xffff) / a
    g = (g * 0xffff) / a
    b = (b * 0xffff) / a
    y := (299 * r + 587 * g + 114 * b + 500) / 1000
    return GrayNA32 { uint16(y), uint16(a) }
}

var GrayNA32Model color.Model = color.ModelFunc(grayna32Model)

/*****************************************************************************/

type GrayA32 struct {
    Y uint16
    A uint16
}

func (c GrayA32) RGBA() (r, g, b, a uint32) {
    y := uint32(c.Y)
    a  = uint32(c.A)
    r  = y
    g  = y
    b  = y
    return
}

func graya32Model(c color.Color) color.Color {
    if _, ok := c.(GrayA32); ok {
        return c
    }
    r, g, b, a := c.RGBA()
    y := (299 * r + 587 * g + 114 * b + 500) / 1000
    return GrayA32 { uint16(y), uint16(a) }
}

var GrayA32Model color.Model = color.ModelFunc(graya32Model)

/*****************************************************************************/
