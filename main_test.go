package main

import "testing"

func Test_stripSymbols(t *testing.T) {
	type args struct {
		name string
		str     string
		symbols []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"strings with multiple symbols",
			args{
				str: "R$100.00R$",
				symbols: []string{"R$", ",", "."},
			},
			10000,
		},
		{
			 "only a number",
			args{
				str: "100",
				symbols: []string{"R$", ",", "."},
			},
			100,
		},
		{
			"too expensive",
			args{
				str: "R$1.799,99",
				symbols: []string{"R$", ",", "."},
			},
			179999,
		},
		{
			"cheaper prices",
			args{
				str: "R$549,90",
				symbols: []string{"R$", ",", "."},
			},
			54990,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := stripSymbols(tt.args.str, tt.args.symbols); got != tt.want {
				t.Errorf("stripSymbols() = got  %v, want %v", got, tt.want)
			}
		})
	}
}
