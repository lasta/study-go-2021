package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	hashAlgorithm := flag.String("hashfunc", "SHA256", "hash function [SHA256, SHA384, SHA512]")
	flag.Parse()

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		data := input.Bytes()
		switch *hashAlgorithm {
		case "SHA384":
			fmt.Printf("%x\n", sha512.Sum384(data))
			continue
		case "SHA512":
			fmt.Printf("%x\n", sha512.Sum512(data))
			continue
		case "SHA256":
			fmt.Printf("%x\n", sha256.Sum256(data))
			continue
		default:
			log.Fatalf("unkown hash function: %s", *hashAlgorithm)
		}
	}
}
