package pdf

import (
	"github.com/jung-kurt/gofpdf"
)

func Generate() (*gofpdf.Fpdf, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	
	// Configuration de la police
	pdf.SetFont("Arial", "B", 24)
	
	// Centrage du texte
	pdf.CellFormat(210, 297/2, "Work In Progress", "", 0, "C", false, 0, "")
	
	return pdf, nil
}
