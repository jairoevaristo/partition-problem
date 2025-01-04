package util

import (
	"runtime"
	"time"
)

func MeasureExecution(algorithm string, matrixSize string, fn func() bool) TestResult {
	beforeUsedMemory := GetMemoryUsage()
	start := time.Now()

	fn()

	duration := time.Since(start).Seconds() * 1000
	afterUsedMemory := GetMemoryUsage()

	var memoryUsage float64

	if afterUsedMemory > beforeUsedMemory {
		memoryUsage = float64(afterUsedMemory-beforeUsedMemory) / (1024 * 1024)
	}
	if afterUsedMemory == beforeUsedMemory {
		memoryUsage = float64(afterUsedMemory) / (1024 * 1024)
	} else {
		memoryUsage = float64(beforeUsedMemory) / (1024 * 1024)
	}

	return TestResult{
		Algorithm:  algorithm,
		MatrixSize: matrixSize,
		TimeMs:     duration,
		Memory:     memoryUsage,
	}
}

func GetMemoryUsage() uint64 {
	runtime.GC()

	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	return memStats.TotalAlloc
}
