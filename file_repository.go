package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func readFile(filename string) string {
	filename = fmt.Sprintf("articles/%s", filename)

	file, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	return string(file)
}

func writeToFile(body string) string {
	filename := getNextFilename()
	result := writeFile(filename, body)

	return result
}

func getNextFilename() string {
	filenames, err := ioutil.ReadDir("articles")

	if err != nil {
		panic(err)
	}

	var highest int64

	for _, v := range filenames {
		i, err := strconv.ParseInt(v.Name(), 10, 64)

		if err != nil {
			panic(err)
		}

		if i <= highest {
			continue
		}

		highest = i
	}

	highest = highest + 1

	return fmt.Sprintf("articles/%s", strconv.FormatInt(highest, 10))
}

func writeFile(filename string, body string) string {

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	b, err := file.WriteString(body)

	if err != nil {
		panic(err)
	} else if b == 0 {
		panic("No data to write.")
	}

	file.Sync()

	return filename
}
