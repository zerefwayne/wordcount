package main

import (
	"fmt"
	"time"
)

func computeSimpleWordCount(words []string) (map[string]int, int) {

	now := time.Now()

	frequency := make(map[string]int)

	for _, word := range words {
		frequency[word]++
	}

	exec := time.Now().Sub(now)

	return frequency, int(exec.Milliseconds())

}

func executeSimpleWordCount(words []string) {

	fmt.Printf("\nSimple Word Count\n\n")

	simpleWordCountMap, time := computeSimpleWordCount(words)

	rankedCount := rankByWordCount(simpleWordCountMap)

	for idx, pair := range rankedCount[40:50] {
		fmt.Println("Rank:", 40+idx, pair.Key, pair.Value)
	}

	fmt.Printf("\nExecution Time: %dms\n", time)

}
