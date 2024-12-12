package reader

import (
	"fmt"
)

type Reader struct {
	dataCh chan interface{}
}

func NewReader(dataCh chan interface{}) *Reader {
	return &Reader{
		dataCh: dataCh,
	}
}

func (r *Reader) StartRead() {

	for data := range r.dataCh {
		fmt.Println(data)
	}
}
