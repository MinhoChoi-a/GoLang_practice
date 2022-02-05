package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const URL = ""

func main() {

	resp, err := http.Get(URL)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	content := string(bytes)

	somethings := someFromJson(content)

	for _, some := range somethings {
		fmt.Println(some.Name)
	}

}

func someFromJson(content string) []Sometin {
	somethings := make([]Sometin, 0, 20)

	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	if err != nil {
		panic(err)
	}

	var some Sometin
	for decoder.More() {
		err := decoder.Decode(&some)
		if err != nil {
			panic(err)
		}
		somethings = append(somethings, some)
	}

	return somethings
}

type Sometin struct {

	//For json parse, it is not case sensitive : JSON decoder knows it
	Name, Some string
}
