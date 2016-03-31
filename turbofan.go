package turbofan

import "reflect"

type Turbofan struct {
	Blast func(bool)
}

// New just creates a new turbofan
func New(chans ...chan bool) *Turbofan {
	t := &Turbofan{}

	// http://stackoverflow.com/questions/19992334/how-to-listen-to-n-channels-dynamic-select-statement
	cases := make([]reflect.SelectCase, len(chans))
	for i, ch := range chans {
		cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch)}
	}

	go func() {
		for {
			chosen, value, _ := reflect.Select(cases)
			for j, ch := range chans {
				if j != chosen {
					ch <- value.Bool()
				}
			}
		}
	}()

	t.Blast = func(b bool) {
		for _, ch := range chans {
			ch <- b
		}
	}

	return t
}
