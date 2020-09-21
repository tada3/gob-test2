package pgs

import (
	"bytes"
	"encoding/gob"

	"github.com/tada3/gob-test2/domain"
)

type PretrainedGobSerializer struct {
	buf *bytes.Buffer
	enc *gob.Encoder
}

type PretrainedGobDeserializer struct {
	buf *bytes.Buffer
	dec *gob.Decoder
}

func NewPretrainedGobSerializer() (*PretrainedGobSerializer, error) {
	s, _, err := NewPretrainedGobSAndD()
	return s, err
}

func NewPretrainedGobDeserializer() (*PretrainedGobDeserializer, error) {
	_, d, err := NewPretrainedGobSAndD()
	return d, err
}

func NewPretrainedGobSAndD() (*PretrainedGobSerializer, *PretrainedGobDeserializer, error) {
	buf1 := &bytes.Buffer{}
	enc := gob.NewEncoder(buf1)
	p1 := domain.Person{}
	err := enc.Encode(p1)
	if err != nil {
		return nil, nil, err
	}

	buf2 := &bytes.Buffer{}
	buf2.ReadFrom(buf1)
	dec := gob.NewDecoder(buf2)
	p2 := &domain.Person{}
	err = dec.Decode(p2)
	if err != nil {
		return nil, nil, err
	}

	buf1.Reset()
	buf2.Reset()

	return &PretrainedGobSerializer{
			buf: buf1,
			enc: enc,
		},
		&PretrainedGobDeserializer{
			buf: buf2,
			dec: dec,
		}, nil
}

func (pte *PretrainedGobSerializer) Serialize(p *domain.Person) ([]byte, error) {
	defer pte.buf.Reset()
	err := pte.enc.Encode(p)
	if err != nil {
		return nil, err
	}
	b := pte.buf.Bytes()
	ret := make([]byte, len(b))
	copy(ret, b)
	return ret, nil
}

func (pgd *PretrainedGobDeserializer) Deserialize(b []byte) (*domain.Person, error) {
	defer pgd.buf.Reset()
	pgd.buf.Write(b)
	p := &domain.Person{}
	err := pgd.dec.Decode(p)
	if err != nil {
		return nil, err
	}
	return p, nil
}
