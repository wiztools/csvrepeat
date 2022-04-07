package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		os.Stderr.WriteString("Enter file name as argument\n")
		os.Exit(1)
	}

	filename := os.Args[1]

	readFile, err := os.Open(filename)
	if err != nil {
		os.Stderr.WriteString("Err opening file:" + filename + "\n")
		os.Exit(1)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		arr := strings.Split(line, ",")
		if len(arr) != 2 {
			os.Stderr.WriteString("Err line does not match signature:" + line + "\n")
			os.Exit(1)
		}
		tag := arr[0]
		count, err := strconv.Atoi(arr[1])
		if err != nil {
			os.Stderr.WriteString("Err CSV col-2 NaN:" + line + "\n")
			os.Exit(1)
		}
		for i := 0; i < count; i++ {
			fmt.Println(tag)
		}
	}
}
