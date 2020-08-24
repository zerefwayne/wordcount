package main

import (
	"bufio"
	"os"
	"strings"
)

func readTxtFile(name string) ([]string, error) {

	f, err := os.Open(name)

	if err != nil {
		return nil, err
	}

	s := bufio.NewScanner(f)

	words := []string{}

	for s.Scan() {

		line := strings.TrimSpace(s.Text())

		lineWords := strings.Split(line, " ")

		for _, word := range lineWords {
			words = append(words, word)
		}

	}

	if s.Err() != nil {
		return nil, s.Err()
	}

	if err = f.Close(); err != nil {
		return nil, err
	}

	return words, nil

}
