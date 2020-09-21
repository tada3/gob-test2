package pgs

type PGSPool struct {
	queueS chan *PretrainedGobSerializer
	queueD chan *PretrainedGobDeserializer
}

func NewPGSPool(n int) (*PGSPool, error) {
	qs := make(chan *PretrainedGobSerializer, n)
	qd := make(chan *PretrainedGobDeserializer, n)

	for i := 0; i < n; i++ {
		s, d, err := NewPretrainedGobSAndD()
		if err != nil {
			return nil, err
		}
		qs <- s
		qd <- d
	}

	return &PGSPool{
		queueS: qs,
		queueD: qd,
	}, nil
}

func (pool *PGSPool) TakeS() *PretrainedGobSerializer {
	return <-pool.queueS

}

func (pool *PGSPool) PutS(enc *PretrainedGobSerializer) {
	pool.queueS <- enc
}

func (pool *PGSPool) TakeD() *PretrainedGobDeserializer {
	return <-pool.queueD

}

func (pool *PGSPool) PutD(enc *PretrainedGobDeserializer) {
	pool.queueD <- enc
}
