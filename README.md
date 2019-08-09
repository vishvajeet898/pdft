## PDFT

PDFT is a Go library for creating PDF documents using existing PDFs as template. This library depend on [gopdf](https://github.com/signintech/gopdf). 

Tested with PDF template files created from Libre office, Google Docs, Microsoft Word.

### Usage

```go
var pt pdft.PDFt
pdft.FontKeyword = "F"

err := pt.Open(pdfsource)
if err != nil {
	log.Println(err.Error())
	return
}

err = pt.AddFont("arial", "./arial.ttf")
if err != nil {
    log.Println(err.Error())
	return
}

err = pt.SetFont("arial", "", 14)
if err != nil {
    log.Println(err.Error())
	return
}

err = pt.Insert("Hi", 1, 10, 10, 100, 100, gopdf.Center|gopdf.Bottom)
if err != nil {
    log.Println(err.Error())
	return
}

err = pt.Save("./output.pdf")
if err != nil {
	log.Println(err.Error())
	return
}
```

### Unit tests

```
go test -v ./...
```