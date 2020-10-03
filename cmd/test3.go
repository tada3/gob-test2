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
	n := 3
	ps := make([]domain.Person, n)
	for i := 0; i < n; i++ {
		ps[i].Name = "Taro" + strconv.Itoa(i)
		ps[i].Id = i
	}
	cs := make([]domain.Car, n)
	for i := 0; i < n; i++ {
		cs[i].Model = "Pajero" + strconv.Itoa(i)
		cs[i].Number = strconv.Itoa(i)
		cs[i].Color = i
	}

	checkSize3(ps, cs)
}

func checkSize3(ps []domain.Person, cs []domain.Car) {
	b, err := serializeJSON3(ps, cs)
	if err != nil {
		panic(err)
	}
	fmt.Println("JSON:", len(b))

	b1, err := serializeGob3(ps, cs)
	if err != nil {
		panic(err)
	}
	fmt.Println("Gob:", len(b1))
}

func serializeJSON3(ps []domain.Person, cs []domain.Car) ([]byte, error) {
	buf := &bytes.Buffer{}
	for i := 0; i < len(ps); i++ {
		b, err := json.Marshal(ps[i])
		if err != nil {
			return nil, err
		}
		buf.Write(b)
		b, err = json.Marshal(cs[i])
		if err != nil {
			return nil, err
		}
		buf.Write(b)
	}
	return buf.Bytes(), nil
}

func serializeGob3(ps []domain.Person, cs []domain.Car) ([]byte, error) {
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)

	for i := 0; i < len(ps); i++ {
		err := enc.Encode(ps[i])
		if err != nil {
			return nil, err
		}
		err = enc.Encode(cs[i])
		if err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}
