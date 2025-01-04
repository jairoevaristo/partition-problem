package test

import (
	"os"
	"partition-problem/internal/service"
	"partition-problem/internal/util"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var results []util.TestResult

var arrPartition1 = []int{1, 4, 3, 2}
var arrPartition2 = []int{1, 4, 3, 2, 1, 4, 3, 2}
var arrPartition3 = []int{1, 4, 3, 2, 1, 4, 3, 2, 1, 4, 3, 2}

func measureExecution(algorithm string, matrixSize string, fn func() bool) {
	start := time.Now()
	fn()
	duration := time.Since(start).Seconds() * 1000
	results = append(results, util.TestResult{
		Algorithm:  algorithm,
		MatrixSize: matrixSize,
		TimeMs:     duration,
	})
}

func TestFindPartitionDynamic(t *testing.T) {
	measureExecution("Programacao Dinamica", "4 elementos", func() bool {
		result := service.FindPartitionDynamic(arrPartition1)
		assert.True(t, result, "Expected true")
		return result
	})

	measureExecution("Programacao Dinamica", "8 elementos", func() bool {
		result := service.FindPartitionDynamic(arrPartition2)
		assert.True(t, result, "Expected true")
		return result
	})

	measureExecution("Programacao Dinamica", "12 elementos", func() bool {
		result := service.FindPartitionDynamic(arrPartition3)
		assert.True(t, result, "Expected true")
		return result
	})
}

func TestHasPartitionGreedy(t *testing.T) {
	measureExecution("Algoritmo Guloso", "4 elementos", func() bool {
		result := service.HasPartitionGreedy(arrPartition1)
		assert.True(t, result, "Expected true")
		return result
	})

	measureExecution("Algoritmo Guloso", "8 elementos", func() bool {
		result := service.HasPartitionGreedy(arrPartition2)
		assert.True(t, result, "Expected true")
		return result
	})

	measureExecution("Algoritmo Guloso", "12 elementos", func() bool {
		result := service.HasPartitionGreedy(arrPartition3)
		assert.True(t, result, "Expected true")
		return result
	})
}

func TestFindPartitionBruteForce(t *testing.T) {
	measureExecution("Forca Bruta", "4 elementos", func() bool {
		result := service.FindPartitionBruteForce(arrPartition1)
		assert.True(t, result, "Expected true")
		return result
	})

	measureExecution("Forca Bruta", "8 elementos", func() bool {
		result := service.FindPartitionBruteForce(arrPartition2)
		assert.True(t, result, "Expected true")
		return result
	})

	measureExecution("Forca Bruta", "12 elementos", func() bool {
		result := service.FindPartitionBruteForce(arrPartition3)
		assert.True(t, result, "Expected true")
		return result
	})
}

func TestHasPartitionDC(t *testing.T) {
	measureExecution("Divisao e Conquista", "4 elementos", func() bool {
		result := service.HasPartitionDC(arrPartition1)
		assert.True(t, result, "Expected true")
		return result
	})

	measureExecution("Divisao e Conquista", "8 elementos", func() bool {
		result := service.HasPartitionDC(arrPartition2)
		assert.True(t, result, "Expected true")
		return result
	})

	measureExecution("Divisao e Conquista", "12 elementos", func() bool {
		result := service.HasPartitionDC(arrPartition3)
		assert.True(t, result, "Expected true")
		return result
	})
}

func TestMain(m *testing.M) {
	code := m.Run()

	util.SaveResults(results)
	util.GenerateTableTimeCompare()

	os.Exit(code)
}
