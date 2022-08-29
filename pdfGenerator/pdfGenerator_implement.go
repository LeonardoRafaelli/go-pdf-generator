package pdfGenerator

import (
	"os"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/google/uuid"
)

type wkhtml struct {
	rootPath string
}

func NewWkHtmlToPdf(rootPath string) PDFGeneratorInterface {
	return &wkhtml{rootPath: rootPath}
}

func (w *wkhtml) Create(htmlFile string) (string, error) {

	file, err := os.Open(htmlFile)

	if err != nil {
		return "1", err
	}

	pdfg, err := wkhtmltopdf.NewPDFGenerator()

	if err != nil {
		return "2", err
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(file))

	if err := pdfg.Create(); err != nil {
		return "3", err
	}

	fileName := w.rootPath + "/" + uuid.New().String() + ".pdf"

	if err := pdfg.WriteFile(fileName); err != nil {
		return "4", err
	}

	return fileName, nil
}
