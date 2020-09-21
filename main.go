package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"

	"github.com/tada3/gob-test2/domain"
	"github.com/tada3/gob-test2/pgs"
)

var (
	pool *pgs.PGSPool
)

func init() {
	var err error
	pool, err = pgs.NewPGSPool(2)
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Hello, playground")

	p := domain.Person{
		Name: "Taro",
		Id:   "123",
	}

	test(p)

}

func test(p domain.Person) {
	fmt.Println(p)

	b, err := serializeJSON(p)
	if err != nil {
		panic(err)
	}
	fmt.Println("bytes:", string(b))
	fmt.Println("size:", len(b))
	pj, err := deserializeJSON(b)
	if err != nil {
		panic(err)
	}
	fmt.Println("deserialized:", pj)

	b1, err := serializeGob(p)
	if err != nil {
		panic(err)
	}
	fmt.Println("\nbytes:", b1)
	fmt.Println("size:", len(b1))
	p1, err := deserializeGob(b1)
	if err != nil {
		panic(err)
	}
	fmt.Println("deserialized:", p1)

	b2, err := serializePGS(p)
	if err != nil {
		panic(err)
	}
	fmt.Println("\nbytes:", b2)
	fmt.Println("size:", len(b2))
	p2, err := deserializePGS(b2)
	if err != nil {
		panic(err)
	}
	fmt.Println("deserialized:", p2)

	b3, err := serializePool(p)
	if err != nil {
		panic(err)
	}
	fmt.Println("\nbytes:", b3)
	fmt.Println("size:", len(b3))
	p3, err := deserializePool(b3)
	if err != nil {
		panic(err)
	}
	fmt.Println("deserialized:", p3)
}

func serializeJSON(p domain.Person) ([]byte, error) {
	b, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func deserializeJSON(b []byte) (*domain.Person, error) {
	p := &domain.Person{}
	err := json.Unmarshal(b, p)
	if err != nil {
		return nil, err
	}
	return p, nil
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

func deserializeGob(b []byte) (*domain.Person, error) {
	r := bytes.NewReader(b)
	dec := gob.NewDecoder(r)
	p := &domain.Person{}
	err := dec.Decode(p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func serializePGS(p domain.Person) ([]byte, error) {
	enc, err := pgs.NewPretrainedGobSerializer()
	if err != nil {
		return nil, err
	}
	b, err := enc.Serialize(&p)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func deserializePGS(b []byte) (*domain.Person, error) {
	dec, err := pgs.NewPretrainedGobDeserializer()
	if err != nil {
		return nil, err
	}
	p, err := dec.Deserialize(b)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func serializePool(p domain.Person) ([]byte, error) {
	s := pool.TakeS()
	defer pool.PutS(s)
	b, err := s.Serialize(&p)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func deserializePool(b []byte) (*domain.Person, error) {
	dec := pool.TakeD()
	defer pool.PutD(dec)
	p, err := dec.Deserialize(b)
	if err != nil {
		return nil, err
	}
	return p, nil
}
