package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	// define CLI-Parameters
	url := flag.String("url", "", "target-url with placeholder")
	wordlist := flag.String("wordlist", "", "path to wordlist")

	// flags parsen
	flag.Parse()

	// error handling
	if *url == "" || *wordlist == "" {
		fmt.Println("[MISSING OPTIONS]...go run . -url https://example.com/api/FUZZ -wordlist wordlist.txt")
		return
	}

	fmt.Println("[WORDLIST] ", *wordlist)
	fmt.Println("[URL] ", *url)

	fmt.Println("[+]...loading up fuzzer-lite")

	// open wordlist
	file, err := os.Open(*wordlist)
	if err != nil {
		fmt.Println("Error while open wordlist:", err)
		return
	}
	defer file.Close()

	// read line for line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		finalURL := strings.Replace(*url, "FUZZ", word, 1)
		// fmt.Println("[TEST] URL:", finalURL)
		resp, err := http.Get(finalURL)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}
		fmt.Println(finalURL, resp.StatusCode)
		defer resp.Body.Close()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error while reading file:", err)
	}

	// send get request for every finalURL

}
