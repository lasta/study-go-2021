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
			// redundant
			name: "when given 5 digits then put a comma",
			args: "12345",
			want: "12,345",
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
			// redundant
			name: "when given 8 digits then put 2 commas",
			args: "12345678",
			want: "12,345,678",
		},
		{
			// redundant
			name: "when given 9 digits then put 2 commas",
			args: "123456789",
			want: "123,456,789",
		},
		{
			// redundant
			name: "when given 10 digits then put 3 commas",
			args: "1234567890",
			want: "1,234,567,890",
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
