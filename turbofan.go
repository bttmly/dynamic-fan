package turbofan

import (
	"reflect"
)

// New just creates a new turbofan
func New(chans ...chan bool) {
	cases := make([]reflect.SelectCase, len(chans))
	go func() {
		for i, ch := range chans {
			cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch)}
		}

		chosen, value, ok := reflect.Select(cases)

		if !ok {
			return
		}

		for j, ch := range chans {
			if j != chosen {
				ch <- value.Bool()
			}
		}
	}()
}
