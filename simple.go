package main

func computeSimpleWordCount(words []string) map[string]int {

	frequency := make(map[string]int)

	for _, word := range words {
		frequency[word]++
	}

	return frequency

}
