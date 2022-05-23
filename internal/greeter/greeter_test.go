package greeter

import "testing"

func TestGreeter_What(t *testing.T) {
	type args struct {
		lang     string
		expected string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "en español",
			args: args{
				lang:     "es",
				expected: "hola",
			},
		},
		{
			name: "in english",
			args: args{
				lang:     "en",
				expected: "hello",
			},
		},
		{
			name: "default",
			args: args{
				lang:     "",
				expected: "01001000 01100101 01101100",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := New(tt.args.lang)
			res := g.What()
			if res != tt.args.expected {
				t.Errorf("wrong greet. Want: %v, Got: %v", tt.args.expected, res)
			}
		})
	}
}
