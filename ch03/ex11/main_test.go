package ex10

import "testing"

func Test_comma(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "when given empty then returns empty",
			args: "",
			want: "",
		},
		{
			name: "when given 1 digit then returns itself",
			args: "1",
			want: "1",
		},
		{
			name: "when given 3 digits then returns itself",
			args: "123",
			want: "123",
		},
		{
			name: "when given 4 digits then put a comma",
			args: "1234",
			want: "1,234",
		},
		{
			name: "when given 6 digits then put a comma",
			args: "123456",
			want: "123,456",
		},
		{
			name: "when given 7 digits then put 2 commas",
			args: "1234567",
			want: "1,234,567",
		},
		{
			name: "when given minus 3 digits then returns itself",
			args: "-123",
			want: "-123",
		},
		{
			name: "when given plus 3 digits then returns itself",
			args: "+123",
			want: "+123",
		},
		{
			name: "when given minus 4 digits then put a comma",
			args: "-1234",
			want: "-1,234",
		},
		{
			name: "when given plus 4 digits then put a comma",
			args: "+1234",
			want: "+1,234",
		},
		{
			name: "when given minus 6 digits then put a comma",
			args: "-123456",
			want: "-123,456",
		},
		{
			name: "when given plus 6 digits then put a comma",
			args: "+123456",
			want: "+123,456",
		},
		{
			name: "when given 2 digits has decimal point then returns itself",
			args: "1.2",
			want: "1.2",
		},
		{
			name: "when given number has forth decimal point then returns itself",
			args: "1.2345",
			want: "1.2345",
		},
		{
			name: "when given number's integer part has 4 digits then put a comma",
			args: "1234.5678",
			want: "1,234.5678",
		},
		{
			name: "when given number is negative and its integer part has 4 digits then put a comma",
			args: "-1234.5678",
			want: "-1,234.5678",
		},
		{
			name: "when given number is positive and its integer part has 4 digits then put a comma",
			args: "+1234.5678",
			want: "+1,234.5678",
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := comma(tt.args); got != tt.want {
				t.Errorf("comma() = %v, want %v", got, tt.want)
			}
		})
	}
}
