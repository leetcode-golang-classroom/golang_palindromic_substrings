package sol

import "testing"

func BenchmarkTest(b *testing.B) {
	name := "abc"
	for idx := 0; idx < b.N; idx++ {
		countSubstrings(name)
	}
}

func Test_countSubstrings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "abc",
			args: args{s: "abc"},
			want: 3,
		},
		{
			name: "aaa",
			args: args{s: "aaa"},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSubstrings(tt.args.s); got != tt.want {
				t.Errorf("countSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
