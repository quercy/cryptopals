package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func readCliArgs() (string, string, string) {
	inputText := flag.String("i", "", "Text to encrypt")
	inputFile := flag.String("f", "", "File to encrypt")
	key := flag.String("k", "", "Key to use for encryption")
	flag.Parse()
	if ((*inputText == "") == (*inputFile == "")) || *key == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	return *inputText, *inputFile, *key
}

func main() {
	inputText, inputFile, key := readCliArgs()
	if inputFile != "" {
		file, err := ioutil.ReadFile(inputFile)
		if err != nil {
			log.Fatal("Error reading input file.")
			os.Exit(1)
		}
		inputText = string(file)
	}
	output := encryptString(inputText, key)
	fmt.Println(output)
}

func encryptString(input string, key string) string {
	output := make([]byte, len(input))
	for i := 0; i < len(input); i++ {
		output[i] = input[i] ^ key[i%len(key)]
	}
	return string(hex.EncodeToString(output))
}
