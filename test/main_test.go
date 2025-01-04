package test

import (
	"os"
	"partition-problem/internal/service"
	"partition-problem/internal/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

var results []util.TestResult

var arrPartition1 = []int{1, 4, 3, 2}
var arrPartition2 = []int{1, 4, 3, 2, 1, 4, 3, 2}
var arrPartition3 = []int{1, 4, 3, 2, 1, 4, 3, 2, 1, 4, 3, 2}

func TestFindPartitionDynamic(t *testing.T) {
	measure1 := util.MeasureExecution("Programacao Dinamica", "4 elementos", func() bool {
		result := service.FindPartitionDynamic(arrPartition1)
		assert.True(t, result, "Expected true")
		return result
	})

	results = append(results, measure1)

	measure2 := util.MeasureExecution("Programacao Dinamica", "8 elementos", func() bool {
		result := service.FindPartitionDynamic(arrPartition2)
		assert.True(t, result, "Expected true")
		return result
	})

	results = append(results, measure2)

	measure3 := util.MeasureExecution("Programacao Dinamica", "12 elementos", func() bool {
		result := service.FindPartitionDynamic(arrPartition3)
		assert.True(t, result, "Expected true")
		return result
	})

	results = append(results, measure3)

}

func TestHasPartitionGreedy(t *testing.T) {
	measure1 := util.MeasureExecution("Algoritmo Guloso", "4 elementos", func() bool {
		result := service.HasPartitionGreedy(arrPartition1)
		assert.True(t, result, "Expected true")
		return result
	})

	results = append(results, measure1)

	measure2 := util.MeasureExecution("Algoritmo Guloso", "8 elementos", func() bool {
		result := service.HasPartitionGreedy(arrPartition2)
		assert.True(t, result, "Expected true")
		return result
	})

	results = append(results, measure2)

	measure3 := util.MeasureExecution("Algoritmo Guloso", "12 elementos", func() bool {
		result := service.HasPartitionGreedy(arrPartition3)
		assert.True(t, result, "Expected true")
		return result
	})
	results = append(results, measure3)
}

func TestFindPartitionBruteForce(t *testing.T) {
	measure1 := util.MeasureExecution("Forca Bruta", "4 elementos", func() bool {
		result := service.FindPartitionBruteForce(arrPartition1)
		assert.True(t, result, "Expected true")
		return result
	})

	results = append(results, measure1)

	measure2 := util.MeasureExecution("Forca Bruta", "8 elementos", func() bool {
		result := service.FindPartitionBruteForce(arrPartition2)
		assert.True(t, result, "Expected true")
		return result
	})

	results = append(results, measure2)

	measure3 := util.MeasureExecution("Forca Bruta", "12 elementos", func() bool {
		result := service.FindPartitionBruteForce(arrPartition3)
		assert.True(t, result, "Expected true")
		return result
	})

	results = append(results, measure3)

}

func TestHasPartitionDC(t *testing.T) {
	measure1 := util.MeasureExecution("Divisao e Conquista", "4 elementos", func() bool {
		result := service.HasPartitionDC(arrPartition1)
		assert.True(t, result, "Expected true")
		return result
	})

	results = append(results, measure1)

	measure2 := util.MeasureExecution("Divisao e Conquista", "8 elementos", func() bool {
		result := service.HasPartitionDC(arrPartition2)
		assert.True(t, result, "Expected true")
		return result
	})

	results = append(results, measure2)

	measure3 := util.MeasureExecution("Divisao e Conquista", "12 elementos", func() bool {
		result := service.HasPartitionDC(arrPartition3)
		assert.True(t, result, "Expected true")
		return result
	})

	results = append(results, measure3)
}

func TestMain(m *testing.M) {
	code := m.Run()

	util.SaveResults(results)
	util.GenerateTableTimeCompare()

	os.Exit(code)
}
