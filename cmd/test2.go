package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/tada3/gob-test2/domain"
)

func main() {
	ps := make([]domain.Person, 10)
	for i := 0; i < 10; i++ {
		ps[i].Name = "Taro" + strconv.Itoa(i)
		ps[i].Id = i
	}

	checkSize(ps)
}

func checkSize(p []domain.Person) {
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

func serializeJSON(p []domain.Person) ([]byte, error) {
	b, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func serializeGob(p []domain.Person) ([]byte, error) {
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	err := enc.Encode(p)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
