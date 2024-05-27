package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
	"log"
	"os"
	"sync"
)

const (
	iterations = 600000             // CHANGE THIS
	keyLen     = 32                 // CHANGE THIS
	salt       = "MSok34zBufo9d1tc" // CHANGE THIS
)

type Result struct {
	Password string
	Hash     string
}

func main() {
	passwordsFile, err := os.Open("password_list.txt") // CHANGE THIS
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer func(passwordsFile *os.File) {
		err := passwordsFile.Close()
		if err != nil {
			log.Fatal("Error closing file:", err)
		}
	}(passwordsFile)

	decodedSalt := []byte(salt)
	expectedHash := "b2adfafaeed459f903401ec1656f9da36f4b4c08a50427ec7841570513bf8e57" // CHANGE THIS

	results := make(chan Result)
	var wg sync.WaitGroup

	scanner := bufio.NewScanner(passwordsFile)
	for scanner.Scan() {
		wg.Add(1)
		go func(password string) {
			defer wg.Done()
			hash := pbkdf2.Key([]byte(password), decodedSalt, iterations, keyLen, sha256.New)
			hashString := fmt.Sprintf("%x", hash)
			results <- Result{Password: password, Hash: hashString}
		}(scanner.Text())
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		if result.Hash == expectedHash {
			fmt.Printf("The hash of the password '%s' matches the expected hash.\n", result.Password)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	} else {
		fmt.Println("No matching password found.")
	}
}
