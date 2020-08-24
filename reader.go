package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

func readTxtFile(name string) ([]string, error) {

	f, err := os.Open(name)

	if err != nil {
		return nil, err
	}

	s := bufio.NewScanner(f)

	words := []string{}

	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}

	for s.Scan() {

		line := strings.TrimSpace(s.Text())

		lineWords := strings.Split(line, " ")

		for _, word := range lineWords {
			word = reg.ReplaceAllString(word, "")
			word = strings.ToLower(word)
			if len(word) > 0 {
				words = append(words, word)
			}
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
