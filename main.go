package main

import (
	"fmt"
	"log"
	"sort"
)

// Pair ...
type Pair struct {
	Key   string
	Value int
}

// PairList ...
type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func rankByWordCount(wordFrequencies map[string]int) PairList {
	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

func main() {

	fileName := "shakespeare.txt"

	fmt.Println("Reading from", fileName)

	words, err := readTxtFile(fileName)

	if err != nil {
		log.Fatalln(err)
	}

	simpleWordCountMap := computeSimpleWordCount(words)

	rankedCount := rankByWordCount(simpleWordCountMap)

	for _, pair := range rankedCount[20:100] {
		fmt.Println(pair)
	}

}
