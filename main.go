package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strings"
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

//Exists reports whether the named file or directory exists.
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
	if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("Wrong number of arguments!")
	}

	//Source filename
	source_filename := args[0]
	//Result filename
	result_filename := fmt.Sprintf("%s.result.txt", source_filename)

	fmt.Println("Reading lines from source")
	//Read the lines of the source file into a slice
	lines, er := readLines(source_filename)
	if er != nil {
		log.Fatal(er)
	}

	if Exists(result_filename) {
		log.Fatal("Resulting filename exists: ", result_filename)
	}

	fmt.Println("Proceeding to create result...")
	f, err := os.OpenFile(result_filename, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	defer f.Close()
	if err != nil {
		return
	}
	w := bufio.NewWriter(f)
	for _, line := range lines {
		words := strings.Split(line, " ")
		statement := fmt.Sprintf("*%d\r\n", len(words))
		for i := 0; i < len(words); i++ {
			word := words[i]
			statement += fmt.Sprintf("$%d\r\n%s\r\n", len(word), word)
		}
		w.WriteString(statement)
	}
	w.Flush()
	
	fmt.Println("Result created:", result_filename)
	//Part ways amicably
	fmt.Println("Have a good day!")
}