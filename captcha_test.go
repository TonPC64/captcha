package captcha_test

import (
	"testing"

	"github.com/TonPC64/captcha"
)

func TestCaptcha(t *testing.T) {
	type args struct {
		pat int
		n1  int
		op  int
		n2  int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "pattern 1,op 1,num 1 and 1",
			args: args{1, 1, 1, 1},
			want: "1 + ONE",
		},
		{
			name: "pattern 2,op 1,num 1 and 1",
			args: args{2, 1, 1, 1},
			want: "ONE + 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := captcha.Captcha(tt.args.pat, tt.args.n1, tt.args.op, tt.args.n2); got != tt.want {
				t.Errorf("Captcha() = %v, want %v", got, tt.want)
			}
		})
	}
}
