package ex12

import "testing"

func Test_isAnagram(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "when given empty strings then returns false",
			args: args{},
			want: false,
		},
		{
			name: "when either one of strings is empty then returns false",
			args: args{s1: "a"},
			want: false,
		},
		{
			name: "when either one of strings is empty then returns false",
			args: args{s2: "a"},
			want: false,
		},
		{
			name: "when given strings is not anagram then returns false (ASCII)",
			args: args{
				s1: "egg",
				s2: "pie",
			},
			want: false,
		},
		{
			name: "when given strings is not anagram then returns false (multi-byte)",
			args: args{
				s1: "たまご",
				s2: "りんご",
			},
			want: false,
		},
		{
			name: "when given strings is anagram then returns true (ASCII)",
			args: args{
				s1: "evil",
				s2: "live",
			},
			want: true,
		},
		{
			name: "when given strings is anagram then returns true (multi-byte)",
			args: args{
				s1: "せかいのななふしぎ", // 世界の七不思議
				s2: "ふしぎのないせなか", // 不思議のない背中
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("isAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
