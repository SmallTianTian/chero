package core

import "testing"

func TestStringUnderscored(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "no any.",
			args: args{"test"},
			want: "test",
		},
		{
			name: "sigle upper",
			args: args{"testAbc"},
			want: "test_abc",
		},
		{
			name: "serial upper",
			args: args{"testABCdef"},
			want: "test_a_b_cdef",
		},
		{
			name: "last upper",
			args: args{"testA"},
			want: "test_a",
		},
		{
			name: "with _",
			args: args{"test_aA"},
			want: "test_a_a",
		},
		{
			name: "_ next is upper",
			args: args{"test_A"},
			want: "test_a",
		},
		{
			name: "with number",
			args: args{"test123A"},
			want: "test123_a",
		},
		{
			name: "with number2.",
			args: args{"test_123A"},
			want: "test_123_a",
		},
		{
			name: "with number3.",
			args: args{"test123_A"},
			want: "test123_a",
		},
		{
			name: "with chinese.",
			args: args{"test你好"},
			want: "test你好",
		},
		{
			name: "with chinese2.",
			args: args{"test你好World"},
			want: "test你好_world",
		},
		{
			name: "path split",
			args: args{"/api/v2/testAbc"},
			want: "/api/v2/test_abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringUnderscored(tt.args.str); got != tt.want {
				t.Errorf("StringUnderscored() = %v, want %v", got, tt.want)
			}
		})
	}
}
