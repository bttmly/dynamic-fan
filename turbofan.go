package turbofan

import "reflect"

type Turbofan struct {
	chans []chan bool
	closed bool
}

func (t *Turbofan) Broadcast(b bool) {
	for _, ch := range t.chans {
		ch <- b
	}
}

func (t *Turbofan) Close() {
	t.closed = true
	for _, ch := range t.chans {
		close(ch)
	}
}

func (t *Turbofan) init() {
	// http://stackoverflow.com/questions/19992334/how-to-listen-to-n-channels-dynamic-select-statement
	cases := make([]reflect.SelectCase, len(t.chans))
	for i, ch := range t.chans {
		cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch)}
	}

	go func() {
		for {
			chosen, value, _ := reflect.Select(cases)

			if (t.closed) { 
				return 
			}

			val := value.Bool()
			for j, ch := range t.chans {
				if j != chosen {
					ch <- val
				}
			}
		}
	}()
}

func New(chans ...chan bool) *Turbofan {
	t := &Turbofan{chans, false}
	t.init()
	return t
}
