package main

import "github.com/jung-kurt/gofpdf"
import(
"os"
)

/* using https://github.com/jung-kurt/gofpdf */

func main() {
    pdf := gofpdf.New("P", "mm", "A4", "../font")
    pdf.AddPage()
    pdf.SetFont("Arial", "B", 16)
    pdf.Cell(40, 10, "Hello, world")
    pdf.Output(os.Stdout)
}

