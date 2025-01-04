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

	var memoryUsage int64

	if afterUsedMemory > beforeUsedMemory {
		memoryUsage = int64(afterUsedMemory - beforeUsedMemory)
	}
	if afterUsedMemory == beforeUsedMemory {
		memoryUsage = int64(afterUsedMemory)
	} else {
		memoryUsage = int64(beforeUsedMemory)
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

	return memStats.Alloc
}
