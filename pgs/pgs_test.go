package pgs

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"strconv"
	"testing"

	"github.com/tada3/gob-test2/domain"
)

func initData(n int) []domain.Person {
	data := make([]domain.Person, n)
	for i := 0; i < n; i++ {
		data[i].Id = i
		data[i].Name = "Person" + strconv.Itoa(i)
	}
	return data
}

func BenchmarkSerializeAndDeserialize(b *testing.B) {
	pool, err := NewPGSPool(2)
	if err != nil {
		b.Fatal(err)
	}

	data := initData(1000)
	b.ResetTimer()
	// Nはコマンド引数から与えられたベンチマーク時間から自動で計算される
	for _, p := range data {
		s := pool.TakeS()
		image, err := s.Serialize(&p)
		if err != nil {
			b.Fatal(err)
		}
		pool.PutS(s)

		d := pool.TakeD()
		p1, err := d.Deserialize(image)
		if err != nil {
			b.Fatal(err)
		}
		if p1.Name != p.Name {
			b.Fatal("Wrong data", p1)
		}
		pool.PutD(d)
	}
}

func BenchmarkJSON(b *testing.B) {
	data := initData(1000)
	b.ResetTimer()
	// Nはコマンド引数から与えられたベンチマーク時間から自動で計算される
	for _, p := range data {
		image, err := json.Marshal(&p)
		if err != nil {
			b.Fatal(err)
		}

		var p1 domain.Person
		json.Unmarshal(image, &p1)
		if p1.Name != p.Name {
			b.Fatal("Wrong data", p1)
		}
	}
}

func BenchmarkGob(b *testing.B) {
	data := initData(1000)
	b.ResetTimer()
	// Nはコマンド引数から与えられたベンチマーク時間から自動で計算される
	for _, p := range data {
		buf := &bytes.Buffer{}
		enc := gob.NewEncoder(buf)
		err := enc.Encode(p)
		if err != nil {
			b.Fatal(err)
		}
		image := buf.Bytes()

		r := bytes.NewReader(image)
		dec := gob.NewDecoder(r)
		var p1 domain.Person
		err = dec.Decode(&p1)
		if err != nil {
			b.Fatal(err)
		}

		if p1.Name != p.Name {
			b.Fatal("Wrong data", p1)
		}
	}
}
