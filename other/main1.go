package main

import (
	"fmt"
	"time"
)

const t = time.RFC3339
const uD = time.UnixDate

//	func main() {
//		rd := bufio.NewReader(os.Stdin)
//		timeStr, err := rd.ReadString('\n')
//		if err != nil && err != io.EOF {
//			panic(err)
//		}
//
//		//timeStr = strings.TrimSuffix(timeStr, "\n")
//		dateSlice := strings.Split(timeStr, ",")
//
//		dateSliceTwo := strings.TrimSpace(dateSlice[1])
//
//		timeOne, _ := time.Parse("2006-01-02 15:04:05", dateSlice[0])
//		timeTwo, _ := time.Parse("2006-01-02 15:04:05", dateSliceTwo)
//
//		fmt.Println(timeOne, "\n", timeTwo)
//
//		fmt.Println(time.Since(timeOne).Round(time.Second) - time.Since(timeTwo).Round(time.Second))
//	}
func main() {
	inputStream := make(chan string)
	outputStream := make(chan string)
	go removeDuplicates(inputStream, outputStream)

	go func() {
		defer close(inputStream)

		for _, r := range "112334456" {
			inputStream <- string(r)
		}
	}()

	for x := range outputStream {
		fmt.Print(x)
	}
}
func removeDuplicates(inputStream, outputStream chan string) {
	var prevD string

	for d := range inputStream {
		if d != prevD {
			outputStream <- d
			prevD = d
		}
	}
	close(outputStream)
}
