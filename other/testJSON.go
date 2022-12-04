package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type MyStruct struct {
	Id int `json:"global_id"`
}

func main() {
	file, err := os.Open("C:\\Users\\User\\Desktop\\Разное\\test2.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	dec := json.NewDecoder(file)
	var dst []MyStruct

	dec.Decode(&dst)

	fmt.Println(dst)

	sum := 0
	for _, el := range dst {
		sum += el.Id
	}
	fmt.Println(sum)
}
