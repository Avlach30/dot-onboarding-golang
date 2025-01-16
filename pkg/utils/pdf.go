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
		return wkhtmltopdf.OrientationPortrait
	}
}

func ExportToPDF(htmlContentOrUrl string, filePath string, orientation OrientationType) {
	// Create new PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	// Set global options
	pdfg.Orientation.Set(orientation.String())

	page := wkhtmltopdf.NewPage(htmlContentOrUrl)

	// Add to document
	pdfg.AddPage(page)

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	// Write buffer contents to file on disk
	err = pdfg.WriteFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")
	// Output: Done
}
