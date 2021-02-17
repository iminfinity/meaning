package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Dictionary struct {
	Word       string `json:"word,omitempty"`
	Definition string `json:"definition,omitempty"`
}

func main() {
	file, err := os.Open("dictionary.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	list := []Dictionary{}
	for scanner.Scan() {
		currentLine := scanner.Text()
		info := strings.Split(currentLine, "\",")
		if len(info) == 3 {
			dict := Dictionary{}
			word := strings.Split(info[0], "\"")
			definition := strings.Split(info[2], "\"")

			dict.Word = word[1]
			dict.Definition = definition[1]
			list = append(list, dict)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	data, _ := json.Marshal(list)
	fmt.Println(string(data))

	saveToJSON, err := json.MarshalIndent(list, " ", "  ")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("dictionary.json", saveToJSON, 0644)
}
