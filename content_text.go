package pdft

import (
	"bytes"

	"github.com/plutov/pdft/minigopdf"
)

//ContentText text in pdf
type ContentText struct {
	text          string
	fontName      string
	fontStyle     int
	fontSize      int
	pageNum       int
	x             float64
	y             float64
	pdfFontData   *PDFFontData
	w             float64
	h             float64
	align         int
	lineWidth     float64
	pdfProtection *gopdf.PDFProtection
}

func (c *ContentText) setProtection(p *gopdf.PDFProtection) {
	c.pdfProtection = p
}

func (c *ContentText) protection() *gopdf.PDFProtection {
	return c.pdfProtection
}

func (c *ContentText) toSteram() (*bytes.Buffer, error) {

	var border = 0
	if c.lineWidth > 0 {
		border = Left | Right | Top | Bottom
	}

	var rgb gopdf.Rgb
	rgb.SetR(1)
	rgb.SetG(1)
	rgb.SetB(1)

	var cc gopdf.CacheContent
	cc.Setup(
		&gopdf.Rect{
			W: c.w,
			H: c.h,
		},
		rgb,
		1.0,
		c.pdfFontData.fontIndex(),
		c.fontSize,
		c.fontStyle,
		0,
		c.x,
		c.y,
		&c.pdfFontData.font,
		pageHeight(),
		gopdf.ContentTypeText,
		gopdf.CellOption{
			Align:  c.align,
			Border: border,
		},
		c.lineWidth,
	)

	cc.WriteTextToContent(c.text)
	buff, err := cc.ToStream(c.protection())
	if err != nil {
		return nil, err
	}
	buff.Write([]byte("\r\n"))

	return buff, nil
}

func (c *ContentText) page() int {
	return c.pageNum
}
