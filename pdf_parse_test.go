package pdft

import (
	"fmt"
	"testing"
)

func TestPDFParse(t *testing.T) {
	FontKeyword = "F"
	//pdfParseNoTrailer(t)
	pdfParseFromGoPdf(t)
}

func pdfParseNoTrailer(t *testing.T) {
	filename := "no_trailer.pdf"
	var ipdf PDFt
	err := ipdf.Open("test/pdf/" + filename)
	if err != nil {
		t.Error(err)
		return
	}
}

func pdfParseFromGoPdf(t *testing.T) {
	filename := "pdf_from_gopdf.pdf"
	var ipdf PDFt
	err := ipdf.Open("test/pdf/" + filename)
	if err != nil {
		t.Error(err)
		return
	}
}

func pdfParseChrome50Win10(t *testing.T) {
	filename := "pdf_from_chrome_50_win10.pdf"
	fmt.Printf("\n\n\n####Open %s ####\n\n", filename)
	var ipdf PDFt
	err := ipdf.Open("test/pdf/" + filename)
	if err != nil {
		t.Error(err)
		return
	}

	var props PDFObjPropertiesData
	pdfObj := ipdf.pdf.getObjByID(84)
	err = readProperties(&pdfObj.data, &props)
	if err != nil {
		t.Error(err)
		return
	}
}
