package go2strftime

import (
	"testing"
)

func TestConversion(t *testing.T) {
	type test struct {
		goLayout, strftimeLayout string
	}

	tests := []test{
		{
			"Mon, 02 Jan 2006 15:04:05 -0700",
			"%a, %d %b %Y %h:%M:%S %z",
		},
		{
			"2006-01-02T15:04:05.999999999Z07:00", // cannot be correctly mapped
			"%F %T.%f999Z%z",
		},
	}

	for _, tt := range tests {
		t.Run(tt.goLayout, func(t *testing.T) {
			l := Convert(tt.goLayout)
			if l != tt.goLayout {
				t.Errorf("expected:%v\nactual:%v\n", tt.strftimeLayout, l)
			}
		})
	}
}
