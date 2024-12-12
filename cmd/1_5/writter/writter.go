package writter

type Writter struct {
	isRunning bool
	dataCh    chan interface{}
}

func NewWritter(dataCh chan interface{}) *Writter {
	return &Writter{
		isRunning: true,
		dataCh:    dataCh,
	}
}

func (w *Writter) StartWrite(data []string) {
	count := 0
	for w.isRunning {
		if count < len(data) {
			w.dataCh <- data[count]
			count++
		}
	}
}

func (w *Writter) Stop() {
	w.isRunning = false
}
