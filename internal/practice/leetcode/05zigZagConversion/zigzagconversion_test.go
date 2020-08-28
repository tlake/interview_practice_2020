package zigzagconversion

import (
	"testing"
)

// func TestOscillator_Next(t *testing.T) {
// 	type fields struct {
// 		Range   []int
// 		Current int
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		want   int
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			o := &Oscillator{
// 				Range:   tt.fields.Range,
// 				Current: tt.fields.Current,
// 			}
// 			if got := o.Next(); got != tt.want {
// 				t.Errorf("Oscillator.Next() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestNewOscillator(t *testing.T) {
// 	type args struct {
// 		numItems int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want *Oscillator
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := NewOscillator(tt.args.numItems); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("NewOscillator() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func Test_convert(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "'PAYPALISHIRING', 3",
			args: args{s: "PAYPALISHIRING", numRows: 3},
			want: "PAHNAPLSIIGYIR",
		},
		{
			name: "'', 3",
			args: args{s: "", numRows: 3},
			want: "",
		},
		{
			name: "'PAYPALISHRING', 1",
			args: args{s: "PAYPALISHRING", numRows: 1},
			want: "PAYPALISHRING",
		},
		{
			name: "'PAYPALISHRING', 2",
			args: args{s: "PAYPALISHRING", numRows: 2},
			want: "PYAIHIGAPLSRN",
		},
		{
			name: "'PAYPALISHRING', 4",
			args: args{s: "PAYPALISHRING", numRows: 4},
			want: "PIGALSNYAHIPR",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
