package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	input1, err := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	if err != nil {
		log.Fatal(err)
	}
	input2, err := hex.DecodeString("686974207468652062756c6c277320657965")
	if err != nil {
		log.Fatal(err)
	}
	output := hexXOR(input1, input2)
	fmt.Println(hex.EncodeToString(output))
}

func hexXOR(input1 []byte, input2 []byte) []byte {
	if len(input1) != len(input2) {
		log.Fatal("Two inputs must be of equal length.")
	}
	output := make([]byte, len(input1))
	for i := 0; i < len(input1); i++ {
		output[i] = input1[i] ^ input2[i]
	}
	return output
}
