package main

import (
	"testing"
)

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "01",
			args: args{
				s: "aab",
				p: "c*a*b",
			},
			want: true,
		},
		{
			name: "02",
			args: args{
				s: "ab",
				p: ".*",
			},
			want: true,
		},
		{
			name: "03",
			args: args{
				s: "aa",
				p: "a*",
			},
			want: true,
		},
		{
			name: "04",
			args: args{
				s: "aa",
				p: "a",
			},
			want: false,
		},
		{
			name: "05",
			args: args{
				s: "ab",
				p: ".*c",
			},
			want: false,
		},
		{
			name: "06",
			args: args{
				s: "aaa",
				p: "ab*a*c*a",
			},
			want: true,
		},
		{
			name: "07",
			args: args{
				s: "mississipp",
				p: "mis*is*ip*",
			},
			want: true,
		},
		{
			name: "08",
			args: args{
				s: "mississippi",
				p: "mis*is*ip*.",
			},
			want: true,
		},
		{
			name: "09",
			args: args{
				s: "i",
				p: ".",
			},
			want: true,
		},
		{
			name: "10",
			args: args{
				s: "aaa",
				p: "aaa",
			},
			want: true,
		},
		{
			name: "11",
			args: args{
				s: "aaa",
				p: "aaaa",
			},
			want: false,
		},
		{
			name: "12",
			args: args{
				s: "aaba",
				p: "ab*a*c*a",
			},
			want: false,
		},
		{
			name: "13",
			args: args{
				s: "aaca",
				p: "ab*a*c*a",
			},
			want: true,
		},
		{
			name: "14",
			args: args{
				s: "aaa",
				p: ".*",
			},
			want: true,
		},
		//{
		//	name: "15",
		//	args: args{
		//		s: "a",
		//		p: "ab*a",
		//	},
		//	want: false,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
