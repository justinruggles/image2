package color

import (
    "image/color"
)

/*****************************************************************************/

type RGB [3]uint8

func (c RGB) RGBA() (r, g, b, a uint32) {
    r = uint32(c[0])
    r |= r << 8
    g = uint32(c[1])
    g |= g << 8
    b = uint32(c[2])
    b |= b << 8
    a = 0xFFFF
    return
}

func rgbModel(c color.Color) color.Color {
    if _, ok := c.(RGB); ok {
        return c
    }
    r, g, b, _ := c.RGBA()
    return RGB { uint8(r >> 8), uint8(g >> 8), uint8(b >> 8) }
}

var RGBModel color.Model = color.ModelFunc(rgbModel)

/*****************************************************************************/

type RGB48 [3]uint16

func (c RGB48) RGBA() (r, g, b, a uint32) {
    r = uint32(c[0])
    g = uint32(c[1])
    b = uint32(c[2])
    a = 0xFFFF
    return
}

func rgb48Model(c color.Color) color.Color {
    if _, ok := c.(RGB48); ok {
        return c
    }
    r, g, b, _ := c.RGBA()
    return RGB48 { uint16(r), uint16(g), uint16(b) }
}

var RGB48Model color.Model = color.ModelFunc(rgb48Model)

/*****************************************************************************/
