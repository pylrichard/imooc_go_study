package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func getSocreReult(score int) string {
	result := ""
	//switch后可以没有表达式，默认自动添加break，否则需要fallthrough
	switch {
	case score < 0 || score > 100:
		result = "range error"
	case score < 60:
		result = "D"
	case score < 80:
		result = "C"
	case score < 90:
		result = "B"
	case score < 100:
		result = "A"
	case score == 100:
		result = "A+"
	}
	return result
}

func main() {
	const fileName = "scores.txt"
	_, err := ioutil.ReadFile(fileName)
	if (err != nil) {
		fmt.Println(err)
	} else {
		file, _ := os.Open(fileName)
		fileScanner := bufio.NewScanner(file)
		for fileScanner.Scan() {
			score, err := strconv.Atoi(fileScanner.Text())
			result := ""
			if (err == nil) {
				result = strconv.Itoa(score) + " -> " + getSocreReult(score)
			} else {
				result = "illegal num : " + fileScanner.Text()
			}
			fmt.Println(result)
		}
	}
}
