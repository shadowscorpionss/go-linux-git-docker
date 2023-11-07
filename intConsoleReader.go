package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//errors constants
const (
	BREAK_WORD_RECIEVED_ERROR="Break Word Recieved"
)


type IntConsoleReader struct {
	scanner *bufio.Scanner
}

func NewIntConsoleReader() *IntConsoleReader {
	return &IntConsoleReader{scanner: bufio.NewScanner(os.Stdin)}
}

func (ir *IntConsoleReader) Read(welcomeString, breakWord string) (res int, err error) {
	//default results
	res, err = 0, nil
	
	fmt.Println(welcomeString)	
	for {
		//reading console and check data
		ir.scanner.Scan()
		consoleData := ir.scanner.Text()
		if strings.EqualFold(consoleData, breakWord) {
			err = errors.New(BREAK_WORD_RECIEVED_ERROR)
			return
		}

		res, err = strconv.Atoi(consoleData)
		if err != nil {
			fmt.Println("Error. Enter integer number.")
			continue
		}
		return
	}
	
}
