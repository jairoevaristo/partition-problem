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
	Memory     int64   `json:"memory"`
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
	pdf.Cell(80, 10, "RELATORIO COM OS TEMPOS DE EXECUCAO")
	pdf.Ln(20)
	pdf.Line(10, 30, 200, 30)

	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(70, 10, "Algoritmo")
	pdf.Cell(50, 10, "Array de elementos")
	pdf.Cell(40, 10, "Tempo (ms)")
	pdf.Cell(60, 10, "Memoria (bytes)")
	pdf.Ln(10)
	pdf.SetFont("Arial", "", 12)

	var currentAlgorithmName string
	for _, result := range results {
		if currentAlgorithmName != result.Algorithm {
			pdf.Ln(10)
		}

		pdf.Cell(70, 10, result.Algorithm)
		pdf.Cell(50, 10, result.MatrixSize)
		pdf.Cell(40, 10, fmt.Sprintf("%.5f", result.TimeMs))
		pdf.Cell(60, 10, fmt.Sprintf("%d", result.Memory))
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
