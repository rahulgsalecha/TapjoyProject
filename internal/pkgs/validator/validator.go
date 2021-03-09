package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func splitLineToSubstrings(line string) []string {
	line = strings.TrimSpace(line)
	line = strings.ReplaceAll(line, "]", "[")
	return strings.Split(line, "[")
}

func isValidSubstring(str string) bool {
	if len(str) < 4 {
		return false
	}

	return (str[3:4] == str[0:1] && str[2:3] == str[1:2]  && str[0:1] != str[1:2]) || isValidSubstring(str[1:])
}

func checkExternalSubstringsValidity(results []bool) bool {
	if len(results) == 0 {
		return false
	}

	return results[0] || checkExternalSubstringsValidity(results[1:])
}

func checkInternalSubstringsValidity(results []bool) bool {
	if len(results) == 0 {
		return true
	}

	return results[0] && checkInternalSubstringsValidity(results[1:])
}

func lineValidator(line string) bool {
	// Split the line into substrings
	substringsSlice := splitLineToSubstrings(line)

	// validate substrings outside/inside the brackets
	validExternalSubstrings := make([]bool, len(substringsSlice)/2)
	validInternalSubstrings := make([]bool, len(substringsSlice)/2)
	i := 0
	j := 0
	for index, substring := range substringsSlice {
		if index % 2 == 0 {
			// validate substrings outside the brackets
			if isValidSubstring(substring) {
				validExternalSubstrings[i] = true
				i++
			}
		} else {
			// validate substrings inside the brackets
			if !isValidSubstring(substring) {
				validInternalSubstrings[j] = true
				j++
			}
		}
	}
	// check if any one of the outside the bracket substrings is valid
	externalValidity := checkExternalSubstringsValidity(validExternalSubstrings)
	// check if all of the inside bracket substrings are valid or not
	internalValidity := checkInternalSubstringsValidity(validInternalSubstrings)

	return externalValidity && internalValidity
}

func countAndPrintValidLines(inputFile string) error {
	count := 0
    file, err := os.Open(inputFile)
    if err != nil {
        return err

    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    fmt.Printf("\n==== VALID LINES ====\n")
    for scanner.Scan() {
		if lineValidator(scanner.Text()) {
			fmt.Println(scanner.Text())
			count++
		}
    }
	fmt.Printf("\n==== TOTAL VALID LINES: %d ====\n", count)

    if err := scanner.Err(); err != nil {
        return err
    }

    return nil

}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
		return
	}

	// Process arguments
	inputFile := os.Args[1]
	if err := countAndPrintValidLines(inputFile) ; err != nil {
		fmt.Printf("Unable to test : error : %s", err.Error())
	}
}


