package lengthoflongestsubstring

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "abcabcbb",
			args: args{s: "abcabcbb"},
			want: 3,
		},
		{
			name: "aaaa",
			args: args{s: "aaaa"},
			want: 1,
		},
		{
			name: "abbbddd",
			args: args{s: "abbbddd"},
			want: 2,
		},
		{
			name: "",
			args: args{s: ""},
			want: 0,
		},
		{
			name: "a",
			args: args{s: "a"},
			want: 1,
		},
		{
			name: "aa",
			args: args{s: "aa"},
			want: 1,
		},
		{
			name: "ab",
			args: args{s: "ab"},
			want: 2,
		},
		{
			name: "aab",
			args: args{s: "aab"},
			want: 2,
		},
		{
			name: "dvdf",
			args: args{s: "dvdf"},
			want: 3,
		},
		{
			name: "abcasdfga",
			args: args{s: "abcasdfga"},
			want: 7,
		},
		{
			name: "aabaab!bb",
			args: args{s: "aabaab!bb"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
