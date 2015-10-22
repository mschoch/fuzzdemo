package split

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		in  []byte
		out []string
		err error
	}{
		{
			in:  []byte("the test input"),
			out: []string{"the", "test", "input"},
			err: nil,
		},
	}

	for _, test := range tests {
		out, err := SplitSpace(test.in)
		if err != test.err {
			t.Errorf("expected err: %v, got %v", test.err, err)
		}
		if !reflect.DeepEqual(out, test.out) {
			t.Errorf("expected: %#v, got %#v", test.out, out)
		}
	}
}
