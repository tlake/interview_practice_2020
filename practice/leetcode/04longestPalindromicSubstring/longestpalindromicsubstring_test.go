package longestpalindromicsubstring

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "'qwertytrewq'",
			args: args{s: "qwertytrewq"},
			want: "qwertytrewq",
		},
		{
			name: "'qwertyytrewq'",
			args: args{"qwertyytrewq"},
			want: "qwertyytrewq",
		},
		{
			name: "'babad'",
			args: args{s: "babad"},
			want: "bab",
		},
		{
			name: "'cbbd'",
			args: args{s: "cbbd"},
			want: "bb",
		},
		{
			name: "'bb'",
			args: args{s: "bb"},
			want: "bb",
		},
		{
			name: "'a'",
			args: args{s: "a"},
			want: "a",
		},
		{
			name: "''",
			args: args{s: ""},
			want: "",
		},
		{
			name: "'aaaaaaaaaaaaa'",
			args: args{s: "aaaaaaaaaaaaa"},
			want: "aaaaaaaaaaaaa",
		},
		{
			name: "'qwertyuiop'",
			args: args{s: "qwertyuiop"},
			want: "q",
		},
		{
			name: "'asdfgfdas'",
			args: args{"asdfgfdas"},
			want: "dfgfd",
		},
		{
			name: "'asdfggfdas'",
			args: args{s: "asdfggfdas"},
			want: "dfggfd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
