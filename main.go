package main

import (
	"PDFGenerator/htmlParser"
	"PDFGenerator/pdfGenerator"
	"fmt"
)

type Data struct {
	Name string
}

func main() {
	h := htmlParser.New("tmp")
	wk := pdfGenerator.NewWkHtmlToPdf("tmp")

	htmlData := Data{
		Name: "Leonardo Rafaelli",
	}

	generatedHTML, err := h.Create("templates/index.html", htmlData)
	if err != nil {
		fmt.Println("1main", err)
		return
	}
	fmt.Println("Generated HTML: ", generatedHTML)

	pdfFileName, err := wk.Create(generatedHTML)
	if err != nil {
		fmt.Println("2main", err)
		return
	}
	fmt.Println("PDF gerado", pdfFileName)

}
