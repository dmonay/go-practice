package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// the Go way to handle errors is to have the function/method return an error value as their sole or last return value,
// or nil if no error occured, and for receivers to always check the error they receive.

// functions can also return named variables. These return variables are set to their zero values when the function is entered,
// and keep their zero values unless explicitly assigned to in the body of the function.

func main() {

	// multiple assignment is possible b/c the function returns three values
	inFilename, outFilename, err := filenamesFromCommandLine()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// files in Go are represented by pointers to values of type os.File. Here we initialize two variables to the standard
	// input and output streams
	inFile, outFile := os.Stdin, os.Stdout
	if inFilename != "" {
		if inFile, err = os.Open(inFilename); err != nil {
			log.Fatal(err)
		}
		defer inFile.Close()
	}
	if outFilename != "" {
		if outFile, err = os.Create(outFilename); err != nil {
			log.Fatal(err)
		}
		// leave the file open to work on, but close it as soon as the enclosing function, in this case main(), returns, thereby ensuring
		// that the file is closed when the program is done with it
		defer outFile.Close()
	}
	if err = americanise(inFile, outFile); err != nil {
		log.Fatal(err)
	}
}

func filenamesFromCommandLine() (inFilename, outFilename string, err error) {
	if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		err = fmt.Errorf("usage: %s [<]infile.txt [>]outfile.txt",
			filepath.Base(os.Args[0]))
		return "", "", err
	}
	if len(os.Args) > 1 {
		inFilename = os.Args[1]
		if len(os.Args) > 2 {
			outFilename = os.Args[2]
		}
	}
	if inFilename != "" && inFilename == outFilename {
		log.Fatal("won't overwrite the infile")
	}

	// if all goes well, return two strings and an error value of nil
	return inFilename, outFilename, nil
}

// for a value to be readable it must satisfy the io.Reader interface, and for it to be writeable it must satisfy the io.Writer interface.
// the bufio package provides buffered input/output
