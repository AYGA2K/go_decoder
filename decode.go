package main

import (
	"encoding/base64"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var hexString string
	var base64Flag string

	flag.StringVar(&hexString, "hex", "", "Input hexadecimal string")
	flag.StringVar(&base64Flag, "base64", "", "Base64-encoded string to replace confi:dential")
	flag.Parse()

	if hexString == "" && base64Flag == "" {
		fmt.Println("Please provide either a hexadecimal string using the -hex flag or a Base64-encoded string using the -base64 flag.")
		os.Exit(1)
	}

	if hexString != "" {
		hexString = removeSpaces(hexString)

		decodedHex, err := hex.DecodeString(hexString)
		if err != nil {
			fmt.Println("Error decoding hex string:", err)
			os.Exit(1)
		}

		asciiString := string(decodedHex)
		fmt.Println("ASCII:", asciiString)
	}
	if base64Flag != "" {
		decodedBase64, err := base64.StdEncoding.DecodeString(base64Flag)
		if err != nil {
			fmt.Println("Error decoding Base64 string:", err)
			os.Exit(1)
		}

		fmt.Println("Decoded Base64:", string(decodedBase64))
	}
}

func removeSpaces(s string) string {
	// Remove spaces from the input string
	return strings.ReplaceAll(s, " ", "")
}
