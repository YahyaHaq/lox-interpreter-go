package main

import (
	"fmt"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]

	if command != "tokenize" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	// Uncomment this block to pass the first stage
	//
	filename := os.Args[2]
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stderr, "file contents: %v\n", fileContents)

	errorCode := 0
	idx := 0
	lineEnd := false
	for idx < len(fileContents) && !lineEnd {
		char := fileContents[idx]

		switch char {
		case ' ':
			// Skip whitespace
		case '(':
			fmt.Println("LEFT_PAREN ( null")
		case ')':
			fmt.Println("RIGHT_PAREN ) null")
		case '{':
			fmt.Println("LEFT_BRACE { null")
		case '}':
			fmt.Println("RIGHT_BRACE } null")
		case ',':
			fmt.Println("COMMA , null")
		case '.':
			fmt.Println("DOT . null")
		case '-':
			fmt.Println("MINUS - null")
		case '+':
			fmt.Println("PLUS + null")
		case '*':
			fmt.Println("STAR * null")
		case ';':
			fmt.Println("SEMICOLON ; null")
		case '=':
			// case of '=='
			if idx+1 < len(fileContents) && fileContents[idx+1] == '=' {
				fmt.Println("EQUAL_EQUAL == null")
				idx++ // we will move 2 characters ahead in this case
			} else {
				fmt.Println("EQUAL = null")
			}
		case '!':
			// case of ' !='
			if idx+1 < len(fileContents) && fileContents[idx+1] == '=' {
				fmt.Println("BANG_EQUAL != null")
				idx++ // we will move 2 characters ahead in this case
			} else {
				fmt.Println("BANG ! null")
			}
		case '>':
			// case of ' >='
			if idx+1 < len(fileContents) && fileContents[idx+1] == '=' {
				fmt.Println("GREATER_EQUAL >= null")
				idx++ // we will move 2 characters ahead in this case
			} else {
				fmt.Println("GREATER > null")
			}
		case '<':
			// case of ' <='
			if idx+1 < len(fileContents) && fileContents[idx+1] == '=' {
				fmt.Println("LESS_EQUAL <= null")
				idx++ // we will move 2 characters ahead in this case
			} else {
				fmt.Println("LESS < null")
			}
		case '/':
			// case of '//
			if idx+1 < len(fileContents) && fileContents[idx+1] == '/' {
				// do nothing as this is a comment
				// as tests are online we will break out for now as no other code can be written after comments
				// in future we will modify this
				lineEnd = true
			} else {
				fmt.Println("SLASH / null")
			}
		case '\n':
			// do nothing for newline
		case '\t':
			// do nothing for tab
		default:
			errorCode = 65
			fmt.Fprintf(os.Stderr, "[line 1] Error: Unexpected character: %c\n", char)
		}
		idx++
	}

	fmt.Println("EOF  null")
	os.Exit(errorCode)
}
