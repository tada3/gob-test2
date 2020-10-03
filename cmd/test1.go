package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"

	"github.com/tada3/gob-test2/domain"
)

func main() {
	p := domain.Person{
		Name: "Taro",
		Id:   123,
	}

	checkSize(p)
}

func checkSize(p domain.Person) {
	b, err := serializeJSON(p)
	if err != nil {
		panic(err)
	}
	fmt.Printf("xxx", b)
	fmt.Println("JSON:", len(b))

	b1, err := serializeGob(p)
	if err != nil {
		panic(err)
	}
	fmt.Println("Gob:", len(b1))
}

func serializeJSON(p domain.Person) ([]byte, error) {
	b, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func serializeGob(p domain.Person) ([]byte, error) {
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	err := enc.Encode(p)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
