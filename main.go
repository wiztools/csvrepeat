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
		fmt.Println("Enter file name as argument")
		os.Exit(1)
	}

	filename := os.Args[1]

	readFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Err opening file:", filename)
		os.Exit(1)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		arr := strings.Split(line, ",")
		if len(arr) != 2 {
			fmt.Println("Err line does not match signature:", line)
			os.Exit(1)
		}
		tag := arr[0]
		count, err := strconv.Atoi(arr[1])
		if err != nil {
			fmt.Println("Err CSV col-2 NaN:", line)
			os.Exit(1)
		}
		for i := 0; i < count; i++ {
			fmt.Println(tag)
		}
	}
}
