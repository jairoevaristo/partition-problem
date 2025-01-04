package util

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jung-kurt/gofpdf"
)

type TestResult struct {
	Algorithm  string  `json:"algorithm"`
	MatrixSize string  `json:"matrix_size"`
	TimeMs     float64 `json:"time_ms"`
}

func SaveResults(results []TestResult) {
	file, err := os.Create("../execution_times.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(results); err != nil {
		panic(err)
	}
}

func GenerateTableTimeCompare() {
	file, err := os.Open("../execution_times.json")
	if err != nil {
		fmt.Printf("Erro ao abrir o arquivo: %v\n", err)
		return
	}
	defer file.Close()

	var results []TestResult
	if err := json.NewDecoder(file).Decode(&results); err != nil {
		fmt.Printf("Erro ao decodificar o JSON: %v\n", err)
		return
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(80, 10, "Algoritmo")
	pdf.Cell(60, 10, "Array de elementos")
	pdf.Cell(40, 10, "Tempo (ms)")
	pdf.Ln(10)
	pdf.SetFont("Arial", "", 12)

	var currentAlgorithmName string
	for _, result := range results {
		if currentAlgorithmName != result.Algorithm {
			pdf.Ln(10)
		}

		pdf.Cell(80, 10, result.Algorithm)
		pdf.Cell(60, 10, result.MatrixSize)
		pdf.Cell(40, 10, fmt.Sprintf("%.5f", result.TimeMs))
		pdf.Ln(10)

		currentAlgorithmName = result.Algorithm
	}

	err = pdf.OutputFileAndClose("../table_output.pdf")
	if err != nil {
		fmt.Printf("Erro ao salvar o PDF: %v\n", err)
		return
	}

	fmt.Println("PDF gerado com sucesso como 'table_output.pdf'.")
}
