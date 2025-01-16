package utils

import (
	"fmt"
	"log"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

// OrientationType is an enumeration of the possible PDF orientations
type OrientationType int

const (
	// OrientationLandscape is the landscape orientation
	OrientationLandscape OrientationType = iota
	// OrientationPotrait is the potrait orientation
	OrientationPotrait
)

func (o OrientationType) String() string {
	switch o {
	case OrientationLandscape:
		return wkhtmltopdf.OrientationLandscape
	case OrientationPotrait:
		return wkhtmltopdf.OrientationPortrait
	default:
		return fmt.Sprintf("Orientation(%d)", o)
	}
}

func ExportHtmlToPDF(htmlContent string, orientation OrientationType) {
	// Create new PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	// Set global options
	pdfg.Dpi.Set(300)
	pdfg.Orientation.Set(orientation.String())
	pdfg.Grayscale.Set(true)

	// Create a new input page from an URL
	page := NewPage("https://godoc.org/github.com/SebastiaanKlippert/go-wkhtmltopdf")

	// Set options for this page
	page.FooterRight.Set("[page]")
	page.FooterFontSize.Set(10)
	page.Zoom.Set(0.95)

	// Add to document
	pdfg.AddPage(page)

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	// Write buffer contents to file on disk
	err = pdfg.WriteFile("./simplesample.pdf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")
	// Output: Done
}
