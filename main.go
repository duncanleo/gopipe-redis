package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	inputFilename string
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// reports whether the named file or directory exists.
func exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func generateRedisScript(lines []string) string {
	var buffer bytes.Buffer

	for _, line := range lines {
		words := strings.Split(line, " ")
		statement := fmt.Sprintf("*%d\r\n", len(words))
		for i := 0; i < len(words); i++ {
			word := words[i]
			statement += fmt.Sprintf("$%d\r\n%s\r\n", len(word), word)
		}
		buffer.WriteString(statement)
	}

	return buffer.String()
}

func main() {
	flag.StringVar(&inputFilename, "i", "", "input filename")
	flag.Parse()

	if len(inputFilename) == 0 {
		log.Fatal("Please specify an input filename using the -i option")
	}

	if !exists(inputFilename) {
		log.Fatal("Specified filename does not exist")
	}

	//Read the lines of the source file into a slice
	lines, er := readLines(inputFilename)
	if er != nil {
		log.Fatal(er)
	}
	fmt.Printf(generateRedisScript(lines))
}
